package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"
)

var genCmd = &cobra.Command{
	Use:  "gen",
	RunE: runGen,
}

func runGen(cmd *cobra.Command, args []string) error {
	b, err := ioutil.ReadFile("assets/wakatime.ico")
	if err != nil {
		return err
	}

	f, err := os.Create("assets/generated.go")
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("assets").Parse(templateStr))
	err = tmpl.Execute(f, map[string]interface{}{
		"Package": "assets",
		"Icon":    genByteArrayStr(b),
	})
	return err
}

func genByteArrayStr(b []byte) string {
	bf := bytes.NewBuffer(nil)
	for i := range b {
		bf.WriteString(strconv.Itoa(int(b[i])))
		bf.WriteRune(',')
	}
	return fmt.Sprintf("[]byte{%v}", bf.String())
}

func init() {
	rootCmd.AddCommand(genCmd)
}

var templateStr = `// Code generated. DO NOT EDIT.
package {{ .Package }}

func Icon() []byte {
  return {{ .Icon }}
}
`
