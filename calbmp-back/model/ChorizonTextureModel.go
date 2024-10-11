package model

type Chorizon_texture struct {
	Objectid    string `gorm:"column:objectid"`
	Texture     string `gorm:"column:texture"`
	Stratextsf  string `gorm:"column:stratextsf"`
	Rvindicato  string `gorm:"column:rvindicato"`
	Texdesc     string `gorm:"column:texdesc"`
	Chkey       string `gorm:"column:chkey"`
	Chtgkey     string `gorm:"column:chtgkey"`
	Oid_1       string `gorm:"column:oid_1"`
	Objectid_1  string `gorm:"column:objectid_1"`
	Texcl       string `gorm:"column:texcl"`
	Lieutex     string `gorm:"column:lieutex"`
	Chtgkey_1   string `gorm:"column:chtgkey_1"`
	Chtkey      string `gorm:"column:chtkey"`
	Ks_m_s_10_6 string `gorm:"column:ks_m_s_10-6"`
	Sav_m       string `gorm:"column:sav_m"`
	Os_m3_m3    string `gorm:"column:os_m3_m3"`
}

func (Chorizon_texture) TableName() string {
	return "chorizon_texture"
}
