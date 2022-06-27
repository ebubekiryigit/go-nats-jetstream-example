package models

type Review struct {
	Id      string `json:"_id"`
	Author  string `json:"author"`
	Store   string `json:"store"`
	Text    string `json:"text"`
	Rating  int    `json:"rating"`
	Created string `json:"created"`
}
