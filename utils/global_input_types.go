package utils

type UpdateUserInput struct {
	Fname string `form:"fname"`
	Lname string `form:"lname"`
	Admin bool   `json:"admin"`
}
