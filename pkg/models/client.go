package model

type Client struct {
	Id         ID     `gorm:"primary_key";"AUTO_INCREMENT"`
	ClientName string `json:"client_name"`
}
