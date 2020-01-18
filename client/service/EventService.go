package service

<<<<<<< HEAD
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
=======


import (
	"bytes"
	"encoding/json"
	"github.com/goEventListing/client/entity"
	"io/ioutil"
	"fmt"
	"net/http"
)

var baseEventURL = "http://localhost:8181/el/event/"


//AllEvents ... handles GET  baseURL/allevents
func AllEvents() (*[]entity.Event,error){
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseEventURL,"allevent")
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238

	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	eventData := []entity.Event{}
=======
	eventData := &[]entity.Event{}
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	err = json.Unmarshal(body, eventData)
	if err != nil {
		return nil, err
	}
	return eventData, nil
}

//Event ... returns all Reviews  /Reviews/
func Event(id uint) (*entity.Event, error) {

	client := &http.Client{}
	fmt.Println("--service--getting url--") //service--getting url
	URL := fmt.Sprintf("%s%s/%s/%d", baseURL, "event", "event", id)
	fmt.Println("--service----set url--") //-service----set url
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)
	fmt.Println("---service---rqstd--for--- res--", res)//service---rqstd--for--- res

	if err != nil {
		return nil, err
	}

	eventData := entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("--entit .evntdata--", eventData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, eventData)
	if err != nil {
		return nil, err
	}
	fmt.Println("--json un marshal--", eventData)
	return &eventData, nil
}
=======
	err = json.Unmarshal(body,eventData)
	if err != nil{
		return nil,err
	}
	return eventData,nil
}
//GetEvent ... request on baseURL/:id
func GetEvent(id uint) (*entity.Event,error){
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d",baseEventURL,id)
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http response
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventdata := &entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,eventdata)
	if err != nil{
		return nil,err
	}
	return eventdata,nil

}
//UpcomingEvent ... request on baseURL/upcoming
func UpcomingEvent()(*[]entity.Event,error){
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseEventURL,"upcoming")
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := &[]entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,eventData)
	if err != nil{
		return nil,err
	}
	return eventData,nil
}
//AddEvent ... handles request on baseURL/create
func AddEvent(event *entity.Event)(*entity.Event, error){
	ouput,err:= json.MarshalIndent(event,"","\t\t")

	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseEventURL,"create")

	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(ouput) )
	res,err := client.Do(req)
	
	if err != nil {
		return nil, err
	}
	evvent := &entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,evvent)
	if err != nil{
		return nil,err
	}
	return evvent,nil
}
//UpdateEvent ... handlers request on baseURL/update
func UpdateEvent(event *entity.Event) (*entity.Event, error){
	ouput,err:= json.MarshalIndent(event,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%d",baseEventURL,"update",event.ID)
	req,_ := http.NewRequest("PUT",URL,bytes.NewBuffer(ouput))

	//DO return an http response
	res,err := client.Do(req)
	
	if err != nil {
		return nil,err
	}
	eventt := &entity.Event{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,eventt)
	if err != nil{
		return nil,err
	}
	return eventt,nil
}
//DeleteEvent ... request on baseURL/remove/:id
func DeleteEvent(id uint) (*entity.Event,error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d",baseEventURL,"remove",id)
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := &entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,eventData)
	if err != nil{
		return nil,err
	}
	return eventData,nil

}

//GetUserSubscribedEvents ... request on baseURL/foru/:id
func GetUserSubscribedEvents(id uint)(*[]entity.Event,error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d",baseEventURL,"foru",id)
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := &[]entity.Event{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,eventData)
	if err != nil{
		return nil,err
	}
	return eventData,nil
}



>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
