package rbac

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/auth/entities"
)

func GetUserRole(user uint) (*entities.Role, []string, []string) {
	var roleId []uint
	database.DB.Model(&entities.UserRole{}).Where("user_id=?", user).Pluck("role_id", &roleId)
	if len(roleId) < 1 {
		return nil, []string{}, []string{}
	}
	var roleItems []*entities.Role
	database.DB.Where("id in ?", roleId).Order("id asc").Find(&roleItems)
	roles := make([]string, len(roleItems))
	isAdministrator := false
	for i := len(roleItems) - 1; i >= 0; i-- {
		roles[i] = roleItems[i].Name
		if roles[i] == "administrator" {
			isAdministrator = true
		}
	}
	var permissions []string
	if isAdministrator {
		database.DB.Model(&entities.Permission{}).Pluck("name", &permissions)
		return roleItems[0], roles, permissions
	}
	var permissionId []uint
	database.DB.Model(&entities.RolePermission{}).Where("role_id in ?", roleId).Pluck("permission_id", &permissionId)
	if len(permissionId) < 1 {
		return roleItems[0], roles, []string{}
	}
	database.DB.Model(&entities.Permission{}).Where("id in ?", permissionId).Pluck("name", &permissions)
	return roleItems[0], roles, permissions
}
