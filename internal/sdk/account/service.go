package account

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/skhanal5/aws-orgs-tryout/internal/sdk/client"
)

func GetAccountClient(ctx context.Context, accountId string) (*account.Client, error) {
	constructor := func (config *aws.Config) (*account.Client) {
		return account.NewFromConfig(*config)
	}
	return client.GetClient(ctx, accountId, "iam/account.json", constructor)
}


func GetAccountName(ctx context.Context, accountId string) (string, error) {
	client,err  := GetAccountClient(ctx, accountId)
	
	if err != nil {
		return "", fmt.Errorf("Failed to build account client: %s", err)
	}

	accountInfoInput := account.GetAccountInformationInput{
		AccountId: &accountId,
	}

	output, err := client.GetAccountInformation(ctx, &accountInfoInput)
	if err != nil {
		return "", fmt.Errorf("Failed to get account information: %s", err)
	}
	return *output.AccountName, nil
}

func UpdateAccountName(ctx context.Context, accountId string, accountName string) error {
	client, err  := GetAccountClient(ctx, accountId)
	
	if err != nil {
		return fmt.Errorf("Failed to build account client: %s", err)
	}

	
	accountInfoInput := account.PutAccountNameInput{
		AccountName: &accountName,
	}

	_, err = client.PutAccountName(ctx, &accountInfoInput)
	if err != nil {
		return fmt.Errorf("Failed to update account: %s", err)
	}
	return nil
}

func PutContactInformation(ctx context.Context, accountId string) error {
	client, err  := GetAccountClient(ctx, accountId)
	
	if err != nil {
		return fmt.Errorf("Failed to build account client: %s", err)
	}

	contactInfo := &types.ContactInformation{
		AddressLine1: aws.String("123 Main St"),
		City:         aws.String("Seattle"),
		CountryCode:  aws.String("US"),
		FullName:     aws.String("John Doe"),
		PhoneNumber:  aws.String("+1-555-555-5555"),
		PostalCode:   aws.String("98101"),
		StateOrRegion: aws.String("WA"),
	}

	accountInfoInput := account.PutContactInformationInput{
		AccountId: &accountId,
		ContactInformation: contactInfo,
	}

	_, err = client.PutContactInformation(ctx, &accountInfoInput)
	if err != nil {
		return fmt.Errorf("Failed to update account: %s", err)
	}
	return nil
}