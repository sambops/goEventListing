package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/goEventListing/client/entity"
)

var baseRoleURL = "http://localhost:8181/el/role/"

//Roles ... request on baseRoleURL/roles
func Roles() (*entity.Role, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s",baseRoleURL,"roles")
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http response
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	roleData :=&entity.Role{}

	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,roleData)
	if err != nil{
		return nil,err
	}
	return roleData,nil
}
//Role ... request on baseRoleURL/
func Role(id uint) (*entity.Role, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%d",baseRoleURL,id)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	roleData := &entity.Role{}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,roleData)
	if err != nil{
		return nil,err
	}
	return roleData,nil
}
//RoleByName ... request on baseRoleURL/rolebyname
func RoleByName(name string) (*entity.Role, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s",baseRoleURL,name)
	req,_ := http.NewRequest("GET",URL,nil)
	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	roleData :=&entity.Role{}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,roleData)
	if err != nil{
		return nil,err
	}
	return roleData,nil

}
//UpdateRole ... request on baseRoleURL/update
func UpdateRole(role *entity.Role) (*entity.Role, error){
	ouput,err:= json.MarshalIndent(role,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%d",baseRoleURL,"update",role.ID)
	req,_ := http.NewRequest("PUT",URL,bytes.NewBuffer(ouput))
	//DO return an http response
	res,err := client.Do(req)
	
	if err != nil {
		return nil,err
	}
	rolee := &entity.Role{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,rolee)
	if err != nil{
		return nil,err
	}
	return rolee,nil
}
//DeleteRole ... request on baseRoleURL/remove
func DeleteRole(id uint) (*entity.Role, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d",baseRoleURL,"remove",id)
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	roleData := &entity.Role{}
	body,err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,roleData)
	if err != nil{
		return nil,err
	}
	return roleData,nil
}
//StoreRole ... request on baseRoleURL/store
func StoreRole(role *entity.Role) (*entity.Role, error){
	ouput,err:= json.MarshalIndent(role,"","\t\t")
	
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s",baseRoleURL,"store")

	//we use bytes.NewBuffer which gives us a bytes buffer based on our bytes slice.
	// This buffer is both readable and writable.
	// It’s “readable” part satisfies the io.Reader interface and serves our purpose.
	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(ouput) )
	//DO return an http responce
	res,err := client.Do(req)
	
	if err != nil {
		return nil, err
	}
	rolee := &entity.Role{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,rolee)
	if err != nil{
		return nil,err
	}
	return rolee,nil

}