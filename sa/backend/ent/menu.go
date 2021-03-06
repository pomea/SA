// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pomea/app/ent/menu"
	"github.com/pomea/app/ent/menugroup"
	"github.com/pomea/app/ent/menutype"
	"github.com/pomea/app/ent/user"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Menuname holds the value of the "menuname" field.
	Menuname string `json:"menuname,omitempty"`
	// Ingredient holds the value of the "ingredient" field.
	Ingredient string `json:"ingredient,omitempty"`
	// Calories holds the value of the "calories" field.
	Calories int `json:"calories,omitempty"`
	// AddedTime holds the value of the "added_time" field.
	AddedTime time.Time `json:"added_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges    MenuEdges `json:"edges"`
	group_id *int
	type_id  *int
	owner_id *int
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Type holds the value of the type edge.
	Type *Menutype
	// Group holds the value of the group edge.
	Group *Menugroup
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TypeOrErr returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) TypeOrErr() (*Menutype, error) {
	if e.loadedTypes[1] {
		if e.Type == nil {
			// The edge type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: menutype.Label}
		}
		return e.Type, nil
	}
	return nil, &NotLoadedError{edge: "type"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) GroupOrErr() (*Menugroup, error) {
	if e.loadedTypes[2] {
		if e.Group == nil {
			// The edge group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: menugroup.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // menuname
		&sql.NullString{}, // ingredient
		&sql.NullInt64{},  // calories
		&sql.NullTime{},   // added_time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Menu) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // group_id
		&sql.NullInt64{}, // type_id
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(values ...interface{}) error {
	if m, n := len(values), len(menu.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field menuname", values[0])
	} else if value.Valid {
		m.Menuname = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ingredient", values[1])
	} else if value.Valid {
		m.Ingredient = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field calories", values[2])
	} else if value.Valid {
		m.Calories = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field added_time", values[3])
	} else if value.Valid {
		m.AddedTime = value.Time
	}
	values = values[4:]
	if len(values) == len(menu.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field group_id", value)
		} else if value.Valid {
			m.group_id = new(int)
			*m.group_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field type_id", value)
		} else if value.Valid {
			m.type_id = new(int)
			*m.type_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			m.owner_id = new(int)
			*m.owner_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Menu.
func (m *Menu) QueryOwner() *UserQuery {
	return (&MenuClient{config: m.config}).QueryOwner(m)
}

// QueryType queries the type edge of the Menu.
func (m *Menu) QueryType() *MenutypeQuery {
	return (&MenuClient{config: m.config}).QueryType(m)
}

// QueryGroup queries the group edge of the Menu.
func (m *Menu) QueryGroup() *MenugroupQuery {
	return (&MenuClient{config: m.config}).QueryGroup(m)
}

// Update returns a builder for updating this Menu.
// Note that, you need to call Menu.Unwrap() before calling this method, if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return (&MenuClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", menuname=")
	builder.WriteString(m.Menuname)
	builder.WriteString(", ingredient=")
	builder.WriteString(m.Ingredient)
	builder.WriteString(", calories=")
	builder.WriteString(fmt.Sprintf("%v", m.Calories))
	builder.WriteString(", added_time=")
	builder.WriteString(m.AddedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu

func (m Menus) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
