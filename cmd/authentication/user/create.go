package user

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	createFullName string
	createEmail    string
	createPassword string
)

func init() {
	createCmd.Flags().StringVar(&createFullName, "full-name", "", "Full name of a new user")
	createCmd.Flags().StringVar(&createEmail, "email", "", "Email of a new user")
	createCmd.Flags().StringVar(&createPassword, "password", "", "Password of a new user")
	createCmd.MarkFlagRequired("full-name")
	createCmd.MarkFlagRequired("email")
	createCmd.MarkFlagRequired("password")
	createCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Aliases: []string{
		"new",
	},
	Short:   "Create a new user",
	PreRunE: config.SetupCobraLogger,
	RunE:    createUser,
}

func createUser(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewUserServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.CreateUserRequest{
		FullName: createFullName,
		Email:    createEmail,
		Password: createPassword,
	})
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to crate user: %w", err)
	}
	zap.L().Info("Created user", zap.Any("user", resp.Msg.User))
	return nil
}
