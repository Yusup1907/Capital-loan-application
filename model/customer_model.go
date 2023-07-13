package model

import "time"

type CustomerModel struct {
	Id           int
	Full_Name    string
	Address      string
	NIK          string
	Phone_number string
	User_Id      int
	Created_at   time.Time
	Updated_at   time.Time
}
