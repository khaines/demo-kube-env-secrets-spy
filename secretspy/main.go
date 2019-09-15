package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		log.Info("msg","found container","id",container.ID[:10],"image",container.Image)

		details,err :=  cli.ContainerInspect(context.Background(),container.ID)
		if err != nil{
			log.Error("msg","could not inspect container", "name",container.Names, "id", container.ID)
		}
		envs:= details.Config.Env
		if len(envs) > 0 {
			log.Info("msg","found environment variables assigned to container", "count",len(envs))
			for _, env := range envs {
				log.Info("msg",env)
			}
		}
	}



}