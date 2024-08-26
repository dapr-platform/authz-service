package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitUserRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user/page", UserPageListHandler)
	r.Get(common.BASE_CONTEXT+"/user", UserListHandler)
	r.Post(common.BASE_CONTEXT+"/user", UpsertUserHandler)
	r.Delete(common.BASE_CONTEXT+"/user/{id}", DeleteUserHandler)
	r.Post(common.BASE_CONTEXT+"/user/batch-delete", batchDeleteUserHandler)
	r.Post(common.BASE_CONTEXT+"/user/batch-upsert", batchUpsertUserHandler)
	r.Get(common.BASE_CONTEXT+"/user/groupby", UserGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags User
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user/groupby [get]
func UserGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_user")
}

// @Summary batch update
// @Description batch update
// @Tags User
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /user/batch-upsert [post]
func batchUpsertUserHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.UserTableInfo.Name, model.User_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags User
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param tenant_id query string false "tenant_id"
// @Param mobile query string false "mobile"
// @Param email query string false "email"
// @Param identity query string false "identity"
// @Param name query string false "name"
// @Param gender query string false "gender"
// @Param address query string false "address"
// @Param password query string false "password"
// @Param type query string false "type"
// @Param org_id query string false "org_id"
// @Param avatar_url query string false "avatar_url"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.User}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user/page [get]
func UserPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.User](w, r, common.GetDaprClient(), "o_user", "id")

}

// @Summary query objects
// @Description query objects
// @Tags User
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param tenant_id query string false "tenant_id"
// @Param mobile query string false "mobile"
// @Param email query string false "email"
// @Param identity query string false "identity"
// @Param name query string false "name"
// @Param gender query string false "gender"
// @Param address query string false "address"
// @Param password query string false "password"
// @Param type query string false "type"
// @Param org_id query string false "org_id"
// @Param avatar_url query string false "avatar_url"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.User} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user [get]
func UserListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.User](w, r, common.GetDaprClient(), "o_user", "id")
}

// @Summary save
// @Description save
// @Tags User
// @Accept       json
// @Param item body model.User true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User} "object"
// @Failure 500 {object} common.Response ""
// @Router /user [post]
func UpsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var val model.User
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("User")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.User)
	}

	err = common.DbUpsert[model.User](r.Context(), common.GetDaprClient(), val, model.UserTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags User
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User} "object"
// @Failure 500 {object} common.Response ""
// @Router /user/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("User")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_user", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags User
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /user/batch-delete [post]
func batchDeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("User")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_user", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
