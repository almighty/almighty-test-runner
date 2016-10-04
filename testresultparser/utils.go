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
	fp, err := os.Open(filepath)
	checkErr(err)
	defer func() {
		err := fp.Close()
		checkErr(err)
	}()
	b, err := ioutil.ReadAll(fp)
	checkErr(err)
	return b
}
