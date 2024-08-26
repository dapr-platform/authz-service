package api

import (
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitUseroleRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/userole/page", UserolePageListHandler)
	r.Get(common.BASE_CONTEXT+"/userole", UseroleListHandler)
	r.Post(common.BASE_CONTEXT+"/userole", UpsertUseroleHandler)
	r.Delete(common.BASE_CONTEXT+"/userole/{id}", DeleteUseroleHandler)
	r.Post(common.BASE_CONTEXT+"/userole/batch-delete", batchDeleteUseroleHandler)
	r.Post(common.BASE_CONTEXT+"/userole/batch-upsert", batchUpsertUseroleHandler)
	r.Get(common.BASE_CONTEXT+"/userole/groupby", UseroleGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Userole
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /userole/groupby [get]
func UseroleGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "r_user_role")
}

// @Summary batch update
// @Description batch update
// @Tags Userole
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /userole/batch-upsert [post]
func batchUpsertUseroleHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.UseroleTableInfo.Name, model.Userole_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Userole
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param user_id query string false "user_id"
// @Param role_id query string false "role_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Userole}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /userole/page [get]
func UserolePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Userole](w, r, common.GetDaprClient(), "r_user_role", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Userole
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param user_id query string false "user_id"
// @Param role_id query string false "role_id"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Userole} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /userole [get]
func UseroleListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Userole](w, r, common.GetDaprClient(), "r_user_role", "id")
}

// @Summary save
// @Description save
// @Tags Userole
// @Accept       json
// @Param item body model.Userole true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Userole} "object"
// @Failure 500 {object} common.Response ""
// @Router /userole [post]
func UpsertUseroleHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Userole
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Userole")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Userole)
	}

	err = common.DbUpsert[model.Userole](r.Context(), common.GetDaprClient(), val, model.UseroleTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Userole
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Userole} "object"
// @Failure 500 {object} common.Response ""
// @Router /userole/{id} [delete]
func DeleteUseroleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Userole")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "r_user_role", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Userole
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /userole/batch-delete [post]
func batchDeleteUseroleHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Userole")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "r_user_role", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
