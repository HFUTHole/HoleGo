package dao

import (
	"testing"
)

func TestCreateVote(t *testing.T) {
	db := GetDB()
	err := CreateVote(db, 1603690087504154624, []string{"option 1", "option 2"})

	if err != nil {
		t.Error(err)
	}
}

func TestGetContentVoting(t *testing.T) {
	db := GetDB()

	voting, err := GetContentVoting(db, 1603690087504154624)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(voting)
}

func TestVoteContent(t *testing.T) {
	db := GetDB()
	err := VoteContent(db, 4, 1603642475187015680, 3)
	if err != nil {
		t.Fatal(err)
	}
}
