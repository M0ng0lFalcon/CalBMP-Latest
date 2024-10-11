package FileUtil

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(file string, result *[]string) error {
	fp, err := os.Open(file)
	if err != nil {
		log.Println("ReadFile:", err)
		return err
	}
	defer fp.Close()
	buf := bufio.NewScanner(fp)
	for buf.Scan() {
		line := buf.Text()
		if len(line) > 0 {
			*result = append(*result, line)
		}
	}
	return nil
}
