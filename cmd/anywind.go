package main

import (
	"encoding/json"
	"io/ioutil"
)

// Library structure
type lib struct {
	Name    string
	Version string
	// Arch string
}

func getLibs() (libs []lib) {

	fileBytes, err := ioutil.ReadFile("./libs.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &libs)

	if err != nil {
		panic(err)
	}

	return libs
}

func saveLibs(libs []lib) {

	libBytes, err := json.Marshal(libs)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./libs.json", libBytes, 0644)
	if err != nil {
		panic(err)
	}

}
