package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	
	docmodule "github.com/weisurya/go-openapi-converter/modules/document"

)

var (
	convertCmd = &cobra.Command{
		Use:     "convert",
		Short:   "Convert OpenAPI spec",
		Long:    "Convert OpenAPI spec into other formats. Currently, it only supports .docx",
		Run: runConverter,
		Example: "./go-openapi-converter convert -s sample.v1.yaml -t template/standard.docx -o result.sample.docx",
	}
)

const (
	defaultOASPath = "sample.v1.yaml"
	defaultOutputPath = "result.sample.docx"
	defaultTemplatePath = "template/standard.docx"
)

var (
	oasPath string
	outputPath string
	templatePath string
)

func init() {
	convertCmd.Flags().StringVarP(&oasPath, "spec", "s", defaultOASPath,"Open API spec file path")

	convertCmd.Flags().StringVarP(&outputPath, "output", "o", defaultOutputPath, "Output file path")

	convertCmd.Flags().StringVarP(&templatePath, "template", "t", defaultTemplatePath, "Template file path")

	rootCmd.AddCommand(convertCmd)
}

func runConverter(cmd *cobra.Command, args []string) {

	doc, err := docmodule.ReadDocTemplate(templatePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	swagger, err := docmodule.ReadOpenAPI(oasPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	docService := docmodule.NewDocumentHandler(doc, swagger)

	docService.BuildFrontpage()

	docService.BuildTOC()

	docService.BuildPaths()

	if err := doc.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc.SaveToFile(outputPath)

	return
}

