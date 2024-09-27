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


Table: v_role_rel_detail
[ 0] id                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] role_id                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] resource_id                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] op                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] filter_conditions                              TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] role_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] resource_name                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] module                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] service_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] service_name_cn                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] type                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[11] api_url                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] data_conditions                                VARCHAR(10240)       null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10240   default: []
[13] support_ops                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "kPiFikmVrZMVALlhkxnSmtBIM",    "role_id": "XhKeHRiWEqEgZEPgBMRydlmDu",    "resource_id": "kEHpeSHOYleNuKugyoULypGHV",    "op": "gdApLCBFAdJvVLXvTscmiiEtt",    "filter_conditions": "ROhHQXanUtQjfnxxgnGxTtjQS",    "role_name": "TYcPrvBfDRfAnMrcjNlWFIrbx",    "resource_name": "jKqWsYddrCDYFTZkuCNxQReQG",    "module": "cTxfbZhpRWTKWwmQwwRFkPEtJ",    "service_name": "rEtOIxTqOtDyhDPDvYLmClEvT",    "service_name_cn": "mOCTRJvdeXsCCZvfpRfWjmjtO",    "type": 63,    "api_url": "DGSZdVDyqEGSSSeFTIhefPxKD",    "data_conditions": "qdgcUyjlmopYDoHCkwsNZBaNn",    "support_ops": "AyeuYAGoTVpHDKOffUIfMKtwO"}


Comments
-------------------------------------
[ 0] Warning table: v_role_rel_detail does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_role_rel_detail primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Role_rel_detail_FIELD_NAME_id = "id"

	Role_rel_detail_FIELD_NAME_role_id = "role_id"

	Role_rel_detail_FIELD_NAME_resource_id = "resource_id"

	Role_rel_detail_FIELD_NAME_op = "op"

	Role_rel_detail_FIELD_NAME_filter_conditions = "filter_conditions"

	Role_rel_detail_FIELD_NAME_role_name = "role_name"

	Role_rel_detail_FIELD_NAME_resource_name = "resource_name"

	Role_rel_detail_FIELD_NAME_module = "module"

	Role_rel_detail_FIELD_NAME_service_name = "service_name"

	Role_rel_detail_FIELD_NAME_service_name_cn = "service_name_cn"

	Role_rel_detail_FIELD_NAME_type = "type"

	Role_rel_detail_FIELD_NAME_api_url = "api_url"

	Role_rel_detail_FIELD_NAME_data_conditions = "data_conditions"

	Role_rel_detail_FIELD_NAME_support_ops = "support_ops"
)

// Role_rel_detail struct is a row record of the v_role_rel_detail table in the  database
type Role_rel_detail struct {
	ID               string `json:"id"`                //id
	RoleID           string `json:"role_id"`           //role_id
	ResourceID       string `json:"resource_id"`       //resource_id
	Op               string `json:"op"`                //op
	FilterConditions string `json:"filter_conditions"` //filter_conditions
	RoleName         string `json:"role_name"`         //role_name
	ResourceName     string `json:"resource_name"`     //resource_name
	Module           string `json:"module"`            //module
	ServiceName      string `json:"service_name"`      //service_name
	ServiceNameCn    string `json:"service_name_cn"`   //service_name_cn
	Type             int32  `json:"type"`              //type
	APIURL           string `json:"api_url"`           //api_url
	DataConditions   string `json:"data_conditions"`   //data_conditions
	SupportOps       string `json:"support_ops"`       //support_ops

}

var Role_rel_detailTableInfo = &TableInfo{
	Name: "v_role_rel_detail",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_role_rel_detail does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_role_rel_detail primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "role_id",
			Comment:            `role_id`,
			Notes:              ``,
			Nullable:           true,
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
			Comment:            `resource_id`,
			Notes:              ``,
			Nullable:           true,
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
			Comment:            `op`,
			Notes:              ``,
			Nullable:           true,
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
			Comment:            `filter_conditions`,
			Notes:              ``,
			Nullable:           true,
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

		&ColumnInfo{
			Index:              5,
			Name:               "role_name",
			Comment:            `role_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoleName",
			GoFieldType:        "string",
			JSONFieldName:      "role_name",
			ProtobufFieldName:  "role_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "resource_name",
			Comment:            `resource_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ResourceName",
			GoFieldType:        "string",
			JSONFieldName:      "resource_name",
			ProtobufFieldName:  "resource_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "module",
			Comment:            `module`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "service_name",
			Comment:            `service_name`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "service_name_cn",
			Comment:            `service_name_cn`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "type",
			Comment:            `type`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "api_url",
			Comment:            `api_url`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "data_conditions",
			Comment:            `data_conditions`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "support_ops",
			Comment:            `support_ops`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Role_rel_detail) TableName() string {
	return "v_role_rel_detail"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Role_rel_detail) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Role_rel_detail) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Role_rel_detail) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Role_rel_detail) TableInfo() *TableInfo {
	return Role_rel_detailTableInfo
}
