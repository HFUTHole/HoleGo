package vo

import (
	"hole/pkgs/common/utils"
	"hole/pkgs/models"
)

type ContentVO struct {
	ID      int64                       `json:"id"`
	Uid     int64                       `json:"uid"`
	Nick    string                      `json:"nick"`
	Avatar  string                      `json:"avatar"`
	Like    int64                       `json:"like"`
	Real    bool                        `json:"real"`
	Title   string                      `json:"title"`
	Text    string                      `json:"text"`
	JumpUrl map[string]ContentJumpUrlVO `json:"jumpUrl"`
	Tags    []TagVO                     `json:"tags"`
	Images  []ContentImageVO            `json:"images"`
	Voting  []VotingOptionVO            `json:"voting,omitempty"`
}

type ContentImageVO struct {
	ID  int64  `json:"id"`
	Cid int64  `json:"cid"`
	URL string `json:"url"`
}

type ContentJumpUrlVO struct {
	ID      int64  `json:"id"`
	Cid     int64  `json:"cid"`
	Text    string `json:"text"`
	JumpCid int64  `json:"jumpCid"`
}

type TagVO struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

type ContentPage struct {
	MaxId int64       `json:"maxId"`
	List  []ContentVO `json:"list"`
}

func ConvertConvertContentVO(c *models.Content, tags []models.Tag, jumps []models.ContentJumpUrl, imgs []models.ContentImage) *ContentVO {
	if c == nil {
		return nil
	}

	contentVO := &ContentVO{
		ID:      c.ID,
		Uid:     c.Uid,
		Nick:    c.Nick,
		Avatar:  c.Avatar,
		Like:    c.Like,
		Real:    utils.IntToBool(c.Real),
		Title:   c.Title,
		Text:    c.Text,
		Tags:    []TagVO{},
		JumpUrl: map[string]ContentJumpUrlVO{},
		Images:  []ContentImageVO{},
	}

	if tags != nil && len(tags) > 0 {
		tagsVO := make([]TagVO, len(tags))
		for i, tag := range tags {
			tagsVO[i] = TagVO{
				ID:   tag.ID,
				Text: tag.Text,
			}
		}
		contentVO.Tags = tagsVO
	}

	if jumps != nil && len(jumps) > 0 {
		jumpsVO := make(map[string]ContentJumpUrlVO, len(jumps))
		for _, jump := range jumps {
			jumpsVO[jump.Text] = ContentJumpUrlVO{
				ID:      jump.ID,
				Cid:     jump.Cid,
				Text:    jump.Text,
				JumpCid: jump.JumpCid,
			}
		}
		contentVO.JumpUrl = jumpsVO
	}

	if imgs != nil && len(imgs) > 0 {
		images := make([]ContentImageVO, len(tags))
		for i, img := range imgs {
			images[i] = ContentImageVO{
				ID:  img.ID,
				Cid: img.Cid,
				URL: img.URL,
			}
		}
		contentVO.Images = images
	}
	return contentVO
}

func ConvertContentVOList(contents []models.Content) []ContentVO {
	contentVO := make([]ContentVO, len(contents))

	for i, content := range contents {
		contentVO[i] = *ConvertConvertContentVO(&content, nil, nil, nil)
	}
	return contentVO
}
