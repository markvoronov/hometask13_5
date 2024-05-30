package main

import (
	"fmt"
	"log"
	"os"
	v1_0_0 "patietnfiles/v1.0.0"
	v1_1_0 "patietnfiles/v1.1.0"
	v2_0_0 "patietnfiles/v2.0.0"
	v2_1_0 "patietnfiles/v2.1.0"
)

// Определим константы для хранения версий
const (
	Version1_0_0 = "1.0.0"
	Version1_1_0 = "1.1.0"
	Version2_0_0 = "2.0.0"
	Version2_1_0 = "2.1.0"
)

func main() {

	// Если решать задачу строго как указано, то здесь нужно написать что-то типа этого
	err := v1_0_0.Do("patients", "result1")
	if err != nil {
		fmt.Println(err)
	}

	// Однако по хорошему, как было обозначено, нужно иметь возможность подключать требуемую версию.
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("Нужно передать версию обработки файла")
		return
	}

	version := arg[1]

	// Определим карту, которая ассоциирует переданную версию с выполняемой функцией
	versionFuncs := map[string]func(string, string) error{
		Version1_0_0: v1_0_0.Do,
		Version1_1_0: v1_1_0.Do,
		Version2_0_0: v2_0_0.Do,
		Version2_1_0: v2_1_0.Do,
	}

	// Проверим, что для переданной версии определена функция
	if doFunc, exists := versionFuncs[version]; exists {

		err := doFunc("patients", "result")
		if err != nil {
			log.Fatalf("Ошибка при работе с версией %s: %v", version, err)
		}
	} else {
		fmt.Printf("Неизвестная версия обработки: %s\n", version)
	}

}
