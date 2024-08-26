package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitRoleRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/role/page", RolePageListHandler)
	r.Get(common.BASE_CONTEXT+"/role", RoleListHandler)
	r.Post(common.BASE_CONTEXT+"/role", UpsertRoleHandler)
	r.Delete(common.BASE_CONTEXT+"/role/{id}", DeleteRoleHandler)
	r.Post(common.BASE_CONTEXT+"/role/batch-delete", batchDeleteRoleHandler)
	r.Post(common.BASE_CONTEXT+"/role/batch-upsert", batchUpsertRoleHandler)
	r.Get(common.BASE_CONTEXT+"/role/groupby", RoleGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Role
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role/groupby [get]
func RoleGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_role")
}

// @Summary batch update
// @Description batch update
// @Tags Role
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /role/batch-upsert [post]
func batchUpsertRoleHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.RoleTableInfo.Name, model.Role_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Role
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param sort_index query string false "sort_index"
// @Param status query string false "status"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Role}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role/page [get]
func RolePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Role](w, r, common.GetDaprClient(), "o_role", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Role
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param sort_index query string false "sort_index"
// @Param status query string false "status"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Role} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role [get]
func RoleListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Role](w, r, common.GetDaprClient(), "o_role", "id")
}

// @Summary save
// @Description save
// @Tags Role
// @Accept       json
// @Param item body model.Role true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role} "object"
// @Failure 500 {object} common.Response ""
// @Router /role [post]
func UpsertRoleHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Role
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Role")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Role)
	}

	err = common.DbUpsert[model.Role](r.Context(), common.GetDaprClient(), val, model.RoleTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Role
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role} "object"
// @Failure 500 {object} common.Response ""
// @Router /role/{id} [delete]
func DeleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Role")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_role", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Role
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /role/batch-delete [post]
func batchDeleteRoleHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Role")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_role", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
