package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitRole_with_resource_idsRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/role-with-resource-ids/page", Role_with_resource_idsPageListHandler)
	r.Get(common.BASE_CONTEXT+"/role-with-resource-ids", Role_with_resource_idsListHandler)
	r.Post(common.BASE_CONTEXT+"/role-with-resource-ids", UpsertRole_with_resource_idsHandler)
	r.Delete(common.BASE_CONTEXT+"/role-with-resource-ids/{id}", DeleteRole_with_resource_idsHandler)
	r.Post(common.BASE_CONTEXT+"/role-with-resource-ids/batch-delete", batchDeleteRole_with_resource_idsHandler)
	r.Post(common.BASE_CONTEXT+"/role-with-resource-ids/batch-upsert", batchUpsertRole_with_resource_idsHandler)
	r.Get(common.BASE_CONTEXT+"/role-with-resource-ids/groupby", Role_with_resource_idsGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Role_with_resource_ids
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids/groupby [get]
func Role_with_resource_idsGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "v_role_with_resource_ids")
}

// @Summary batch update
// @Description batch update
// @Tags Role_with_resource_ids
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids/batch-upsert [post]
func batchUpsertRole_with_resource_idsHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Role_with_resource_idsTableInfo.Name, model.Role_with_resource_ids_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Role_with_resource_ids
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
// @Param menu_resource_ids query string false "menu_resource_ids"
// @Param func_resource_ids query string false "func_resource_ids"
// @Param api_resource_ids query string false "api_resource_ids"
// @Param data_resource_ids query string false "data_resource_ids"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Role_with_resource_ids}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids/page [get]
func Role_with_resource_idsPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Role_with_resource_ids](w, r, common.GetDaprClient(), "v_role_with_resource_ids", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Role_with_resource_ids
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param sort_index query string false "sort_index"
// @Param status query string false "status"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param remark query string false "remark"
// @Param menu_resource_ids query string false "menu_resource_ids"
// @Param func_resource_ids query string false "func_resource_ids"
// @Param api_resource_ids query string false "api_resource_ids"
// @Param data_resource_ids query string false "data_resource_ids"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Role_with_resource_ids} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids [get]
func Role_with_resource_idsListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Role_with_resource_ids](w, r, common.GetDaprClient(), "v_role_with_resource_ids", "id")
}

// @Summary save
// @Description save
// @Tags Role_with_resource_ids
// @Accept       json
// @Param item body model.Role_with_resource_ids true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role_with_resource_ids} "object"
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids [post]
func UpsertRole_with_resource_idsHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Role_with_resource_ids
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Role_with_resource_ids")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Role_with_resource_ids)
	}

	err = common.DbUpsert[model.Role_with_resource_ids](r.Context(), common.GetDaprClient(), val, model.Role_with_resource_idsTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Role_with_resource_ids
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role_with_resource_ids} "object"
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids/{id} [delete]
func DeleteRole_with_resource_idsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Role_with_resource_ids")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "v_role_with_resource_ids", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Role_with_resource_ids
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /role-with-resource-ids/batch-delete [post]
func batchDeleteRole_with_resource_idsHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Role_with_resource_ids")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "v_role_with_resource_ids", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
