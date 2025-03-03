/*
Copyright (c) 2019 Red Hat, Inc.

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

package generate

import (
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-api-metamodel/cmd/metamodel/generate/docs"
	"github.com/openshift-online/ocm-api-metamodel/cmd/metamodel/generate/golang"
	"github.com/openshift-online/ocm-api-metamodel/cmd/metamodel/generate/openapi"
)

// Cmd is the definition of the command:
var Cmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code",
	Long:  "Generate code.",
}

func init() {
	// Register the sub-commands:
	Cmd.AddCommand(docs.Cmd)
	Cmd.AddCommand(golang.Cmd)
	Cmd.AddCommand(openapi.Cmd)
}
