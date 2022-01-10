package dto

type Param struct {
	Message string
}

type Reply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
