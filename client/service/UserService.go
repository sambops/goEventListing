package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"github.com/goEventListing/client/entity"
)

var baseUserURL = "http://localhost:8181/el/user/"



//GetUser ... request on baseUserURL/id
func GetUser(id uint) (*entity.User,error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%d",baseUserURL,id)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http response
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	userdata := &entity.User{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userdata)
	if err != nil{
		return nil,err
	}
	return userdata,nil
}
//GetUsers ... request on baseUserURL
func GetUsers() (*[]entity.User,error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s",baseUserURL)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http response
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	userdata := &[]entity.User{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userdata)
	if err != nil{
		return nil,err
	}
	return userdata,nil
}

//GetUserByUserName ... request on baseUserURL/userName
func GetUserByUserName(userName string) (*entity.User,error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s",baseUserURL,"username",userName)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		fmt.Println("unexpected")
		return nil, err
	}
	userdata := &entity.User{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("another one")
		return nil, err
	}
	err = json.Unmarshal(body,userdata)
	if err != nil{
		return nil,err
	}
	return userdata,nil
}
//DeleteUser ... request on baseUserURL/remove
func DeleteUser(id uint)(*entity.User,error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d",baseUserURL,"remove",id)
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	userdata := &entity.User{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userdata)
	if err != nil{
		return nil,err
	}
	return userdata,nil

}

//RegisterUser ... request on baseUserURL/register
func RegisterUser(user *entity.User)(*entity.User,error){
	ouput,err:= json.MarshalIndent(user,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseUserURL,"register")

	//we use bytes.NewBuffer which gives us a bytes buffer based on our bytes slice.
	// This buffer is both readable and writable.
	// It’s “readable” part satisfies the io.Reader interface and serves our purpose.
	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(ouput) )
	//DO return an http responce
	res,err := client.Do(req)
	
	if err != nil {
		return nil, err
	}

	userr := &entity.User{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userr)
	if err != nil{
		return nil,err
	}
	return userr,nil
}
//EditUser ... request on baseUserURL/edit/:id
func EditUser(user *entity.User)(*entity.User,error){
	ouput,err:= json.MarshalIndent(user,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%d",baseUserURL,"edit",user.ID)
	req,_ := http.NewRequest("PUT",URL,bytes.NewBuffer(ouput))
	//DO return an http response
	res,err := client.Do(req)
	
	if err != nil {
		return nil,err
	}
	userr := &entity.User{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userr)
	if err != nil{
		return nil,err
	}
	return userr,nil
}
//AuthenticateUser .... request on baseUserURL/login
func AuthenticateUser(userNamee string,password string)(*entity.User,error){

	authenticate := &entity.Authenticate{UserName:userNamee,Password:password}
	
	ouput,err:= json.MarshalIndent(authenticate,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseUserURL,"login")
	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(ouput))
	//DO return an http responce
	res,err := client.Do(req)
	
	if err != nil {
		return nil,err
	}

	userr := &entity.User{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userr)
	if err != nil{
		return nil,err
	}
	return userr,nil
}

//UserRoles ... baseUserURL/role/:user
func UserRoles(user *entity.User)([]entity.Role,error){
	ouput,err:= json.MarshalIndent(user,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseUserURL,"role")

	//we use bytes.NewBuffer which gives us a bytes buffer based on our bytes slice.
	// This buffer is both readable and writable.
	// It’s “readable” part satisfies the io.Reader interface and serves our purpose.
	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(ouput) )
	//DO return an http responce
	res,err := client.Do(req)
	
	if err != nil {
		return nil, err
	}

	role := []entity.Role{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,role)
	if err != nil{
		return nil,err
	}
	return role,nil
}
//UserByEmail ... baseUserURL/:email
func UserByEmail(email string) (*entity.User, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s",baseUserURL,email)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)
	if err != nil {
		return nil, err
	}
	userdata := &entity.User{}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,userdata)
	if err != nil{
		return nil,err
	}
	return userdata,nil
}

//PhoneExists ... /check/phone/:phone
func PhoneExists(phone string) *bool{
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s",baseUserURL,phone)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)
	if err != nil {
		return nil
	}
	var boolvalue *bool
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body,boolvalue)
	if err != nil{
		return nil
	}
	return boolvalue
}
//EmailExists ... /check/email/:email
func EmailExists(email string) *bool{
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s",baseUserURL,email)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)
	if err != nil {
		return nil
	}
	var boolvalue *bool
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body,boolvalue)
	if err != nil{
		return nil
	}
	return boolvalue
}






