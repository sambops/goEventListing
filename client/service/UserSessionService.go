package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/goEventListing/API/entity"
)

var baseSessionURL = "http://localhost:8181/el/session/"

//Session ... request on  baseUserURL/session/:sid
func Session(sessionID string) (*entity.Session, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s",baseSessionURL,sessionID)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http response
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	sessionData := &entity.Session{}

	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,sessionData)
	if err != nil{
		return nil,err
	}
	return sessionData,nil	
}
//StoreSession ... request on baseSessionURL/store
func StoreSession(session *entity.Session) (*entity.Session, error){
	ouput,err:= json.MarshalIndent(session,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseSessionURL,"store")
	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(ouput) )
	//DO return an http responce
	res,err := client.Do(req)
	
	if err != nil {
		return nil, err
	}
	sess := &entity.Session{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,sess)
	if err != nil{
		return nil,err
	}
	return sess,nil
}
//DeleteSession ... request on baseSessionURL/remove/:id
func DeleteSession(sessionID string) (*entity.Session, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s",baseUserURL,"remove",sessionID)
	req,_ := http.NewRequest("POST",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	sessionData := &entity.Session{}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,sessionData)
	if err != nil{
		return nil,err
	}
	return sessionData,nil

}
