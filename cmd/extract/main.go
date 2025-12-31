package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

var classRegex = regexp.MustCompile(`(?s)public\s+class\s+([A-Za-z_][A-Za-z0-9_]*)\s*:\s*TBase,\s*TAbstractBase\s*\{.*?Properties(.*?)\n\n.+?\n\}`)
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

const BAD_ARGS_CODE = 1
const BAD_INPUT_FILE_CODE = 2
const BAD_CLASS_MATCHING_CODE = 3

func main() {
	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "dump thrift classes to json file")
	flag.Parse()

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		os.Exit(BAD_ARGS_CODE)
	}

	inputFilepath := argsWithoutProg[0]
	dumpFile, err := os.Open(inputFilepath)
	if err != nil {
		os.Exit(BAD_INPUT_FILE_CODE)
	}

	dumpFileContent, _ := io.ReadAll(dumpFile)
	classes := []ThriftClass{}

	classMatches := extractClasses(dumpFileContent)
	for _, class := range classMatches {
		className := string(class[1])

		properties := extractProperties(class[2])
		classes = append(classes, ThriftClass{
			Name:       className,
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

	generateGoModels(classes)
}

func generateGoModels(classes []ThriftClass) error {
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name < classes[j].Name
	})

	goModelFile, err := os.Create("thrift_model.go")
	if err != nil {
		return err
	}
	defer goModelFile.Close()

	goModelFile.WriteString("package model\n")

	addedClasses := map[string]bool{}

	for _, class := range classes {
		goStructName := strings.ToUpper(class.Name[0:1]) + class.Name[1:]
		if _, ok := addedClasses[goStructName]; ok {
			continue
		}

		addedClasses[goStructName] = true

		var structLine strings.Builder

		fmt.Fprintf(&structLine, "type %s struct {\n", goStructName)
		for _, prop := range class.Properties {
			_type := csGoTypeMapping[prop.Type]
			if _type == "" {
				_type = strings.ToUpper(prop.Type[0:1]) + prop.Type[1:]
			}
			fmt.Fprintf(&structLine, "\t%s %s\n", prop.Name, _type)
		}

		structLine.WriteString("}\n")
		goModelFile.WriteString(structLine.String())
	}

	fmt.Println("Created thrift data models at ./thrift_model.go")
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
	Name       string
	Properties []Property
}

type Property struct {
	Name string
	Type string
}
