package dto

type Args struct {
	Message string `json:"message"`
}

type Reply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
