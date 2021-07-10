package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	directory := flag.String("dir", wd, "direcotory where to search else cirrent wd will be taken")
	filename := flag.String("file", "", "file name to be searched")

	flag.Parse()

	r := search(*directory, *filename)
	fmt.Println(r)

}

func search(directory string, filenmae string) string {

	var resarr []string
	var res string
	if directory != "" {

		fp, err := ioutil.ReadDir(directory)

		if err != nil {
			fmt.Printf("Error %s", err)
		}
		for _, f := range fp {

			resarr = append(resarr, f.Name())
		}

	}

	for _, file := range resarr {

		if file == filenmae {

			res = file

		}

	}

	return res
}
