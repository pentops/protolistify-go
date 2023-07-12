package main

import (
	"fmt"

	"github.com/pentops/protoc-gen-listify/listify/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func genFilterTypeValidators(gen *protogen.Plugin, g *protogen.GeneratedFile) {
	g.P("func ListifyFilter_ValidDouble(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseFloat"), "(s, 64)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidFloat(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseFloat"), "(s, 32)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidInt32(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseInt"), "(s, 10, 32)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidInt64(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseInt"), "(s, 10, 64)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidUint32(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseUint"), "(s, 10, 32)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidUint64(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseUint"), "(s, 10, 64)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidBool(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("s := f.GetValue()")
	g.P("_, err := ", strconvPackage.Ident("ParseBool"), "(s)")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, s)`)
	g.P("}")
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidUniqueString(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: range", ErrFilterValidationInvalidValue)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidUUID(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: range", ErrFilterValidationInvalidValue)`)
	g.P("}")
	g.P()
	g.P("if !uuidMatch.MatchString(f.GetValue()) {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, f.GetValue())`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidDate(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("if !dateMatch.MatchString(f.GetValue()) {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, f.GetValue())`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func ListifyFilter_ValidTimestamp(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("_, err := ", timePackage.Ident("Parse"), "(time.RFC3339, f.GetValue())")
	g.P("if err != nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, f.GetValue())`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()
}

func genFilterableContent(gen *protogen.Plugin, method *protogen.Method, g *protogen.GeneratedFile) error {
	// Generate the ValidateFilters method for the children of the output message
	filterables, err := genFilterables(gen, method.Output, g, true)
	if err != nil {
		return err
	}

	// Generate the ValidateFilters method for the input message
	g.P("func (r *", method.Input.GoIdent.GoName, ") ValidateFilters() error{")
	g.P("validFilters := map[string]func(*", listifyPackage.Ident("Filter"), ")error{")
	for filterable, validFunc := range filterables {
		g.P(`"`, filterable, `": `, validFunc, ",")
	}
	g.P("}")
	g.P()
	g.P("failedFilters := []string{}")
	g.P()
	g.P("for _, filter := range r.Filters {")
	g.P("validator, ok := validFilters[filter.Field]")
	g.P("if !ok {")
	g.P("failedFilters = append(failedFilters, filter.Field)")
	g.P("continue")
	g.P("}")
	g.P()
	g.P("if err := validator(filter); err != nil {")
	g.P("failedFilters = append(failedFilters, filter.Field)")
	g.P("}")
	g.P("}")
	g.P()
	g.P("if len(failedFilters) > 0 {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %q", ErrFilterValidationInvalidField, failedFilters)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()

	// Generate the FiltersToSql method for the input message
	g.P("func (r *", method.Input.GoIdent.GoName, ") FiltersToSql() (string, []interface{}, error){")
	g.P(`return "", nil, nil`)
	g.P("}")
	g.P()

	return nil
}

