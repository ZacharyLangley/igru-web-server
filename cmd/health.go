package cmd

import (
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use: "health",
	Run: func(cmd *cobra.Command, args []string) {
		cli := http.Client{
			Timeout: time.Second * 5,
		}
		resp, err := cli.Get("http://localhost/health")
		if err != nil {
			os.Exit(2)
		}
		if resp.StatusCode >= 300 || resp.StatusCode < 200 {
			os.Exit(3)
		}
	},
}
