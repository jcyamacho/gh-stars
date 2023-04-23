package domain

type Repo struct {
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"`
	URL         string   `json:"url"`
	Homepage    string   `json:"homepage"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
}
