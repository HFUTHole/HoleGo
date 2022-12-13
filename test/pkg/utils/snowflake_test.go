package utils

import (
	"hole/pkgs/common/utils"
	"testing"
	"time"
)

func TestNextSnowflake(t *testing.T) {
	utils.InitSnowflakeNode()

	for i := 0; i < 100; i++ {
		snowflake := utils.NextSnowflake()
		t.Log(snowflake)
		time.Sleep(time.Microsecond * 100)
	}
}
