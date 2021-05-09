module github.com/AlexBaily/ecs-helper

go 1.15

require (
	github.com/AlexBaily/ecs-helper/cmd v0.0.0-00010101000000-000000000000
	github.com/AlexBaily/ecs-helper/ecs v0.0.0-20210508214335-c450657d3983 // indirect
	github.com/aws/aws-sdk-go v1.38.36 // indirect
	github.com/spf13/cobra v1.1.3
)

replace github.com/AlexBaily/ecs-helper/cmd => ./cmd
