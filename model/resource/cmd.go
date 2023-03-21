package resource

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CmdOptions struct {
	Namespace string
	Label     string
	Type      string
}

const (
	defaultNamespace = "default"
	defaultLabel     = "all"
)


var (
	createCmdOpts CmdOptions
	deleteCmdOpts CmdOptions
	modifyCmdOpts CmdOptions
	listCmdOpts CmdOptions
)

func init() {
	Cmd.AddCommand(createCmd, deleteCmd, modifyCmd, listCmd)
	modifyCmd.Flags().StringVarP(&createCmdOpts.Namespace, "namespace", "n", defaultNamespace, "Namespace of the kubernetes")
	modifyCmd.Flags().StringVarP(&createCmdOpts.Label, "label", "l", defaultLabel, "request or limit")
	modifyCmd.Flags().StringVarP(&createCmdOpts.Type, "type", "t", "", "cpu or memory")

	deleteCmd.Flags().StringVarP(&deleteCmdOpts.Namespace, "namespace", "n", defaultNamespace, "Namespace of the kubernetes")

	listCmd.Flags().StringVarP(&listCmdOpts.Namespace, "namespace", "n", defaultNamespace, "Namespace of the kubernetes")

	createCmd.Flags().StringVarP(&modifyCmdOpts.Namespace, "namespace", "n", defaultNamespace, "Namespace of the kubernetes")

}


var Cmd = &cobra.Command{
	Use:   "resource [modify] [deploy] [100Mi]",
	Short: "modify, create or delete resource",
	Long:  `可通过简短的命令行，来创建,修改或删除deploy的Resource配置`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var modifyCmd = &cobra.Command{
	Use:   "modify [deploy] [100Mi]",
	Short: "modify resource",
	Long:  `可以修改某个命名空间下某个deploy的资源限制`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := modify(args[1], modifyCmdOpts.Namespace, args[2], modifyCmdOpts.Label, modifyCmdOpts.Type); err != nil{
			fmt.Println(err.Error())
			return
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create [deploy] [100Mi]",
	Short: "create resource",
	Long:  `可以创建某个命名空间下某个deploy的资源限制`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := create(args[1], createCmdOpts.Namespace, args[2]); err != nil{
			fmt.Println(err.Error())
			return
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [deploy]",
	Short: "delete resource",
	Long:  `可以删除某个命名空间下某个deploy的资源限制`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := delete(args[1], deleteCmdOpts.Namespace); err != nil{
			fmt.Println(err.Error())
			return
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list [namespace]",
	Short: "list resource",
	Long:  `可以查询某个命名空间下所有deploy的资源限制`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := list(listCmdOpts.Namespace); err != nil{
			fmt.Println(err.Error())
			return
		}
	},
}


