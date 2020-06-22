package module

import "time"

type Server struct {
	Id int
	IP string
	User User
	ServerGrup KubernetesCluster
	Kubernetes string
	CreateTime time.Time
	RunTime time.Time
}

type KubernetesCluster struct {
	Id int
	User  User
	Master []Server
	Node []Server
	ConfigDir string
	CreateTime time.Time
	RunTime time.Time
}
