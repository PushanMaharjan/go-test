package models

type User struct {
	Fname string `json:"fname" form:"fname"`
	Lname string `json:"lname" form:"lname"`
}

func (u User) TableName() string {
	return "users"
}
