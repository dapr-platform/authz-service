package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitRole_resource_operateRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/role-resource-operate/page", Role_resource_operatePageListHandler)
	r.Get(common.BASE_CONTEXT+"/role-resource-operate", Role_resource_operateListHandler)
	r.Post(common.BASE_CONTEXT+"/role-resource-operate", UpsertRole_resource_operateHandler)
	r.Delete(common.BASE_CONTEXT+"/role-resource-operate/{id}", DeleteRole_resource_operateHandler)
	r.Post(common.BASE_CONTEXT+"/role-resource-operate/batch-delete", batchDeleteRole_resource_operateHandler)
	r.Post(common.BASE_CONTEXT+"/role-resource-operate/batch-upsert", batchUpsertRole_resource_operateHandler)
	r.Get(common.BASE_CONTEXT+"/role-resource-operate/groupby", Role_resource_operateGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Role_resource_operate
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate/groupby [get]
func Role_resource_operateGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "r_role_resource_operate")
}

// @Summary batch update
// @Description batch update
// @Tags Role_resource_operate
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate/batch-upsert [post]
func batchUpsertRole_resource_operateHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Role_resource_operateTableInfo.Name, model.Role_resource_operate_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Role_resource_operate
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param role_id query string false "role_id"
// @Param resource_id query string false "resource_id"
// @Param op query string false "op"
// @Param filter_conditions query string false "filter_conditions"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Role_resource_operate}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate/page [get]
func Role_resource_operatePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Role_resource_operate](w, r, common.GetDaprClient(), "r_role_resource_operate", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Role_resource_operate
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param role_id query string false "role_id"
// @Param resource_id query string false "resource_id"
// @Param op query string false "op"
// @Param filter_conditions query string false "filter_conditions"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Role_resource_operate} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate [get]
func Role_resource_operateListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Role_resource_operate](w, r, common.GetDaprClient(), "r_role_resource_operate", "id")
}

// @Summary save
// @Description save
// @Tags Role_resource_operate
// @Accept       json
// @Param item body model.Role_resource_operate true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role_resource_operate} "object"
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate [post]
func UpsertRole_resource_operateHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Role_resource_operate
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Role_resource_operate")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Role_resource_operate)
	}

	err = common.DbUpsert[model.Role_resource_operate](r.Context(), common.GetDaprClient(), val, model.Role_resource_operateTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Role_resource_operate
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role_resource_operate} "object"
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate/{id} [delete]
func DeleteRole_resource_operateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Role_resource_operate")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "r_role_resource_operate", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Role_resource_operate
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /role-resource-operate/batch-delete [post]
func batchDeleteRole_resource_operateHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Role_resource_operate")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "r_role_resource_operate", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
