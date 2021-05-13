package cmd

import (
	"os"
	"strconv"

	"ecs-helper/ecs"
	"github.com/olekukonko/tablewriter"

)

//describeClustersCmd takes a cluster id and renders a table of results.
//Returns: void
func describeClustersCmd(cluid string) {

	table := standardTable()
	table.SetHeader([]string{"Cluster Arn", "Cluster Name", "Status"})
	clusters := ecs.EcsClient.DescribeEcsClusters(cluid)
	var data [][]string //Used to contain the data for the table.
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

	table := standardTable()
	table.SetHeader([]string{"Service Name", "Launch Type", "Running Count",
		"Pending Count", "Desired Count"})


	clusters := ecs.EcsClient.DescribeEcsClusters(cluid)
	for _, cluster := range clusters {
		var data [][]string //Used to contain the data for the table. 
		services := ecs.EcsClient.DescribeServices(cluster.ClusterArn, service)
		for _, service := range services {
			data = append(data, []string{
				service.ServiceName,
				service.LaunchType,
				strconv.FormatInt(service.RunningCount, 32),
				strconv.FormatInt(service.PendingCount, 32),
				strconv.FormatInt(service.DesiredCount, 32),
			})
		}
		table.AppendBulk(data)
		table.Render()
		table.ClearRows() //Clear rows ready to add data for the next cluster.
	}


}


//standardTable generates a new table with the set parameters we want.
//Returns: *tablewriter.Table
func standardTable() (*tablewriter.Table) {
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
	return table
}