package ecs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

//ecsInt is the ecs interface struct that holds the ECSAPI, used for mocking.
type EcsInt struct {
	Client ecsiface.ECSAPI
}

//EcsClient holds the reference to the ECS interface we are using in the package.
var EcsClient *EcsInt

//Configure the session with the ECS service client.
func configureECS() {
	EcsClient = new(EcsInt)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
    svc := ecs.New(sess)
	EcsClient.Client = ecsiface.ECSAPI(svc)
}

func init() {
	configureECS()
}