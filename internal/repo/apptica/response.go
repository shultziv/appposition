package apptica

type dateToPosition map[string]uint32
type subCategoryToDate map[uint32]dateToPosition
type categoryToSubCategory map[uint32]subCategoryToDate

type topHistoryResponse struct {
	StatusCode uint32                `json:"status_code"`
	Message    string                `json:"message"`
	Data       categoryToSubCategory `json:"data"`
}
