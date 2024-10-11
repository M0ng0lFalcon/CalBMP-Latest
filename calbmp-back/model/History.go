package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Username    string
	ProjectName string
	CreatedDate string
	CountyName  string
	ZipCode     string
	Soil        string
	Muname      string
	HistoryPath string
	Step1       string
	Step2       string
	EchartList  int
}
