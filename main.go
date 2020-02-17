package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)



func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need module name")
	}
	moduleName := os.Args[1]
	log.Printf("Start building module %s", moduleName)

	err := os.Mkdir(moduleName, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create a directory %s", err.Error())
	}

	mainImport := "import SwiftUI \n\n"

	factoryData := mainImport + fmt.Sprintf(
		    "protocol %sViewFactory { \n" +
			"    func make() -> AnyView \n" +
			"} \n\n" +


			"class Default%sViewFactory: %sViewFactory { \n" +
			"    func make() -> AnyView { \n" +
			"        let viewModel = %sViewModel()\n" +
			"        return AnyView(%sView(viewModel: viewModel))\n" +
			"     }\n" +
			"}\n" +

			"", moduleName, moduleName, moduleName, moduleName, moduleName)

	writeFile(moduleName, "ViewFactory", factoryData)

	viewData := mainImport + fmt.Sprintf(
		"struct %sView: View { \n\n"+
			    "    @ObservedObject var viewModel: %sViewModel\n\n" +
			    "    var body: some View {\n" +
				"        Text(\"%s\").frame(maxWidth: .infinity, maxHeight: .infinity)\n" +
			    "    }\n" +
			    "}\n\n" +
				"struct %sView_Previews: PreviewProvider {\n" +
				"    static var previews: some View {\n\n" +
				"	 }\n" +
				"}" +
				"", moduleName, moduleName, moduleName, moduleName)

	writeFile(moduleName, "View", viewData)


	viewModelData := mainImport + fmt.Sprintf(
		"class %sViewModel: ObservableObject { \n\n\n"+
				"}\n" +
			"", moduleName)

	writeFile(moduleName, "ViewModel", viewModelData)

	log.Printf("Finished building module %s", moduleName)

}

func writeFile(moduleName string, fileRoot string,  data string) {
	filePath := fmt.Sprintf("%s/%s%s.swift", moduleName, moduleName, fileRoot)
	err := ioutil.WriteFile(filePath, []byte(data), os.ModePerm)

	if err != nil {
		log.Fatalf("Unable to write %s: %s",filePath, err.Error())
	}  else {
		log.Printf("Successfully built %s%s.swift", moduleName, fileRoot)
	}
}
