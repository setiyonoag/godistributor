package service

import (
	"godistributor/internal/entity"
	"godistributor/internal/repo"
)

type SDService struct {
	sDRepo repo.InterfaceSDRepo
}

type InterfaceSDService interface {
	GetSD() ([]entity.SDModelView, error)
	GetOneSD(int) (*entity.SDModelView, error)
}

func NewSDService(sdRepo repo.InterfaceSDRepo) *SDService {
	var sDService = SDService{}
	sDService.sDRepo = sdRepo
	return &sDService
}

func (s *SDService) GetSD() ([]entity.SDModelView, error) {
	result, err := s.sDRepo.GetSD()
	if err != nil {
		return nil, err
	}

	var sds []entity.SDModelView
	for _, items := range result {
		var sd entity.SDModelView
		sd.IdSecDistributor = items.IdSecDistributor
		sd.NameSD = items.NameSD
		sd.LocationSD = items.LocationSD
		sd.StockSD = items.StockSD
		sd.IsDelete = items.IsDelete
		sds = append(sds, sd)
	}

	return sds, nil
}

func (r *SDService) GetOneSD(id int) (*entity.SDModelView, error) {
	result, err := r.sDRepo.GetOneSD(id)

	if err != nil {
		return nil, err
	}

	var sdVM entity.SDModelView

	if result != nil {
		sdVM = entity.SDModelView{
			IdSecDistributor: result.IdSecDistributor,
			NameSD: result.NameSD,
			LocationSD: result.LocationSD,
			StockSD: result.StockSD,
			IsDelete: result.IsDelete,
		}
	}

	return &sdVM, nil
}