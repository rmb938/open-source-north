package github_events

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	ghtransport "github.com/bored-engineer/github-conditional-http-transport"
	"github.com/bored-engineer/github-conditional-http-transport/memory"
	"github.com/google/go-github/v72/github"
	"github.com/warpstreamlabs/bento/public/service"
)

func inputSpec() *service.ConfigSpec {
	return service.NewConfigSpec().Fields(service.NewAutoRetryNacksToggleField())
}

func newInput(conf *service.ParsedConfig) (service.BatchInput, error) {
	return service.AutoRetryNacksBatchedToggled(conf, &githubInput{})
}

func init() {
	err := service.RegisterBatchInput("github_events", inputSpec(),
		func(pConf *service.ParsedConfig, res *service.Resources) (service.BatchInput, error) {
			return newInput(pConf)
		})
	if err != nil {
		panic(err)
	}

}

type githubInput struct {
	githubClient *github.Client

	nextAvailableBatch time.Time
}

func (i *githubInput) Connect(ctx context.Context) error {
	i.githubClient = github.NewClient(&http.Client{
		Transport: ghtransport.NewTransport(
			memory.NewStorage(),
			http.DefaultTransport,
		),
	})

	return nil
}

func (i *githubInput) ReadBatch(ctx context.Context) (service.MessageBatch, service.AckFunc, error) {
	// Block until the next batch is available
	time.Sleep(time.Until(i.nextAvailableBatch))

	batch := make(service.MessageBatch, 0)

	listOpts := &github.ListOptions{
		PerPage: 100,
		Page:    1,
	}
	events, resp, err := i.githubClient.Activity.ListEvents(ctx, listOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("error listing first page of github events: %w", err)
	}

	for _, event := range events {
		jsonData, err := json.Marshal(event)
		if err != nil {
			return nil, nil, fmt.Errorf("error json marshaling github event %d: %w", event.ID, err)
		}

		msg := service.NewMessage(jsonData)
		batch = append(batch, msg)
	}

	for resp.NextPage != 0 {
		listOpts.Page = resp.NextPage
		events, resp, err = i.githubClient.Activity.ListEvents(ctx, listOpts)
		if err != nil {
			return nil, nil, fmt.Errorf("error listing page %d of github events: %w", listOpts.Page, err)
		}

		for _, event := range events {
			jsonData, err := json.Marshal(event)
			if err != nil {
				return nil, nil, fmt.Errorf("error json marshaling github event %d: %w", event.ID, err)
			}

			msg := service.NewMessage(jsonData)
			batch = append(batch, msg)
		}
	}

	pollIntervalStr := resp.Header.Get("X-Poll-Interval")

	pollInterval, err := strconv.ParseInt(pollIntervalStr, 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing poll interval: %w", err)
	}

	i.nextAvailableBatch = time.Now().Add(time.Duration(pollInterval) * time.Second)

	return batch, func(context.Context, error) error { return nil }, nil
}

func (i *githubInput) Close(ctx context.Context) error {
	return nil
}
