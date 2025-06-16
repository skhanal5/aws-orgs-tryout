package client

import (
    "context"
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/sts"
    "github.com/aws/aws-sdk-go-v2/service/sts/types"
)

var (
	roleFormatSpec = "arn:aws:iam::%s:role/%s"
)

func readPolicy(policyPath string) (string, error) {
    if policyPath == "" {
        return "", nil
    }
    data, err := os.ReadFile(policyPath)
    if err != nil {
        return "", fmt.Errorf("failed to read policy file %q: %w", policyPath, err)
    }
    return string(data), nil
}

func loadDefaultAWSConfig(ctx context.Context) (aws.Config, error) {
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        return aws.Config{}, fmt.Errorf("failed to get default AWS config: %w", err)
    }
    return cfg, nil
}
func assumeRole(ctx context.Context, cfg aws.Config, roleArn, policyStr string) (*types.Credentials, error) {
    stsClient := sts.NewFromConfig(cfg)
    input := &sts.AssumeRoleInput{
        RoleArn:         aws.String(roleArn),
        RoleSessionName: aws.String("sts-session"),
        DurationSeconds: aws.Int32(600),
        Policy:          aws.String(policyStr),
    }

    resp, err := stsClient.AssumeRole(ctx, input)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve credentials: %w", err)
    }
    return resp.Credentials, nil
}

func buildAssumedRoleConfig(ctx context.Context, creds *types.Credentials) (aws.Config, error) {
    provider := credentials.NewStaticCredentialsProvider(
        aws.ToString(creds.AccessKeyId),
        aws.ToString(creds.SecretAccessKey),
        aws.ToString(creds.SessionToken),
    )
    newCfg, err := config.LoadDefaultConfig(
        ctx,
        config.WithDefaultRegion("us-east-1"),
        config.WithCredentialsProvider(provider),
    )
    if err != nil {
        return aws.Config{}, fmt.Errorf("failed to load config with assumed credentials: %w", err)
    }
    return newCfg, nil
}


func GetClient[T any](
    ctx context.Context,
    accountId string, 
    policyPath string,
    constructor func(cfg *aws.Config) T,
) (T, error) {
    var empty T

    cfg, err := loadDefaultAWSConfig(ctx)
    if err != nil {
        return empty, err
    }

    policyStr, err := readPolicy(policyPath)
    if err != nil {
        return empty, err
    }

	roleArn := fmt.Sprintf(roleFormatSpec, accountId, "OrganizationAccountAccessRole")

    creds, err := assumeRole(ctx, cfg, roleArn, policyStr)
    if err != nil {
        return empty, err
    }

    newCfg, err := buildAssumedRoleConfig(ctx, creds)
    if err != nil {
        return empty, err
    }

    return constructor(&newCfg), nil
}