func genFilterables(gen *protogen.Plugin, msg *protogen.Message, g *protogen.GeneratedFile, output bool) (map[string]string, error) {
	filterables := map[string]string{}

	for _, field := range msg.Fields {
		fieldAnnotations, ok := proto.GetExtension(field.Desc.Options(), listify.E_Rules).(*listify.FieldRulesOptions)
		if ok && fieldAnnotations != nil {
			switch field.Desc.Kind().String() {
			case "double":
				if fieldAnnotations.GetDouble() != nil && fieldAnnotations.GetDouble().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidDouble"
				}
			case "float":
				if fieldAnnotations.GetFloat() != nil && fieldAnnotations.GetFloat().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidFloat"
				}
			case "int32", "sint32", "sfixed32":
				if fieldAnnotations.GetInt32() != nil && fieldAnnotations.GetInt32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidInt32"
				}
			case "int64", "sint64", "sfixed64":
				if fieldAnnotations.GetInt64() != nil && fieldAnnotations.GetInt64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidInt64"
				}
			case "uint32", "fixed32":
				if fieldAnnotations.GetUint32() != nil && fieldAnnotations.GetUint32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidUint32"
				}
			case "uint64", "fixed64":
				if fieldAnnotations.GetUint64() != nil && fieldAnnotations.GetUint64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidUint64"
				}
			case "bool":
				if fieldAnnotations.GetBool() != nil && fieldAnnotations.GetBool().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidBool"
				}
			case "enum":
				if fieldAnnotations.GetEnum() != nil && fieldAnnotations.GetEnum().Filterable {
					valids := map[string]struct{}{}
					for _, v := range field.Enum.Values {
						valids[string(v.Desc.Name())] = struct{}{}
					}

					g.P("func ", msg.GoIdent.GoName, "_Valid", field.GoName, "(f *", listifyPackage.Ident("Filter"), ") error{")
					g.P("if f.GetRange() != nil {")
					g.P("return ", fmtPackage.Ident("Errorf"), `("%w: range", ErrFilterValidationInvalidValue)`)
					g.P("}")
					g.P()
					g.P("switch f.GetValue() {")
					for v := range valids {
						g.P("case \"", v, "\":")
						g.P("return nil")
					}
					g.P("}")
					g.P()
					g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, f.GetValue())`)
					g.P("}")
					g.P()

					filterables[field.Desc.TextName()] = fmt.Sprintf("%s_Valid%s", msg.GoIdent.GoName, field.GoName)
				}
			case "string":
				if fieldAnnotations.GetString_() != nil {
					if fieldAnnotations.GetString_().GetForeignKey() != nil {
						if fieldAnnotations.GetString_().GetForeignKey().GetUniqueString() != nil &&
							fieldAnnotations.GetString_().GetForeignKey().GetUniqueString().Filterable {
							filterables[field.Desc.TextName()] = "ListifyFilter_ValidUniqueString"
						}

						if fieldAnnotations.GetString_().GetForeignKey().GetUuid() != nil &&
							fieldAnnotations.GetString_().GetForeignKey().GetUuid().Filterable {
							filterables[field.Desc.TextName()] = "ListifyFilter_ValidUUID"
						}
					}

					if fieldAnnotations.GetString_().GetDate() != nil && fieldAnnotations.GetString_().GetDate().Filterable {
						filterables[field.Desc.TextName()] = "ListifyFilter_ValidDate"
					}
				}
			case "message":
				if field.Desc.Message().FullName() == "google.protobuf.Timestamp" {
					if fieldAnnotations.GetTimestamp() != nil && fieldAnnotations.GetTimestamp().Filterable {
						filterables[field.Desc.TextName()] = "ListifyFilter_ValidTimestamp"
					}
				}
			}
		}

		if field.Message != nil {
			nested, err := genFilterables(gen, field.Message, g, false)
			if err != nil {
				return filterables, err
			}

			for k, v := range nested {
				filterables[k] = v
			}
		}
	}

	for _, oneof := range msg.Oneofs {
		oneofAnnotations, ok := proto.GetExtension(oneof.Desc.Options(), listify.E_OneofRules).(*listify.OneofRulesOptions)
		if ok && oneofAnnotations != nil {
			if oneofAnnotations.Filterable {
				valids := map[string]struct{}{}
				for _, v := range oneof.Fields {
					valids[string(v.Desc.Message().Name())] = struct{}{}
				}

				g.P("func ", msg.GoIdent.GoName, "_Valid", oneof.GoName, "(f *", listifyPackage.Ident("Filter"), ") error{")
				g.P("if f.GetRange() != nil {")
				g.P("return ", fmtPackage.Ident("Errorf"), `("%w: range", ErrFilterValidationInvalidValue)`)
				g.P("}")
				g.P()
				g.P("switch f.GetValue() {")
				for v := range valids {
					g.P("case \"", v, "\":")
					g.P("return nil")
				}
				g.P("}")
				g.P()
				g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %s", ErrFilterValidationInvalidValue, f.GetValue())`)
				g.P("}")
				g.P()

				filterables[string(oneof.Desc.Name())] = fmt.Sprintf("%s_Valid%s", msg.GoIdent.GoName, oneof.GoName)
			}
		}
	}

	if len(filterables) > 0 && !output {
		genValidateFiltersMethod(gen, msg, g, filterables)
	}

	return filterables, nil
}

func genValidateFiltersMethod(gen *protogen.Plugin, msg *protogen.Message, g *protogen.GeneratedFile, filterables map[string]string) {
	g.P("func (r *", msg.GoIdent.GoName, ") ValidateFilters(filters []*", listifyPackage.Ident("Filter"), ") error{")
	g.P("validFilters := map[string]func(*", listifyPackage.Ident("Filter"), ")error{")
	for filterable, validFunc := range filterables {
		g.P(`"`, filterable, `": `, validFunc, ",")
	}
	g.P("}")
	g.P()
	g.P("failedFilters := []string{}")
	g.P()
	g.P("for _, filter := range filters {")
	g.P("validator, ok := validFilters[filter.Field]")
	g.P("if !ok {")
	g.P("failedFilters = append(failedFilters, filter.Field)")
	g.P("continue")
	g.P("}")
	g.P()
	g.P("if err := validator(filter); err != nil {")
	g.P("failedFilters = append(failedFilters, filter.Field)")
	g.P("}")
	g.P("}")
	g.P()
	g.P("if len(failedFilters) > 0 {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: %q", ErrFilterValidationInvalidField, failedFilters)`)
	g.P("}")
	g.P()
	g.P("return nil")
	g.P("}")
	g.P()
}

func genFiltersToSqlMethod(gen *protogen.Plugin, msg *protogen.Message, g *protogen.GeneratedFile) {
}
