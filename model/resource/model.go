package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/barrettzjh/kutil/model"
	v12 "k8s.io/api/core/v1"
	resource2 "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func modify(deployname, namespace, value, label, tp string)error{
	deploy, err := model.Client.AppsV1().Deployments(namespace).Get(context.TODO(), deployname, v1.GetOptions{})
	if err != nil{
		return err
	}

	if label == "request"{
		deploy.Spec.Template.Spec.Containers[0].Resources.Requests[v12.ResourceName(tp)] = resource2.MustParse(value)
	}else if label == "limit"{
		deploy.Spec.Template.Spec.Containers[0].Resources.Limits[v12.ResourceName(tp)] = resource2.MustParse(value)
	}else if label == "all" {
		deploy.Spec.Template.Spec.Containers[0].Resources.Requests[v12.ResourceName(tp)] = resource2.MustParse(value)
		deploy.Spec.Template.Spec.Containers[0].Resources.Limits[v12.ResourceName(tp)] = resource2.MustParse(value)
	}

	_, err = model.Client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, v1.UpdateOptions{})
	if err != nil{
		return err
	}
	return nil
}
func create(deployname, namespace, value string)error{
	deploy, err := model.Client.AppsV1().Deployments(namespace).Get(context.TODO(), deployname, v1.GetOptions{})
	if err != nil{
		return err
	}

	var resource Resource
	err = json.Unmarshal([]byte(value), &resource)
	if err != nil{
		return err
	}

	resourcerequirment := v12.ResourceRequirements{
		Limits: v12.ResourceList{
			"cpu": resource2.MustParse(resource.Limits.CPU),
			"memory": resource2.MustParse(resource.Limits.Memory),
		},
		Requests: v12.ResourceList{
			"cpu": resource2.MustParse(resource.Requests.CPU),
			"memory": resource2.MustParse(resource.Requests.Memory),
		},
	}
	deploy.Spec.Template.Spec.Containers[0].Resources = resourcerequirment

	_, err = model.Client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, v1.UpdateOptions{})
	if err != nil{
		return err
	}
	return nil
}
func delete(deployname, namespace string)error{
	deploy, err := model.Client.AppsV1().Deployments(namespace).Get(context.TODO(), deployname, v1.GetOptions{})
	if err != nil{
		return err
	}

	deploy.Spec.Template.Spec.Containers[0].Resources =v12.ResourceRequirements{}

	_, err = model.Client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, v1.UpdateOptions{})
	if err != nil{
		return err
	}
	return nil
}

func list(namespace string)error{
	deploy, err := model.Client.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil{
		return err
	}

	for _, v := range deploy.Items{
		data, err := json.Marshal(v.Spec.Template.Spec.Containers[0].Resources)
		if err != nil{
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("%s\t%s\n", v.Name, string(data))
	}
	return nil
}
