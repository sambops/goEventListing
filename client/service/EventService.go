package service

import (
	"bytes"
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
	fmt.Println("@@--service/EventService/Eent line:45\n url==", URL) //-service----set url
	req, _ := http.NewRequest("GET", URL, nil)

	fmt.Println("@@ service/EventService/Event line 48: req==", req)
	//DO return an http responce
	res, err := client.Do(req)
	fmt.Println("@@---service/EventService/Event line:51\n---rqstd--for--- res==", res) //service---rqstd--for--- res

	if err != nil {
		return nil, err
	}

	eventData := &entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("@@---service/EventService/Eent line:60\n---Readed the body--- body==", body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, eventData)
	if err != nil {
		return nil, err
	}
	fmt.Println("@@---service/EventService/Eent line:68--json un marshal--", eventData, "##returned- succesfully!!!")
	return eventData, nil
}

//UpcomingEvent ... request on baseURL/upcoming
func UpcomingEvent() (*[]entity.Event, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s", baseURL, "upcoming")
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := &[]entity.Event{}

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

//AddEvent ... handles request on baseURL/create
func AddEvent(event *entity.Event) (*entity.Event, error) {
	ouput, err := json.MarshalIndent(event, "", "\t\t")

	client := &http.Client{}
	URL := fmt.Sprintf("%s%s", baseURL, "create")

	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(ouput))
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	evvent := &entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, evvent)
	if err != nil {
		return nil, err
	}
	return evvent, nil
}

//UpdateEvent ... handlers request on baseURL/update
func UpdateEvent(event *entity.Event) (*entity.Event, error) {
	ouput, err := json.MarshalIndent(event, "", "\t\t")

	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%d", baseURL, "update", event.ID)
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(ouput))

	//DO return an http response
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventt := &entity.Event{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, eventt)
	if err != nil {
		return nil, err
	}
	return eventt, nil
}

//DeleteEvent ... request on baseURL/remove/:id
func DeleteEvent(id uint) (*entity.Event, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d", baseURL, "remove", id)
	req, _ := http.NewRequest("GET", URL, nil)

	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := &entity.Event{}

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

//GetUserSubscribedEvents ... request on baseURL/foru/:id
func GetUserSubscribedEvents(id uint) (*[]entity.Event, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d", baseURL, "foru", id)
	req, _ := http.NewRequest("GET", URL, nil)

	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := &[]entity.Event{}

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
