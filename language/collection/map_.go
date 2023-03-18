package collection

import "fmt"

func MapSolution() {
	mapValueIsList()
}

func mapValueIsList() {

	strMap := map[string][]string{
		"str1": {"eric", "mao"},
		"str2": {"tom", "ken"},
	}

	strList := strMap["str1"]
	strList = append(strList, "jane")
	strMap["str1"] = strList

	fmt.Printf("map: %+v", strMap)
}
