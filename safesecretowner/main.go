package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/signal"
)


func main(){
	// still using the same ENV name as before, but now we expect the value to be a file path
	// to the location the actual secret data is residing.
	secretFilePath := os.Getenv("KUBE_SECRET")

	if secretFilePath != ""{
		log.Info(fmt.Sprintf("found value in KUBE_SECRET ENV: %s",secretFilePath))
	}else{
		log.Fatal(fmt.Sprintf("Could not find a value in KUBE_SECRET. Check kube spec settings"))
		return
	}

	if data,err:= ioutil.ReadFile(secretFilePath); err != nil {
		log.Fatal(fmt.Sprintf("Could not read path specified in ENV (%s): %s",secretFilePath,err.Error()))
		return
	} else if data != nil && len(data)> 0{
		log.Info(fmt.Sprintf("Secret data loaded from %s",secretFilePath))
	} else{
		log.Error(fmt.Sprintf("Found no data in %s", secretFilePath))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,os.Kill)

	// Block until a signal is received.
	<-c
}