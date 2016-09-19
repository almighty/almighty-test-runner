package testresultparser

import (
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filepath string) []byte {
	fp, err := os.Open(filepath)
	check(err)
	defer func() {
		check(fp.Close())
	}()
	b, err := ioutil.ReadAll(fp)
	check(err)
	return b
}
