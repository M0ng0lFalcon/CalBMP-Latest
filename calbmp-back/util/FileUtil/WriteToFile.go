package FileUtil

import (
	"fmt"
	"log"
	"os"
)

func WriteToFile(path string, content []string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// write lines to file
		for _, v := range content {
			_, errWrite := f.WriteString(v)
			if errWrite != nil {
				log.Fatal(errWrite)
			}
		}
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return err
}
