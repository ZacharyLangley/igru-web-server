package session

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	deleteSessionID string
)

func init() {
	config.Must(deleteCmd.MarkFlagRequired("config"))
	deleteCmd.Flags().StringVar(&deleteSessionID, "id", "", "ID of an existing session")
	config.Must(deleteCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing session",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteSession,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		for _, arg := range args {
			_, err := uuid.Parse(arg)
			if err != nil {
				return fmt.Errorf("invalid user ID specified: %s", arg)
			}
		}
		return nil
	},
}

func deleteSession(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	cli := authenticationv1connect.NewSessionServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.DeleteSessionRequest{
		Id: deleteSessionID,
	})
	_, err := cli.DeleteSession(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}
