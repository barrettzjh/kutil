package main

import (
	"github.com/barrettzjh/kutil/model/resource"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "kutil",
		Short: "kubernetes cli util",
		Long:  `一个日常工作时会用到的小工具,可以解决复杂一点的工作的相对自动化的命令行工具`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().Help()
		},
	}
	rootCmd.AddCommand(resource.Cmd)
	rootCmd.Execute()
}
