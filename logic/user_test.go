package logic

import (
	"context"
	"local/test-container-demo/domain"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGenerateValues(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	testCases := []struct {
		name  string
		users domain.Users
	}{
		{
			name: "Test 1",
		},
		{
			name: "Test 1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			addr, _ := redisC.ContainerIP(ctx)
			users, tracker := GenerateValues(addr)
			assert.NotEmptyf(t, users, "users can not be empty")
			assert.NotEmptyf(t, tracker, "tracker can not be empty")
		})
	}
}
