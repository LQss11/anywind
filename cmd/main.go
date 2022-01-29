package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//'anywind get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `anywind get` command
	getAll := getCmd.Bool("all", false, "Get all libs")
	getLib := getCmd.String("Lib", "", "Get lib By name")

	//'anywind add' subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for `anywind add` command
	addName := addCmd.String("name", "", "lib Name")
	addVersion := addCmd.String("version", "", "lib Version")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "get": // if its the 'get' command
		HandleGet(getCmd, getAll, getLib)
	case "add": // if its the 'add' command
		HandleAdd(addCmd, addName, addVersion)
	default: // if we don't understand the input
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, name *string) {

	getCmd.Parse(os.Args[2:])

	if *all == false && *name == "" {
		fmt.Print("name is required or specify --all for all libs")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		//return all libs
		libs := getLibs()

		fmt.Printf("Name \t Version\n")
		for _, lib := range libs {
			fmt.Printf("%v \t %v \n", lib.Name, lib.Version)
		}

		return
	}

	// return by name
	if *name != "" {
		libs := getLibs()
		Name := *name
		for _, lib := range libs {
			if Name == lib.Name {
				fmt.Printf("Name \t Version\n")
				fmt.Printf("%v \t %v \n", lib.Name, lib.Version)
			}
		}
	}

}

// Validate Anywind add
func ValidateLib(addCmd *flag.FlagSet, name *string, version *string) {

	addCmd.Parse(os.Args[2:])

	if *name == "" || *version == "" {
		fmt.Print("all fields are required for adding a lib")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, name *string, version *string) {

	ValidateLib(addCmd, name, version)

	lib := lib{
		Name:    *name,
		Version: *version,
	}

	libs := getLibs()
	libs = append(libs, lib)

	saveLibs(libs)

}
