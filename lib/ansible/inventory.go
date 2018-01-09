/*
Copyright 2017 Maximilien Richer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ansible

import (
	"encoding/json"
	"fmt"

	"github.com/gravitational/teleport/lib/services"
)

// Inventory matches the JSON struct needed for DynamicInventoryList
type Inventory struct {
	Groups map[string]Group
}

// Group gather hosts and variables common to them
type Group struct {
	Hosts []string
	Vars  map[string]string
}

// DynamicInventoryList returns a JSON-formated ouput compatible with Ansible --list flag
//
// The JSON output SHOULD HAVE the following format:
// ```json
// {
//     "group_name": {
//         "hosts": ["host1.example.com", "host2.example.com"],
//         "vars": {
//             "a": true
//         }
//     },
// }
// ```
// TODO: Implement `_meta` and host variables?
func DynamicInventoryList(nodes []services.Server) (string, error) {
	hostsByLabels := bufferLabels(nodes)

	var inventory = Inventory{
		Groups: make(map[string]Group),
	}
	for labelDashValue, hosts := range hostsByLabels {
		inventory.Groups[labelDashValue] = Group{
			Hosts: hosts,
			Vars:  make(map[string]string),
		}
	}
	out, err := json.Marshal(inventory)
	if err != nil {
		return "", fmt.Errorf("cannot encode JSON objet: %s", err)
	}
	return string(out), nil
}

// DynamicInventoryHost returns a JSON-formated ouput compatible with Ansible --host <string> flag
//
// (From ansible ref. doc)
// When called with the arguments --host <hostname>, the script must print either an empty JSON hash/dictionary,
// or a hash/dictionary of variables to make available to templates and playbooks.
func DynamicInventoryHost(nodes []services.Server, host string) {
	// filter only the required node
}

// StaticInventory returns an INI-formated ouput compatible with Ansible static inventory format
//
// It crafts groups using the labels associated with each nodes. Each label is build in the form
// <label>-<value> (with a dash in the middle).
func StaticInventory(nodes []services.Server) {
	inventory := make(map[string][]string)
	// get all keys
	for _, n := range nodes {
		// get labels and add to groups
		for label, val := range n.GetAllLabels() {
			// groupName is of the form apache-2.2
			groupName := label + "-" + val
			inventory[groupName] = append(inventory[groupName], n.GetAddr())
		}
	}
	// write one tulpe by keys
	for groupName, nodeIPs := range inventory {
		fmt.Println("[" + groupName + "]")
		for _, IP := range nodeIPs {
			fmt.Println(IP)
		}
	}
}

// bufferLabels gather labels values and create groups associating hosts with identical labels values
func bufferLabels(nodes []services.Server) map[string][]string {
	labelBuffer := make(map[string][]string)
	// get all keys
	for _, n := range nodes {
		// get labels and add to groups
		for label, val := range n.GetAllLabels() {
			// groupName is of the form apache-2.2
			groupName := label + "-" + val
			labelBuffer[groupName] = append(labelBuffer[groupName], n.GetAddr())
		}
	}
	return labelBuffer
}
