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


Table: v_role_detail
[ 0] id                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] sort_index                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] create_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] update_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 6] remark                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] menu_resources                                 JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 8] func_resources                                 JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 9] api_resources                                  JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[10] data_resources                                 JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "PotXLcEjMBKAhSTqRfELShPiY",    "name": "eecDZZesvbpSnNLUeoKnYypjP",    "sort_index": 53,    "status": 77,    "create_at": 27,    "update_at": 67,    "remark": "yQlmPANanetglEluNquGBOTbu",    "menu_resources": 77,    "func_resources": 11,    "api_resources": 23,    "data_resources": 8}


Comments
-------------------------------------
[ 0] Warning table: v_role_detail does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_role_detail primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Role_detail_FIELD_NAME_id = "id"

	Role_detail_FIELD_NAME_name = "name"

	Role_detail_FIELD_NAME_sort_index = "sort_index"

	Role_detail_FIELD_NAME_status = "status"

	Role_detail_FIELD_NAME_create_at = "create_at"

	Role_detail_FIELD_NAME_update_at = "update_at"

	Role_detail_FIELD_NAME_remark = "remark"

	Role_detail_FIELD_NAME_menu_resources = "menu_resources"

	Role_detail_FIELD_NAME_func_resources = "func_resources"

	Role_detail_FIELD_NAME_api_resources = "api_resources"

	Role_detail_FIELD_NAME_data_resources = "data_resources"
)

// Role_detail struct is a row record of the v_role_detail table in the  database
type Role_detail struct {
	ID            string           `json:"id"`             //id
	Name          string           `json:"name"`           //name
	SortIndex     int32            `json:"sort_index"`     //sort_index
	Status        int32            `json:"status"`         //status
	CreateAt      common.LocalTime `json:"create_at"`      //create_at
	UpdateAt      common.LocalTime `json:"update_at"`      //update_at
	Remark        string           `json:"remark"`         //remark
	MenuResources any              `json:"menu_resources"` //menu_resources
	FuncResources any              `json:"func_resources"` //func_resources
	APIResources  any              `json:"api_resources"`  //api_resources
	DataResources any              `json:"data_resources"` //data_resources

}

var Role_detailTableInfo = &TableInfo{
	Name: "v_role_detail",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_role_detail does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_role_detail primary key column id is nullable column, setting it as NOT NULL
`,
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
			Comment:            `name`,
			Notes:              ``,
			Nullable:           true,
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
			Name:               "sort_index",
			Comment:            `sort_index`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
			Comment:            `status`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "create_at",
			Comment:            `create_at`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreateAt",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "create_at",
			ProtobufFieldName:  "create_at",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "update_at",
			Comment:            `update_at`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdateAt",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "update_at",
			ProtobufFieldName:  "update_at",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "remark",
			Comment:            `remark`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "menu_resources",
			Comment:            `menu_resources`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "MenuResources",
			GoFieldType:        "any",
			JSONFieldName:      "menu_resources",
			ProtobufFieldName:  "menu_resources",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "func_resources",
			Comment:            `func_resources`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "FuncResources",
			GoFieldType:        "any",
			JSONFieldName:      "func_resources",
			ProtobufFieldName:  "func_resources",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "api_resources",
			Comment:            `api_resources`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "APIResources",
			GoFieldType:        "any",
			JSONFieldName:      "api_resources",
			ProtobufFieldName:  "api_resources",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "data_resources",
			Comment:            `data_resources`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "DataResources",
			GoFieldType:        "any",
			JSONFieldName:      "data_resources",
			ProtobufFieldName:  "data_resources",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Role_detail) TableName() string {
	return "v_role_detail"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Role_detail) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Role_detail) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Role_detail) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Role_detail) TableInfo() *TableInfo {
	return Role_detailTableInfo
}
