package api

import (
	"authz-service/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
)

var (
	_ = null.String{}
)

func InitUser_roleRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user-role", User_roleListHandler)
	r.Post(common.BASE_CONTEXT+"/user-role", AddUser_roleHandler)
	r.Post(common.BASE_CONTEXT+"/user-role/batch", BatchAddUser_roleHandler)
	r.Put(common.BASE_CONTEXT+"/user-role", UpdateUser_roleHandler)
	r.Delete(common.BASE_CONTEXT+"/user-role", DeleteUser_roleHandler)
}

// @Summary 查询所有用户角色关系
// @Description 查询所有用户角色关系, 可设置page, page_size, order, 以及查询条件等，例如 status=1, name=$like.%25%25CAM%25%25 等
// @Tags User_role
// @Param _page query int false "current page"
// @Param _page_size query int false "page size"
// @Param _order query string false "order"
// @Param _select query string false "select"
// @Param id query string false "id"
// @Param user_id query string false "user_id"
// @Param role_id query string false "role_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Userole}} "用户角色关系的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-role [get]
func User_roleListHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page != "" && pageSize != "" {
		common.CommonPageQuery[model.Userole](w, r, common.GetDaprClient(), "r_user_role", "id")
	} else {
		common.CommonQuery[model.Userole](w, r, common.GetDaprClient(), "r_user_role", "id")
	}

}

// @Summary 添加用户角色关系
// @Description 添加用户角色关系
// @Tags User_role
// @Accept       json
// @Param item body model.Userole true "用户角色关系全部信息"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Userole} "用户角色关系"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-role [post]
func AddUser_roleHandler(w http.ResponseWriter, r *http.Request) {
	var info model.Userole
	err := common.ReadRequestBody(r, &info)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("read body error ").AppendMsg(err.Error()))
		return
	}
	if info.ID == "" {
		info.ID = common.NanoId()
	}
	err = common.DbUpsertIg[model.Userole](r.Context(), common.GetDaprClient(), info, model.UseroleTableInfo.Name, model.Userole_FIELD_NAME_user_id+","+model.Userole_FIELD_NAME_role_id, model.Userole_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("DbUpsertIg error ").AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK.WithData(info))
}

// @Summary 批量添加用户角色关系
// @Description 批量添加用户角色关系,参数包含全部的角色
// @Tags User_role
// @Accept       json
// @Param delete_before query int false "是否先删除再添加，1:是，0：否 默认为0， "
// @Param user_ids query string false "用户id, 多个用,分隔"
// @Param item body []model.Userole true "用户角色关系全部信息"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Userole} "用户角色关系"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-role/batch [post]
func BatchAddUser_roleHandler(w http.ResponseWriter, r *http.Request) {
	deleteBefore := 0
	dstr := r.URL.Query().Get("delete_before")

	if dstr != "" {
		deleteBefore, _ = strconv.Atoi(dstr)
	}
	var infos []model.Userole
	err := common.ReadRequestBody(r, &infos)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("read body error ").AppendMsg(err.Error()))
		return
	}

	iinfos := make([]model.Userole, 0)
	userIds := make(map[string]string, 0)
	for _, info := range infos {
		if info.ID == "" {
			info.ID = common.NanoId()
		}
		userIds[info.UserID] = info.UserID
		iinfos = append(iinfos, info)
	}
	userIdsStr := r.URL.Query().Get("user_ids")
	if userIdsStr != "" {
		userIdsArr := strings.Split(userIdsStr, ",")
		for _,u :=range userIdsArr{
			userIds[u] = u
		}
	}

	if deleteBefore == 1 {
		for k, _ := range userIds {
			err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), model.UseroleTableInfo.Name, []string{"user_id"}, []string{"=="}, []any{k})
			if err != nil {
				common.HttpResult(w, common.ErrParam.AppendMsg("DbDeleteByOps error ").AppendMsg(err.Error()))
				return
			}
		}
	}
	if len(iinfos) == 0 {
		common.HttpResult(w, common.OK)
		return
	} else {
		err = common.DbBatchUpsertIg[model.Userole](r.Context(), common.GetDaprClient(), iinfos, model.UseroleTableInfo.Name, model.Userole_FIELD_NAME_user_id+","+model.Userole_FIELD_NAME_role_id, model.Userole_FIELD_NAME_id)
		if err != nil {
			common.HttpResult(w, common.ErrParam.AppendMsg("DbBatchInsert error ").AppendMsg(err.Error()))
			return
		}
		common.HttpResult(w, common.OK.WithData(iinfos))
	}

}

// @Summary 更新用户角色关系
// @Description 更新用户角色关系
// @Tags User_role
// @Accept       json
// @Param item body model.Userole true "用户角色关系全部信息"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Userole} "用户角色关系"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-role [put]
func UpdateUser_roleHandler(w http.ResponseWriter, r *http.Request) {

	var info model.Userole
	err := common.ReadRequestBody(r, &info)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("read body error ").AppendMsg(err.Error()))
		return
	}
	if info.ID == "" {
		info.ID = common.NanoId()
	}
	err = common.DbUpsertIg[model.Userole](r.Context(), common.GetDaprClient(), info, model.UseroleTableInfo.Name, model.Userole_FIELD_NAME_user_id+","+model.Userole_FIELD_NAME_role_id, model.Userole_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("DbUpsertIg error ").AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK.WithData(info))
}

// @Summary 删除用户角色关系
// @Description 删除用户角色关系
// @Tags User_role
// @Param item body []string true "id集合"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Userole} "用户角色关系"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-role [delete]
func DeleteUser_roleHandler(w http.ResponseWriter, r *http.Request) {
	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.Logger.Error("read body error ", err)
	}
	errStr := ""

	for _, id := range ids {

		err = common.DbDelete(r.Context(), common.GetDaprClient(), model.UseroleTableInfo.Name, "id", id)
		if err != nil {
			errStr += id + " delete error " + err.Error() + "\n"
			continue
		}

	}
	if errStr != "" {
		common.HttpResult(w, common.ErrService.AppendMsg("db delete error").AppendMsg(errStr))
		return
	}
	common.HttpResult(w, common.OK)
}
