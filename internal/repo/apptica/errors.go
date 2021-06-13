package apptica

import (
	"fmt"
	"time"
)

type ErrInvalidDateRange struct {
	DateFrom time.Time
	DateTo   time.Time
}

func (e ErrInvalidDateRange) Error() string {
	return fmt.Sprintf("Invalid date range [%v, %v]", e.DateFrom, e.DateTo)
}

type ErrInvalidResponseFromApi struct {
	ResponseData []byte
}

func (e ErrInvalidResponseFromApi) Error() string {
	return fmt.Sprintf("Invalid response from api. Response: %s", string(e.ResponseData))
}
