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
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/dark-shade/go-setup/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	full       bool
	ops        bool
	license    string
	location   string
	author     string
	modulePath string
	profiles   []string
	config     bool
)

//go:embed data/*
var f embed.FS

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a project",
	Long:  `Initializes a project by adding recommended directory structure and files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check location exists
		locationExists, err := utils.Exists(location)
		if err != nil {
			utils.CheckErrFatal(err)
		}

		if !locationExists {
			utils.CheckErrFatal(errors.New("location to initialize project doesn't exist"))
		}

		// create .go-setup directory structure in user home
		homeDirPath, err := os.UserHomeDir()
		if err != nil {
			utils.CheckErrNonFatal(err)
		}

		if err := os.MkdirAll(filepath.Join(homeDirPath, ".go-setup", "profiles"), os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		} else {
			fmt.Println("Config and profiles path setup up at " + filepath.Join(homeDirPath, ".go-setup", "profiles"))
		}

		if config {
			return
		}

		// location for file definitions and data
		var profileLoc string

		if len(profiles) > 1 || profiles[0] != "default" {
			// get home directory
			homeDirPath, err := os.UserHomeDir()
			if err != nil {
				utils.CheckErrNonFatal(err)
			}

			fmt.Println("Copying profiles...")

			for _, profile := range profiles {
				fmt.Println("Copying profile: " + profile)

				profileLoc = filepath.Join(homeDirPath, ".go-setup", "profiles", profile)

				utils.CopyDirectory(profileLoc, location)
			}
		}

		if full {
			fmt.Println("Setting up full-scale project structure...")
			bareSetup()
			opsSetup()
			remainderSetup()
			fmt.Println("Finished full-scale project structure setup")
		} else {
			fmt.Println("Setting up bare-minimum project structure...")
			bareSetup()
			fmt.Println("Finished bare-minimum project structure setup")
		}

		if ops {
			fmt.Println("Setting up operations project structure...")
			opsSetup()
			fmt.Println("Finished operations project structure setup")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// local flags for initCmd
	initCmd.Flags().BoolVarP(&full, "full", "f", false, "initializes all files and directories in the recommend layout")
	initCmd.Flags().BoolVarP(&ops, "ops", "o", false, "initializes all the operations related files (also initializes bare-minimum setup)")
	initCmd.Flags().StringVarP(&license, "license", "i", "mit", "initializes the license")
	initCmd.Flags().StringVarP(&location, "location", "l", ".", "location for project structure setup")
	initCmd.Flags().StringVarP(&author, "author", "a", "", "author name and email, e.g. Jane Doe jane.doe@gmail.com")
	initCmd.Flags().StringVarP(&modulePath, "moduleP-path", "m", ".", "module path for go mod init")
	initCmd.Flags().StringSliceVarP(&profiles, "profile", "p", []string{"default"}, "profile to use for project setup")
	initCmd.Flags().BoolVarP(&config, "config", "c", false, "initializes the ~/.go-setup/profiles path")

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
	// we cannot just return errors since these are all non-fatal errors
	// setup directories
	if err := os.Mkdir(filepath.Join(location, "bin"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "configs"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "docs"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "examples"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "pkg"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "scripts"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.MkdirAll(filepath.Join(location, "test", "data"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	// setup files
	/// main.go file
	mainData, err := f.ReadFile(filepath.Join(".", "data", "main.go"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, "main.go"), mainData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}

	/// .gitignore file
	gitignoreData, err := f.ReadFile(filepath.Join(".", "data", ".gitignore"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, ".gitignore"), gitignoreData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}

	/// Makefile file
	makefileData, err := f.ReadFile(filepath.Join(".", "data", "Makefile"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, "Makefile"), makefileData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}

	/// README.md file
	readmeData, err := f.ReadFile(filepath.Join(".", "data", "README.md"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, "README.md"), readmeData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}

	/// LICENSE file
	if license == "mit" {
		licenseData, err := f.ReadFile(filepath.Join(".", "data", "licenses", "mit"))
		if err != nil {
			utils.CheckErrNonFatal(err)
		} else {
			if err := utils.CreateFile(filepath.Join(location, "LICENSE"), licenseData, os.ModePerm); err != nil {
				utils.CheckErrNonFatal(err)
			}
		}
	} else if license == "apache" {
		licenseData, err := f.ReadFile(filepath.Join(".", "data", "licenses", "apache"))
		if err != nil {
			utils.CheckErrNonFatal(err)
		} else {
			if err := utils.CreateFile(filepath.Join(location, "LICENSE"), licenseData, os.ModePerm); err != nil {
				utils.CheckErrNonFatal(err)
			}
		}
	} else {
		fmt.Println("Error: Invalid license: " + license + ". Valid values are mit or apache")
	}

	/// go.mod file
	r, _ := regexp.Compile(`\d.\d+`)

	goModData := "module github.com/dark-shade/temp\ngo " + r.FindString(runtime.Version())
	if err := utils.CreateFile(filepath.Join(location, "go.mod"), []byte(goModData), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	/// CHANGELOG.md file
	changelogData, err := f.ReadFile(filepath.Join(".", "data", "CHANGELOG.md"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, "CHANGELOG.md"), changelogData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}
}

// opsSetup setups the ops structure
func opsSetup() {
	// we cannot just return errors since these are all non-fatal errors

	// create deployment related files
	/// Dockerfile file
	dockerfileData, err := f.ReadFile(filepath.Join(".", "data", "Dockerfile"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, "Dockerfile"), dockerfileData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}

	/// Jenkinsfile file
	jenkinsfileData, err := f.ReadFile(filepath.Join(".", "data", "Jenkinsfile"))
	if err != nil {
		utils.CheckErrNonFatal(err)
	} else {
		if err := utils.CreateFile(filepath.Join(location, "Jenkinsfile"), jenkinsfileData, os.ModePerm); err != nil {
			utils.CheckErrNonFatal(err)
		}
	}
}

// remainderSetup setups up the full structure excluding the bareSetup and opsSetup
func remainderSetup() {
	// we cannot just return errors since these are all non-fatal errors
	// setup directories
	if err := os.Mkdir(filepath.Join(location, "api"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "assets"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "build"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "cmd"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "deployments"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "githooks"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "init"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "internal"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "third_party"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "tools"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "web"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}

	if err := os.Mkdir(filepath.Join(location, "website"), os.ModePerm); err != nil {
		utils.CheckErrNonFatal(err)
	}
}
