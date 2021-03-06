package usecase

import (
	"shalust/api/pkg/infra"
	"shalust/api/pkg/model"
)

func GetAllIllustratio(data *[]model.ContentData) error {

	client, err := infra.Init_mysql()
	if err != nil {
		return err
	}

	client.
		From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", "testid").
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", "testid").
		Where("illustratio = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetAllRough(data *[]model.ContentData) error {

	client, err := infra.Init_mysql()
	if err != nil {
		return err
	}

	client.
		From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", "testid").
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", "testid").
		Where("rough = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}
func GetAllCommic(data *[]model.ContentData) error {

	client, err := infra.Init_mysql()
	if err != nil {
		return err
	}

	client.
		From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", "testid").
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", "testid").
		Where("commic = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}
func GetAllGraffiti(data *[]model.ContentData) error {

	client, err := infra.Init_mysql()
	if err != nil {
		return err
	}

	client.
		From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", "testid").
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", "testid").
		Where("graffiti = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserIllustratio(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("illustratio = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserCommic(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url,  likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("commic = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserGraffiti(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("graffiti = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}
func GetUserRough(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("rough = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserRoughManagement(user_id string, data *[]model.ContentHandling) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Where("user_id = ?", user_id).
		Where("rough = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserGraffitiManagement(user_id string, data *[]model.ContentHandling) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Where("user_id = ?", user_id).
		Where("graffiti = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserIllustratioManagement(user_id string, data *[]model.ContentHandling) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Where("user_id = ?", user_id).
		Where("illustratio = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserCommicManagement(user_id string, data *[]model.ContentHandling) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Where("user_id = ?", user_id).
		Where("commic = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserLikedIllustratio(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("illustratio = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserLikedCommic(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url,  likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("commic = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserLikedGraffiti(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("graffiti = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}
func GetUserLikedRough(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("rough = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserBookmarkedIllustratio(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("illustratio = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserBookmarkedCommic(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url,  likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("commic = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserBookmarkedGraffiti(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("graffiti = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}
func GetUserBookmarkedRough(user_id string, data *[]model.ContentData) error {
	client, err := infra.Init_mysql()
	if err != nil {

		return err
	}
	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", user_id).
		Join("JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", user_id).
		Where("content_handling.user_id = (SELECT user_id FROM user WHERE account_id =?)", user_id).
		Where("rough = 1").
		Where("disclose = 1").
		Scan(data)

	defer client.Close()
	return err
}
