package cmd

import (
	"fmt"

	"github.com/AlexBaily/ecs-helper/ecs"

)


func describeClustersCmd(cluid string) {
	output := ecs.EcsClient.DescribeEcsClusters(cluid)
	fmt.Println(output)
}