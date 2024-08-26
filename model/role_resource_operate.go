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


Table: r_role_resource_operate
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] role_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] resource_id                                    VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] op                                             VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] filter_conditions                              TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "VxGAkiyuWwcqBEKsbECalVIbk",    "role_id": "xPyQknNLjrMZWwgjVQCDPOCWb",    "resource_id": "grGVDcBQrQkdmAmsSxDnQwhTl",    "op": "NwQPJLmlqRPONoqlKlDysuxSF",    "filter_conditions": "FSLlGcsDevZJmHOZtwxUOQpsg"}



*/

var (
	Role_resource_operate_FIELD_NAME_id = "id"

	Role_resource_operate_FIELD_NAME_role_id = "role_id"

	Role_resource_operate_FIELD_NAME_resource_id = "resource_id"

	Role_resource_operate_FIELD_NAME_op = "op"

	Role_resource_operate_FIELD_NAME_filter_conditions = "filter_conditions"
)

// Role_resource_operate struct is a row record of the r_role_resource_operate table in the  database
type Role_resource_operate struct {
	ID               string `json:"id"`                //唯一标识
	RoleID           string `json:"role_id"`           //角色id
	ResourceID       string `json:"resource_id"`       //资源id
	Op               string `json:"op"`                //操作（字典值）
	FilterConditions string `json:"filter_conditions"` //过滤条件json字符串

}

var Role_resource_operateTableInfo = &TableInfo{
	Name: "r_role_resource_operate",
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
			Name:               "role_id",
			Comment:            `角色id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoleID",
			GoFieldType:        "string",
			JSONFieldName:      "role_id",
			ProtobufFieldName:  "role_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "resource_id",
			Comment:            `资源id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ResourceID",
			GoFieldType:        "string",
			JSONFieldName:      "resource_id",
			ProtobufFieldName:  "resource_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "op",
			Comment:            `操作（字典值）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Op",
			GoFieldType:        "string",
			JSONFieldName:      "op",
			ProtobufFieldName:  "op",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "filter_conditions",
			Comment:            `过滤条件json字符串`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "FilterConditions",
			GoFieldType:        "string",
			JSONFieldName:      "filter_conditions",
			ProtobufFieldName:  "filter_conditions",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Role_resource_operate) TableName() string {
	return "r_role_resource_operate"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Role_resource_operate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Role_resource_operate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Role_resource_operate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Role_resource_operate) TableInfo() *TableInfo {
	return Role_resource_operateTableInfo
}
