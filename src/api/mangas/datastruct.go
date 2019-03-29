package mangas

type subitem struct {
	Rating string   `yaml:"rating"`
	Titles []string `yaml:"titles"`
}

type Items struct {
	Item []subitem `yaml:"item"`
}
