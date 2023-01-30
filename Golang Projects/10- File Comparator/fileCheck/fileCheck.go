package fileCheck

import (
	"io/ioutil"
)

func FileCheck(filePath string, dataBase *[][]string) {
	liste := []string{}
	files, _ := ioutil.ReadDir(filePath)

	for _, f := range files {
		if f.IsDir() == false {
			liste = append(liste, f.Name())
		} else {
			newPath := filePath + "/" + f.Name()
			FileCheck(newPath, dataBase)
		}
	}

	*dataBase = append(*dataBase, liste)
}
