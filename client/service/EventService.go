package service



import (
	"encoding/json"
	"github.com/goEventListing/client/entity"
	"io/ioutil"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080/el/"


//AllEvents ... returns all events  /event/allevents
func AllEvents() ([]entity.Event,error){
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%s",baseURL,"event","allevent")
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	eventData := []entity.Event{}

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