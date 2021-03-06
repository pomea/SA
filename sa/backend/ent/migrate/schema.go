// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "menuname", Type: field.TypeString},
		{Name: "ingredient", Type: field.TypeString},
		{Name: "calories", Type: field.TypeInt},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "group_id", Type: field.TypeInt, Nullable: true},
		{Name: "type_id", Type: field.TypeInt, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "menus_menugroups_menus",
				Columns: []*schema.Column{MenusColumns[5]},

				RefColumns: []*schema.Column{MenugroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "menus_menutypes_menus",
				Columns: []*schema.Column{MenusColumns[6]},

				RefColumns: []*schema.Column{MenutypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "menus_users_menus",
				Columns: []*schema.Column{MenusColumns[7]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MenugroupsColumns holds the columns for the "menugroups" table.
	MenugroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group", Type: field.TypeString, Unique: true},
	}
	// MenugroupsTable holds the schema information for the "menugroups" table.
	MenugroupsTable = &schema.Table{
		Name:        "menugroups",
		Columns:     MenugroupsColumns,
		PrimaryKey:  []*schema.Column{MenugroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MenutypesColumns holds the columns for the "menutypes" table.
	MenutypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
	}
	// MenutypesTable holds the schema information for the "menutypes" table.
	MenutypesTable = &schema.Table{
		Name:        "menutypes",
		Columns:     MenutypesColumns,
		PrimaryKey:  []*schema.Column{MenutypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MenusTable,
		MenugroupsTable,
		MenutypesTable,
		UsersTable,
	}
)

func init() {
	MenusTable.ForeignKeys[0].RefTable = MenugroupsTable
	MenusTable.ForeignKeys[1].RefTable = MenutypesTable
	MenusTable.ForeignKeys[2].RefTable = UsersTable
}
