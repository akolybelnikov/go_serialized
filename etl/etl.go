package etl

import (
	"encoding/json"
	"github.com/akolybelnikov/go_serialized/utils"
	"io/ioutil"
)

type Message struct {
	Body string `json:"body"`
}

func ReadData(filename string) (message Message) {
	data, err := ioutil.ReadFile(filename)
	utils.LogErrors(err)

	err = json.Unmarshal(data, &message)
	utils.LogErrors(err)

	return message
}
