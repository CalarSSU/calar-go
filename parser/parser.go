package parser

import (
	"github.com/spf13/cobra"
)

type Config struct {
	Department string
	Education  string
	Group      string
	Subgroups  []string
}

func ParseArguments(cfg *Config) {
	var cmdConfig = &cobra.Command{
		Use:   "",
		Short: "calar arguments",
		Run: func(cmd *cobra.Command, args []string) {
			cfg.Department, _ = cmd.Flags().GetString("department")
			cfg.Education, _ = cmd.Flags().GetString("education")
			cfg.Group, _ = cmd.Flags().GetString("group")
			cfg.Subgroups, _ = cmd.Flags().GetStringSlice("subgroups")
		},
	}

	cmdConfig.Flags().StringP("department", "d", "", "department token")
	cmdConfig.Flags().StringP("education", "e", "full", "type of eduction")
	cmdConfig.Flags().StringP("group", "g", "", "group number")
	cmdConfig.Flags().StringSliceP("subgroups", "s", []string{},
		"list of subgroups")
	cmdConfig.Execute()
}
