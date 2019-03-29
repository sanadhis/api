package computers

type subitem struct {
	Type string `yaml:"name"`
	Year string `yaml:"year"`
}

type Items struct {
	Item []subitem `yaml:"item"`
}
