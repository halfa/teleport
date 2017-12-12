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
	"time"

	"github.com/gravitational/teleport/lib/services"
)

var serverFixture = []services.Server{
	&services.ServerV2{
		Metadata: services.Metadata{
			Labels: map[string]string{
				"os": "gentoo",
			},
		},
		Spec: services.ServerSpecV2{
			CmdLabels: map[string]services.CommandLabelV2{
				"time": services.CommandLabelV2{
					Period:  services.NewDuration(time.Second),
					Command: []string{"time"},
					Result:  "now",
				},
			},
		},
	},
	&services.ServerV2{
		Metadata: services.Metadata{
			Labels: map[string]string{
				"os":   "coreos",
				"role": "database",
			},
		},
		Spec: services.ServerSpecV2{
			CmdLabels: map[string]services.CommandLabelV2{
				"time": services.CommandLabelV2{
					Period:  services.NewDuration(time.Second),
					Command: []string{"time"},
					Result:  "now",
				},
			},
		},
	},
	&services.ServerV2{
		Metadata: services.Metadata{
			Labels: map[string]string{
				"os":   "plan9",
				"role": "database",
			},
		},
		Spec: services.ServerSpecV2{
			CmdLabels: map[string]services.CommandLabelV2{
				"time": services.CommandLabelV2{
					Period:  services.NewDuration(time.Second),
					Command: []string{"time"},
					Result:  "now",
				},
			},
		},
	},
}

func DynamicInventoryHostTest() {

}
