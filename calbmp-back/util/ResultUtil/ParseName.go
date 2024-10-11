package ResultUtil

import (
	"fmt"
	"regexp"
)

// ParseName : parse names from request body and get index in column name
func ParseName(names, ztsColumns []string) []int {
	// return index of name in column list
	res := make([]int, 0)

	// traverse names
	for _, name := range names {
		// 1. check concentration mode
		if CheckConcentration(name) {
			// -1 represent concentration mode
			res = append(res, -1)
			break
		}

		// 2. else continue normal mode
		// match name with columns
		for i, column := range ztsColumns {
			// match column name
			flag := FuzzyMatchColumnName(column, name)
			if flag == true {
				// add index of column to index result
				res = append(res, i)
				// have cumulative data
				if HaveTCUM(name) {
					res = append(res, i+1)
				}
				break
			}

		}
	}
	return res
}

func CheckConcentration(name string) bool {
	if name == "concentration" {
		return true
	}
	return false
}

func HaveTCUM(name string) bool {
	var haveTCUMNameList = []string{"RFLX", "EFLX", "VFLX"}
	for _, v := range haveTCUMNameList {
		if v == name {
			return true
		}
	}
	return false
}

func FuzzyMatchColumnName(ColumnName, param string) bool {
	pattern := fmt.Sprintf(`^%s`, param)
	flag, _ := regexp.MatchString(pattern, ColumnName)
	return flag
}
