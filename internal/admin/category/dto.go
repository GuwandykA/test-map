package category

type ReqDTO struct {
	Value []int `json:"value"`
}

type PaginationDTO struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type CategoryDTO struct {
	UUID  int           `json:"id"`
	Value []interface{} `json:"value"`
}
