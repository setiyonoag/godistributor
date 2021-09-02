package service

import (
	"godistributor/internal/entity"
	"godistributor/internal/repo"
)

type MDService struct {
	mDRepo repo.InterfaceMDRepo
}

type InterfaceMDService interface {
	GetMD() ([]entity.MDModelView, error)
	GetOneMD(int) (*entity.MDModelView, error)
}

func NewMDService(mDRepo repo.InterfaceMDRepo) *MDService {
	var mDService = MDService{}
	mDService.mDRepo = mDRepo
	return &mDService
}

func (s *MDService) GetMD() ([]entity.MDModelView, error) {
	result, err := s.mDRepo.GetMD()
	if err != nil {
		return nil, err
	}

	var mds []entity.MDModelView
	for _, items := range result {
		var md entity.MDModelView
		md.IdMainDistributor = items.IdMainDistributor
		md.NameMD = items.NameMD
		md.LocationMD = items.LocationMD
		md.StockMD = items.StockMD
		md.IsDelete = items.IsDelete
		mds = append(mds, md)
	}

	return mds, nil
}

func (r *MDService) GetOneMD(id int) (*entity.MDModelView, error) {
	result, err := r.mDRepo.GetOneMD(id)

	if err != nil {
		return nil, err
	}

	var mdVM entity.MDModelView

	if result != nil {
		mdVM = entity.MDModelView{
			IdMainDistributor: result.IdMainDistributor,
			NameMD: result.NameMD,
			LocationMD: result.LocationMD,
			StockMD: result.StockMD,
			IsDelete: result.IsDelete,
		}
	}

	return &mdVM, nil
}