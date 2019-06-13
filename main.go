package main

import (
	"fmt"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
)

func main() {
	fmt.Println(v1beta1.InitConfiguration{})
	fmt.Println(v1beta1.JoinConfiguration{})
}