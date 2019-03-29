package animes

type subitem struct {
	Title  string `yaml:"name"`
	Rating string `yaml:"rating"`
}

type Items struct {
	Item []subitem `yaml:"item"`
}
