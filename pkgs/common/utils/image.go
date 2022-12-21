package utils

import (
	"errors"
	"fmt"
	"hole/pkgs/config/base"
	"strings"
)

func FileNameToContentType(filename string) (string, error) {
	suffix := filename[strings.LastIndex(filename, ".")+1:]

	switch suffix {
	case "png":
		return "image/png", nil
	case "jpg":
		return "image/jpg", nil
	case "jpeg":
		return "image/jpeg", nil
	default:
		return "", errors.New("不支持该文件格式")
	}
}

func ImageIdToUrl(id string, bucket string) string {
	return fmt.Sprintf("%s:%d/image/%s/%s", base.GetDomain(), base.GetPort(), bucket, id)
}
