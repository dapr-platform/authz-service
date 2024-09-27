package api

import (
	"authz-service/entity"
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"net/http"
)

var (
	_ = null.String{}
)

func InitCustomRole_resource_operateRoute(r chi.Router) {

	r.Post(common.BASE_CONTEXT+"/role-resource-operate/batch", BatchAddRole_resource_operateHandler)
	r.Get(common.BASE_CONTEXT+"/role/detail", roleDetailHandler)

}

// @Summary 批量更新角色的资源关系
// @Description 批量更新角色的资源关系, 需要传role_id, 先全部删除再增加
// @Tags Role_resource_operate
// @Accept       json
// @Param role_id query string true "role_id"
// @Param items body []model.Role_resource_operate true "实例全部信息"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Role_resource_operate} "实例"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /role-resource-operate/batch [post]
func BatchAddRole_resource_operateHandler(w http.ResponseWriter, r *http.Request) {
	roleId := r.URL.Query().Get("role_id")
	if roleId == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("role id is blank"))
		return
	}
	err := common.DbDeleteByOps(r.Context(), common.GetDaprClient(), model.Role_resource_operateTableInfo.Name, []string{"role_id"}, []string{"=="}, []interface{}{roleId})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("DbDeleteByOps error"))
		return
	}
	var data []model.Role_resource_operate
	err = common.ReadRequestBody(r, &data)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("read body error"))
		return
	}
	updateData := make([]model.Role_resource_operate, 0)
	for _, d := range data {
		d.ID = common.GetMD5Hash(d.RoleID + "_" + d.ResourceID + "_" + d.Op)
		updateData = append(updateData, d)
	}
	if len(data) > 0 {

		err = common.DbBatchUpsert(r.Context(), common.GetDaprClient(), updateData, model.Role_resource_operateTableInfo.Name, "id")
		if err != nil {
			common.HttpResult(w, common.ErrService.AppendMsg("DbBatchUpsert error"))
			return
		}
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary 查询角色详情列表
// @Description 查询角色详情列表
// @Tags Role
// @Param _page query int false "current page"
// @Param _page_size query int false "page size"
// @Param _order query string false "order"
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.RoleDetailVo} "实例"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /role/detail [get]
func roleDetailHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page != "" && pageSize != "" {
		common.CommonPageQuery[entity.RoleDetailVo](w, r, common.GetDaprClient(), model.Role_detailTableInfo.Name, model.Role_detail_FIELD_NAME_id)
	} else {
		common.CommonQuery[entity.RoleDetailVo](w, r, common.GetDaprClient(), model.Role_detailTableInfo.Name, model.Role_detail_FIELD_NAME_id)
	}

}
