package domain

import "time"

type AppPosition struct {
	AppId         uint32
	CountryId     uint32
	CategoryId    uint32
	SubCategoryId uint32
	Date          time.Time
	Position      uint32
}
