package ecs

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
)
//DescribeEcsClusters will take an input of a string and return a list of clusters.
func (e EcsInt) DescribeEcsClusters(cluster string) (*ecs.DescribeClustersOutput) {

	input := &ecs.DescribeClustersInput{}
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
	fmt.Println(output)
	return output

}