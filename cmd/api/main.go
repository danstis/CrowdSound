package main

import (
	"context"
	"log"
	"os"

	"github.com/danstis/CrowdSound/internal/version"
	"github.com/danstis/CrowdSound/pkg/api"
	"github.com/uptrace/uptrace-go/uptrace"
)

const (
	ADDRESS string = ":9200"
)

func main() {
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

func initTracer() {
	serviceName := "CrowdSound API"
	serviceVersion := version.Version
	deploymentEnvironment := getDeploymentEnvironment()

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(serviceName),
		uptrace.WithServiceVersion(serviceVersion),
		uptrace.WithDeploymentEnvironment(deploymentEnvironment),
	)
}

func getDeploymentEnvironment() string {
	env := os.Getenv("UPTRACE_ENVIRONMENT")
	if env != "" {
		return env
	}
	return "Development"
}
