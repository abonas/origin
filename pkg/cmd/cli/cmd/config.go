package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd/config"

	cmdconfig "github.com/openshift/origin/pkg/cmd/cli/config"
)

func NewCmdConfig(parentName, name string) *cobra.Command {
	pathOptions := &config.PathOptions{
		LocalFile:  cmdconfig.OpenShiftConfigFileName,
		GlobalFile: path.Join(os.Getenv("HOME"), cmdconfig.OpenShiftConfigHomeDirFileName),
		EnvVar:     cmdconfig.OpenShiftConfigPathEnvVar,
		EnvVarFile: os.Getenv(cmdconfig.OpenShiftConfigPathEnvVar),

		ExplicitFileFlag: cmdconfig.OpenShiftConfigFlagName,

		LoadingRules: cmdconfig.NewOpenShiftClientConfigLoadingRules(),
	}
	pathOptions.LoadingRules.DoNotResolvePaths = true

	cmd := config.NewCmdConfig(pathOptions, os.Stdout)
	cmd.Short = "Change configuration files for the client"
	cmd.Long = fmt.Sprintf(`Manages the OpenShift config files using subcommands like:

%[1]s %[2]s use-context my-context
%[1]s %[2]s set preferences.some true

Reference: https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/kubeconfig-file.md`, parentName, name)

	return cmd
}
