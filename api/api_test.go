package api

import (
	"github.com/akolybelnikov/go_serialized/etl"
	"github.com/akolybelnikov/go_serialized/utils"
	"github.com/google/uuid"
	"log"
	"path/filepath"
	"testing"
)

func TestApi(t *testing.T) {
	messageUUID := uuid.New().String()
	url := "https://api.serialized.io/aggregates/message/" + messageUUID + "/events"

	filename := filepath.Join("../inputs", "message_one.json")
	msg := etl.ReadData(filename)
	passphrase := utils.MakePassphrase()
	encryptedMessage := msg.EncodeData(passphrase)

	eventUUID := uuid.New().String()
	eventBody := PostBody{
		EventId:   eventUUID,
		EventType: "SaveMessageEvent",
		Data: PostMessage{
			EncryptedMessage: string(encryptedMessage.Message),
			Passphrase:       string(encryptedMessage.Passphrase),
		},
	}

	response := eventBody.Call("POST", url)

	if response.Result != "SUCCESS" {
		log.Fatalf("received result %s for message with id %s", response.Result, messageUUID)
	}
}
