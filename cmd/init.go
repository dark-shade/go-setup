/*
Copyright © 2021 Sankul Rawat sankul.rawat.28@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	full     bool
	ops      bool
	license  string
	location string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a project",
	Long:  `Initializes a project by adding recommended directory structure and files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debug: init called")

		// TODO check location values

		if full {
			fmt.Println("debug: init full called")
			bareSetup()
			opsSetup()
			remainderSetup()
		} else {
			fmt.Println("debug: init full not called, init bare (default)")
			bareSetup()
		}

		if ops {
			fmt.Println("debug: init ops called")
			opsSetup()
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// local flags for initCmd
	initCmd.Flags().BoolVarP(&full, "full", "f", false, "initializes all files and directories in the recommend layout")
	initCmd.Flags().BoolVarP(&ops, "ops", "o", false, "initializes all the operations related files")
	initCmd.Flags().StringVarP(&license, "license", "lic", "mit", "initializes the license")
	initCmd.Flags().StringVarP(&location, "location", "loc", location, "location for project structure setup")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// bareSetup setups the bare structure
func bareSetup() {
	// setup directories
	os.Mkdir(filepath.Join(location, "bin"), os.ModePerm)
	os.Mkdir(filepath.Join(location, "configs"), os.ModePerm)
	os.Mkdir(filepath.Join(location, "docs"), os.ModePerm)
	os.Mkdir(filepath.Join(location, "examples"), os.ModePerm)
	os.Mkdir(filepath.Join(location, "pkg"), os.ModePerm)
	os.Mkdir(filepath.Join(location, "scripts"), os.ModePerm)
	os.MkdirAll(filepath.Join(location, "test", "data"), os.ModePerm)

	// setup files
	mainData, _ := os.ReadFile(filepath.Join(".", "data", "main"))
	os.WriteFile(filepath.Join(location, "main.go"), mainData, os.ModePerm)

	gitignoreData, _ := os.ReadFile(filepath.Join(".", "data", "gitignore"))
	os.WriteFile(filepath.Join(location, ".gitignore"), gitignoreData, os.ModePerm)

	makefileData, _ := os.ReadFile(filepath.Join(".", "data", "makefile"))
	os.WriteFile(filepath.Join(location, "Makefile"), makefileData, os.ModePerm)

	readmeData, _ := os.ReadFile(filepath.Join(".", "data", "readme"))
	os.WriteFile(filepath.Join(location, "README.md"), readmeData, os.ModePerm)

	if license == "mit" {
		licenseData, _ := os.ReadFile(filepath.Join(".", "data", "license", "mit"))
		os.WriteFile(filepath.Join(location, "LICENSE"), licenseData, os.ModePerm)
	} else if license == "apache" {
		licenseData, _ := os.ReadFile(filepath.Join(".", "data", "license", "apache"))
		os.WriteFile(filepath.Join(location, "LICENSE"), licenseData, os.ModePerm)
	} else {
		fmt.Println("Error: Invalid license: " + license + ". Valid values are mit or apache")
	}

}

// opsSetup setups the ops structure
func opsSetup() {

}

// remainderSetup setups up the full structure excluding the bareSetup and opsSetup
func remainderSetup() {

}
