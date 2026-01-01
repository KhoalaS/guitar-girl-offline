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

var classRegex = regexp.MustCompile(`(?s)\/\/ Namespace: (.+?)\n\[Serializable\]\npublic\s+class\s+([A-Za-z_][A-Za-z0-9_]*)\s*:\s*TBase,\s*TAbstractBase\s*\{.*?Properties(.*?)\n\n.+?\n\}`)
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
		generateGoModels(classes, destDir)
	case "python":
		generatePythonModels(classes)
	default:
		fmt.Println("Invalid target language:", targetFlag)
	}
}

func generateGoModels(classes []ThriftClass, destDir string) error {

	builderMap := map[string]*strings.Builder{}

	goModelFile, err := os.Create("thrift_model.go")
	if err != nil {
		return err
	}
	defer goModelFile.Close()

	goModelFile.WriteString("package model\n\n")

	for _, class := range classes {
		if _, ok := builderMap[class.Namespace]; !ok {
			var packageBuilder strings.Builder
			fmt.Fprintf(&packageBuilder, "package %s_model\n\n", class.Namespace)
			builderMap[class.Namespace] = &packageBuilder
		}

		goStructName := strings.ToUpper(class.Name[0:1]) + class.Name[1:]
		var structLine strings.Builder

		fmt.Fprintf(&structLine, "type %s struct {\n", goStructName)
		for propIndex, prop := range class.Properties {
			_type := csGoTypeMapping[prop.Type]
			if _type == "" {
				_type = strings.ToUpper(prop.Type[0:1]) + prop.Type[1:]
			}
			fmt.Fprintf(&structLine, "\t%s %s `thrift:\",%d,omitempty\"`", prop.Name, _type, propIndex+1)
			if strings.Contains(prop.Type, "`") {
				structLine.WriteString(" // TODO")
			}
			structLine.WriteString("\n")
		}

		structLine.WriteString("}\n\n")
		packageBuilder := builderMap[class.Namespace]
		packageBuilder.WriteString(structLine.String())
	}

	for namespace, builder := range builderMap {
		packageName := fmt.Sprintf("%s_model", namespace)
		destfolder := filepath.Join(destDir, packageName)
		os.Mkdir(destfolder, 0775)

		destfile := filepath.Join(destfolder, fmt.Sprintf("%s.go", namespace))
		f, _ := os.Create(destfile)
		f.WriteString(builder.String())
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
