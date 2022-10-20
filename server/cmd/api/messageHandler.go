package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type MessagePayload struct {
	Msg_id int    `json:"Msg_id"`
	Sender string `json:"Sender"`
	Msg    string `json:"Msg"`
}

type jsonResp struct {
	Code          string
	Received_time time.Time
}

func (app *application) messageHandler(w http.ResponseWriter, r *http.Request) {
	var payload MessagePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	// Here to process MessagePayload

	ok := jsonResp{
		Code:          "OK",
		Received_time: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
