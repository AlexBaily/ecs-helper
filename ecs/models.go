package ecs

//EcsCluster struct for holding information on an ECS cluster.
type EcsCluster struct {
	ClusterArn 	string `json:ClusterArn`
	ClusterName string `json:ClusterName`
	Status		string `json:Status`
	Services	[]EcsService `json:Services`
}

type EcsService struct {
	ServiceName	 string `json:ServiceName`
	LaunchType   string `json:LaunchType`
	RunningCount int64 `json:RunningCount`
	DesiredCount int64 `json:DesiredCount`
	PendingCount int64 `json:PendingCount`
}