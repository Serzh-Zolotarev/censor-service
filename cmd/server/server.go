package main

import (
	"censor-service/pkg/api"
	"censor-service/pkg/censor"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	HostURL string `json:"host_url"`
}

func main() {
	confFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var conf config
	err = json.Unmarshal(confFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	apiDb := api.New(censor.Validate)

	err = http.ListenAndServe(conf.HostURL, apiDb.Router())
	if err != nil {
		log.Fatal(err)
	}
}
