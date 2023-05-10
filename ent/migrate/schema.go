// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CasbinRulesColumns holds the columns for the "casbin_rules" table.
	CasbinRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ptype", Type: field.TypeString, Size: 10, Default: ""},
		{Name: "v0", Type: field.TypeString, Size: 25, Default: ""},
		{Name: "v1", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "v2", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "v3", Type: field.TypeString, Size: 10, Default: ""},
		{Name: "v4", Type: field.TypeString, Size: 10, Default: ""},
		{Name: "v5", Type: field.TypeString, Size: 10, Default: ""},
	}
	// CasbinRulesTable holds the schema information for the "casbin_rules" table.
	CasbinRulesTable = &schema.Table{
		Name:       "casbin_rules",
		Columns:    CasbinRulesColumns,
		PrimaryKey: []*schema.Column{CasbinRulesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CasbinRulesTable,
	}
)

func init() {
}
