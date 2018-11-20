package core

import (
	"reflect"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

type ConfigMap struct {
	Version     string            `json:"version,omitempty"`
	Cluster     string            `json:"cluster,omitempty"`
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Data        map[string]string `json:"data,omitempty"`
}

func NewConfigMapFromKubeConfigMap(cm interface{}) (*ConfigMap, error) {
	switch reflect.TypeOf(cm) {
	case reflect.TypeOf(v1.ConfigMap{}):
		obj := cm.(v1.ConfigMap)
		return fromKubeV1(&obj)
	case reflect.TypeOf(&v1.ConfigMap{}):
		return fromKubeV1(cm.(*v1.ConfigMap))
	default:
		return fromKubeV1(cm.(*v1.ConfigMap))
	}
}

func fromKubeV1(kubeConfigMap *v1.ConfigMap) (*ConfigMap, error) {
	cm := &ConfigMap{
		Name:        kubeConfigMap.Name,
		Namespace:   kubeConfigMap.Namespace,
		Version:     kubeConfigMap.APIVersion,
		Cluster:     kubeConfigMap.ClusterName,
		Labels:      kubeConfigMap.Labels,
		Annotations: kubeConfigMap.Annotations,
		Data:        kubeConfigMap.Data,
	}

	return cm, nil
}

func (cm *ConfigMap) ToKube() (runtime.Object, error) {
	switch cm.Version {
	case "v1":
		return cm.toKubeV1()
	default:
		return cm.toKubeV1()
	}
}

func (cm *ConfigMap) toKubeV1() (*v1.ConfigMap, error) {
	kubeConfigMap := &v1.ConfigMap{}

	kubeConfigMap.Name = cm.Name
	kubeConfigMap.Namespace = cm.Namespace
	kubeConfigMap.APIVersion = cm.Version
	kubeConfigMap.ClusterName = cm.Cluster
	kubeConfigMap.Kind = "ConfigMap"
	kubeConfigMap.Labels = cm.Labels
	kubeConfigMap.Annotations = cm.Annotations
	kubeConfigMap.Data = cm.Data

	return kubeConfigMap, nil
}