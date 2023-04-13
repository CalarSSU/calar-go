package parser

import (
	"github.com/spf13/cobra"
)

type Request struct {
	Department string
	Education  string
	Group      string
	Translator bool
	Subgroups  []string
}

func ParseArguments(request *Request) error {
	// var retError error
	var cmdRequest = &cobra.Command{
		Use:   "./calar-go",
		Short: "calar arguments",
		Run: func(cmd *cobra.Command, args []string) {
			request.Department, _ = cmd.Flags().GetString("department")
			request.Education, _ = cmd.Flags().GetString("education")
			request.Group, _ = cmd.Flags().GetString("group")
			request.Subgroups, _ = cmd.Flags().GetStringSlice("subgroups")
			request.Translator, _ = cmd.Flags().GetBool("translator")
		},
	}
	cmdRequest.Flags().StringP("department", "d", "", "department token")
	cmdRequest.Flags().StringP("education", "e", "full", "type of eduction")
	cmdRequest.Flags().StringP("group", "g", "", "group number")
	cmdRequest.Flags().StringSliceP("subgroups", "s", []string{},
		"list of subgroups")
	cmdRequest.Flags().BoolP("translator", "t", false,
		"additional education")
	return cmdRequest.Execute()

}
