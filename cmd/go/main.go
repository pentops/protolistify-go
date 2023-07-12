package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pentops/protoc-gen-listify/listify/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	version = "1.0"

	errorPackage   = protogen.GoImportPath("errors")
	fmtPackage     = protogen.GoImportPath("fmt")
	regexpPackage  = protogen.GoImportPath("regexp")
	strconvPackage = protogen.GoImportPath("strconv")
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

			_, err := genFile(gen, f)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// genFile generates a _listify.pb.go file containing helper methods on services
// requests for generalizing sorting, filtering, searching, and pagination
func genFile(gen *protogen.Plugin, file *protogen.File) (*protogen.GeneratedFile, error) {
	if len(file.Services) == 0 {
		return nil, nil
	}

	filename := file.GeneratedFilenamePrefix + ".pb.listify.go"

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-listify. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("var (")
	g.P("ErrFilterValidationInvalidField = ", errorPackage.Ident("New"), `("invalid field name(s) for filter")`)
	g.P("ErrFilterValidationInvalidValue = ", errorPackage.Ident("New"), `("invalid value(s) for filter")`)
	g.P()
	g.P(`uuidMatch = `, regexpPackage.Ident("MustCompile"), `("(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")`)
	g.P(`dateMatch = `, regexpPackage.Ident("MustCompile"), `("^([0-9]{4}-?([0-9]{2})?-?([0-9]{2})?)$")`)
	g.P(")")
	g.P()

	generated, err := genFileContent(gen, file, g)
	if err != nil {
		return nil, err
	}

	if !generated {
		g.Skip()
		return nil, nil
	}

	return g, nil
}

func genFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) (bool, error) {
	genFilterTypeValidators(gen, g)

	for _, service := range file.Services {
		for _, method := range service.Methods {
			// Skip streaming methods
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}

			// Only generate for methods with the listify.api option set to true
			methodAnnotations, ok := proto.GetExtension(method.Desc.Options(), listify.E_Api).(*listify.MethodOptions)
			if ok {
				if methodAnnotations == nil || !methodAnnotations.Enable {
					continue
				}
			}

			err := genFilterableContent(gen, method, g)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
