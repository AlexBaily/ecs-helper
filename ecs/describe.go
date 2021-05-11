package ecs

import (
	"fmt"
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
		//Todo Add list-clusters to find all available clusters.
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
				log.Fatal("Invalid parameter for given API request.")
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
	fmt.Println(ecsClusters)
	return ecsClusters

}