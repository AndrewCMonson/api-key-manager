package awsconfig

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	awsCfg	  aws.Config
	cfgErr 		error
	once      sync.Once
)

func GetAWSConfig(region string) (aws.Config, error){
	once.Do(func() {
		awsCfg, cfgErr = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	})
	
	return awsCfg, cfgErr
}
