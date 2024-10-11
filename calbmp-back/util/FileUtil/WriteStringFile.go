package FileUtil

import (
	"fmt"
	"log"
	"os"
)

func WriteString2File(path, str string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		_, errWrite := f.WriteString(str)
		if errWrite != nil {
			log.Println("[WriteString2File] :", errWrite)
		}
	}
}
