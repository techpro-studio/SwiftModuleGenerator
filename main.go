package main

import (
	"log"
	"os"
)



func main() {
	if len(os.Args) != 3 {
		log.Fatal("Need module name and type of module")
	}
	moduleType := os.Args[1]
	moduleName := os.Args[2]

	const SwiftUIModule  = "swiftui"
	const UIKitModule = "uikit"

	if moduleName == "" || moduleType == "" {
		log.Fatal("Module name or type of should not be empty")
	}

	if moduleType != SwiftUIModule && moduleType != UIKitModule {
		log.Fatalf("Module type should be swiftui or uikit")
	}

	log.Printf("Start building module %s", moduleName)

	err := os.Mkdir(moduleName, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create a directory %s", err.Error())
	}

	if moduleType == SwiftUIModule {
		genSwiftUIModule(moduleName)
	} else {
		genUIKitModule(moduleName)
	}

	log.Printf("Finished building module %s", moduleName)

}
