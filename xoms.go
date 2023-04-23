package xoms

//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

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"path/filepath"
)

func Hello() {
	fmt.Println("Hello, World!")
}

func listDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}
