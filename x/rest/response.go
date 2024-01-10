package rest

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}
