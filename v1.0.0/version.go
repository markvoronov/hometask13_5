package v1_0_0

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(sourcePath string, receiverPath string) error {
	fmt.Println("Модуль первой версии")

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

	str, err := json.Marshal(res)

	receiverFile, err := os.Create(receiverPath)
	if err != nil {
		log.Fatalln(err)
	}

	receiverFile.WriteString(string(str))

	return nil
}
