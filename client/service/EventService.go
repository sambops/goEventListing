package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goEventListing/client/entity"
)

var baseURL = "http://localhost:8081/el/" //DevSkim: ignore DS137138 until 2020-02-10

//AllEvents ... returns all events  /event/allevents
func AllEvents() ([]entity.Event, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%s", baseURL, "event", "allevent")
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := []entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, eventData)
	if err != nil {
		return nil, err
	}
	return eventData, nil
}

//Event ... returns all Reviews  /Reviews/
func Event(id uint) (*entity.Event, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s/%d", baseURL, "event", "event", id)
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, eventData)
	if err != nil {
		return nil, err
	}
	return &eventData, nil
}
