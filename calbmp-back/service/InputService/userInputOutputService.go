package InputService

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/util/GlobalValueUtil"
	"fmt"
)

func GetRecordU1(rec InputParams.UserInputStepReceiver) string {
	pesticideCnt := rec.CntPesticide
	return fmt.Sprintf("%d", 8*pesticideCnt+5+10)
}

func GetRecordU2(
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
) string {
	/*
			Used variables:
			1. RUNF -> runoff
			2. PRCP -> precipitation
			3. IRRG -> irrigation
			4. ESLS -> eroded solids
		    ------- chemical -----------
			5. RFLX -> loading in runoff
			6. EFLX -> loading in erosion
			7. VFLX(FPVL, VFLX) -> loading in volatilization
	*/
	res := ""
	// chemical part
	pesticideCnt := rec.CntPesticide
	templates := []string{
		"RFLX,%d,TSER,   0,  0,  1",
		"RFLX,%d,TCUM,   0,  0,  1",
		"EFLX,%d,TSER,   0,  0,  1",
		"EFLX,%d,TCUM,   0,  0,  1",
		"FPVL,%d,TSER,   0,  0,  1",
		"FPVL,%d,TCUM,   0,  0,  1",
		"VFLX,%d,TSER,   1,  1,  1",
		"VFLX,%d,TCUM,   1,  1,  1",
	}
	for i := 1; i <= pesticideCnt; i++ {
		for _, v := range templates {
			res += fmt.Sprintf(v, i)
			res += "\n"
		}
	}
	// unused variable
	res += fmt.Sprintf("INFL,0,TSER,   %3d,  %3d,  1.0\n", gv.U2Num, gv.U2Num)

	var lines []string
	lines = append(lines, "COFX,1,TSER,   %3d,  %3d,  100000\n")
	lines = append(lines, "COFX,1,TCUM,   %3d,  %3d,  100000\n")
	lines = append(lines, "TCON,1,TSER,   %3d,  %3d,  1000\n")
	lines = append(lines, "TCON,1,TCUM,   %3d   %3d,  1000\n")
	lines = append(lines, "SWTR,1,TSER,   %3d   %3d,  1.0")

	for _, v := range lines {
		res += fmt.Sprintf(v, gv.U2Num-1, gv.U2Num)
	}
	return res
}
