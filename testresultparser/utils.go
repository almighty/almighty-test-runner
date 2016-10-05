package testresultparser

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

// ReportParser returns byte object to the interface parse method
func readFile(filepath string) []byte {
	fp, errOpenFile := os.Open(filepath)
	checkErr(errOpenFile)
	defer func() {
		errCloseFile := fp.Close()
		checkErr(errCloseFile)
	}()
	b, errReadFile := ioutil.ReadAll(fp)
	checkErr(errReadFile)
	return b
}
