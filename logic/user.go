package logic

import (
	"context"
	"fmt"
	"local/test-container-demo/config/cache"
	"local/test-container-demo/domain"
	"strconv"
)

func GenerateValues(addr string) (domain.Users, []string) {
	ctx := context.Background()
	rCache := cache.SetUp(addr)

	var tracker []string
	users := domain.GetUsers()

	if len(users) > 0 {
		for i := 0; i < 3; i++ {
			randomUser := users.GetRandomUser()
			userID := strconv.Itoa(randomUser.ID)
			err := rCache.Set(ctx, userID, randomUser.ToString(), 0).Err()
			if err != nil {
				panic(err)
			}
			tracker = append(tracker, userID)
		}
	}

	for _, index := range tracker {
		val, err := rCache.Get(ctx, index).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("cache value: %s\n", val)
	}

	return users, tracker
}
