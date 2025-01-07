package resp

// SuccessRes success response struct.
type SuccessRes struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// ListRes pagination response struct.
type ListRes struct {
	Total int `json:"total"`
	List  any `json:"list"`
}

// ErrorRes error response struct.
type ErrorRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
	Msg    string `json:"msg"`
}
