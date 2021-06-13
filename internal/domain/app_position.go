package domain

import "time"

type AppPosition struct {
	AppId         uint32
	CountryId     uint8
	CategoryId    uint8
	SubCategoryId uint8
	Date          time.Time
	Position      uint8
}
