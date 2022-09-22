package terraform

import "text/template"

func shortRegionName(i string) string {
	switch i {
	case "us-east-1":
		return "east"
	case "us-west-2":
		return "west"
	default:
		panic("invalid region name " + i)
	}
}
func quotedList(val []string) string {
	result := "["
	for i, j := range val {
		result += "\"" + j + "\""
		if i != len(val)-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}

var templateFunctionMap = template.FuncMap{
	"shortRegionName": shortRegionName,
	"quotedList":      quotedList,
}
