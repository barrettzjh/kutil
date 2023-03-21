package resource

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "resource [modify] [deploy] [100Mi]",
	Short: "modify, create or delete resource",
	Long:  `可通过简短的命令行，来创建,修改或删除deploy的Resource配置`,
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if !ArgsCheck(args){
			fmt.Println("args Error")
			return
		}
		ns, _ := cmd.Flags().GetString("namespace")
		lb, _ := cmd.Flags().GetString("label")
		tp, _ := cmd.Flags().GetString("type")

		if args[0] == "modify"{
			if err := Modify(args[1], ns, args[2], lb, tp);err != nil{
				fmt.Println(err.Error())
				return
			}
		}else if args[0] == "create"{
			if err := Create(args[1], ns, args[2], lb, tp);err != nil{
				fmt.Println(err.Error())
				return
			}

		}else if args[0] == "delete"{
			if err := Delete(args[1], ns, args[2], lb, tp);err != nil{
				fmt.Println(err.Error())
				return
			}
		}

		fmt.Printf("deploy %s %s Success!\n", args[1], args[0])
	},
}

func init(){
	Cmd.Flags().StringP("namespace", "n", "default", "Namespace of the kubernetes")
	Cmd.Flags().StringP("label", "l", "all", "request or limit")
	Cmd.Flags().StringP("type", "t", "", "cpu or memory")
}
