package service

import (
	"authz-service/entity"
	"authz-service/model"
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
)

func ChangeUserPassword(ctx context.Context, info entity.ChangePasswordInfo) (err error) {
	qstr := "_select=" + model.User_FIELD_NAME_id + "," + model.User_FIELD_NAME_password + "&" + model.User_FIELD_NAME_id + "=" + info.UserId + "&" + model.User_FIELD_NAME_password + "=" + info.OldPassword
	user, err := common.DbGetOne[map[string]any](ctx, common.GetDaprClient(), model.UserTableInfo.Name, qstr)
	if err != nil {
		return errors.Wrap(err, "DbGetOne")
	}
	if user == nil {
		return errors.New("旧密码错误")
	}
	m := *user
	m[model.User_FIELD_NAME_password] = info.NewPassword
	m[model.User_FIELD_NAME_update_at] = common.LocalTime(time.Now())
	return common.DbUpsert[map[string]any](ctx, common.GetDaprClient(), m, model.UserTableInfo.Name, model.User_FIELD_NAME_id)

}
func GetUserByIdAndPassword(ctx context.Context, id, password string) (user *model.User, err error) {

	users, err := common.DbQuery[model.User](ctx, common.GetDaprClient(), model.UserTableInfo.Name, "id="+id+"&password="+password)
	if err != nil {
		common.Logger.Error("db query error", err)
		return nil, nil
	}
	if len(users) == 0 {
		common.Logger.Error("user not found", id)
		return nil, nil
	}
	user = &users[0]
	return
}

func GetUserInfoByIdentityAndPassword(ctx context.Context, identity, password string) (user *entity.UserInfo, err error) {

	user, err = common.DbGetOne[entity.UserInfo](ctx, common.GetDaprClient(), model.User_with_menuTableInfo.Name, "identity="+identity+"&password="+password)
	if err != nil {
		common.Logger.Error("db query error", err)
		return nil, errors.Wrap(err, "db query error")
	}
	if user == nil {
		return
	}

	if user.Type == 1 {
		user.IsAdmin = 1
		allMenuResources, err := common.DbQuery[model.Resource](ctx, common.GetDaprClient(), model.ResourceTableInfo.Name, "type=1") //菜单权限
		if err != nil {
			common.Logger.Error("db query all resource error", err)
			return nil, errors.Wrap(err, "db query all resource error")
		}
		for _, v := range allMenuResources {
			user.MenuIds = append(user.MenuIds, entity.MenuId{
				ResourceId: v.ID,
			})
		}

	} else {
		user.IsAdmin = 0
	}
	menuIds := make([]entity.MenuId, 0)
	tmpMap := make(map[string]string, 0)
	for _, m := range user.MenuIds {
		_, exist := tmpMap[m.ResourceId]
		if !exist { //去重
			tmpMap[m.ResourceId] = m.ResourceId
			menuIds = append(menuIds, m)
		}
	}
	user.MenuIds = menuIds
	return user, nil
}

func GetUserByFieldName(ctx context.Context, field, value string) (user *model.User, err error) {
	value = url.QueryEscape(value)
	qstr := ""

	user, err = common.DbGetOne[model.User](ctx, common.GetDaprClient(), model.UserTableInfo.Name, field+"="+value+qstr)
	if err != nil {
		err = errors.WithMessage(err, "db query error field="+field+" value="+value)
		return
	}
	return
}

func getUserFromStateStore(ctx context.Context, id string) (user *entity.UserInfo, err error) {
	item, err := common.GetDaprClient().GetState(ctx, common.GLOBAL_STATESTOR_NAME, common.USER_STATESTORE_KEY_PREFIX+id, nil)
	if err != nil {
		common.Logger.Error("getUserFromStateStore", err)
		return
	}
	if len(item.Value) == 0 { //没有这个key
		return nil, nil
	}
	user = &entity.UserInfo{}
	err = json.Unmarshal(item.Value, user)
	if err != nil {
		common.Logger.Error("user unmarshal error:", err)
		return
	}
	return
}

func GetCurrentUserInfo(ctx context.Context, val string) (user *entity.UserInfo, err error) {
	user, err = getUserFromStateStore(ctx, val)
	if err != nil {
		common.Logger.Error("getUserFromStateStore error", err.Error())
		return
	}
	if user == nil {
		user, err = getUserInfoByIdFromDb(ctx, val)
		if err != nil {
			common.Logger.Error("getUserInfoByIdFromDb error", err.Error())
			return
		}
		if user == nil {
			err = errors.New("can't find user " + val + " in db. ")
		} else {
			buf, _ := json.Marshal(user)
			return user, common.SaveInStateStore(ctx, common.GetDaprClient(), common.GLOBAL_STATESTOR_NAME, common.USER_STATESTORE_KEY_PREFIX+user.ID, buf, true, time.Second*time.Duration(common.USER_EXPIRED_SECONDS))
		}

	}

	return
}

func getUserInfoByIdFromDb(ctx context.Context, id string) (user *entity.UserInfo, err error) {
	users, err := common.DbQuery[entity.UserInfo](ctx, common.GetDaprClient(), model.User_with_menuTableInfo.Name, "id="+id)
	if err != nil {
		common.Logger.Error("db query error", err)
		return nil, errors.Wrap(err, "db query error")
	}
	if len(users) == 0 {
		common.Logger.Error("user not found", id)
		return nil, errors.Wrap(err, "user not found "+id)
	}
	user = &users[0]
	if user.Type == 1 {
		user.IsAdmin = 1
		allMenuResources, err := common.DbQuery[model.Resource](ctx, common.GetDaprClient(), model.ResourceTableInfo.Name, "type=1") //菜单权限
		if err != nil {
			common.Logger.Error("db query all resource error", err)
			return nil, errors.Wrap(err, "db query all resource error")
		}
		for _, v := range allMenuResources {
			user.MenuIds = append(user.MenuIds, entity.MenuId{
				ResourceId:   v.ID,
				ResourceName: v.Name,
			})
		}
	}
	return user, nil
}
