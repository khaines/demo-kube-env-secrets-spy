package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)


func main(){
	secretValue := os.Getenv("KUBE_SECRET")

	if secretValue != ""{
		log.Info("msg","found value in KUBE_SECRET ENV")
	}else{
		log.Fatal("msg","Could not find a value in KUBE_SECRET. Check kube spec settings")
	}
}