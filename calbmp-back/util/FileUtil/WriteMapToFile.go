package FileUtil

import (
	"bufio"
	"calbmp-back/util/StringUtil"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func WriteMapToFile(path string, TemplateFilePath string, content map[string]string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// regex variable
		reg, _ := regexp.Compile(`\{.+\}`)

		// open template file
		TemplateFile, errTemplate := os.OpenFile(TemplateFilePath, os.O_RDONLY, 0666)
		if errTemplate != nil {
			log.Println("Open file error! ", err)
			_ = TemplateFile.Close()
		}

		// read template file line by line
		buf := bufio.NewReader(TemplateFile)
		for {
			// get one line of template file
			line, errRead := buf.ReadString('\n')

			// flag = true : matched record item
			flag := reg.MatchString(line)
			if flag {
				// convert line to like : 1,2,3,4,C1,C2
				RecordId := StringUtil.ReLineCmd(line)

				// write specific line
				RecordLine := content[RecordId]
				_, errWrite := f.WriteString(RecordLine)
				if errWrite != nil {
					log.Fatal(errWrite)
				}
			} else {
				_, errWrite := f.WriteString(line)
				if errWrite != nil {
					log.Fatal(errWrite)
				}
			}

			if errRead != nil {
				if errRead == io.EOF {
					break
				} else {
					log.Println("Read file error : ", errRead)
				}
			}
		}

	}
	err = f.Close()
	if err != nil {
		return err
	}
	return err
}
