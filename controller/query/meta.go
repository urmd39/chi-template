package query

type Pagination struct {
	Page     int `json:"page" schema:"page"`
	PageSize int `json:"page_size" schema:"page_size"`
}
