package entity

type MainDistributor struct {
	IdMainDistrib uint   `json:"IdMainDistrib" gorm:"primary_key;column:id_main_distrib"`
	NameMD        string `json:"nameMD"`
	AddressMD     string `json:"addressMD"`
	LocationMD    string `json:"locationMD"`
	StockMD       int    `json:"stockMD"`
	IsDelete      bool   `json:"isDelete"`
}

type MDModelView struct {
	IdMainDistrib uint   `json:"id" gorm:"primary_key"`
	NameMD        string `json:"nameMD"`
	AddressMD     string `json:"addressMD"`
	LocationMD    string `json:"locationMD"`
	StockMD       int    `json:"stockMD"`
	IsDelete      bool   `json:"isDelete"`
}

type SecDistributor struct {
	IdSecDistributor int               `json:"id_sd" gorm:"primary_key"`
	NameSD           string            `json:"nameSD"`
	AddressSD        string            `json:"addressSD"`
	LocationSD       string            `json:"locationSD"`
	StockSD          int               `json:"stockSD"`
	IsDelete         bool              `json:"isDelete"`
	IdMD             int               `json:"IdMainDistrib"`
	MainDistributor  []MainDistributor `gorm:"foreignKey:IdMainDistrib;"`
}

type SDModelView struct {
	IdSecDistributor int    `json:"id_sd"`
	NameSD           string `json:"nameSD"`
	AddressSD        string `json:"addressSD"`
	LocationSD       string `json:"locationSD"`
	StockSD          int    `json:"stockSD"`
	IsDelete         bool   `json:"isDelete"`
}
