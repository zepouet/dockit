package docker

import (
	"fmt"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"log"
	"os"
)

// Load docker env variables
func LoadConfig2() string {
	dockerClient := os.Getenv("DOCKER_HOST")

	// if var not defined, change for defaults values
	if dockerClient == "" {
		dockerClient = "unix:///var/run/docker.sock"
	}

	return dockerClient
}

// Load docker env variables
func Connect2() string {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	client, err := client.NewClient("unix:///var/run/docker.sock", "v1.24", nil, defaultHeaders)

	if err != nil {
		log.Println("lol")
		panic(err)
	}

	log.Println("Connected to docker")
	options := types.ContainerListOptions{All: true}
	containers, err := client.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Println(c.ID + " = " + c.Names[0])
	}
	return ""
}


