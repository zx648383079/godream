package dao

import (
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"zodream.cn/godream/database"
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/modules/seo/emoji"
	"zodream.cn/godream/utils"
	"zodream.cn/godream/utils/rule"
)

const (
	MESSAGE_TEXT  = 0
	MESSAGE_IMAGE = 1
	MESSAGE_VIDEO = 2
	MESSAGE_VOICE = 3
	MESSAGE_FILE  = 4

	RULE_IMAGE_TAG = "[图片]"
	RULE_VIDEO_TAG = "[视频]"
	RULE_VOICE_TAG = "[语音]"
)

func GetMessageList(user uint, startTime uint, itemType uint, itemId uint) []*models.MessageItem {
	var data []*entities.Message
	var query *gorm.DB
	if itemType > 0 {
		query = database.DB.Where("group_id=?", itemId)
	} else {
		query = database.DB.Where("group_id=0 AND ((user_id=? AND receive_id=?) OR (user_id=? AND receive_id=?))", user, itemId, itemId, user)
	}
	if startTime < 1 {
		query.Order("created_at desc").Limit(10).Find(&data)
	} else {
		query.Where("created_at>=?", startTime).Order("created_at asc").Find(&data)
	}
	l := len(data)
	items := make([]*models.MessageItem, l)
	if l < 1 {
		return items
	}
	userId := []uint{user}
	for _, v := range data {
		exsit := false
		newId := v.UserId
		if newId == user {
			newId = v.ReceiveId
		}
		for _, u := range userId {
			if u == newId {
				exsit = true
				break
			}
		}
		if !exsit {
			userId = append(userId, newId)
		}
	}
	var userItems []*auth_models.UserSimple
	database.DB.Find(&userItems, userId)
	for i := 0; i < l; i++ {
		v := data[i]
		item := models.MessageItem{
			Message: *v,
		}
		for _, u := range userItems {
			if u.ID == v.UserId {
				item.User = u
				break
			}
		}
		j := i
		if startTime < 1 {
			j = l - i - 1
		}
		items[j] = &item
	}
	return items
}

func GetPing(user int, startTime uint, itemType uint, itemId uint) {

}

func SendText(user uint, itemType uint, itemId uint, content string) (*models.MessageItem, error) {
	rules := emoji.RenderRule(content)
	if itemType > 0 {
		rules = append(rules, renderAt(user, itemId, content)...)
	}
	return Send(user, itemType, itemId, MESSAGE_TEXT, content, rules)
}

func SendImage(user uint, itemType uint, itemId uint, image string) (*models.MessageItem, error) {
	return Send(user, itemType, itemId, MESSAGE_IMAGE, RULE_IMAGE_TAG, []rule.RuleItem{rule.FormatImage(RULE_IMAGE_TAG, image)})
}

func SendVideo(user uint, itemType uint, itemId uint, file string) (*models.MessageItem, error) {
	return Send(user, itemType, itemId, MESSAGE_VIDEO, RULE_VIDEO_TAG, []rule.RuleItem{rule.FormatFile(RULE_VIDEO_TAG, file)})
}

func SendFile(user uint, itemType uint, itemId uint, fileName string, file string) (*models.MessageItem, error) {
	tag := fmt.Sprintf("[%s]", fileName)
	return Send(user, itemType, itemId, MESSAGE_FILE, tag, []rule.RuleItem{rule.FormatFile(tag, file)})
}

func SendVoice(user uint, itemType uint, itemId uint, file string) (*models.MessageItem, error) {
	return Send(user, itemType, itemId, MESSAGE_VOICE, RULE_VOICE_TAG, []rule.RuleItem{rule.FormatFile(RULE_VOICE_TAG, file)})
}

func Send(user uint, itemType uint, itemId uint, messageType uint32, content string, extraRule []rule.RuleItem) (*models.MessageItem, error) {
	if itemType < 1 && user == itemId {
		return nil, errors.New("不能自己发送给自己")
	}
	ruleByte, _ := json.Marshal(extraRule)
	now := uint(utils.Now())
	model := entities.Message{
		Type:      messageType,
		Content:   content,
		UserId:    user,
		ExtraRule: string(ruleByte),
		UpdatedAt: now,
		CreatedAt: now,
	}
	if itemType < 1 {
		model.ReceiveId = itemId
	} else {
		model.GroupId = itemId
	}
	smt := database.DB.Create(&model)
	if model.ID < 1 {
		return nil, smt.Error
	}
	addHistory(user, itemType, itemId, model.ID, 0)
	if itemType < 1 {
		addHistory(itemId, itemType, user, model.ID, 1)
	}
	var userModel auth_models.UserSimple
	database.DB.First(&user, user)
	return &models.MessageItem{
		Message: model,
		User:    &userModel,
	}, nil
}
