package BMPController

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/Res"
	"calbmp-back/service/BMPService"
	"github.com/gin-gonic/gin"
	"log"
)

func BmpScenario(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	// get params
	var rec BmpParams.BmpBasicParams
	errBind := ctx.BindJSON(&rec)
	if errBind != nil {
		log.Fatalln(errBind)
	}

	// service fun
	BMPService.BmpScenario(rec, username.(string))

	Res.SuccessMsg(ctx, "success")
}
