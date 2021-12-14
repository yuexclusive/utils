package auth

import (
	"context"

	"github.com/yuexclusive/utils/rpc/client"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	authService "github.com/yuexclusive/utils/srv/auth/proto/auth"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	var authServiceName string

	validToken(token, authServiceName)

	// WARNING: in production define your own type to avoid context collisions
	newCtx := context.WithValue(ctx, "tokenInfo", token)

	return newCtx, nil

}

// validToken validates the authorization.
func validToken(token, authSrviceName string) error {
	closer, client, err := client.Dial(authSrviceName, "")

	if err != nil {
		return err
	}

	defer closer.Close()

	_, err = authService.NewAuthClient(client).Validate(context.Background(), &authService.ValidateRequest{Token: token})

	if err != nil {
		return err
	}

	return nil
}
