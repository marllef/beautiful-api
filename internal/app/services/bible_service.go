package services

import (
	"bytes"
	"encoding/json"
	"io"
	entities "marllef/beautiful-api/internal/app/models/entity"
	"marllef/beautiful-api/pkg/path"
	"os"
	"strconv"
	"strings"

)

type BibleService interface {
	GetAllProducts() ([]*entities.Product, error)
	GetOneProduct(id int64) (*entities.Product, error)
	GetSingleVerse(book string, chapter int64, verse int64) string
}

type bibleService struct {
	books    []entities.BibleBook
	chapters map[string]string
	BibleService
}

func LoadBible() (*bibleService, error) {
	var books []entities.BibleBook

	file, err := os.Open(path.Root + "/resources/bibles/nvi.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileJson, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	value := bytes.TrimPrefix(fileJson, []byte("\xef\xbb\xbf"))

	json.Unmarshal(value, &books)

	chapters := make(map[string]string)
	for _, book := range books {
		chapters[strings.ToLower(book.Name)] = book.Abbrev
	}

	return &bibleService{
		books:    books,
		chapters: chapters,
	}, nil
}

func (service *bibleService) GetSingleVerse(book string, chapter int64, verse int64) string {

	if bookIndex, err := strconv.ParseInt(book, 10, 32); err == nil && len(service.books[bookIndex].Chapters) >= int(chapter) && len(service.books[bookIndex].Chapters[chapter-1]) >= int(verse) {
		return service.books[bookIndex].Chapters[chapter-1][verse-1]
	}

	if abbrev, exists := service.chapters[book]; exists {
		book = abbrev
	}

	if index := service.findByAbbrev(book); index != nil && len(service.books[*index].Chapters) >= int(chapter) && len(service.books[*index].Chapters[chapter-1]) >= int(verse) {
		verse_txt := service.books[*index].Chapters[chapter-1][verse-1]
		
		return verse_txt
	}
	
	return "Não encontrado."
}

func (service *bibleService) findByAbbrev(abbrev string) *int {
	for i, book := range service.books {
		if book.Abbrev == abbrev {
			return &i
		}
	}
	return nil
}
