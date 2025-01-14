package awsconfig

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	awsConfig	  aws.Config
	configErr error
	once      sync.Once
)

func GetAWSConfig(region string) (aws.Config, error){
	once.Do(func() {
		awsConfig, configErr = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if configErr != nil {
			configErr = fmt.Errorf("error loading AWS config: %w", configErr)
		}
	})

	if configErr != nil {
		return aws.Config{}, configErr
	}
	return awsConfig, nil
}
