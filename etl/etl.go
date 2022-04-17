package etl

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"github.com/akolybelnikov/go_serialized/utils"
	"io"
	"io/ioutil"
)

type Message struct {
	Body string `json:"body"`
}

type EncryptedMessage struct {
	Passphrase []byte `json:"passphrase"`
	Message    []byte `json:"message"`
}

func ReadData(filename string) (message Message) {
	data, err := ioutil.ReadFile(filename)
	utils.LogErrors(err)

	err = json.Unmarshal(data, &message)
	utils.LogErrors(err)

	return message
}

func (m Message) EncodeData(salt []byte) EncryptedMessage {
	text, key := []byte(m.Body), salt

	c, err := aes.NewCipher(key)
	utils.LogErrors(err)

	gcm, err := cipher.NewGCM(c)
	utils.LogErrors(err)

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	utils.LogErrors(err)

	return EncryptedMessage{
		Passphrase: key,
		Message:    gcm.Seal(nonce, nonce, text, nil),
	}
}

func (m EncryptedMessage) DecodeData() string {
	key := m.Passphrase
	cipherText := m.Message

	c, err := aes.NewCipher(key)
	utils.LogErrors(err)

	gcm, err := cipher.NewGCM(c)
	utils.LogErrors(err)

	nonceSize := gcm.NonceSize()

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	utils.LogErrors(err)

	return string(plaintext)
}
