package FileUtil

import (
	"log"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Mkdir(path string) {
	dir := path
	// mkdir
	exist, err := PathExists(dir)
	if err != nil {
		log.Println(err.Error())
	} else {
		if exist {
			//log.Println(dir + " [!] dir already existed!")
		} else {
			// dirname, auth
			errMkdir := os.Mkdir(dir, os.ModePerm)
			if errMkdir != nil {
				//log.Println(dir+" [!] make dir failed", errMkdir.Error())
			} else {
				//log.Println(dir + " [*] make dir success")
			}
		}
	}
}
