package v2_1_0

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sort"
)

type Patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type patients struct {
	List []Patient `xml:"Patient"`
}

func Do(sourcePath string, receiverPath string) error {
	fmt.Println("Модуль четверой версии")

	sourcefile, err := os.Open(sourcePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer sourcefile.Close()

	dec := json.NewDecoder(sourcefile)

	res := make([]Patient, 0, 3)
	for dec.More() {

		var p Patient
		err := dec.Decode(&p)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, p)

	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Age < res[j].Age
	})

	d := patients{}
	d.List = res

	receiverFile, err := os.Create(receiverPath)
	if err != nil {
		log.Fatalln(err)
	}

	receiverFile.WriteString(xml.Header)
	enc := xml.NewEncoder(receiverFile)
	enc.Indent("", "  ")
	err = enc.Encode(d)
	if err != nil {
		log.Fatalln(err)
	}
	receiverFile.Close()

	return nil
}
