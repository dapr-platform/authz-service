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


Table: o_role
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] sort_index                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 3] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]
[ 4] create_at                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [now()]
[ 5] update_at                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [now()]
[ 6] remark                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "krafltUZAAlktmJwbgxpeqXCp",    "name": "FVfBJMRlEdeBfRiKCdCUpUQjP",    "sort_index": 27,    "status": 85,    "create_at": 3,    "update_at": 11,    "remark": "ZwJLVcDnQMRaUpmiDRwuhdotQ"}



*/

var (
	Role_FIELD_NAME_id = "id"

	Role_FIELD_NAME_name = "name"

	Role_FIELD_NAME_sort_index = "sort_index"

	Role_FIELD_NAME_status = "status"

	Role_FIELD_NAME_create_at = "create_at"

	Role_FIELD_NAME_update_at = "update_at"

	Role_FIELD_NAME_remark = "remark"
)

// Role struct is a row record of the o_role table in the  database
type Role struct {
	ID        string           `json:"id"`         //唯一标识
	Name      string           `json:"name"`       //名称
	SortIndex int32            `json:"sort_index"` //sort_index
	Status    int32            `json:"status"`     //status
	CreateAt  common.LocalTime `json:"create_at"`  //create_at
	UpdateAt  common.LocalTime `json:"update_at"`  //update_at
	Remark    string           `json:"remark"`     //remark

}

var RoleTableInfo = &TableInfo{
	Name: "o_role",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `唯一标识`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
			Comment:            `status`,
			Notes:              ``,
			Nullable:           false,
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
			Nullable:           false,
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
			Nullable:           false,
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
			Nullable:           false,
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
	},
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "o_role"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Role) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Role) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Role) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Role) TableInfo() *TableInfo {
	return RoleTableInfo
}
