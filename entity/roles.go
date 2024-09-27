package entity

import "authz-service/model"

type RoleDetailVo struct {
	ID            string                  `json:"id"`             //唯一标识
	Name          string                  `json:"name"`           //名称
	MenuResources []model.Role_rel_detail `json:"menu_resources"` //菜单权限
	ApiResources  []model.Role_rel_detail `json:"api_resources"`  //api权限
	FuncResources []model.Role_rel_detail `json:"func_resources"` //功能点权限
	DataResources []model.Role_rel_detail `json:"data_resources"` //数据权限
}

type RoleWithResourceIds struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	SortIndex       int      `json:"sort_index"`
	Status          int      `json:"status"`
	CreateAt        string   `json:"create_at"`
	UpdateAt        string   `json:"update_at"`
	Remark          string   `json:"remark"`
	MenuResourceIds []string `json:"menu_resource_ids"`
	FuncResourceIds []string `json:"func_resource_ids"`
	ApiResourceIds  []string `json:"api_resource_ids"`
	DataResourceIds []string `json:"data_resource_ids"`
}

type ResourceTreeVo struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	ParentId string            `json:"parent_id"`
	Children []*ResourceTreeVo `json:"children"`
}
