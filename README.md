# ecs-helper
A command line tool to help navigate ECS clusters, troubleshooting and maintenance.

AWS Credentials are loaded from (and in order from first to last):

* Environment Variables
* Shared Credentials file (~/.aws/config)
* Shared Configuration file (~/.aws/credentials)
* EC2 Instance Metadata



## Example usage
* * *

```
$ ecs-helper describe-clusters -i prod-cluster
CLUSTER ARN                                             CLUSTER NAME    STATUS 
arn:aws:ecs:eu-west-1:360577070949:cluster/prod-cluster prod-cluster    ACTIVE
```



## Installation
* * *

Binary downloads can be found on the Releases page.

You can also clone this repo and compile the binary.