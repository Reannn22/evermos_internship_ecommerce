package responder

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
	Keyword    string      `json:"keyword"`
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}
