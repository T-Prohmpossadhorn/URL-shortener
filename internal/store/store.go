package store

type Service interface {
	Save(string, string) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*Item, error)
	Delete(string) error
	Close() error
}

type Item struct {
	Link    string `json:"link" redis:"link"`
	URL     string `json:"url" redis:"url"`
	Expires string `json:"expires" redis:"expires"`
	Visits  int    `json:"visits" redis:"visits"`
}

type PostItem struct {
	URL     string `json:"url"`
	Expires string `json:"expires"`
}
