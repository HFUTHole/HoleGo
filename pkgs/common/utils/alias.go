package utils

import "errors"

var alias = map[string]int64{
	"a": -1,
	"b": -2,
}

func AliasID(nick string) (int64, error) {

	id := alias[nick]
	if id == 0 {
		return id, errors.New("nick not exists")
	}
	return id, nil
}
