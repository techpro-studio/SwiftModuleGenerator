package main

import "fmt"

func genUIKitModule(moduleName string) {
	genUIKitRoutes(moduleName)
	genUIKitFactory(moduleName)
	genUIKitController(moduleName)
	genUIKitPresenter(moduleName)
}


func genUIKitRoutes(moduleName string) {
	routes := fmt.Sprintf(
		"import Foundation\n" +
			"import RCKit\n" +
			"\n" +
			"protocol %sModuleRoutes: ModuleRoutes { \n"+
			"\n" +
			"} \n" +
			"", moduleName)

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
			"		fatalError(\"implement me\")\n" +
			"	}\n" +
			"}" +
			"", moduleName, moduleName, moduleName, moduleName, moduleName,
	)
	writeFile(moduleName, "Factory", factory)
}

func genUIKitController(moduleName string) {
	viewController := fmt.Sprintf(
		"import Foundation\n" +
			"import UIKit\n" +
			"\n" +
			"class %sViewController: UIViewController, %sModuleRoutes {\n" +
			"	var presenter: %sPresenter!\n" +
			"}" +
			"", moduleName, moduleName, moduleName,
	)
	writeFile(moduleName, "ViewController", viewController)
}

func genUIKitPresenter(moduleName string) {
	presenter := fmt.Sprintf(
		"import Foundation \n" +
			"import RCKit\n" +
			"import RxSwift\n" +
			"\n" +
			"protocol %sPresenter {\n" +
			"\n" +
			"}\n" +
			"\n" +
			"class Default%sPresenter: %sPresenter {\n" +
			"\n" +
			"}\n" +
			"", moduleName, moduleName, moduleName,
	)
	writeFile(moduleName, "Presenter", presenter)
}
