package model

import (
	"database/sql"
	"time"
)

type CustomerModel struct {
	Id       			int
	FullName 			string
	Address  			string
	NIK      			string
	Phone    			string
	NoKK	 			string
	EmergencyName		string
	EmergencyContact 	string
	LastSalary			float64
	UserId   			int
	CreateAt 			time.Time
	UpdateAt 			time.Time
}

type ValidasiCustomerModel struct{
	Id					int
	NIK      			sql.NullString
	NoKK	 			sql.NullString
	EmergencyName		sql.NullString
	EmergencyContact 	sql.NullString
	LastSalary			sql.NullFloat64
}