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
	"strings"

	"github.com/absolutedevops/civo/api"
	"github.com/absolutedevops/civo/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var instanceCreateName string
var instanceCreateSize string
var instanceCreateRegion string
var instanceCreateSSHKeyID string
var instanceCreatePublicIP bool
var instanceCreateTemplate string
var instanceCreateFirewallID string
var instanceCreateNetwork string
var instanceCreateInitialUser string
var instanceCreateTags string

var instanceCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "build"},
	Short:   "Create a new instance",
	Example: "civo instance create --name test1.example.com --size g1.small --region svg1 --ssh-key-id 1234567890",
	Long:    `Create a new instance with the described specification under your current account`,
	Run: func(cmd *cobra.Command, args []string) {
		if instanceCreateName == "" {
			instanceCreateName = api.InstanceSuggestHostname()
		}
		if instanceCreateRegion == "" {
			instanceCreateRegion = config.DefaultRegion()
		}
		if instanceCreateSSHKeyID != "" {
			instanceCreateSSHKeyID = api.SshKeyFind(instanceCreateSSHKeyID)
		}

		params := api.InstanceParams{
			Name:        instanceCreateName,
			Size:        instanceCreateSize,
			Region:      instanceCreateRegion,
			SSHKeyID:    instanceCreateSSHKeyID,
			Template:    instanceCreateTemplate,
			InitialUser: instanceCreateInitialUser,
			PublicIP:    instanceCreatePublicIP,
			NetworkID:   instanceCreateNetwork,
			FirewallID:  instanceCreateFirewallID,
			Tags:        instanceCreateTags,
		}
		res, err := api.InstanceCreate(params)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		hostname := res.S("hostname").Data().(string)
		ID := res.S("id").Data().(string)
		parts := strings.Split(ID, "-")
		fmt.Printf("Building instance called `%s` with ID %s\n", hostname, parts[0])
	},
}

func init() {
	instanceCreatePublicIP = true
	instanceCmd.AddCommand(instanceCreateCmd)
	instanceCreateCmd.Flags().StringVarP(&instanceCreateName, "name", "n", "", "Name of the instance; lowercase, hyphen separated. If you don't specify one, a random one will be used.")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateSize, "size", "s", "g1.small", "The size from 'civo sizes'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateTags, "tags", "g", "", "A space-separated list of tags to use")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateRegion, "region", "r", "", "The region from 'civo regions'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateSSHKeyID, "ssh-key-id", "k", "default", "The SSH key ID from 'civo sshkeys'")
	instanceCreateCmd.Flags().BoolVarP(&instanceCreatePublicIP, "public-ip", "p", true, "Should a public IP address be allocated")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateTemplate, "template", "t", "ubuntu-16.04", "The template from 'civo templates'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateFirewallID, "firewall", "f", "default", "The firewall ID or name from 'civo firewalls'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateNetwork, "network", "w", "Default", "The network ID or name from 'civo networks'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateInitialUser, "initial-user", "u", "civo", "The default user to create")
}
