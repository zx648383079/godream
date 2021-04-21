package models

import (
	auth_models "zodream.cn/godream/modules/auth/models"
)

type Profile struct {
	Name      string                  `json:"name"`
	User      *auth_models.UserSimple `json:"user"`
	Signature string                  `json:"signature"`
	NewCount  uint                    `json:"new_count"`
}
