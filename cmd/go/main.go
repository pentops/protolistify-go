package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	version = "1.0"

	bytesPackage   = protogen.GoImportPath("bytes")
	errorPackage   = protogen.GoImportPath("errors")
	fmtPackage     = protogen.GoImportPath("fmt")
	regexpPackage  = protogen.GoImportPath("regexp")
	strconvPackage = protogen.GoImportPath("strconv")
	stringsPackage = protogen.GoImportPath("strings")
	timePackage    = protogen.GoImportPath("time")

	listifyPackage = protogen.GoImportPath("github.com/pentops/protoc-gen-listify/listify/v1")
)

func main() {
	showVersion := flag.Bool("version", false, "Print the version information and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-listify %v\n", version)
		os.Exit(0)
	}

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			_, err := genListifyFile(gen, f)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
