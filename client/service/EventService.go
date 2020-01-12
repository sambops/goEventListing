package service



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



