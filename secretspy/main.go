package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"os/signal"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv,client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		log.Info(fmt.Sprintf("found container. id: %v image: %v",container.ID[:10],container.Image))

		details,err :=  cli.ContainerInspect(context.Background(),container.ID)
		if err != nil{
			log.Error(fmt.Sprintf("could not inspect container. id: %v image: %v",container.ID[:10],container.Image))
		}
		envs:= details.Config.Env
		if len(envs) > 0 {
			log.Info(fmt.Sprintf("found environment variables assigned to container: count=%v", len(envs)))
			for _, env := range envs {
				log.Info(env)
			}
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,os.Kill)

	// Block until a signal is received.
	<-c

}