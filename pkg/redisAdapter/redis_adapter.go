package redisAdapter

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

type testdata struct {
	TestString string `json: testString`
	TestInt    int    `json: testInt`
}

func ConnectAndStore() {
	keyAsBytes, _ := json.Marshal(testdata{
		TestInt:    1,
		TestString: "testString",
	})
	fmt.Println("json: ", string(keyAsBytes))
	redisSetKeyErr := redisClient.Set(ctx, "test key", keyAsBytes, 0).Err()
	if redisSetKeyErr != nil {
		fmt.Println("error while setting test key: ", redisSetKeyErr)
	} else {
		fmt.Println("successfully set test key")
	}
}

func GetAllKeys() map[string]string {
	keyValueMap := make(map[string]string)
	keys, _ := redisClient.Keys(ctx, "*").Result()
	for _, key := range keys {
		fmt.Println("key: ", key)
		keyValueMap[key] = redisClient.Get(ctx, key).Val()
	}
	return keyValueMap
}
