package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)


func main(){
	secretValue := os.Getenv("KUBE_SECRET")

	if secretValue != ""{
		log.Info(fmt.Sprintf("found value in KUBE_SECRET ENV"))
	}else{
		log.Fatal(fmt.Sprintf("Could not find a value in KUBE_SECRET. Check kube spec settings"))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,os.Kill)

	// Block until a signal is received.
	<-c
}