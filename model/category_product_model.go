package model

import "time"

type CategoryProductModel struct {
	Id int
	CategoryProductName  string
	CreateAt time.Time
	UpdateAt time.Time
}