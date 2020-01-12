package entity

//Authenticate ... used for sending username and password to the server for authentication
type Authenticate struct{
	Name string
	Pass string
}