package mutate

import (
    "github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	"context"
	"fmt"
)


type Ik8sManager interface {
	CreateConfigMap(name string, namespace string, data map[string]string) bool
}

type K8sManager struct {
	Client *k8s.Client
}

func (k  K8sManager) CreateConfigMap(name string, namespace string, data map[string]string) bool {
	configMap := &corev1.ConfigMap{
		Metadata: &metav1.ObjectMeta{
			Name:      k8s.String(name),
			Namespace: k8s.String(namespace),
		},
		Data: data,
	}
	if err := k.Client.Create(context.Background(), configMap); err != nil {
		fmt.Println(err)
		if err := k.Client.Update(context.Background(), configMap); err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return true
}
