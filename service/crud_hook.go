package service

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"net/http"
)

func init() {
	common.RegisterDeleteBeforeHook("User", DeleteUserBeforeHook)
	common.RegisterDeleteBeforeHook("Role", DeleteRoleBeforeHook)
}

func DeleteUserBeforeHook(r *http.Request, in any) (out any, err error) {
	userId := in.(string)

	userRoleRels, err := common.DbQuery[model.Userole](r.Context(), common.GetDaprClient(), model.UseroleTableInfo.Name, "user_id="+userId)
	if err != nil {
		common.Logger.Error("query userole error", err)
		return nil, err
	}
	for _, userRoleRel := range userRoleRels {
		err := common.DbDelete(r.Context(), common.GetDaprClient(), model.UseroleTableInfo.Name, model.Userole_FIELD_NAME_id, userRoleRel.ID)
		if err != nil {
			common.Logger.Error("delete userole error", err)
			return nil, err
		}
	}

	return in, nil
}

func DeleteRoleBeforeHook(r *http.Request, in any) (out any, err error) {
	roleId := in.(string)

	rels, err := common.DbQuery[model.Role_resource_operate](r.Context(), common.GetDaprClient(), model.Role_resource_operateTableInfo.Name, "role_id="+roleId)
	if err != nil {
		common.Logger.Error("query role resource operate error", err)
		return nil, err
	}
	for _, rel := range rels {
		err := common.DbDelete(r.Context(), common.GetDaprClient(), model.Role_resource_operateTableInfo.Name, model.Role_resource_operate_FIELD_NAME_id, rel.ID)
		if err != nil {
			common.Logger.Error("delete role resource operate error", err)
			return nil, err
		}
	}

	return in, nil
}
