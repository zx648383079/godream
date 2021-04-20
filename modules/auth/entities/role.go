package entities

// Permission Model
type Permission struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	UpdatedAt   uint   `json:"updated_at"`
	CreatedAt   uint   `json:"created_at"`
}

// TableName 表名
func (Permission) TableName() string {
	return "rbac_permission"
}

// Role Model
type Role struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	UpdatedAt   uint   `json:"updated_at"`
	CreatedAt   uint   `json:"created_at"`
}

// TableName 表名
func (Role) TableName() string {
	return "rbac_role"
}

// RolePermission Model
type RolePermission struct {
	RoleId       uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
}

// TableName 表名
func (RolePermission) TableName() string {
	return "rbac_role_permission"
}

// UserRole Model
type UserRole struct {
	UserId uint `json:"user_id"`
	RoleId uint `json:"role_id"`
}

// TableName 表名
func (UserRole) TableName() string {
	return "rbac_user_role"
}
