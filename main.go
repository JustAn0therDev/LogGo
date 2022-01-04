package main

import (
	"fmt"
	"os"

	"github.com/JustAn0therDev/LogGo/funcs"
)

func main() {
	errorMessage := "testing error"
	warningMessage := "testing warning"
	infoMessage := "testing info"
	funcs.LogError(&errorMessage)
	funcs.LogWarning(&warningMessage)
	funcs.LogInfo(&infoMessage)

	var response string

	fmt.Print("Do you want to read the log file into the terminal buffer? (y/n) ")

	fmt.Scanln(&response)

	if response != "n" {
		bytes, err := os.ReadFile("2022-01-03.logGo")

		if err != nil {
			panic(err)
		}
	
		fmt.Println(convertByteArrayToString(bytes))
	}
}

func convertByteArrayToString(buf []byte) string {
	var toReturn []rune

	for _, byt := range buf {
		toReturn = append(toReturn, rune(byt))
	}

	return string(toReturn)
}