package docker

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"log"
	"os"
)

// Load docker env variables
func LoadConfig() string {
	dockerClient :=			os.Getenv("DOCKER_HOST")

	// if var not defined, change for defaults values
	if dockerClient == "" {
		dockerClient = "unix:///var/run/docker.sock"
	}

	return dockerClient
}

// Load docker env variables
func Connect() string {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	client, err := client.NewClient("unix:///var/run/docker.sock", "v1.24", nil, defaultHeaders)

	log.Println("Connected to docker")
	if err != nil {
		panic(err)
	}

	options := types.ContainerListOptions{All: true}
	containers, err := client.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	//At a minimum specifiy a message to display to end-user.
	note := gosxnotifier.NewNotification("Check your Apple Stock!")
	note.Push()

	for _, c := range containers {
		fmt.Println(c.ID + " = " + c.Names[0])
	}
	return ""
}


