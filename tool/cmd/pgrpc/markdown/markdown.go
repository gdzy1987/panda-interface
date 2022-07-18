package markdown

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/XM-GO/panda-interface/tool/pkg/markdown"
	"github.com/spf13/cobra"
)

// CmdMarkdown represents the new command.
var CmdMarkdown = &cobra.Command{
	Use:   "markdown",
	Short: "Create a markdown document.",
	Long:  "Create a markdown document using the repository template. Example: panda-tool markdown -f ../pandademo/api/apidocs.swagger.json -t ../pandademo/third_party/markdown-templates/ -o ../panda-docs/docs/api/Greeter -m all",
	Run:   run,
}

var (
	mode         string
	swaggerPath  string
	templatePath string
	outputPath   string
)

func init() {
	CmdMarkdown.Flags().StringVarP(&swaggerPath, "swagger", "f", "", "swagdown requires exactly one argument, a path to a local swagger file(.json|.yml|.yaml)")
	CmdMarkdown.Flags().StringVarP(&templatePath, "template", "t", "", "template path, like (markdown/templates/)")
	CmdMarkdown.Flags().StringVarP(&outputPath, "output", "o", "docs", "output path,  default is stdio")
	CmdMarkdown.Flags().StringVarP(&mode, "mode", "m", "tag", "mode(all | tag | method)")

	CmdMarkdown.MarkFlagRequired("swaggerPath")

}

func run(cmd *cobra.Command, args []string) {

	switch mode {
	case "tag": // 文档目录
		Render(swaggerPath, templatePath, outputPath, "tag")
	case "method":
		Render(swaggerPath, templatePath, outputPath, "method")
	case "all": // 文档目录
		Render(swaggerPath, templatePath, outputPath, "tag")
		Render(swaggerPath, templatePath, outputPath, "method")
	default:
		exit(errors.New("error mode"))
	}
}

func Render(sPath, tPath, oPath, docsMode string) {
	f, err := os.Open(sPath)
	if err != nil {
		exit(fmt.Errorf("error find %s: %w", sPath, err))
	}
	tmpl := markdown.New(tPath, docsMode)

	switch filepath.Ext(f.Name()) {
	case ".json":
		if err := markdown.RenderFromJSON(markdown.NewWriter(oPath), f, tmpl); err != nil {
			exit(err)
		}
	case ".yml", ".yaml":
		if err := markdown.RenderFromYAML(markdown.NewWriter(oPath), f, tmpl); err != nil {
			exit(err)
		}
	default:
		exit(errors.New("swagdown only recognizes JSON or YAML Swagger files"))
	}
	fmt.Printf("Rendering %s markdown from '%s' is complete.\n", docsMode, sPath)
}

func exit(err error) {
	fmt.Println("error: ", err)
	os.Exit(1)
}
