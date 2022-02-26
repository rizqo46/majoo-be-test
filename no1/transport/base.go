package transport

type BaseResponse struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}
