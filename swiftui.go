package main

import (
	"fmt"
)

func genSwiftUIModule(moduleName string) {
	mainImport := "import SwiftUI \n\n"

	genSwiftUIFactory(mainImport, moduleName)

	genSwiftUIView(mainImport, moduleName)

	genSwiftUIViewModel(mainImport, moduleName)
}

func genSwiftUIViewModel(mainImport string, moduleName string) {
	viewModelData := mainImport + fmt.Sprintf(
		"class %sViewModel: ObservableObject { \n\n\n"+
			"}\n"+
			"", moduleName)

	writeFile(moduleName, "ViewModel", viewModelData)
}

func genSwiftUIView(mainImport string, moduleName string) {
	viewData := mainImport + fmt.Sprintf(
		"struct %sView: View { \n\n"+
			"    @ObservedObject var viewModel: %sViewModel\n"+
			"\n"+
			"    var body: some View {\n"+
			"        Text(\"%s\").frame(maxWidth: .infinity, maxHeight: .infinity)\n"+
			"    }\n"+
			"}\n"+
			"\n"+
			"struct %sView_Previews: PreviewProvider {\n"+
			"    static var previews: some View {\n"+
			"	fatalError(\"implement me if need\")\n"+
			"	 }\n"+
			"}"+
			"", moduleName, moduleName, moduleName, moduleName)

	writeFile(moduleName, "View", viewData)
}

func genSwiftUIFactory(mainImport string, moduleName string) {
	factoryData := mainImport + fmt.Sprintf(
		"protocol %sViewFactory { \n"+
			"    func make() -> AnyView \n"+
			"} \n"+
			"\n"+
			"struct Default%sViewFactory: %sViewFactory { \n"+
			"    func make() -> AnyView { \n"+
			"        let viewModel = %sViewModel()\n"+
			"        return AnyView(%sView(viewModel: viewModel))\n"+
			"     }\n"+
			"}\n"+
			"", moduleName, moduleName, moduleName, moduleName, moduleName)

	writeFile(moduleName, "ViewFactory", factoryData)
}
