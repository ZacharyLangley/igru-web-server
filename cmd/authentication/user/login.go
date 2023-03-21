package user

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var getTokenEmail string
var getTokenPassword string

func init() {
	loginCmd.Flags().StringVar(&getTokenEmail, "email", "", "Email of a new user")
	loginCmd.Flags().StringVar(&getTokenPassword, "password", "", "Password of a new user")
	config.Must(loginCmd.MarkFlagRequired("email"))
	config.Must(loginCmd.MarkFlagRequired("password"))
}

var loginCmd = &cobra.Command{
	Use: "login",
	Aliases: []string{
		"create-token",
	},
	Short:   "Get Session token for user",
	PreRunE: config.SetupCobraLogger,
	RunE:    login,
}

func login(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	clientConfig, err := config.NewClient()
	if err != nil {
		return fmt.Errorf("failed to open new client config: %w", err)
	}
	userClient := authenticationv1connect.NewSessionServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.CreateSessionRequest{
		Email:    getTokenEmail,
		Password: getTokenPassword,
	})
	resp, err := userClient.CreateSession(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	token, err := authentication.ExtractSessionToken(resp.Header())
	if err != nil {
		zap.L().Info("Failed to find auth header", zap.Error(err))
		for key, values := range resp.Header() {
			zap.L().Info("Found header", zap.String("key", key), zap.Strings("values", values))
		}
	}
	if err := clientConfig.Set(token); err != nil {
		zap.L().Error("failed to set token to config", zap.Error(err))
	} else {
		fmt.Print(token)
	}
	return nil
}
