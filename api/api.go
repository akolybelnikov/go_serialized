package api

import (
	"bytes"
	"encoding/json"
	"github.com/akolybelnikov/go_serialized/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
)

type PostMessage struct {
	EncryptedMessage string `json:"encryptedMessage"`
	Passphrase       string `json:"passphrase"`
}

type PostBody struct {
	EventId   string      `json:"eventId"`
	EventType string      `json:"eventType"`
	Data      PostMessage `json:"data"`
}

type PostData struct {
	Events []PostBody `json:"events"`
}

type SerializedResponse struct {
	AggregateVersion int    `json:"aggregateVersion"`
	Result           string `json:"result"`
	TaskId           string `json:"taskId"`
}

func (b PostBody) Call(method string, url string) *SerializedResponse {
	err := godotenv.Load("../.env")
	utils.LogErrors(err)

	myEnv, err := godotenv.Read("../.env")
	utils.LogErrors(err)

	var events []PostBody
	events = append(events, b)

	postData := &PostData{Events: events}

	jsonData, err := json.Marshal(postData)
	utils.LogErrors(err)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonData))
	utils.LogErrors(err)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Serialized-Access-Key", myEnv["ACCESS_KEY"])
	req.Header.Add("Serialized-Secret-Access-Key", myEnv["SECRET_ACCESS_KEY"])

	resp, err := client.Do(req)
	utils.LogErrors(err)

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	utils.LogErrors(err)

	var response SerializedResponse
	err = json.Unmarshal(bodyBytes, &response)
	utils.LogErrors(err)

	return &response
}
