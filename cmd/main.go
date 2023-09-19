package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EdsonHTJ/jtos"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	var outputDir string
	var packageName string
	var mainStructName string
	var gen string

	rootCmd := &cobra.Command{
		Use:   "jtos [path/to/file]",
		Short: "jtos is a json to struct generator",
		Long:  "jtos is a json to struct generator",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inputFile := args[0]
			if packageName == "" {
				filenameWithExt := filepath.Base(inputFile)
				extension := filepath.Ext(filenameWithExt)
				packageName = filenameWithExt[0 : len(filenameWithExt)-len(extension)]

			}

			if mainStructName == "" {
				mainStructName = packageName
			}

			response, err := jtos.ParseJsonFile(packageName, mainStructName, inputFile, jtos.GOLANG_GENERATOR)
			if err != nil {
				fmt.Println("error on parse:", err)
				os.Exit(1)
			}

			outPath := filepath.Join(outputDir, response.RecomendedPath)
			os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
			f, err := os.Create(outPath)
			if err != nil {
				fmt.Println("error on create file:", err)
				os.Exit(1)
			}

			_, err = f.WriteString(response.Output)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&outputDir, "out", "o", "./", "path to the output dir")
	rootCmd.PersistentFlags().StringVarP(&gen, "gen", "g", "go", "generator type")
	rootCmd.PersistentFlags().StringVarP(&mainStructName, "struct", "s", "", "name of output structure")
	rootCmd.PersistentFlags().StringVarP(&packageName, "package", "p", "", "name of output package")
	return rootCmd
}

func main() {
	root := Root()
	root.Execute()
}
