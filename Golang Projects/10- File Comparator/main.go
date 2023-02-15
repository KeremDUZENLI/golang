package main

import (
	"FileComparator/fileCheck"
	"FileComparator/fileCompare"
	"FileComparator/listConvert"
)

func main() {
	// PATH
	var myPath1, myPath2 string
	myPath1 = ""
	myPath2 = ""
	// tools.Input(&myPath1)
	// tools.Input(&myPath2)

	// DATABASE
	dataBase1 := [][]string{}
	dataBase2 := [][]string{}

	// BASIC LIST
	basicList1 := []string{}
	basicList2 := []string{}

	// ---------------------------------------------------------------

	// FILECHECK
	fileCheck.FileCheck(myPath1, &dataBase1)
	fileCheck.FileCheck(myPath2, &dataBase2)

	// LISTCONVERTER
	listConvert.ListConvert(dataBase1, &basicList1)
	listConvert.ListConvert(dataBase2, &basicList2)

	// FILECOMPARE
	fileCompare.FileCompare(basicList1, basicList2)
}
