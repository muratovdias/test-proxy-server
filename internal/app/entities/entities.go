package entities

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      string              `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}
