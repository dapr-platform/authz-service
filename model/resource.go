package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: o_resource
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] parent_id                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] module                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] service_name                                   VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] service_name_cn                                VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] api_url                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] sort_index                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 9] data_conditions                                VARCHAR(10240)       null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 10240   default: []
[10] support_ops                                    VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "KXGdnYVWVasYrPbPRnWBcOCEW",    "name": "sRjlwGdiaQncwhwuktJMqXHLo",    "parent_id": "kAsPNjyMmDBnhZjGMLbMgCVXl",    "module": "spucETjcmOUoygWQqGnMhXMNM",    "service_name": "DKMsLeIvwYcALYjoeSIBCcofo",    "service_name_cn": "xVCBTxfXdlwoWhUFrragaVyyR",    "type": 50,    "api_url": "ntorVEVRwxDaOqFVhxVFpOSOD",    "sort_index": 72,    "data_conditions": "WJiftOkDUpMxqJxhSLSNksuJt",    "support_ops": "OnVIUmmZCVbxVkXudgpgPebKP"}



*/

var (
	Resource_FIELD_NAME_id = "id"

	Resource_FIELD_NAME_name = "name"

	Resource_FIELD_NAME_parent_id = "parent_id"

	Resource_FIELD_NAME_module = "module"

	Resource_FIELD_NAME_service_name = "service_name"

	Resource_FIELD_NAME_service_name_cn = "service_name_cn"

	Resource_FIELD_NAME_type = "type"

	Resource_FIELD_NAME_api_url = "api_url"

	Resource_FIELD_NAME_sort_index = "sort_index"

	Resource_FIELD_NAME_data_conditions = "data_conditions"

	Resource_FIELD_NAME_support_ops = "support_ops"
)

// Resource struct is a row record of the o_resource table in the  database
type Resource struct {
	ID             string `json:"id"`              //唯一标识,例如UI-001-001
	Name           string `json:"name"`            //名称
	ParentID       string `json:"parent_id"`       //parent_id
	Module         string `json:"module"`          //模块
	ServiceName    string `json:"service_name"`    //服务名
	ServiceNameCn  string `json:"service_name_cn"` //中文服务名
	Type           int32  `json:"type"`            //分类（1:菜单,2:api,3:功能点,4:数据)
	APIURL         string `json:"api_url"`         //api鉴权用到
	SortIndex      int32  `json:"sort_index"`      //sort_index
	DataConditions string `json:"data_conditions"` //数据条件（数据鉴权用到，存放对象的json字符串)
	SupportOps     string `json:"support_ops"`     //支持的操作(字典值,1,2,3)

}

var ResourceTableInfo = &TableInfo{
	Name: "o_resource",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `唯一标识,例如UI-001-001`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            `名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "parent_id",
			Comment:            `parent_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ParentID",
			GoFieldType:        "string",
			JSONFieldName:      "parent_id",
			ProtobufFieldName:  "parent_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "module",
			Comment:            `模块`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Module",
			GoFieldType:        "string",
			JSONFieldName:      "module",
			ProtobufFieldName:  "module",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "service_name",
			Comment:            `服务名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ServiceName",
			GoFieldType:        "string",
			JSONFieldName:      "service_name",
			ProtobufFieldName:  "service_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "service_name_cn",
			Comment:            `中文服务名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ServiceNameCn",
			GoFieldType:        "string",
			JSONFieldName:      "service_name_cn",
			ProtobufFieldName:  "service_name_cn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "type",
			Comment:            `分类（1:菜单,2:api,3:功能点,4:数据)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "api_url",
			Comment:            `api鉴权用到`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "APIURL",
			GoFieldType:        "string",
			JSONFieldName:      "api_url",
			ProtobufFieldName:  "api_url",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "sort_index",
			Comment:            `sort_index`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SortIndex",
			GoFieldType:        "int32",
			JSONFieldName:      "sort_index",
			ProtobufFieldName:  "sort_index",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "data_conditions",
			Comment:            `数据条件（数据鉴权用到，存放对象的json字符串)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10240)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10240,
			GoFieldName:        "DataConditions",
			GoFieldType:        "string",
			JSONFieldName:      "data_conditions",
			ProtobufFieldName:  "data_conditions",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "support_ops",
			Comment:            `支持的操作(字典值,1,2,3)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "SupportOps",
			GoFieldType:        "string",
			JSONFieldName:      "support_ops",
			ProtobufFieldName:  "support_ops",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Resource) TableName() string {
	return "o_resource"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Resource) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Resource) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Resource) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Resource) TableInfo() *TableInfo {
	return ResourceTableInfo
}
