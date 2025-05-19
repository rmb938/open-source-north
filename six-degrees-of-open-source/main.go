package main

import (
	"context"

	// import all components with:
	_ "github.com/warpstreamlabs/bento/public/components/all"

	"github.com/warpstreamlabs/bento/public/service"
	// import your plugins:

	_ "github.com/rmb938/open-source-north/six-degrees-of-open-source/plugins/github_events"
	_ "github.com/rmb938/open-source-north/six-degrees-of-open-source/plugins/github_events_archive"
)

func main() {
	// RunCLI accepts a number of optional functions:
	// https://pkg.go.dev/github.com/warpstreamlabs/bento/public/service#CLIOptFunc
	service.RunCLI(context.Background())
}
