package model

type Admin struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AdminRole int    `json:"adminRole"`
}
