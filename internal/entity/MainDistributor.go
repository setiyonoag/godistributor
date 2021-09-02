package entity

import (
	"github.com/google/uuid"
	"time"
)

type MainDistributor struct {
	IdMainDistrib       uint    `json:"IdMainDistrib" gorm:"primary_key;column:id_main_distrib"`
	NameMD				string 	`json:"nameMD"`
	AddressMD			string	`json:"addressMD"`
	LocationMD			string	`json:"locationMD"`
	StockMD				int		`json:"stockMD"`
	IsDelete         	bool  	`json:"isDelete"`
}

type MDModelView struct {
	IdMainDistrib       uint    `json:"id" gorm:"primary_key"`
	NameMD				string 	`json:"nameMD"`
	AddressMD			string	`json:"addressMD"`
	LocationMD			string	`json:"locationMD"`
	StockMD				int		`json:"stockMD"`
	IsDelete         	bool    `json:"isDelete"`
}

type SecDistributor struct {
	IdSecDistributor	int		`json:"id_sd" gorm:"primary_key"`
	NameSD				string 	`json:"nameSD"`
	AddressSD			string	`json:"addressSD"`
	LocationSD			string	`json:"locationSD"`
	StockSD				int		`json:"stockSD"`
	IsDelete         	bool  	`json:"isDelete"`
	IdMD				int		`json:"IdMainDistrib"`
	MainDistributor		[]MainDistributor	`gorm:"foreignKey:IdMainDistrib;"`
}

type SDModelView struct {
	IdSecDistributor	int		`json:"id_sd"`
	NameSD				string 	`json:"nameSD"`
	AddressSD			string	`json:"addressSD"`
	LocationSD			string	`json:"locationSD"`
	StockSD				int		`json:"stockSD"`
	IsDelete         	bool	 `json:"isDelete"`
}

type Artikel struct {
	ID        uuid.UUID  `gorm:"column:id" json:"artikel_id"`
	Title     string     `gorm:"column:title" json:"title"`
	Kontent   string     `gorm:"column:kontent" json:"kontent"`
	Komentar  []Komentar `gorm:"Foreignkey:Artikel_ID;association_foreignkey:ID;" json:"komentar"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

type Komentar struct {
	ID         uuid.UUID  `gorm:"column:id" json:"komentar_id"`
	Artikel_ID string     `gorm:"column:artikel_id" json:"artikel_id"`
	Komentar   string     `gorm:"column:komentar" json:"komentar"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"-"`
}