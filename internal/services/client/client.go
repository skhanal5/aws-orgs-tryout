package client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func readPolicy(policyPath string) string {
	return ""
}

func GetClient[T any](ctx context.Context, accountId string, roleArn string, policy string, constructor func(cfg aws.Config) T) (T, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	stsClient := sts.NewFromConfig(cfg)

	input := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String("sts session"),
		DurationSeconds: aws.Int32(600),
		Policy:          aws.String(policy),
	}
	response, err := stsClient.AssumeRole(ctx, &input)
	if err != nil {
		var empty T
		return empty, fmt.Errorf("Failed to retrieve credentials")
	}
	creds := response.Credentials
	cache := aws.NewCredentialsCache(
		credentials.NewStaticCredentialsProvider(
			*creds.AccessKeyId,
			*creds.AccessKeyId,
			*creds.SessionToken,
		),
	)
	cfg.Credentials = cache
	return constructor(cfg), nil
}
