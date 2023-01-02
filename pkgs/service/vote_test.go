package service

import (
	"hole/pkgs/config"
	"testing"
	"time"
)

func TestCreateContentVoting(t *testing.T) {
	config.InitConfigFileWithTest()
	//logger.InitLogger()

	voting, err := CreateContentVoting(1, 1603642730045509632, []string{"option 1", "option 1"}, time.Now())
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(voting)
}
