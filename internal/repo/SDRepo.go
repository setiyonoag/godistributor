package repo

import (
	"godistributor/internal/entity"
	"gorm.io/gorm"
)

type SDRepo struct {
	db *gorm.DB
}

type InterfaceSDRepo interface {
	GetSD() ([]entity.SecDistributor, error)
	GetOneSD(int) (*entity.SecDistributor, error)
}

func NewSDRepo(db *gorm.DB) *SDRepo {
	var sDRepo = SDRepo{}
	sDRepo.db = db
	return &sDRepo
}

func (r *SDRepo) GetSD() ([]entity.SecDistributor, error) {
	var sD []entity.SecDistributor
	err := r.db.Order("id_sec_distributor desc").Find(&sD).Error
	if err != nil {
		return nil, err
	}
	return sD, err
}

func (r *SDRepo) GetOneSD(id int) (*entity.SecDistributor, error) {
	var sD *entity.SecDistributor
	err := r.db.Where("id_sec_distributor = ?", id).Find(&sD).Error
	if err != nil {
		return nil, err
	}
	return sD, err
}
