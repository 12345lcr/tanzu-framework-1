// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	cliapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/cli/v1alpha1"
)

func newInfoCmd(desc *cliapi.PluginDescriptor) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "info",
		Short:  "Plugin info",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := json.Marshal(desc)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			return nil
		},
	}

	return cmd
}
