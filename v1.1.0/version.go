package v1_1_0

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(sourcePath string, receiverPath string) error {
	//fmt.Println("Модуль версии 1.1.0")

	sourcefile, err := os.Open(sourcePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer sourcefile.Close()

	dec := json.NewDecoder(sourcefile)
	res := make([]Person, 0, 3)
	for dec.More() {

		var p Person
		err := dec.Decode(&p)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, p)

	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Age < res[j].Age
	})

	str, err := json.Marshal(res)

	receiverFile, err := os.Create(receiverPath)
	if err != nil {
		log.Fatalln(err)
	}

	receiverFile.WriteString(string(str))

	return nil
}
