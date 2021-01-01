package utils

type GithubErrorResponse struct {
	StatusCode       int    `json:"status_code"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}
