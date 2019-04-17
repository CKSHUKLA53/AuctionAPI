package model

type Client struct {
	Id         int    `gorm:"primary_key";"AUTO_INCREMENT"`
	ClientName string `json:"client_name"`
}
