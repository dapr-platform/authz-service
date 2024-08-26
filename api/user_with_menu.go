package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitUser_with_menuRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user-with-menu/page", User_with_menuPageListHandler)
	r.Get(common.BASE_CONTEXT+"/user-with-menu", User_with_menuListHandler)
	r.Post(common.BASE_CONTEXT+"/user-with-menu", UpsertUser_with_menuHandler)
	r.Delete(common.BASE_CONTEXT+"/user-with-menu/{id}", DeleteUser_with_menuHandler)
	r.Post(common.BASE_CONTEXT+"/user-with-menu/batch-delete", batchDeleteUser_with_menuHandler)
	r.Post(common.BASE_CONTEXT+"/user-with-menu/batch-upsert", batchUpsertUser_with_menuHandler)
	r.Get(common.BASE_CONTEXT+"/user-with-menu/groupby", User_with_menuGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags User_with_menu
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu/groupby [get]
func User_with_menuGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "v_user_with_menu")
}

// @Summary batch update
// @Description batch update
// @Tags User_with_menu
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu/batch-upsert [post]
func batchUpsertUser_with_menuHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.User_with_menuTableInfo.Name, model.User_with_menu_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags User_with_menu
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
// @Param menu_ids query string false "menu_ids"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.User_with_menu}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu/page [get]
func User_with_menuPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.User_with_menu](w, r, common.GetDaprClient(), "v_user_with_menu", "id")

}

// @Summary query objects
// @Description query objects
// @Tags User_with_menu
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
// @Param menu_ids query string false "menu_ids"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.User_with_menu} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu [get]
func User_with_menuListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.User_with_menu](w, r, common.GetDaprClient(), "v_user_with_menu", "id")
}

// @Summary save
// @Description save
// @Tags User_with_menu
// @Accept       json
// @Param item body model.User_with_menu true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User_with_menu} "object"
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu [post]
func UpsertUser_with_menuHandler(w http.ResponseWriter, r *http.Request) {
	var val model.User_with_menu
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("User_with_menu")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.User_with_menu)
	}

	err = common.DbUpsert[model.User_with_menu](r.Context(), common.GetDaprClient(), val, model.User_with_menuTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags User_with_menu
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User_with_menu} "object"
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu/{id} [delete]
func DeleteUser_with_menuHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("User_with_menu")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "v_user_with_menu", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags User_with_menu
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /user-with-menu/batch-delete [post]
func batchDeleteUser_with_menuHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("User_with_menu")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "v_user_with_menu", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
