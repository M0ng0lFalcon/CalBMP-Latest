package HistoryController

import (
	"calbmp-back/Params/HistoryParams"
	"calbmp-back/Res"
	"calbmp-back/service/HistoryService"
	"github.com/gin-gonic/gin"
	"log"
)

func AddHistory(ctx *gin.Context) {
	// get history params
	var historyRec HistoryParams.HistoryRec
	err := ctx.BindJSON(&historyRec)
	if err != nil {
		log.Println("[!] Add history err:", err)
	}

	// add history
	HistoryService.AddHistoryServ(historyRec)

	Res.SuccessMsg(ctx, "[*] Add history for:"+historyRec.Username)
}

func CheckHistory(ctx *gin.Context) {
	// get username
	Username := ctx.Query("username")

	HistoryList := HistoryService.CheckHistoryServ(Username)

	Res.Success(ctx, gin.H{"history_list": HistoryList}, "[*] Get Histories")
}
