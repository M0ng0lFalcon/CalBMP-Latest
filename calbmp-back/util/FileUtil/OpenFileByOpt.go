package FileUtil

import (
	"log"
	"os"
)

func OpenFileAsRead(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		log.Println("[!] Open file as read error :", err, "path:", path)
		_ = f.Close()
	}
	//defer func(fi *os.File) {
	//	errClose := fi.Close()
	//	fmt.Println("[!] Cole file :", path)
	//	// using the function
	//	if errClose != nil {
	//		log.Println("[!] Close file error :", errClose)
	//	}
	//}(f)

	return f
}

func OpenFileAsWrite(path string) *os.File {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("[!] Open file as write error :", err)
		_ = f.Close()
	}
	//defer func(fi *os.File) {
	//	errClose := fi.Close()
	//	if errClose != nil {
	//		log.Println("[!] Close file error :", errClose)
	//	}
	//}(f)

	return f
}
