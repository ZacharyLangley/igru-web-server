package gardens_test

import (
	"context"
	"os"
	"testing"
	"time"
)

func initContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}

func init() {

}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}
