// Copyright © 2016 Absolute DevOps Ltd <info@absolutedevops.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/absolutedevops/civo/api"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// dnsDomainCmd represents the accounts command
var dnsDomainCmd = &cobra.Command{
	Use:     "domain",
	Aliases: []string{"dns", "domains"},
	Short:   "List all DNS domains",
	Long:    `List the DNS domains for the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.DnsDomainsList()
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"ID", "Name"})
		items, _ := result.Children()
		for _, child := range items {
			table.Append([]string{
				child.S("id").Data().(string),
				child.S("name").Data().(string),
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(dnsDomainCmd)
}
