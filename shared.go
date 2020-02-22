package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writeFile(moduleName string, fileRoot string,  data string) {
	filePath := fmt.Sprintf("%s/%s%s.swift", moduleName, moduleName, fileRoot)
	err := ioutil.WriteFile(filePath, []byte(data), os.ModePerm)

	if err != nil {
		log.Fatalf("Unable to write %s: %s",filePath, err.Error())
	}  else {
		log.Printf("Successfully built %s%s.swift", moduleName, fileRoot)
	}
}

