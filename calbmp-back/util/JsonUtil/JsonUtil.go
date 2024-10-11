package JsonUtil

import (
	"encoding/json"
	"log"
	"os"
)

func WriteJson(path string, info interface{}) {
	filePtr, err := os.Create(path)
	if err != nil {
		log.Println("[WriteJson] Path:", path, "->err:", err)
		return
	}

	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		log.Println("[WriteJson] Path:", path, "->err:", err)
	}
}

func ReadJson(path string, info interface{}) {
	filePtr, err := os.Open(path)
	if err != nil {
		log.Println("[ReadJson] Path:", path, "->err:", err)
		return
	}

	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		log.Println("[ReadJson] Path:", path, "->err:", err)
	}

}

func ReadJson2Map(path string, mp *map[string]interface{}) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	} else {
		err = json.Unmarshal(b, mp)
		if err != nil {
			log.Println(err)
		}
	}

}
