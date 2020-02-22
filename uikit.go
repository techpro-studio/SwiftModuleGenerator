package main

import "fmt"

func uiKitGen(moduleName string) {
	genRoutes(moduleName)
}

func genRoutes(moduleName string) {
	factoryData := fmt.Sprintf(
		"import Foundation\n" +
			"import RCKit\n\n" +
			"protocol %sModuleRoutes: ModuleRoutes { \n\n"+
			"" +
			"} \n\n" +
			"", moduleName)

	writeFile(moduleName, "ModuleRoutes", factoryData)
}

func genFactory(moduleName string) {

}
