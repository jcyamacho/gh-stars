package domain

type Repo struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"`
	URL         string   `json:"url"`
	HtmlURL     string   `json:"html_url"`
	Homepage    string   `json:"homepage"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
	Language    string   `json:"language"`
}
