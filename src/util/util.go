package util

import (
	"io/ioutil"
	"log"

	"../api/animes"
	"../api/computers"
	"../api/handphones"
	"../api/mangas"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func CompileJSON(serviceName string, responsePayload *[]byte) {
	var t interface{}

	switch serviceName {
	case "animes":
		t = animes.Items{}
	case "computers":
		t = computers.Items{}
	case "handphones":
		t = handphones.Items{}
	case "mangas":
		t = mangas.Items{}
	}

	data, err := ioutil.ReadFile("src/data/" + serviceName + ".yml")
	Check(err)

	err = yaml.Unmarshal([]byte(data), &(t))
	Check(err)

	jsonData, err := jsoniter.Marshal(t)
	Check(err)

	*responsePayload = jsonData
}
