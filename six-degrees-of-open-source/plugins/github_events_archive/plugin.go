package github_events_archive

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-github/v72/github"
	"github.com/warpstreamlabs/bento/public/service"
)

func inputSpec() *service.ConfigSpec {
	return service.NewConfigSpec().Fields(service.NewAutoRetryNacksToggleField())
}

func newInput(conf *service.ParsedConfig) (service.BatchInput, error) {
	return service.AutoRetryNacksBatchedToggled(conf, &githubArchiveInput{
		// TODO: hard coded for now
		nextArchiveTime: time.Now().Add(-1 * 7 * 24 * time.Hour).UTC(),
	})
}

func init() {
	err := service.RegisterBatchInput("github_events_archive", inputSpec(),
		func(pConf *service.ParsedConfig, res *service.Resources) (service.BatchInput, error) {
			return newInput(pConf)
		})
	if err != nil {
		panic(err)
	}

}

type githubArchiveInput struct {
	nextArchiveTime time.Time

	trackedBatchesLock sync.Mutex
	trackedBatches     []string
}

func (i *githubArchiveInput) Connect(ctx context.Context) error {
	return nil
}

func (i *githubArchiveInput) ReadBatch(ctx context.Context) (service.MessageBatch, service.AckFunc, error) {
	for {
		i.trackedBatchesLock.Lock()
		if len(i.trackedBatches) == 0 {
			i.trackedBatchesLock.Unlock()
			break
		}
		fmt.Println("There is back preasure, waiting...")
		i.trackedBatchesLock.Unlock()
		time.Sleep(5 * time.Second)
	}

	fmt.Println("No backpreasure starting")

	i.trackedBatchesLock.Lock()
	defer i.trackedBatchesLock.Unlock()

	fmt.Println("Got the lock")

	batch := make(service.MessageBatch, 0)

	archiveURL := fmt.Sprintf("https://data.gharchive.org/%s-%d.json.gz", i.nextArchiveTime.Format("2006-01-02"), i.nextArchiveTime.Hour())
	fmt.Printf("Downloading archive from %s\n", archiveURL)
	resp, err := http.Get(archiveURL)
	if err != nil {
		return nil, nil, fmt.Errorf("error making http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// archive not found, so sleep 1 min
		time.Sleep(1 * time.Minute)
		return batch, func(context.Context, error) error { return nil }, nil
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading body during error: %w", err)
		}
		return nil, nil, fmt.Errorf("error response from server (status code %d): %s", resp.StatusCode, string(bodyBytes))
	}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating gzip reader: %w", err)
	}
	defer gzipReader.Close()

	jsonDecoder := json.NewDecoder(gzipReader)

	var startID *string
	var endID string

	for {
		var event github.Event
		if err := jsonDecoder.Decode(&event); err != nil {
			if err == io.EOF {
				break // End of stream
			}

			return nil, nil, fmt.Errorf("error decoding json object: %w", err)
		}

		jsonData, err := json.Marshal(event)
		if err != nil {
			return nil, nil, fmt.Errorf("error json marshaling github event %d: %w", event.ID, err)
		}

		msg := service.NewMessage(jsonData)
		batch = append(batch, msg)

		if startID == nil {
			startID = event.ID
		}

		endID = *event.ID

		i.trackedBatches = append(i.trackedBatches, *event.ID)
		// fmt.Printf("Read Event with ID %s and create time %s\n", *event.ID, event.CreatedAt.String())
		// break
	}

	i.nextArchiveTime = i.nextArchiveTime.Add(1 * time.Hour)

	fmt.Printf("Sending batch with size %d Start ID: %s End ID: %s\n", len(batch), *startID, endID)
	return batch, func(ctx context.Context, err error) error {
		fmt.Printf("starting ack func: %v\n", err)

		i.trackedBatchesLock.Lock()
		defer i.trackedBatchesLock.Unlock()

		i.trackedBatches = i.trackedBatches[:0]

		fmt.Println("finished ack func")
		return nil
	}, nil
}

func (i *githubArchiveInput) Close(ctx context.Context) error {
	return nil
}
