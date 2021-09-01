package response

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}
