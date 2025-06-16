package main

import (
	"context"
	"fmt"

	"github.com/skhanal5/aws-orgs-tryout/internal/sdk/account"
)

func main() {
	ctx := context.Background()
	err := account.PutContactInformation(ctx, "804373361712")
	if err != nil {
		fmt.Printf("\nFailed to get account name: \n%s", err)
	}
}