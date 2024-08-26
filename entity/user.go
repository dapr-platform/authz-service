package entity

import "authz-service/model"

type MenuId struct {
	ResourceId string `json:"resource_id"` //for example proj-001
}
type UserInfo struct {
	model.User
	IsAdmin int      `json:"is_admin"` //is super admin
	MenuIds []MenuId `json:"menu_ids"`
}

type ChangePasswordInfo struct {
	UserId      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserLoginReq struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
