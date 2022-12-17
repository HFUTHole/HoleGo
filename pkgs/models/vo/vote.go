package vo

import (
	"hole/pkgs/models"
)

type VotingOptionVO struct {
	ID    int64  `json:"id"`
	Cid   int64  `json:"cid"`
	Text  string `json:"text"`
	Total int64  `json:"total" `
}

func ConvertVotingOption(options []models.VotingOption) []VotingOptionVO {
	if options == nil && len(options) <= 0 {
		return []VotingOptionVO{}
	}

	optionsVO := make([]VotingOptionVO, len(options))

	for i, option := range options {
		optionsVO[i] = VotingOptionVO{
			ID:    option.ID,
			Cid:   option.Cid,
			Text:  option.Text,
			Total: option.Total,
		}
	}

	return optionsVO
}

type VotingInfoVO struct {
	ID  int64 `json:"id"`
	Vid int64 `json:"vid"`
	Uid int64 `json:"uid"`
}
