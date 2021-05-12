package ecs

//EcsCluster struct for holding information on an ECS cluster.
type EcsCluster struct {
	ClusterArn 	string `json:ClusterArn`
	ClusterName string `json:ClusterName`
	Status		string `json:Status`
	Services	[]EcsService `json:Services`
}

type EcsService struct {
	DesiredCount int `json:DesiredCount`
	PendingCount int `json:PendingCount`
}