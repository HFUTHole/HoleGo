package dao

import (
	"gorm.io/gorm"
	"hole/pkgs/models"
	"time"
)

func CreateContent(tx *gorm.DB, content *models.Content) error {
	err := tx.Model(content).Create(content).Error
	return err
}

func SetContentVoteEndTime(tx *gorm.DB, cid int64, enable int, endTime time.Time) error {
	res := map[string]interface{}{"enable_voting": enable, "end_time": endTime}
	err := tx.Model(&models.Content{}).Where("id = ?", cid).Updates(res).Error
	return err
}

func GetContent(tx *gorm.DB, cid int64) (models.Content, error) {
	var content models.Content
	err := tx.Model(&models.Content{}).Where("id = ?", cid).First(&content).Error
	return content, err
}

func GetContentFormID(tx *gorm.DB, id int64) (*models.Content, error) {
	var content models.Content
	err := tx.Model(&content).Where("id = ?", id).First(&content).Error
	return &content, err
}

func GetContentNextPage(tx *gorm.DB, maxID int64, limit int) ([]models.Content, error) {
	var list []models.Content
	err := tx.Model(&models.Content{}).Where("id < ?", maxID).Order("id DESC").Limit(limit).Find(&list).Error
	return list, err
}

func GetContentPage(tx *gorm.DB, limit int) ([]models.Content, error) {
	var list []models.Content
	err := tx.Model(&models.Content{}).Order("id DESC").Limit(limit).Find(&list).Error
	return list, err
}

func GetContentOffset10ByUid(tx *gorm.DB, uid int64) (*models.Content, error) {
	var content models.Content
	err := tx.Model(&models.Content{}).Where("uid = ?", uid).Offset(9).Limit(1).First(&content).Error
	return &content, err
}

func GetContentOneDay(tx *gorm.DB, t time.Time) ([]models.Content, error) {
	year, month, day := t.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	var content []models.Content
	err := tx.Model(&models.Content{}).Where("created_at > ?", date).Find(&content).Error
	return content, err
}
func GetContentOneDayCount(tx *gorm.DB, t time.Time) (int64, error) {
	year, month, day := t.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	var count int64
	err := tx.Model(&models.Content{}).Where("created_at > ?", date).Count(&count).Error
	return count, err
}

func CreateContentJumpUrl(tx *gorm.DB, c *models.ContentJumpUrl) error {
	err := tx.Model(&models.ContentJumpUrl{}).Create(c).Error
	return err
}

func CreateContentJumpUrls(tx *gorm.DB, list []models.ContentJumpUrl) error {
	err := tx.Model(&models.ContentJumpUrl{}).CreateInBatches(&list, len(list)).Error
	return err
}

func GetContentJumpUrls(tx *gorm.DB, cid int64) ([]models.ContentJumpUrl, error) {
	var list []models.ContentJumpUrl
	err := tx.Model(&models.ContentJumpUrl{}).Where("cid = ?", cid).Find(&list).Error
	return list, err
}

func CreateTag(tx *gorm.DB, tag *models.Tag) error {
	return tx.Model(&models.Tag{}).Create(tag).Error
}

func GetTag(tx *gorm.DB, id int64) (*models.Tag, error) {
	var tag models.Tag
	err := tx.Model(&models.Tag{}).Where("id = ?", id).First(&tag).Error
	return &tag, err
}

func GetTagByText(tx *gorm.DB, text string) (*models.Tag, error) {
	var tag models.Tag
	err := tx.Model(&models.Tag{}).Where("text = ?", text).First(&tag).Error
	return &tag, err
}

func LinkContentTag(tx *gorm.DB, table *models.ContentTag) error {
	return tx.Model(&models.ContentTag{}).Create(table).Error
}

func GetContentTags(tx *gorm.DB, cid int64) ([]models.Tag, error) {
	var tags []models.Tag
	err := tx.Model(&models.Tag{}).Where(
		"id in (?)", tx.Model(&models.ContentTag{}).Select("tid").Where("cid = ?", cid),
	).Find(&tags).Error
	return tags, err
}

func CreateAndLinkTags(tx *gorm.DB, cid int64, tags []string) error {
	for _, tag := range tags {
		_ = CreateTag(tx, &models.Tag{Text: tag})
	}
	err := tx.Exec("insert into content_tags(cid, tid)  select ?, id from tags where text in ?", cid, tags).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateContentImages(tx *gorm.DB, cid int64, imgs []string) error {
	images := make([]models.ContentImage, len(imgs))
	for i, img := range imgs {
		images[i] = models.ContentImage{
			Cid: cid,
			URL: img,
		}
	}
	return tx.Model(&models.ContentImage{}).CreateInBatches(&images, len(images)).Error
}

func GetContentImages(tx *gorm.DB, cid int64) ([]models.ContentImage, error) {
	var images []models.ContentImage
	err := tx.Model(&models.ContentImage{}).Where("cid = ?", cid).Find(&images).Error
	return images, err
}
