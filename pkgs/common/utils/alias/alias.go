package alias

import "errors"

var alias = map[string]int64{
	"nick1": -1,
	"nick2": -2,
}

var avatar = map[int64]string{
	-1: "1605473820834009088",
	-2: "1605473820834009088",
}

func GetID(nick string) (int64, error) {

	id := alias[nick]
	if id == 0 {
		return id, errors.New("nick not exists")
	}
	return id, nil
}

func GetAvatarByNick(nick string) (string, error) {
	id, err := GetID(nick)
	if err != nil {
		return "", err
	}
	return GetAvatarByID(id)
}

func GetAvatarByID(id int64) (string, error) {
	ava := avatar[id]
	if ava == "" {
		return "", errors.New("avatar not exists")
	}
	return ava, nil
}
