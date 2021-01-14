// Code generated by entc, DO NOT EDIT.

package menu

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMenuname holds the string denoting the menuname field in the database.
	FieldMenuname = "menuname"
	// FieldIngredient holds the string denoting the ingredient field in the database.
	FieldIngredient = "ingredient"
	// FieldCalories holds the string denoting the calories field in the database.
	FieldCalories = "calories"
	// FieldAddedTime holds the string denoting the added_time field in the database.
	FieldAddedTime = "added_time"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"

	// Table holds the table name of the menu in the database.
	Table = "menus"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "menus"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// TypeTable is the table the holds the type relation/edge.
	TypeTable = "menus"
	// TypeInverseTable is the table name for the Menutype entity.
	// It exists in this package in order to avoid circular dependency with the "menutype" package.
	TypeInverseTable = "menutypes"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "type_id"
	// GroupTable is the table the holds the group relation/edge.
	GroupTable = "menus"
	// GroupInverseTable is the table name for the Menugroup entity.
	// It exists in this package in order to avoid circular dependency with the "menugroup" package.
	GroupInverseTable = "menugroups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_id"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldMenuname,
	FieldIngredient,
	FieldCalories,
	FieldAddedTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Menu type.
var ForeignKeys = []string{
	"group_id",
	"type_id",
	"owner_id",
}

var (
	// MenunameValidator is a validator for the "menuname" field. It is called by the builders before save.
	MenunameValidator func(string) error
	// CaloriesValidator is a validator for the "calories" field. It is called by the builders before save.
	CaloriesValidator func(int) error
)