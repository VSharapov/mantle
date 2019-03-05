package secret

import (
  "fmt"
	v1 "k8s.io/api/core/v1"
)

type SecretWrapper struct {
	Secret Secret `json:"secret,omitempty"`
}

type Secret struct {
	Version     string            `json:"version,omitempty"`
	Cluster     string            `json:"cluster,omitempty"`
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	StringData  map[string]string `json:"string_data,omitempty"`
	Data        map[string][]byte `json:"data,omitempty"`
	SecretType  SecretType        `json:"type,omitempty"`
}

type SecretType string

const (
	SecretTypeOpaque              SecretType = "Opaque"
	SecretTypeServiceAccountToken SecretType = "kubernetes.io/service-account-token"
	SecretTypeDockercfg           SecretType = "kubernetes.io/dockercfg"
	SecretTypeDockerConfigJson    SecretType = "kubernetes.io/dockerconfigjson"
	SecretTypeBasicAuth           SecretType = "kubernetes.io/basic-auth"
	SecretTypeSSHAuth             SecretType = "kubernetes.io/ssh-auth"
	SecretTypeTLS                 SecretType = "kubernetes.io/tls"
)

func (s SecretType) ToString() string {
  return string(s)
}

func CompareSecretTypes(secret1, secret2 interface{}) (bool, error) {
  s1str, s2str := "", ""
  sV1, ok1 := secret1.(v1.SecretType)
  if ok1 { s1str = string(sV1) }
  sMantle, ok2 := secret1.(SecretType)
  if ok2 { s1str = string(sMantle) }
  if !ok1 && !ok2 {
    return false, fmt.Errorf("%s is not a k8s.v1 secretType nor a mantle SecretType", secret1)
  }

  sV1, ok1 = secret2.(v1.SecretType)
  if ok1 { s2str = string(sV1) }
  sMantle, ok2 = secret2.(SecretType)
  if ok2 { s2str = string(sMantle) }
  if !ok1 && !ok2 {
    return false, fmt.Errorf("%s is not a k8s.v1 secretType nor a mantle SecretType", secret2)
  }
  
  return s1str == s2str, nil
}


