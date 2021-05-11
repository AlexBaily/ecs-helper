package cmd

import (
	"os"

	"github.com/AlexBaily/ecs-helper/ecs"
	"github.com/olekukonko/tablewriter"

)

//describeClustersCmd takes a cluster id and renders a table of results.
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
	table.SetHeader([]string{"ClusterArn", "ClusterName", "Status"})
	clusters := ecs.EcsClient.DescribeEcsClusters(cluid)
	var data [][]string
	for _, cluster := range clusters {
		data = append(data, []string{cluster.ClusterArn, 
			cluster.ClusterName, cluster.Status})
	}
	table.AppendBulk(data)
	table.Render()
}