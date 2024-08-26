package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitResourceRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/resource/page", ResourcePageListHandler)
	r.Get(common.BASE_CONTEXT+"/resource", ResourceListHandler)
	r.Post(common.BASE_CONTEXT+"/resource", UpsertResourceHandler)
	r.Delete(common.BASE_CONTEXT+"/resource/{id}", DeleteResourceHandler)
	r.Post(common.BASE_CONTEXT+"/resource/batch-delete", batchDeleteResourceHandler)
	r.Post(common.BASE_CONTEXT+"/resource/batch-upsert", batchUpsertResourceHandler)
	r.Get(common.BASE_CONTEXT+"/resource/groupby", ResourceGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Resource
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /resource/groupby [get]
func ResourceGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_resource")
}

// @Summary batch update
// @Description batch update
// @Tags Resource
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /resource/batch-upsert [post]
func batchUpsertResourceHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.ResourceTableInfo.Name, model.Resource_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Resource
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param parent_id query string false "parent_id"
// @Param module query string false "module"
// @Param service_name query string false "service_name"
// @Param service_name_cn query string false "service_name_cn"
// @Param type query string false "type"
// @Param api_url query string false "api_url"
// @Param sort_index query string false "sort_index"
// @Param data_conditions query string false "data_conditions"
// @Param support_ops query string false "support_ops"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Resource}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /resource/page [get]
func ResourcePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Resource](w, r, common.GetDaprClient(), "o_resource", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Resource
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param parent_id query string false "parent_id"
// @Param module query string false "module"
// @Param service_name query string false "service_name"
// @Param service_name_cn query string false "service_name_cn"
// @Param type query string false "type"
// @Param api_url query string false "api_url"
// @Param sort_index query string false "sort_index"
// @Param data_conditions query string false "data_conditions"
// @Param support_ops query string false "support_ops"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Resource} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /resource [get]
func ResourceListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Resource](w, r, common.GetDaprClient(), "o_resource", "id")
}

// @Summary save
// @Description save
// @Tags Resource
// @Accept       json
// @Param item body model.Resource true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Resource} "object"
// @Failure 500 {object} common.Response ""
// @Router /resource [post]
func UpsertResourceHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Resource
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Resource")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Resource)
	}

	err = common.DbUpsert[model.Resource](r.Context(), common.GetDaprClient(), val, model.ResourceTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Resource
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Resource} "object"
// @Failure 500 {object} common.Response ""
// @Router /resource/{id} [delete]
func DeleteResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Resource")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_resource", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Resource
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /resource/batch-delete [post]
func batchDeleteResourceHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Resource")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_resource", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
