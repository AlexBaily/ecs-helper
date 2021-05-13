package cmd

import (
	"os"

	"ecs-helper/ecs"
	"github.com/olekukonko/tablewriter"

)

//describeClustersCmd takes a cluster id and renders a table of results.
//Returns: void
func describeClustersCmd(cluid string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.SetHeader([]string{"Cluster Arn", "Cluster Name", "Status"})
	clusters := ecs.EcsClient.DescribeEcsClusters(cluid)
	var data [][]string
	for _, cluster := range clusters {
		data = append(data, []string{cluster.ClusterArn, 
			cluster.ClusterName, cluster.Status})
	}
	table.AppendBulk(data)
	table.Render()
}


// describeServicesCmd takes a cluster id or a service naame and outputs information.
// A cluster ID and service is required to perform this operation.
// If not cluster id or service is supplied then all services on all clusters are returned.
// Returns: void
func describeServicesCmd(cluid string, service []string) {
	clusters := ecs.EcsClient.DescribeEcsClusters(cluid)
	for _, cluster := range clusters {
		ecs.EcsClient.DescribeServices(cluster.ClusterArn, service)
	}


}