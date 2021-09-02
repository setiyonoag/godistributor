package entity

type SecDistributor struct {
	IdSecDistributor int    `json:"id_sd" gorm:"primary_key"`
	NameSD           string `json:"nameSD"`
	AddressSD        string `json:"addressSD"`
	LocationSD       string `json:"locationSD"`
	StockSD          int    `json:"stockSD"`
	IsDelete         bool   `json:"isDelete"`
	IdMD             int    `json:"IdMainDistrib" gorm:"uniqueIndex"`
	//MainDistributor
	//Refer    uint
	//MainDistributor	 MainDistributor	`json:"MainDistributor"`
	//MainDistributor  []MainDistributor `gorm:"foreignKey:IdMD"`
	//Profiles []MainDistributor `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:ProfileRefer"`
	//Refer    uint
}

type MainDistributor struct {
	IdMainDistrib  uint             `gorm:"primary_key;column:id_main_distrib"`
	NameMD         string           `json:"nameMD"`
	AddressMD      string           `json:"addressMD"`
	LocationMD     string           `json:"locationMD"`
	StockMD        int              `json:"stockMD"`
	IsDelete       bool             `json:"isDelete"`
	SecDistributor []SecDistributor `gorm:"foreignKey:IdMD"`
}

type MDModelView struct {
	IdMainDistrib uint   `json:"id" gorm:"primary_key"`
	NameMD        string `json:"nameMD"`
	AddressMD     string `json:"addressMD"`
	LocationMD    string `json:"locationMD"`
	StockMD       int    `json:"stockMD"`
	IsDelete      bool   `json:"isDelete"`
}

type SDModelView struct {
	IdSecDistributor int    `json:"id_sd"`
	NameSD           string `json:"nameSD"`
	AddressSD        string `json:"addressSD"`
	LocationSD       string `json:"locationSD"`
	StockSD          int    `json:"stockSD"`
	IsDelete         bool   `json:"isDelete"`
}
