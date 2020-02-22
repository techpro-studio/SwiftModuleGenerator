package main

import "fmt"

func swiftUIGen(moduleName string) {
	mainImport := "import SwiftUI \n\n"

	factoryData := mainImport + fmt.Sprintf(
		"protocol %sViewFactory { \n"+
			"    func make() -> AnyView \n"+
			"} \n\n"+
			"struct Default%sViewFactory: %sViewFactory { \n"+
			"    func make() -> AnyView { \n"+
			"        let viewModel = %sViewModel()\n"+
			"        return AnyView(%sView(viewModel: viewModel))\n"+
			"     }\n"+
			"}\n"+
			"", moduleName, moduleName, moduleName, moduleName, moduleName)

	writeFile(moduleName, "ViewFactory", factoryData)

	viewData := mainImport + fmt.Sprintf(
		"struct %sView: View { \n\n"+
			"    @ObservedObject var viewModel: %sViewModel\n\n"+
			"    var body: some View {\n"+
			"        Text(\"%s\").frame(maxWidth: .infinity, maxHeight: .infinity)\n"+
			"    }\n"+
			"}\n\n"+
			"struct %sView_Previews: PreviewProvider {\n"+
			"    static var previews: some View {\n\n"+
			"	 }\n"+
			"}"+
			"", moduleName, moduleName, moduleName, moduleName)

	writeFile(moduleName, "View", viewData)

	viewModelData := mainImport + fmt.Sprintf(
		"class %sViewModel: ObservableObject { \n\n\n"+
			"}\n"+
			"", moduleName)

	writeFile(moduleName, "ViewModel", viewModelData)
}
