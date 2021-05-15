package ecs

import (
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
)
//DescribeEcsClusters will take an input of a string and return a list of clusters.
//Returns: []EcsCluster .
func (e EcsInt) DescribeEcsClusters(cluster string) ([]EcsCluster) {
	var clusters []*string
	if cluster == "" {
		clusters = e.ListAllClusterArns()
	} else {
		clusters = append(clusters, aws.String(cluster))
	}
	input := &ecs.DescribeClustersInput{
		Clusters: clusters,
	}
	output, err := e.Client.DescribeClusters(input)


	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException :
				log.Fatal("Server exception occured.")
			case ecs.ErrCodeClientException:
				log.Fatal(`Client exception occurred, please check credential permissions
				 or resource identifiers.`)
			case ecs.ErrCodeInvalidParameterException:
				log.Fatal("Invalid parameter for describe-clusters API request.")
			default:
				log.Fatal(aerr.Error())
			}
		}
	}

	var ecsClusters []EcsCluster
	for _, cluster := range output.Clusters {
		c := EcsCluster{
			ClusterArn: *cluster.ClusterArn, 
			ClusterName: *cluster.ClusterName,
			Status: *cluster.Status,
		}
		ecsClusters = append(ecsClusters, c)
	}
	return ecsClusters

}

//ListAllClusterArns will return a list of all ClusterArns in the region
// Returns: cluster []*string
func (e EcsInt) ListAllClusterArns() ([]*string) {
	var clusters []*string
	err := e.Client.ListClustersPages(&ecs.ListClustersInput{},
		func(page *ecs.ListClustersOutput, lastPage bool) bool {
			for _, cluster := range page.ClusterArns {
				clusters = append(clusters, cluster)
			}
			return page.NextToken != nil
		})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException :
				log.Fatal("Server exception occured.")
			case ecs.ErrCodeClientException:
				log.Fatal(`Client exception occurred, please check credential permissions
					or resource identifiers.`)
			case ecs.ErrCodeInvalidParameterException:
				log.Fatal("Invalid parameter for given API request.")
			default:
				log.Fatal(aerr.Error())
			}
		}
	}
	return clusters
}

//DescribeServices will return a list of services based on a cluster.
//Returns: []EcsService
func (e EcsInt) DescribeServices(cluid string, services []string) ([]EcsService) {
	
	if len(services) < 1 {
		//Convert []*string to []string
		services = aws.StringValueSlice(e.ListAllServicesArns(cluid))
	}
	var ecsServices []EcsService
	//ecs.DescribeServicesInput will only take services in batches of 10.
	batch := 10
	for i := 0; i < len(services); i += batch {
		j := i + batch
		if j > len(services) {
			j = len(services)
		}
		input := &ecs.DescribeServicesInput{
			Cluster: aws.String(cluid),
			Services: aws.StringSlice(services[i:j]),
		}

		output, err := e.Client.DescribeServices(input)


		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ecs.ErrCodeServerException :
					log.Fatal("Server exception occured.")
				case ecs.ErrCodeClientException:
					log.Fatal(`Client exception occurred, please check credential permissions
					 or resource identifiers.`)
				case ecs.ErrCodeInvalidParameterException:
					log.Fatal("Invalid parameter for describe-services API request.")
				default:
					log.Fatal(aerr.Error())
				}
			}
		}

		for _, service := range output.Services {
			s := EcsService{
				ServiceName: *service.ServiceName, 
				LaunchType: *service.LaunchType,
				RunningCount: *service.RunningCount,
				DesiredCount: *service.DesiredCount,
				PendingCount: *service.PendingCount,
			}
			ecsServices = append(ecsServices, s)
		}

	}

	return ecsServices
}


//ListAllServicesArns will return a list of all services in the cluster.
// Returns: services []*string
func (e EcsInt) ListAllServicesArns(cluid string) ([]*string) {
	var services []*string
	err := e.Client.ListServicesPages(&ecs.ListServicesInput{
		Cluster: aws.String(cluid),
	},
		func(page *ecs.ListServicesOutput, lastPage bool) bool {
			for _, service := range page.ServiceArns {
				services = append(services, service)
			}
			return page.NextToken != nil
		})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException :
				log.Fatal("Server exception occured.")
			case ecs.ErrCodeClientException:
				log.Fatal(`Client exception occurred, please check credential permissions
					or resource identifiers.`)
			case ecs.ErrCodeInvalidParameterException:
				log.Fatal("Invalid parameter for given API request.")
			default:
				log.Fatal(aerr.Error())
			}
		}
	}
	return services
}