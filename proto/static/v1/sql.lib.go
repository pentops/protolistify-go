// TODO: Generate this file using a plugin
package listify

import (
	"bytes"
	"fmt"
	"strings"
)

func (x *FilterClauses) ToSql() ([]string, []interface{}) {
	stmts, args := x.FilterStatementsWithPlaceholders()

	ordinal := 1
	for i, stmt := range stmts {
		if !strings.Contains(stmt, "?") {
			continue
		}

		buf := &bytes.Buffer{}
		for {
			p := strings.Index(stmt, "?")
			if p == -1 {
				break
			}

			if len(stmt[p:]) > 1 && stmt[p:p+2] == "??" { // escape ?? => ?
				buf.WriteString(stmt[:p])
				buf.WriteString("?")
				if len(stmt[p:]) == 1 {
					break
				}
				stmt = stmt[p+2:]
			} else {
				buf.WriteString(stmt[:p])
				buf.WriteString(fmt.Sprintf("$%d", ordinal))
				stmt = stmt[p+1:]
				ordinal++
			}
		}

		stmts[i] = buf.String()
	}

	return stmts, args
}

func (x *FilterClauses) ToSqrl() ([]string, []interface{}) {
	return x.FilterStatementsWithPlaceholders()
}

func (x *FilterClauses) FilterStatementsWithPlaceholders() ([]string, []interface{}) {
	stmts := []string{}
	args := []interface{}{}

	for _, f := range x.Clauses {
		stmt, arg := f.ToSql()
		stmts = append(stmts, stmt)
		args = append(args, arg...)
	}

	return stmts, args
}

func (x *FilterClause) ToSql() (string, []interface{}) {
	args := []interface{}{}

	for _, arg := range x.Arguments {
		switch k := arg.GetKind().(type) {
		case *FilterArgument_Double:
			args = append(args, k.Double)
		case *FilterArgument_Fixed32:
			args = append(args, k.Fixed32)
		case *FilterArgument_Fixed64:
			args = append(args, k.Fixed64)
		case *FilterArgument_Float:
			args = append(args, k.Float)
		case *FilterArgument_Int32:
			args = append(args, k.Int32)
		case *FilterArgument_Int64:
			args = append(args, k.Int64)
		case *FilterArgument_Sfixed32:
			args = append(args, k.Sfixed32)
		case *FilterArgument_Sfixed64:
			args = append(args, k.Sfixed64)
		case *FilterArgument_Sint32:
			args = append(args, k.Sint32)
		case *FilterArgument_Sint64:
			args = append(args, k.Sint64)
		case *FilterArgument_Uint32:
			args = append(args, k.Uint32)
		case *FilterArgument_Uint64:
			args = append(args, k.Uint64)
		case *FilterArgument_String_:
			args = append(args, k.String_)
		}
	}

	return x.Predicate, args
}
