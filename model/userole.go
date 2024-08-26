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


Table: r_user_role
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] role_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "ACbtxtbOSLTYjienSpDtusNCg",    "user_id": "FNBFHfHMNLGpJQcUvZYxnDgHH",    "role_id": "yomndcCrhDNVNTlnqVYtlrnia"}



*/

var (
	Userole_FIELD_NAME_id = "id"

	Userole_FIELD_NAME_user_id = "user_id"

	Userole_FIELD_NAME_role_id = "role_id"
)

// Userole struct is a row record of the r_user_role table in the  database
type Userole struct {
	ID     string `json:"id"`      //唯一标识
	UserID string `json:"user_id"` //用户id
	RoleID string `json:"role_id"` //角色id

}

var UseroleTableInfo = &TableInfo{
	Name: "r_user_role",
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
			Name:               "user_id",
			Comment:            `用户id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *Userole) TableName() string {
	return "r_user_role"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Userole) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Userole) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Userole) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Userole) TableInfo() *TableInfo {
	return UseroleTableInfo
}
