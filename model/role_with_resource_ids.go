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


Table: v_role_with_resource_ids
[ 0] id                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] sort_index                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] create_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] update_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 6] remark                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] menu_resource_ids                              JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 8] func_resource_ids                              JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 9] api_resource_ids                               JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[10] data_resource_ids                              JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "GaYBPxZflnXfeKCXiUQbOoukZ",    "name": "TcfclERvHNDvKISmRZHdZaXWx",    "sort_index": 77,    "status": 79,    "create_at": 3,    "update_at": 25,    "remark": "IpgPhCDRrIovgLevfLwJFvysX",    "menu_resource_ids": 75,    "func_resource_ids": 6,    "api_resource_ids": 23,    "data_resource_ids": 3}


Comments
-------------------------------------
[ 0] Warning table: v_role_with_resource_ids does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_role_with_resource_ids primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Role_with_resource_ids_FIELD_NAME_id = "id"

	Role_with_resource_ids_FIELD_NAME_name = "name"

	Role_with_resource_ids_FIELD_NAME_sort_index = "sort_index"

	Role_with_resource_ids_FIELD_NAME_status = "status"

	Role_with_resource_ids_FIELD_NAME_create_at = "create_at"

	Role_with_resource_ids_FIELD_NAME_update_at = "update_at"

	Role_with_resource_ids_FIELD_NAME_remark = "remark"

	Role_with_resource_ids_FIELD_NAME_menu_resource_ids = "menu_resource_ids"

	Role_with_resource_ids_FIELD_NAME_func_resource_ids = "func_resource_ids"

	Role_with_resource_ids_FIELD_NAME_api_resource_ids = "api_resource_ids"

	Role_with_resource_ids_FIELD_NAME_data_resource_ids = "data_resource_ids"
)

// Role_with_resource_ids struct is a row record of the v_role_with_resource_ids table in the  database
type Role_with_resource_ids struct {
	ID              string           `json:"id"`                //id
	Name            string           `json:"name"`              //name
	SortIndex       int32            `json:"sort_index"`        //sort_index
	Status          int32            `json:"status"`            //status
	CreateAt        common.LocalTime `json:"create_at"`         //create_at
	UpdateAt        common.LocalTime `json:"update_at"`         //update_at
	Remark          string           `json:"remark"`            //remark
	MenuResourceIds any              `json:"menu_resource_ids"` //menu_resource_ids
	FuncResourceIds any              `json:"func_resource_ids"` //func_resource_ids
	APIResourceIds  any              `json:"api_resource_ids"`  //api_resource_ids
	DataResourceIds any              `json:"data_resource_ids"` //data_resource_ids

}

var Role_with_resource_idsTableInfo = &TableInfo{
	Name: "v_role_with_resource_ids",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_role_with_resource_ids does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_role_with_resource_ids primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "menu_resource_ids",
			Comment:            `menu_resource_ids`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "MenuResourceIds",
			GoFieldType:        "any",
			JSONFieldName:      "menu_resource_ids",
			ProtobufFieldName:  "menu_resource_ids",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "func_resource_ids",
			Comment:            `func_resource_ids`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "FuncResourceIds",
			GoFieldType:        "any",
			JSONFieldName:      "func_resource_ids",
			ProtobufFieldName:  "func_resource_ids",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "api_resource_ids",
			Comment:            `api_resource_ids`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "APIResourceIds",
			GoFieldType:        "any",
			JSONFieldName:      "api_resource_ids",
			ProtobufFieldName:  "api_resource_ids",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "data_resource_ids",
			Comment:            `data_resource_ids`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "DataResourceIds",
			GoFieldType:        "any",
			JSONFieldName:      "data_resource_ids",
			ProtobufFieldName:  "data_resource_ids",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Role_with_resource_ids) TableName() string {
	return "v_role_with_resource_ids"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Role_with_resource_ids) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Role_with_resource_ids) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Role_with_resource_ids) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Role_with_resource_ids) TableInfo() *TableInfo {
	return Role_with_resource_idsTableInfo
}
