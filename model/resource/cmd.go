package resource

import (
	"fmt"
	"github.com/spf13/cobra"
)

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
		ns, _ := cmd.Flags().GetString("namespace")
		lb, _ := cmd.Flags().GetString("label")
		tp, _ := cmd.Flags().GetString("type")
		if err := Modify(args[1], ns, args[2], lb, tp); err != nil{
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
		ns, _ := cmd.Flags().GetString("namespace")
		if err := Create(args[1], ns, args[2]); err != nil{
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
		ns, _ := cmd.Flags().GetString("namespace")
		if err := Delete(args[1], ns); err != nil{
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
		ns, _ := cmd.Flags().GetString("namespace")
		if err := List(ns); err != nil{
			fmt.Println(err.Error())
			return
		}
	},
}


func init(){
	Cmd.Flags().StringP("namespace", "n", "default", "Namespace of the kubernetes")
	Cmd.Flags().StringP("label", "l", "all", "request or limit")
	Cmd.Flags().StringP("type", "t", "", "cpu or memory")
	Cmd.AddCommand(createCmd, deleteCmd, modifyCmd, listCmd)
}
