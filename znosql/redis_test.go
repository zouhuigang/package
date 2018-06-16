package znosql_test

import (
	"testing"

	"github.com/zouhuigang/package/znosql"
)

func TestHSCAN(t *testing.T) {
	redisClient := znosql.NewRedisClient()
	defer redisClient.Close()

	_, _, err := redisClient.HSCAN("store:691goods", 0)
	if err != nil {
		t.Fatal("HSCAN error:", err)
	}
}
