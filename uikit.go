package main

import "fmt"

func genUIKitModule(moduleName string) {
	genUIKitRoutes(moduleName)
	genUIKitFactory(moduleName)
}

func genUIKitRoutes(moduleName string) {
	routes := fmt.Sprintf(
		"import Foundation\n" +
			"import RCKit\n" +
			"\n" +
			"protocol %sModuleRoutes: ModuleRoutes { \n"+
			"\n" +
			"} \n" +
			"\n", moduleName)

	writeFile(moduleName, "ModuleRoutes", routes)
}

func genUIKitFactory(moduleName string) {
	factory := fmt.Sprintf(
		"import Foundation\n" +
			"import RCKit\n" +
			"\n" +
			"protocol %sFactory {\n" +
			"	func make() -> %sModuleRoutes\n" +
			"}\n" +
			"\n" +
			"class Default%sFactory: BaseFactory, %sFactory {\n" +
			"\n" +
			"	func make() -> %sModuleRoutes { \n" +
			"\n" +
			"	}\n" +
			"}" +
			"", moduleName, moduleName, moduleName, moduleName, moduleName,
	)
	writeFile(moduleName, "Factory", factory)
}
