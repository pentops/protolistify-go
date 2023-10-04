package main

import "google.golang.org/protobuf/compiler/protogen"

func genLibFile(gen *protogen.Plugin, file *protogen.File) (*protogen.GeneratedFile, error) {
	filename := file.GeneratedFilenamePrefix + ".lib.go"

	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	genSqlHelpers(gen, file, g)

	return nil, nil
}

func genSqlHelpers(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
}
