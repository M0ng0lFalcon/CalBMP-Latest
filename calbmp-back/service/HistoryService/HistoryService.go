package HistoryService

import (
	"calbmp-back/Params/HistoryParams"
	"calbmp-back/Repository/HistoryRepository"
	"calbmp-back/model"
)

func AddHistoryServ(historyRec HistoryParams.HistoryRec) {
	HistoryRepository.InsertHistory(historyRec)
}

func CheckHistoryServ(Username string) (HistoryList []model.History) {
	HistoryList = HistoryRepository.FindByUsername(Username)

	return HistoryList
}
