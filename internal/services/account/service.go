package member

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/account"
)


type AccountService struct {
	client account.Client
}

func (a *AccountService) GetAccountName(ctx context.Context, accountId string) (string, error) {
	accountInfoInput := account.GetAccountInformationInput{
		AccountId: &accountId,
	}

	output, err := a.client.GetAccountInformation(ctx, &accountInfoInput)
	if err != nil {
		return "", fmt.Errorf("Failed to get account information: %s", err)
	}
	return *output.AccountName, nil
}

func (a *AccountService) UpdateAccountName(ctx context.Context, accountName string) error {
	accountInfoInput := account.PutAccountNameInput{
		AccountName: &accountName,
	}

	_, err := a.client.PutAccountName(ctx, &accountInfoInput)
	if err != nil {
		return fmt.Errorf("Failed to update account: %s", err)
	}
	return nil
}