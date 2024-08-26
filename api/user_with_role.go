package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitUser_with_roleRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user-with-role/page", User_with_rolePageListHandler)
	r.Get(common.BASE_CONTEXT+"/user-with-role", User_with_roleListHandler)
	r.Post(common.BASE_CONTEXT+"/user-with-role", UpsertUser_with_roleHandler)
	r.Delete(common.BASE_CONTEXT+"/user-with-role/{id}", DeleteUser_with_roleHandler)
	r.Post(common.BASE_CONTEXT+"/user-with-role/batch-delete", batchDeleteUser_with_roleHandler)
	r.Post(common.BASE_CONTEXT+"/user-with-role/batch-upsert", batchUpsertUser_with_roleHandler)
	r.Get(common.BASE_CONTEXT+"/user-with-role/groupby", User_with_roleGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags User_with_role
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user-with-role/groupby [get]
func User_with_roleGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "v_user_with_role")
}

// @Summary batch update
// @Description batch update
// @Tags User_with_role
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /user-with-role/batch-upsert [post]
func batchUpsertUser_with_roleHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.User_with_roleTableInfo.Name, model.User_with_role_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags User_with_role
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
// @Param roles query string false "roles"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.User_with_role}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user-with-role/page [get]
func User_with_rolePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.User_with_role](w, r, common.GetDaprClient(), "v_user_with_role", "id")

}

// @Summary query objects
// @Description query objects
// @Tags User_with_role
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
// @Param roles query string false "roles"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.User_with_role} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user-with-role [get]
func User_with_roleListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.User_with_role](w, r, common.GetDaprClient(), "v_user_with_role", "id")
}

// @Summary save
// @Description save
// @Tags User_with_role
// @Accept       json
// @Param item body model.User_with_role true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User_with_role} "object"
// @Failure 500 {object} common.Response ""
// @Router /user-with-role [post]
func UpsertUser_with_roleHandler(w http.ResponseWriter, r *http.Request) {
	var val model.User_with_role
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("User_with_role")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.User_with_role)
	}

	err = common.DbUpsert[model.User_with_role](r.Context(), common.GetDaprClient(), val, model.User_with_roleTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags User_with_role
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User_with_role} "object"
// @Failure 500 {object} common.Response ""
// @Router /user-with-role/{id} [delete]
func DeleteUser_with_roleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("User_with_role")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "v_user_with_role", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags User_with_role
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /user-with-role/batch-delete [post]
func batchDeleteUser_with_roleHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of ids is 0"))
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("User_with_role")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "v_user_with_role", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
