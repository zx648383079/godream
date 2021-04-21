package server

import (
	"encoding/json"
	"strconv"
	"strings"
)

const (
	MESSAGE_PREFIX      = "ws-message:"
	MESSAGE_SEPARATOR   = ";"
	MESSAGE_TYPE_STRING = 0
	MESSAGE_TYPE_INT    = 1
	MESSAGE_TYPE_BOOL   = 2
	MESSAGE_TYPE_JSON   = 4

	EVENT_AUTH               = "auth"
	EVENT_ERROR              = "error"
	EVENT_PROFILE            = "chat/user"
	EVENT_HISTORY            = "chat/chat"
	EVENT_FRIENDS            = "chat/friend"
	EVENT_FRIEND_SEARCH      = "chat/friend/search"
	EVENT_FRIEND_APPLY       = "chat/friend/apply"
	EVENT_GROUPS             = "chat/group"
	EVENT_MESSAGE            = "chat/message"
	EVENT_MESSAGE_PING       = "chat/message/ping"
	EVENT_MESSAGE_SEND       = "chat/message/send"
	EVENT_MESSAGE_SEND_TEXT  = "chat/message/send_text"
	EVENT_MESSAGE_SEND_IMAGE = "chat/message/send_image"
	EVENT_MESSAGE_SEND_VIDEO = "chat/message/send_video"
	EVENT_MESSAGE_SEND_AUDIO = "chat/message/send_audio"
	EVENT_MESSAGE_SEND_FILE  = "chat/message/send_file"
)

func encodeMessage(event string, messageType int, message string) string {
	return MESSAGE_PREFIX + event + MESSAGE_SEPARATOR + strconv.Itoa(messageType) + MESSAGE_SEPARATOR + message
}

func encodeByte(event string, messageType int, message []byte) []byte {
	sep := []byte(MESSAGE_SEPARATOR)
	data := append([]byte(MESSAGE_PREFIX), []byte(event)...)
	data = append(data, sep...)
	data = append(data, []byte(strconv.Itoa(messageType))...)
	data = append(data, sep...)
	return append(data, message...)
}

func encodeJSON(event string, data interface{}) []byte {
	bytes, _ := json.Marshal(data)
	return encodeByte(event, MESSAGE_TYPE_JSON, bytes)
}

func decodeMessage(message string) (string, int, string) {
	var event string
	var messageType int
	i := strings.Index(message, MESSAGE_PREFIX)
	if i < 0 {
		return event, messageType, message
	}
	i += len(MESSAGE_PREFIX)
	message = message[i:]
	i = strings.Index(message, MESSAGE_SEPARATOR)
	event = message[:i]
	i += len(MESSAGE_SEPARATOR)
	message = message[i:]
	i = strings.Index(message, MESSAGE_SEPARATOR)
	messageType, _ = strconv.Atoi(message[:i])
	i += len(MESSAGE_SEPARATOR)
	return event, messageType, message[i:]
}

func decodeJSON(content string, v interface{}) error {
	return json.Unmarshal([]byte(content), v)
}
