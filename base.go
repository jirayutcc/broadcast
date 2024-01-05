package broadcast

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Base interface {
	Broadcast(message *BroadcastTransaction) (BroadcastTransactionResponse, error)
	Monitor(message *GetTransaction) (GetTransactionResponse, error)
}

type base struct {
}

func InitBaseAPI() Base {
	return base{}
}

func (b base) Broadcast(message *BroadcastTransaction) (response BroadcastTransactionResponse, err error) {
	if message == nil {
		err = ErrorRequestInvalid
		return response, err
	}

	url := fmt.Sprintf("%s/broadcast", NodeUrl)

	postBody, err := json.Marshal(*message)
	if err != nil {
		return
	}

	responseBody := bytes.NewBuffer(postBody)

	rawResponse, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return
	}

	defer rawResponse.Body.Close()

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return response, nil
}

func (b base) Monitor(message *GetTransaction) (response GetTransactionResponse, err error) {
	if message == nil {
		err = ErrorRequestInvalid
		return
	}

	url := fmt.Sprintf("%s/check/%s", NodeUrl, message.TxHash)

	rawResponse, err := http.Get(url)
	if err != nil {
		return
	}

	defer rawResponse.Body.Close()

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}