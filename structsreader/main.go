package structsreader

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"

	"github.com/fatih/structtag"
)

func GetSchemas() map[string]map[string]string {
	f, err := parser.ParseDir(token.NewFileSet(), "models", nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
	}
	var structs = map[*ast.TypeSpec][]*ast.Field{}

	for _, file := range f["models"].Files {
		for _, node := range file.Decls {
			switch node.(type) {
			case *ast.GenDecl:
				genDecl := node.(*ast.GenDecl)
				for _, spec := range genDecl.Specs {
					switch spec.(type) {
					case *ast.TypeSpec:
						typeSpec := spec.(*ast.TypeSpec)
						switch typeSpec.Type.(type) {
						case *ast.StructType:
							structType := typeSpec.Type.(*ast.StructType)
							structs[typeSpec] = structType.Fields.List
						}
					}
				}
			}
		}
	}
	schemas := make(map[string]map[string]string, 0)
	for s := range structs {
		schemas[ToLowerCamel(s.Name.String())] = getSchema(s, structs[s])
	}
	return schemas
}

func getSchema(structure *ast.TypeSpec, fields []*ast.Field) map[string]string {
	m := map[string]string{}
	m["sortColumn"] = "name"
	m["label"] = "name"
	m["value"] = "id"
	for index, field := range fields {
		if field.Tag == nil {
			continue
		}
		tags := parseTags(field.Tag.Value)
		if index == 0 {
			m["tableName"] = getBunSelectTableName(tags)
			continue
		}
		m[getJSONName(tags)] = getColName(tags)
	}
	m["key"] = ToLowerCamel(structure.Name.Name)
	return m
}

func getJSONName(tags *structtag.Tags) string {
	jsonName, err := tags.Get("json")
	if err != nil {
		return ""
	}
	return jsonName.Name
}

func getColName(tags *structtag.Tags) string {
	bunTag, err := tags.Get("bun")
	if err == nil && bunTag.Name != "-" && bunTag.Name != "" && !strings.Contains(bunTag.Name, ":") {
		return bunTag.Name
	}
	return toSnake(getJSONName(tags))
}

func getBunSelectTableName(tags *structtag.Tags) string {
	bunTag, err := tags.Get("bun")
	if err != nil {
		return ""
	}
	tableName := bunTag.Name
	for _, opt := range bunTag.Options {
		parts := strings.Split(opt, ":")
		if len(parts) == 2 && parts[0] == "select" {
			tableName = parts[1]
		}
	}
	return tableName
}

func parseTags(tagString string) *structtag.Tags {
	tag, err := strconv.Unquote(tagString)
	if err != nil {
		panic(err)
	}
	tags, err := structtag.Parse(tag)
	if err != nil {
		panic(err)
	}
	return tags
}
