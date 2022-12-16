package utils

import (
	"hole/pkgs/common/utils"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	t.Log(utils.CheckPassword("123456"))
	t.Log(utils.CheckPassword("123456abc"))
	t.Log(utils.CheckPassword("123456abcABC&"))
	t.Log(utils.CheckPassword("123456abcABC"))
}
