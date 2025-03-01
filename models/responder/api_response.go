package responder

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   *string     `json:"errors"`
	Data    interface{} `json:"data"`
}
