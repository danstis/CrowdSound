package main

import (
	"context"
	"log"
	"os"

	rookout "github.com/Rookout/GoSDK"
	"github.com/danstis/CrowdSound/internal/version"
	"github.com/danstis/CrowdSound/pkg/api"
	"github.com/uptrace/uptrace-go/uptrace"
)

const (
	ADDRESS string = ":9200"
)

var (
	// environment stores the name of the environment we are running in.
	// It can be populated by the environment variable UPTRACE_ENVIRONMENT.
	// It defaults to "Development".
	environment string
)

func init() {
	environment = os.Getenv("UPTRACE_ENVIRONMENT")
	if environment == "" {
		environment = "Development"
	}
}

func main() {
	initRookout()
	initTracer()
	defer func() {
		err := uptrace.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()
	api := api.New()
	api.Address = ADDRESS
	log.Fatal(api.Run())
}

func initRookout() {
	err := rookout.Start(rookout.RookOptions{
		Token: os.Getenv("ROOKOUT_TOKEN"),
		Labels: map[string]string{
			"env":     environment,
			"version": version.Version,
		},
	})

	if err != nil {
		log.Fatalf("failed to start rookout: %v", err)
	}
}

func initTracer() {
	serviceName := "CrowdSound API"
	serviceVersion := version.Version
	deploymentEnvironment := environment

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(serviceName),
		uptrace.WithServiceVersion(serviceVersion),
		uptrace.WithDeploymentEnvironment(deploymentEnvironment),
	)
}
