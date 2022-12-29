package entities

type Bible struct {
	Name string `json:"name"`
	Books []BibleBook
}

type BibleBook struct {
	Abbrev string `json:"abbrev"`
	Chapters BibleChapters `json:"chapters"`
	Name string `json:"name"`
}

type BibleChapters [][]string