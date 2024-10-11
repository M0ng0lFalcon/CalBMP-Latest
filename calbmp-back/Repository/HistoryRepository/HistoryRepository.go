package HistoryRepository

import (
	"calbmp-back/Database"
	"calbmp-back/Params/HistoryParams"
	"calbmp-back/model"
)

func InsertHistory(historyRec HistoryParams.HistoryRec) {
	db := Database.GetDB()

	username := historyRec.Username
	step1 := historyRec.Step1
	step2 := historyRec.Step2
	CreatedTime := historyRec.CreatedTime
	EchartList := historyRec.EchartList

	History := model.History{
		Username:    username,
		ProjectName: historyRec.ProjectName,
		CreatedDate: CreatedTime,
		CountyName:  historyRec.County,
		ZipCode:     historyRec.ZipCode,
		Soil:        historyRec.CompName,
		Muname:      historyRec.Muname,
		Step1:       step1,
		Step2:       step2,
		EchartList:  EchartList,
	}

	db.Save(&History)
}

func FindByUsername(Username string) (HistoryList []model.History) {
	db := Database.GetDB()
	db.Where("username=?", Username).Find(&HistoryList)

	return HistoryList
}
