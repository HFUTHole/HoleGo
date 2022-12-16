package service

import "hole/pkgs/common/utils"

func GenerateToken(uid int64, sub string) (string, error) {
	var id = utils.NextSnowflake()
	token, err := utils.GenerateToken(uid, sub, id)
	if err != nil {
		return "", err
	}

	return token, nil
}
