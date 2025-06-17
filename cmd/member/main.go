package main

import (
	"context"
	"log"

	"github.com/skhanal5/aws-orgs-tryout/internal/sdk/account"
)

func main() {
	ctx := context.Background()
	name, err := account.GetAccountName(ctx, "804373361712")
	if err != nil {
		log.Fatalf("\nFailed to get account name: \n%s", err)
	}
	log.Printf("The name is: %s", name)
}