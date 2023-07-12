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
	g.P("r := f.GetRange().GetDouble()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected double", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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
	g.P("r := f.GetRange().GetFloat()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected float", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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
	g.P("r := f.GetRange().GetInt32()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected int32", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	g.P("func ListifyFilter_ValidSint32(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("r := f.GetRange().GetSint32()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected sint32", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	g.P("func ListifyFilter_ValidSfixed32(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("r := f.GetRange().GetSfixed32()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected sfixed32", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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
	g.P("r := f.GetRange().GetInt64()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected int64", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	g.P("func ListifyFilter_ValidSint64(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("r := f.GetRange().GetSint64()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected sint64", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	g.P("func ListifyFilter_ValidSfixed64(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("r := f.GetRange().GetSfixed64()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected sfixed64", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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
	g.P("r := f.GetRange().GetUint32()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected uint32", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	g.P("func ListifyFilter_ValidFixed32(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("r := f.GetRange().GetFixed32()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected fixed32", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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
	g.P("r := f.GetRange().GetUint64()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected uint64", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	g.P("func ListifyFilter_ValidFixed64(f *", listifyPackage.Ident("Filter"), ") error{")
	g.P("if f.GetRange() != nil {")
	g.P("r := f.GetRange().GetFixed64()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected fixed64", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.Value > r.Max.Value {")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: range", ErrFilterValidationInvalidValue)`)
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
	g.P("r := f.GetRange().GetDate()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected date", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P(`if r.Min == nil && r.Max == nil {`)
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
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
	g.P("r := f.GetRange().GetTimestamp()")
	g.P("if r == nil {")
	g.P("return ", fmtPackage.Ident("Errorf"), `("%w: expected timestamp", ErrFilterValidationInvalidRangeType)`)
	g.P("}")
	g.P()
	g.P("if r.Min == nil && r.Max == nil {")
	g.P("return ErrFilterValidationInvalidRangeMinMaxMissing")
	g.P("}")
	g.P()
	g.P("if r.Min != nil && r.Max != nil && r.Min.AsTime().After(r.Max.AsTime()){")
	g.P("return ErrFilterValidationInvalidRangeMinMax")
	g.P("}")
	g.P()
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

	// Generate the FiltersToStatements method for the input message
	genFilterStatementMethods(gen, g, method.Input.GoIdent.GoName)

	return nil
}

func genFilterStatementMethods(gen *protogen.Plugin, g *protogen.GeneratedFile, receiver string) {
	g.P("func (r *", receiver, ") FilterStatements() ([]string, []interface{}, error){")
	g.P("return r.FilterStatementsWithOrdinal(1)")
	g.P("}")
	g.P()

	g.P("func (r *", receiver, ") FilterStatementsWithOrdinal(startingOrdinal int) ([]string, []interface{}, error){")
	g.P("var statements []string")
	g.P("var args []interface{}")
	g.P()
	g.P("ordinal := startingOrdinal")
	g.P("for _, filter := range r.Filters {")
	g.P("if filter.GetRange() != nil {")
	g.P("switch r := filter.GetRange().GetType().(type){")

	g.P("case *", listifyPackage.Ident("Range_Double"), ":")
	g.P("if r.Double.Min != nil && r.Double.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Double.Min.Value, r.Double.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Double.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Double.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Double.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Double.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Fixed32"), ":")
	g.P("if r.Fixed32.Min != nil && r.Fixed32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Fixed32.Min.Value, r.Fixed32.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Fixed32.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Fixed32.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Fixed32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Fixed32.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Fixed64"), ":")
	g.P("if r.Fixed64.Min != nil && r.Fixed64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Fixed64.Min.Value, r.Fixed64.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Fixed64.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Fixed64.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Fixed64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Fixed64.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Float"), ":")
	g.P("if r.Float.Min != nil && r.Float.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Float.Min.Value, r.Float.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Float.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Float.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Float.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Float.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Int32"), ":")
	g.P("if r.Int32.Min != nil && r.Int32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Int32.Min.Value, r.Int32.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Int32.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Int32.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Int32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Int32.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Int64"), ":")
	g.P("if r.Int64.Min != nil && r.Int64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Int64.Min.Value, r.Int64.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Int64.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Int64.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Int64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Int64.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Sfixed32"), ":")
	g.P("if r.Sfixed32.Min != nil && r.Sfixed32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Sfixed32.Min.Value, r.Sfixed32.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Sfixed32.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sfixed32.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Sfixed32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sfixed32.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Sfixed64"), ":")
	g.P("if r.Sfixed64.Min != nil && r.Sfixed64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Sfixed64.Min.Value, r.Sfixed64.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Sfixed64.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sfixed64.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Sfixed64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sfixed64.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Sint32"), ":")
	g.P("if r.Sint32.Min != nil && r.Sint32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Sint32.Min.Value, r.Sint32.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Sint32.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sint32.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Sint32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sint32.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Sint64"), ":")
	g.P("if r.Sint64.Min != nil && r.Sint64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Sint64.Min.Value, r.Sint64.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Sint64.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sint64.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Sint64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Sint64.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Uint32"), ":")
	g.P("if r.Uint32.Min != nil && r.Uint32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Uint32.Min.Value, r.Uint32.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Uint32.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Uint32.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Uint32.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Uint32.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Uint64"), ":")
	g.P("if r.Uint64.Min != nil && r.Uint64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P(`args = append(args, r.Uint64.Min.Value, r.Uint64.Max.Value)`)
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Uint64.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Uint64.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Uint64.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Uint64.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Date"), ":")
	g.P("if r.Date.Min != nil && r.Date.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Date.Min.Value, r.Date.Max.Value)")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Date.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Date.Min.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Date.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Date.Max.Value)")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("case *", listifyPackage.Ident("Range_Timestamp"), ":")
	g.P("if r.Timestamp.Min != nil && r.Timestamp.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d AND %s <= $%d`, filter.Field, ordinal, filter.Field, ordinal+1))")
	g.P("args = append(args, r.Timestamp.Min.AsTime(), r.Timestamp.Max.AsTime())")
	g.P("ordinal += 2")
	g.P("} else {")
	g.P("if r.Timestamp.Min != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s >= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Timestamp.Min.AsTime())")
	g.P("ordinal++")
	g.P("}")
	g.P("if r.Timestamp.Max != nil {")
	g.P("statements = append(statements, ", fmtPackage.Ident("Sprintf"), "(`%s <= $%d`, filter.Field, ordinal))")
	g.P("args = append(args, r.Timestamp.Max.AsTime())")
	g.P("ordinal++")
	g.P("}")
	g.P("}")

	g.P("}")
	g.P("} else {")
	g.P(`statements = append(statements, `, fmtPackage.Ident("Sprintf"), `("%s = $%d", filter.Field, ordinal))`)
	g.P("args = append(args, filter.GetValue())")
	g.P("ordinal++")
	g.P("}")
	g.P("}")
	g.P()
	g.P("return statements, args, nil")
	g.P("}")
	g.P()
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
			case "int32":
				if fieldAnnotations.GetInt32() != nil && fieldAnnotations.GetInt32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidInt32"
				}
			case "sint32":
				if fieldAnnotations.GetSint32() != nil && fieldAnnotations.GetSint32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidSint32"
				}
			case "sfixed32":
				if fieldAnnotations.GetSfixed32() != nil && fieldAnnotations.GetSfixed32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidSfixed32"
				}
			case "int64":
				if fieldAnnotations.GetInt64() != nil && fieldAnnotations.GetInt64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidInt64"
				}
			case "sint64":
				if fieldAnnotations.GetSint64() != nil && fieldAnnotations.GetSint64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidSint64"
				}
			case "sfixed64":
				if fieldAnnotations.GetSfixed64() != nil && fieldAnnotations.GetSfixed64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidSfixed64"
				}
			case "uint32":
				if fieldAnnotations.GetUint32() != nil && fieldAnnotations.GetUint32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidUint32"
				}
			case "fixed32":
				if fieldAnnotations.GetFixed32() != nil && fieldAnnotations.GetFixed32().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidFixed32"
				}
			case "uint64":
				if fieldAnnotations.GetUint64() != nil && fieldAnnotations.GetUint64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidUint64"
				}
			case "fixed64":
				if fieldAnnotations.GetFixed64() != nil && fieldAnnotations.GetFixed64().Filterable {
					filterables[field.Desc.TextName()] = "ListifyFilter_ValidFixed64"
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
