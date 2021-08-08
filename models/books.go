package models

type Book struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
	Path   string   `json:"path"`
}

type Books struct {
	Books []Book `json:"books"`
}
