package registry

import (
	"marllef/beautiful-api/internal/app/api/controller"
	"marllef/beautiful-api/internal/app/services"
)

func (r *registry) NewBibleController() (controller.BibleController, error) {
	bibleService, err := services.LoadBible()
	if err != nil {
		return nil, err
	}
	return controller.NewBibleController(bibleService), nil
}
