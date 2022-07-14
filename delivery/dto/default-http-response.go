package dto

type DefaultHttpResponse struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Ok         bool                   `json:"ok"`
	Data       map[string]interface{} `json:"data"`
}
