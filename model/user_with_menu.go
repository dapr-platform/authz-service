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


Table: v_user_with_menu
[ 0] id                                             VARCHAR(25)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
[ 1] tenant_id                                      VARCHAR(25)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
[ 2] mobile                                         VARCHAR(15)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
[ 3] email                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] identity                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] gender                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] address                                        VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 8] password                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] type                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[10] org_id                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] avatar_url                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] create_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[13] update_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[14] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] menu_ids                                       JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "KLXfFgSGmkdyNsZDftEVvMXdw",    "tenant_id": "LPYDvFTUycwuIpAUcGETCIBeV",    "mobile": "HrGLqtKTomuRdTZBbZLJcFbpP",    "email": "YmXkFKIaXUtFZRdyBcVvMFFtO",    "identity": "fCxIUJHXdIDLksvVKhNQIghEs",    "name": "pylVkDAvfXlNXAQdPueJmYdya",    "gender": 47,    "address": "SglUxiWGiKAeyuqtYcDydIagf",    "password": "KGPepNdqBGXapCZkYAUdALDZP",    "type": 91,    "org_id": "BnnkcNpvcsnOOySqcQVAnwTxV",    "avatar_url": "IqcsyYpHRIBlloRUZwuXJKAaE",    "create_at": 22,    "update_at": 87,    "status": 76,    "menu_ids": 62}


Comments
-------------------------------------
[ 0] Warning table: v_user_with_menu does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_user_with_menu primary key column id is nullable column, setting it as NOT NULL




*/

var (
	User_with_menu_FIELD_NAME_id = "id"

	User_with_menu_FIELD_NAME_tenant_id = "tenant_id"

	User_with_menu_FIELD_NAME_mobile = "mobile"

	User_with_menu_FIELD_NAME_email = "email"

	User_with_menu_FIELD_NAME_identity = "identity"

	User_with_menu_FIELD_NAME_name = "name"

	User_with_menu_FIELD_NAME_gender = "gender"

	User_with_menu_FIELD_NAME_address = "address"

	User_with_menu_FIELD_NAME_password = "password"

	User_with_menu_FIELD_NAME_type = "type"

	User_with_menu_FIELD_NAME_org_id = "org_id"

	User_with_menu_FIELD_NAME_avatar_url = "avatar_url"

	User_with_menu_FIELD_NAME_create_at = "create_at"

	User_with_menu_FIELD_NAME_update_at = "update_at"

	User_with_menu_FIELD_NAME_status = "status"

	User_with_menu_FIELD_NAME_menu_ids = "menu_ids"
)

// User_with_menu struct is a row record of the v_user_with_menu table in the  database
type User_with_menu struct {
	ID        string           `json:"id"`         //id
	TenantID  string           `json:"tenant_id"`  //tenant_id
	Mobile    string           `json:"mobile"`     //mobile
	Email     string           `json:"email"`      //email
	Identity  string           `json:"identity"`   //identity
	Name      string           `json:"name"`       //name
	Gender    int32            `json:"gender"`     //gender
	Address   string           `json:"address"`    //address
	Password  string           `json:"password"`   //password
	Type      int32            `json:"type"`       //type
	OrgID     string           `json:"org_id"`     //org_id
	AvatarURL string           `json:"avatar_url"` //avatar_url
	CreateAt  common.LocalTime `json:"create_at"`  //create_at
	UpdateAt  common.LocalTime `json:"update_at"`  //update_at
	Status    int32            `json:"status"`     //status
	MenuIds   any              `json:"menu_ids"`   //menu_ids

}

var User_with_menuTableInfo = &TableInfo{
	Name: "v_user_with_menu",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_user_with_menu does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_user_with_menu primary key column id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "tenant_id",
			Comment:            `tenant_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "TenantID",
			GoFieldType:        "string",
			JSONFieldName:      "tenant_id",
			ProtobufFieldName:  "tenant_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "mobile",
			Comment:            `mobile`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       15,
			GoFieldName:        "Mobile",
			GoFieldType:        "string",
			JSONFieldName:      "mobile",
			ProtobufFieldName:  "mobile",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "email",
			Comment:            `email`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "identity",
			Comment:            `identity`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Identity",
			GoFieldType:        "string",
			JSONFieldName:      "identity",
			ProtobufFieldName:  "identity",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "gender",
			Comment:            `gender`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Gender",
			GoFieldType:        "int32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "address",
			Comment:            `address`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "Address",
			GoFieldType:        "string",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "password",
			Comment:            `password`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "org_id",
			Comment:            `org_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "OrgID",
			GoFieldType:        "string",
			JSONFieldName:      "org_id",
			ProtobufFieldName:  "org_id",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "avatar_url",
			Comment:            `avatar_url`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AvatarURL",
			GoFieldType:        "string",
			JSONFieldName:      "avatar_url",
			ProtobufFieldName:  "avatar_url",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "menu_ids",
			Comment:            `menu_ids`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "MenuIds",
			GoFieldType:        "any",
			JSONFieldName:      "menu_ids",
			ProtobufFieldName:  "menu_ids",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *User_with_menu) TableName() string {
	return "v_user_with_menu"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *User_with_menu) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *User_with_menu) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *User_with_menu) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *User_with_menu) TableInfo() *TableInfo {
	return User_with_menuTableInfo
}
