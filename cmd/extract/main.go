package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var classRegex = regexp.MustCompile(`(?s)\/\/ Namespace: ([^\n]+?)\n\[Serializable\]\npublic\s+class\s+([A-Za-z_][A-Za-z0-9_]*)\s*:\s*TBase,\s*TAbstractBase\s*\{.*?Properties(.*?)\n\n.+?\n\}`)
var csGoTypeMapping = map[string]string{
	"String":       "string",
	"Int32":        "int32",
	"Int16":        "int16",
	"Int64":        "int64",
	"Boolean":      "bool",
	"Double":       "float64",
	"Float":        "float32",
	"List`1":       "[]any",
	"Dictionary`2": "map[any]any",
}

var csPythonTypeMapping = map[string]string{
	"String":       "str",
	"Int32":        "int",
	"Int16":        "int",
	"Int64":        "int",
	"Boolean":      "bool",
	"Double":       "float",
	"Float":        "float",
	"List`1":       "list[Any] # TODO",
	"Dictionary`2": "dict[Any, Any] # TODO",
}

const BAD_ARGS_CODE = 1
const BAD_INPUT_FILE_CODE = 2
const BAD_CLASS_MATCHING_CODE = 3

func main() {
	var debugFlag bool
	var targetFlag string
	var dumpFilePath string
	var destDir string
	flag.BoolVar(&debugFlag, "debug", false, "dump thrift classes to json file")
	flag.StringVar(&targetFlag, "target", "go", "target programming language")
	flag.StringVar(&dumpFilePath, "dump", "./dump.cs", "path to the dump.cs file")
	flag.StringVar(&destDir, "dest", "./pkg/model", "path to the destination of the model files")
	flag.Parse()

	dumpFile, err := os.Open(dumpFilePath)
	if err != nil {
		os.Exit(BAD_INPUT_FILE_CODE)
	}

	dumpFileContent, _ := io.ReadAll(dumpFile)
	classes := []ThriftClass{}

	classMatches := extractClasses(dumpFileContent)
	for _, class := range classMatches {
		namespace := string(class[1])
		className := string(class[2])

		properties := extractProperties(class[3])
		classes = append(classes, ThriftClass{
			Name:       className,
			Namespace:  namespace,
			Properties: properties,
		})
	}

	if debugFlag {
		outfile, _ := os.Create("thrift_classes.json")
		outdata, _ := json.Marshal(classes)
		outfile.Write(outdata)
		fmt.Println("Dumped thrift classes to ./thrift_classes.json")
		defer outfile.Close()
	}

	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name < classes[j].Name
	})

	switch targetFlag {
	case "go":
		models := generateGoModels(classes)
		writeGoModels(models, destDir)
	case "python":
		generatePythonModels(classes)
	default:
		fmt.Println("Invalid target language:", targetFlag)
	}
}

func generateGoModels(classes []ThriftClass) map[string]*Model {
	modelMap := map[string]*Model{}

	for _, class := range classes {
		if _, ok := modelMap[class.Namespace]; !ok {
			modelMap[class.Namespace] = &Model{
				Namespace: class.Namespace,
				Classes:   make(map[string]ThriftClass),
			}
		}

		modelMap[class.Namespace].Classes[class.Name] = class
	}

	for _, model := range modelMap {
	Outer:
		for _, class := range model.Classes {
			for _, prop := range class.Properties {
				if _, builtinType := csGoTypeMapping[prop.Type]; !builtinType {
					if _, inNamespace := model.Classes[prop.Type]; !inNamespace {
						model.NeedsCommonImport = true
						break Outer
					}
				}
			}
		}
	}

	return modelMap
}

func writeGoModels(models map[string]*Model, destDir string) error {
	for _, model := range models {
		var modelStringBuilder strings.Builder
		packageName := fmt.Sprintf("%s_model", model.Namespace)
		packageName = strings.ReplaceAll(packageName, ".", "_")

		fmt.Fprintf(&modelStringBuilder, "package %s\n\n", packageName)
		if model.NeedsCommonImport {
			modelStringBuilder.WriteString("import \"github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model\"\n\n")
		}

		for _, class := range model.Classes {

			goStructName := strings.ToUpper(class.Name[0:1]) + class.Name[1:]
			var structLine strings.Builder

			fmt.Fprintf(&structLine, "type %s struct {\n", goStructName)
			for propIndex, prop := range class.Properties {
				_type := csGoTypeMapping[prop.Type]
				if _type == "" {
					_type = strings.ToUpper(prop.Type[0:1]) + prop.Type[1:]
					if _, namespaceHasType := model.Classes[prop.Type]; !namespaceHasType {
						_type = fmt.Sprintf("common_model.%s", _type)
					}
				}

				fmt.Fprintf(&structLine, "\t%s %s `thrift:\",%d,omitempty\"`", prop.Name, _type, propIndex+1)
				if strings.Contains(prop.Type, "`") {
					structLine.WriteString(" // TODO")
				}
				structLine.WriteString("\n")
			}

			structLine.WriteString("}\n\n")
			modelStringBuilder.WriteString(structLine.String())
		}

		destfolder := filepath.Join(destDir, packageName)
		os.Mkdir(destfolder, 0775)

		destfile := filepath.Join(destfolder, fmt.Sprintf("%s.go", model.Namespace))
		f, _ := os.Create(destfile)
		f.WriteString(modelStringBuilder.String())
	}

	fmt.Println("Created thrift data models")
	return nil
}

func generatePythonModels(classes []ThriftClass) error {
	pythonModelFile, err := os.Create("thrift_model.py")
	if err != nil {
		return err
	}
	defer pythonModelFile.Close()
	pythonModelFile.WriteString("from typing import Any\nfrom dataclasses import dataclass\nfrom __future__ import annotations\n\n")

	addedClasses := map[string]bool{}

	for _, class := range classes {
		pythonClassName := strings.ToUpper(class.Name[0:1]) + class.Name[1:]
		if _, ok := addedClasses[pythonClassName]; ok {
			continue
		}
		addedClasses[pythonClassName] = true

		var structLine strings.Builder

		fmt.Fprintf(&structLine, "@dataclass\nclass %s:\n", pythonClassName)
		for _, prop := range class.Properties {
			_type := csPythonTypeMapping[prop.Type]
			if _type == "" {
				_type = strings.ToUpper(prop.Type[0:1]) + prop.Type[1:]
			}
			fmt.Fprintf(&structLine, "\t%s: %s\n", prop.Name, _type)
		}

		if len(class.Properties) == 0 {
			fmt.Fprint(&structLine, "\tpass\n")
		}

		structLine.WriteString("\n\n")
		pythonModelFile.WriteString(structLine.String())
	}

	return nil
}

func extractClasses(data []byte) [][][]byte {
	matches := classRegex.FindAllSubmatch(data, -1)
	if len(matches) == 0 {
		os.Exit(BAD_CLASS_MATCHING_CODE)
	}

	return matches
}

func extractProperties(data []byte) []Property {
	properties := []Property{}
	props := strings.TrimSpace(string(data))
	if props == "" {
		return properties
	}

	for prop := range strings.SplitSeq(props, "\n") {
		parts := strings.Split(prop, " ")
		properties = append(properties, Property{
			Type: parts[1],
			Name: parts[2],
		})
	}

	return properties
}

type ThriftClass struct {
	Namespace  string
	Name       string
	Properties []Property
}

type Property struct {
	Name string
	Type string
}

type Model struct {
	Namespace         string
	NeedsCommonImport bool
	Classes           map[string]ThriftClass
}
