package main

import (
	"github.com/pentops/protoc-gen-listify/listify/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func genSortableContent(gen *protogen.Plugin, method *protogen.Method, g *protogen.GeneratedFile) error {
	// Generate the ValidateSorts method for the children of the output message
	sortables, err := genSortables(gen, method.Output, g, true)
	if err != nil {
		return err
	}

	// Generate the ValidateSorts method for the input message
	g.P("func (r *", method.Input.GoIdent.GoName, ") ValidateSorts() error{")
	g.P("validSorts := map[string]interface{}{")
	for sortable := range sortables {
		g.P(`"`, sortable, `": nil,`)
	}
	g.P("}")
	g.P()
	g.P("failedSorts := []string{}")
	g.P("directions:= map[string]interface{}{}")
	g.P()
	g.P("for _, sort := range r.Sorts {")
	g.P("_, ok := validSorts[sort.Field]")
	g.P("if !ok {")
	g.P("failedSorts = append(failedSorts, sort.Field)")
	g.P("continue")
	g.P("}")
	g.P()
	g.P("if sort.Descending {")
	g.P(`directions["desc"] = nil`)
	g.P("} else {")
	g.P(`directions["asc"] = nil`)
	g.P("}")
	g.P("}")
	g.P()
	g.P("if len(failedSorts) > 0 {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %q", ErrSortValidationInvalidField, failedSorts)`)
	g.P("}")
	g.P()
	g.P("if len(directions) > 1 {")
	g.P("return ErrSortValidationInvalidDirection")
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	genSortSqlMethods(gen, g, method.Input)

	return nil
}

func genSortSqlMethods(gen *protogen.Plugin, g *protogen.GeneratedFile, msg *protogen.Message) {
	g.P("func (r *", msg.GoIdent.GoName, ") GetSortClauses() (*", listifyPackage.Ident("SqlClauses"), "){")
	g.P("resp := &", listifyPackage.Ident("SqlClauses"), "{}")
	g.P()
	g.P("return resp")
	g.P("}")
	g.P()
}

func genSortables(gen *protogen.Plugin, msg *protogen.Message, g *protogen.GeneratedFile, output bool) (map[string]interface{}, error) {
	sortables := map[string]interface{}{}

	for _, field := range msg.Fields {
		fieldAnnotations, ok := proto.GetExtension(field.Desc.Options(), listify.E_Rules).(*listify.FieldRulesOptions)
		if ok && fieldAnnotations != nil {
			switch field.Desc.Kind() {
			case protoreflect.DoubleKind:
				if fieldAnnotations.GetDouble() != nil && fieldAnnotations.GetDouble().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.FloatKind:
				if fieldAnnotations.GetFloat() != nil && fieldAnnotations.GetFloat().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Int32Kind:
				if fieldAnnotations.GetInt32() != nil && fieldAnnotations.GetInt32().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Sint32Kind:
				if fieldAnnotations.GetSint32() != nil && fieldAnnotations.GetSint32().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Sfixed32Kind:
				if fieldAnnotations.GetSfixed32() != nil && fieldAnnotations.GetSfixed32().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Int64Kind:
				if fieldAnnotations.GetInt64() != nil && fieldAnnotations.GetInt64().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Sint64Kind:
				if fieldAnnotations.GetSint64() != nil && fieldAnnotations.GetSint64().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Sfixed64Kind:
				if fieldAnnotations.GetSfixed64() != nil && fieldAnnotations.GetSfixed64().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Uint32Kind:
				if fieldAnnotations.GetUint32() != nil && fieldAnnotations.GetUint32().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Fixed32Kind:
				if fieldAnnotations.GetFixed32() != nil && fieldAnnotations.GetFixed32().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Uint64Kind:
				if fieldAnnotations.GetUint64() != nil && fieldAnnotations.GetUint64().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.Fixed64Kind:
				if fieldAnnotations.GetFixed64() != nil && fieldAnnotations.GetFixed64().Sorting.Sortable {
					sortables[field.Desc.TextName()] = nil
				}
			case protoreflect.MessageKind:
				if field.Desc.Message().FullName() == "google.protobuf.Timestamp" {
					if fieldAnnotations.GetTimestamp() != nil && fieldAnnotations.GetTimestamp().Sorting.Sortable {
						sortables[field.Desc.TextName()] = nil
					}
				}
			}
		}

		if field.Message != nil {
			nested, err := genSortables(gen, field.Message, g, false)
			if err != nil {
				return sortables, err
			}

			for k, v := range nested {
				sortables[k] = v
			}
		}
	}

	if len(sortables) > 0 && !output {
		genValidateSortsMethod(gen, msg, g, sortables)
	}

	return sortables, nil
}

func genValidateSortsMethod(gen *protogen.Plugin, msg *protogen.Message, g *protogen.GeneratedFile, sortables map[string]interface{}) {
	g.P("func (r *", msg.GoIdent.GoName, ") ValidateSorts(sorts []*", listifyPackage.Ident("Sort"), ") error{")
	g.P("validSorts := map[string]interface{}{")
	for sortable := range sortables {
		g.P(`"`, sortable, `": nil,`)
	}
	g.P("}")
	g.P()
	g.P("failedSorts := []string{}")
	g.P("directions:= map[string]interface{}{}")
	g.P()
	g.P("for _, sort := range sorts {")
	g.P("_, ok := validSorts[sort.Field]")
	g.P("if !ok {")
	g.P("failedSorts = append(failedSorts, sort.Field)")
	g.P("continue")
	g.P("}")
	g.P()
	g.P("if sort.Descending {")
	g.P(`directions["desc"] = nil`)
	g.P("} else {")
	g.P(`directions["asc"] = nil`)
	g.P("}")
	g.P("}")
	g.P()
	g.P("if len(failedSorts) > 0 {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %q", ErrSortValidationInvalidField, failedSorts)`)
	g.P("}")
	g.P()
	g.P("if len(directions) > 1 {")
	g.P("return ErrSortValidationInvalidDirection")
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()
}
