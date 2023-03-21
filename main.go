package main

import (
	"context"
	"fmt"
	"github.com/barrettzjh/kutil/model"
	"github.com/barrettzjh/kutil/model/resource"
	"github.com/spf13/cobra"
	v12 "k8s.io/api/core/v1"
	resource2 "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "kutil",
		Short: "kubernetes cli util",
		Long:  `一个日常工作时会用到的小工具,可以解决复杂一点的工作的相对自动化的命令行工具`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	var resourceCmd = &cobra.Command{
		Use:   "resource [modify] [deploy] [100Mi]",
		Short: "modify, create or delete resource",
		Long:  `可通过简短的命令行，来创建,修改或删除deploy的Resource配置`,
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			if !resource.ArgsCheck(args){
				fmt.Println("args Error")
				return
			}
			_namespace, _ := cmd.Flags().GetString("namespace")
			//_label, _ := cmd.Flags().GetString("label")
			_type, _ := cmd.Flags().GetString("type")

			fmt.Println(_namespace,_type)
			deploy, err := model.Client.AppsV1().Deployments(_namespace).Get(context.TODO(), args[1], v1.GetOptions{})
			if err != nil{
				fmt.Println(err.Error())
				return
			}

			//resourcerequirment := v12.ResourceRequirements{
			//
			//}
			//
			//resourcerequirment := v12.ResourceRequirements{
			//	Limits: v12.ResourceList{
			//		"cpu": resource2.MustParse(),
			//		"memory": resource2.MustParse(),
			//	},
			//	Requests: v12.ResourceList{
			//		"cpu": resource2.MustParse(),
			//		"memory": resource2.MustParse(),
			//	},
			//}

			//deploy.Spec.Template.Spec.Containers[0].Resources = resourcerequirment
			deploy.Spec.Template.Spec.Containers[0].Resources.Requests[v12.ResourceName(_type)] = resource2.MustParse("100m")
			_, err = model.Client.AppsV1().Deployments(_namespace).Update(context.TODO(), deploy, v1.UpdateOptions{})
			if err != nil{
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("deploy %s Update Success!\n", args[2])
		},
	}

	rootCmd.Flags().StringP("namespace", "n", "default", "Namespace of the kubernetes")
	resourceCmd.Flags().StringP("label", "l", "all", "request or limit")
	resourceCmd.Flags().StringP("type", "t", "", "cpu or memory")
	rootCmd.AddCommand(resourceCmd)

	rootCmd.Execute()
}
