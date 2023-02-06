package main

import (
	"os"

	"k8s.io/component-base/cli"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"github.com/Fish-pro/sample-scheduler/pkg/sample"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(sample.Name, sample.New),
	)
	code := cli.Run(command)
	os.Exit(code)
}
