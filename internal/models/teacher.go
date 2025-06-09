package models

type Teacher struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
}
