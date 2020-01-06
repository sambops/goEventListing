package entity

//Authenticate ... used for Authentication() in user handler
type Authenticate struct{
UserName string `json:"username"`
Password string	 `json:"password"`
}