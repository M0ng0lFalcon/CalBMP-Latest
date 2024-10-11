package model

import "gorm.io/gorm"

type CommodityDatabaseUSLEC struct {
	gorm.Model
	Classes    string `gorm:"column:Classes"`
	Parameters string `gorm:"column:parameters"`
	C1         string `gorm:"column:c1"`
	C2         string `gorm:"column:c2"`
	C3         string `gorm:"column:c3"`
	C4         string `gorm:"column:c4"`
	C5         string `gorm:"column:c5"`
	C6         string `gorm:"column:c6"`
	C7         string `gorm:"column:c7"`
	C8         string `gorm:"column:c8"`
	C9         string `gorm:"column:c9"`
	C10        string `gorm:"column:c10"`
	C11        string `gorm:"column:c11"`
	C12        string `gorm:"column:c12"`
	C13        string `gorm:"column:c13"`
	C14        string `gorm:"column:c14"`
	C15        string `gorm:"column:c15"`
	C16        string `gorm:"column:c16"`
	C17        string `gorm:"column:c17"`
	C18        string `gorm:"column:c18"`
	C19        string `gorm:"column:c19"`
	C20        string `gorm:"column:c20"`
	C21        string `gorm:"column:c21"`
	C22        string `gorm:"column:c22"`
	C23        string `gorm:"column:c23"`
	C24        string `gorm:"column:c24"`
}

func (CommodityDatabaseUSLEC) TableName() string {
	return "CommodityDatabaseUSLEC"
}
