package apptica

type dateToPosition map[string]uint8
type subCategoryToDate map[uint8]dateToPosition
type categoryToSubCategory map[uint8]subCategoryToDate

type topHistoryResponse struct {
	StatusCode uint8                 `json:"status_code"`
	Message    string                `json:"message"`
	Data       categoryToSubCategory `json:"data"`
}
