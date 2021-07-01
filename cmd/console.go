/*
Copyright © 2021 Jarvib Ding <ding@ibyte.me>
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

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/higker/mapisto/core"
	"github.com/spf13/cobra"
)

var (
	mode          string
	comoandSymbol = "😃:shell>"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Console interaction",
	Long: color.RedString(`
	console: console interactive mode.`),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Banner)
		fmt.Println("Please select table name.")

		me := core.New(core.Golang)

		me.SetDB(core.NewDB(
			&core.DBInfo{
				HostIP:   "45.76.202.255:3306",
				UserName: "emp_db",
				Password: "TsTkHXDK4xPFtCph",
				DBType:   "mysql",
				Charset:  "utf8",
			},
		))
		fmt.Println(me.DataBase.DBInfo)
		t := prompt.Input(comoandSymbol, completer)
		fmt.Println(me.Generate("emp_db", t))
	},
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "user_info"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
func init() {
	consoleCmd.Flags().StringVarP(&mode, "mode", "m", "", "input `ui` to run in web ui mode")
	rootCmd.AddCommand(consoleCmd)
}
