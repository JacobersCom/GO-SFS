package main

import "os"

func check(e error) {
	if e != nil {

		panic(e)
	}
}

func make_dir(dir_name string) {

	err := os.Mkdir(dir_name, 0755)
	check(err)
}
