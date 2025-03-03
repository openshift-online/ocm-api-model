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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/openshift-online/ocm-api-metamodel/cmd/metamodel/check"
	"github.com/openshift-online/ocm-api-metamodel/cmd/metamodel/generate"
	"github.com/openshift-online/ocm-api-metamodel/cmd/metamodel/version"
)

// Root command:
var root = &cobra.Command{
	Use:  "metamodel",
	Long: "OCM metamodel tool.",
}

func init() {
	// Register the options that are managed by the 'flag' package, so that they will also be parsed
	// by the 'pflag' package:
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Register the sub-commands:
	root.AddCommand(check.Cmd)
	root.AddCommand(generate.Cmd)
	root.AddCommand(version.Cmd)
}

func main() {
	// This is needed to make `glog` believe that the flags have already been parsed, otherwise
	// every log messages is prefixed by an error message stating the the flags haven't been
	// parsed.
	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse empty command line to satisfy 'glog': %v\n", err)
		os.Exit(1)
	}

	// Execute the root command:
	root.SetArgs(os.Args[1:])
	err = root.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute root command: %v\n", err)
		os.Exit(1)
	}
}
