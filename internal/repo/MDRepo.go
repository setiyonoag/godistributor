package repo

import (
	"godistributor/internal/entity"
	"gorm.io/gorm"
)

type MDRepo struct {
	db *gorm.DB
}

type InterfaceMDRepo interface {
	GetMD() ([]entity.MainDistributor, error)
	GetOneMD(int) (*entity.MainDistributor, error)
}

func NewMDRepo(db *gorm.DB) *MDRepo {
	var mDRepo = MDRepo{}
	mDRepo.db = db
	return &mDRepo
}

func (r *MDRepo) GetMD() ([]entity.MainDistributor, error) {
	var mD []entity.MainDistributor
	err := r.db.Order("id_main_distributor desc").Find(&mD).Error
	if err != nil {
		return nil, err
	}
	return mD, err
}

func (r *MDRepo) GetOneMD(id int) (*entity.MainDistributor, error) {
	var mD *entity.MainDistributor
	err := r.db.Where("id_main_distributor = ?", id).Find(&mD).Error
	if err != nil {
		return nil, err
	}
	return mD, err
}
