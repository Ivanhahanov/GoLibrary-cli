package models

type Book struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

type Books struct {
	Books []Book `json:"books"`
}

type Description struct {
	ID              string   `yaml:"id"`
	Title           string   `yaml:"title"`
	Author          string   `yaml:"author"`
	Publisher       string   `yaml:"publisher"`
	Tags            []string `yaml:"tags"`
	BookDescription string   `yaml:"description"`
}
