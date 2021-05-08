package ecs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

//ecsInt is the ecs interface struct that holds the ECSAPI, used for mocking.
type ecsInt struct {
	Client ecsiface.ECSAPI
}

//ecsClient holds the reference to the ECS interface we are using in the package.
var ecsClient *ecsInt

//Configure the session with the ECS service client.
func configureECS() {
	ecsClient = new(ecsInt)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
    svc := ecs.New(sess)
	ecsClient.Client = ecsiface.ECSAPI(svc)
}

func init() {
	configureECS()
}