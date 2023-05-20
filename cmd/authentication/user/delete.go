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
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	deleteUserID string
)

func init() {
	deleteCmd.Flags().StringVar(&deleteUserID, "id", "", "ID of an existing user")
	config.Must(deleteCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing user",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteUser,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		for _, arg := range args {
			_, err := uuid.Parse(arg)
			if err != nil {
				return authentication.InvalidUserIDError{
					UserID: arg,
				}
			}
		}
		return nil
	},
}

func deleteUser(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewUserServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.DeleteUserRequest{
		UserId: deleteUserID,
	})
	_, err := userClient.DeleteUser(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
