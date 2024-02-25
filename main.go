package main

import (
	"encoding/json"
	"fmt"
	"golang.design/x/clipboard"
	"io"
	"jsontostruct/enums"
	"jsontostruct/util"
	"jsontostruct/views"
	"log"
	"os"
)

func main() {
	args := os.Args
	filePath := ""
	err := clipboard.Init()
	if err != nil {
		log.Default().Printf("error while initializing clipboard - %v", err)
	}
	if len(args) > 1 {
		filePath = args[1]
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Default().Printf("error opening the file %v", err)
		return
	}
	byteData, err := io.ReadAll(file)
	if err != nil {
		log.Default().Printf("error while reading the file %v", err)
		return
	}
	jsonString := string(byteData)
	mapData := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonString), &mapData)
	if err != nil {
		log.Default().Println("error while unmarshalling", err)
		return
	}
	generate(mapData)
}

func generate(jsonData map[string]interface{}) {
	var stack []views.ObjectInfo
	stack = append(stack, views.ObjectInfo{
		Title:   "Base",
		MapData: jsonData,
	})

	customTypeCount := 0
	fieldString := "%v %v `json:\"%v\"`\n"
	for i := 0; i < len(stack); i++ {
		base := "type " + util.Capitalize(stack[i].Title) + " struct {\n"
		for key, value := range stack[i].MapData {
			dataType, customType := util.GetType(value)
			if customType != nil {
				// push customType to stack till each field is resolved
				stack = append(stack, views.ObjectInfo{
					Title:   key,
					MapData: customType,
				})
				switch dataType {
				case enums.Object:
					customTypeCount++
					base += fmt.Sprintf(fieldString, util.ToCamelCase(key), util.Capitalize(stack[i+1].Title), key)
				case enums.SliceOfObjects:
					customTypeCount++
					base += fmt.Sprintf(fieldString, util.ToCamelCase(key), "[]"+util.Capitalize(stack[i+1].Title), key)
				default:
					base += fmt.Sprintf(fieldString, util.ToCamelCase(key), dataType, key)
				}
			} else {
				base += fmt.Sprintf(fieldString, util.ToCamelCase(key), dataType, key)
			}
		}
		base += "}\n"
		fmt.Println(base)
	}
}
