package InputService

import (
	"bufio"
	"calbmp-back/Params/InputParams"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/StringUtil"
	"io"
	"log"
	"regexp"
)

// RepPlaceHolder2Fun : first step
func RepPlaceHolder2Fun(
	TemplateFilePath string,
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	username string,
) (map[string]string, error) {
	resMap := make(map[string]string)

	// regex variable
	reg, _ := regexp.Compile(`\{.+\}`)

	// open file
	fi := FileUtil.OpenFileAsRead(TemplateFilePath)

	// read template file line by line
	buf := bufio.NewReader(fi)
	for {
		// get one line
		line, errRead := buf.ReadString('\n')

		// flag == true : Matched Record item
		flag := reg.MatchString(line)
		if flag {
			// convert line to like : 1,2,3,4,C1,C2
			RecordId := StringUtil.ReLineCmd(line)

			// run fun to generate record
			line = handleFun(RecordId, rec, gv, username) + "\n"

			// add specific line to resMap
			resMap[RecordId] = line
		}

		if errRead != nil {
			if errRead == io.EOF {
				break
			} else {
				log.Fatalln("Read file error : ", errRead)
				return nil, errRead
			}
		}

	}

	_ = fi.Close()
	return resMap, nil
}

// handleFun : third step
func handleFun(
	funName string,
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	username string,
) string {
	var res string
	switch funName {
	case "CurrentTime":
		res = GetCurrentTime()
	case "A1":
		res = GetRecordA1(rec, username)
	case "A2":
		res = GetRecordA2(rec, username)
	case "A3":
		res = GetRecordA3()
	case "1":
		res = GetRecord1(rec)
	case "2":
		res = GetRecord2()
	case "3":
		res = GetRecord3(rec, username)
	case "4":
		res = GetRecord4()
	case "5":
		res = GetRecord5(rec, gv)
	case "6":
		res = GetRecord6()
	case "7":
		res = GetRecord7(rec)
	case "8":
		res = GetRecord8()
	case "9":
		res = GetRecord9(rec, gv)
	case "14":
		res = GetRecord14(rec)
	case "15":
		res = GetRecord15(rec, gv, username)
	case "16":
		res = GetRecord16(rec)
	case "17":
		res = GetRecord17()
	case "C1":
		res = GetRecordC1(rec)
	case "C2":
		res = GetRecordC2(rec)
	case "C3":
		res = GetRecordC3(rec)
	case "C4":
		res = GetRecordC4(rec)
	case "C5":
		res = GetRecordC5(rec)
	case "C6":
		res = GetRecordC6(rec)
	case "C7":
		res = GetRecordC7(rec)
	case "C7A":
		res = GetRecordC7A(rec)
	case "C7B":
		res = GetRecordC7B(rec)
	case "C7C":
		res = GetRecordC7C(rec)
	case "C7D":
		res = GetRecordC7D()
	case "C7E":
		res = GetRecordC7E(rec)
	case "C8":
		res = GetRecordC8(rec)
	case "C9":
		res = GetRecordC9(rec)
	case "U1":
		res = GetRecordU1(rec)
	case "U2":
		res = GetRecordU2(rec, gv)
	}
	return res
}
