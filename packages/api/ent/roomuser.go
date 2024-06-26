// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/ent/room"
	"github.com/billc-dev/tuango-go/ent/roomuser"
	"github.com/billc-dev/tuango-go/ent/user"
)

// RoomUser is the model entity for the RoomUser schema.
type RoomUser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// RoomID holds the value of the "room_id" field.
	RoomID *string `json:"room_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID *string `json:"user_id,omitempty"`
	// LastReadMessageID holds the value of the "last_read_message_id" field.
	LastReadMessageID *string `json:"last_read_message_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomUserQuery when eager-loading is set.
	Edges        RoomUserEdges `json:"-"`
	selectValues sql.SelectValues
}

// RoomUserEdges holds the relations/edges for other nodes in the graph.
type RoomUserEdges struct {
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomUserEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[0] {
		if e.Room == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roomuser.FieldID, roomuser.FieldRoomID, roomuser.FieldUserID, roomuser.FieldLastReadMessageID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomUser fields.
func (ru *RoomUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roomuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ru.ID = value.String
			}
		case roomuser.FieldRoomID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field room_id", values[i])
			} else if value.Valid {
				ru.RoomID = new(string)
				*ru.RoomID = value.String
			}
		case roomuser.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ru.UserID = new(string)
				*ru.UserID = value.String
			}
		case roomuser.FieldLastReadMessageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_read_message_id", values[i])
			} else if value.Valid {
				ru.LastReadMessageID = new(string)
				*ru.LastReadMessageID = value.String
			}
		default:
			ru.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoomUser.
// This includes values selected through modifiers, order, etc.
func (ru *RoomUser) Value(name string) (ent.Value, error) {
	return ru.selectValues.Get(name)
}

// QueryRoom queries the "room" edge of the RoomUser entity.
func (ru *RoomUser) QueryRoom() *RoomQuery {
	return NewRoomUserClient(ru.config).QueryRoom(ru)
}

// QueryUser queries the "user" edge of the RoomUser entity.
func (ru *RoomUser) QueryUser() *UserQuery {
	return NewRoomUserClient(ru.config).QueryUser(ru)
}

// Update returns a builder for updating this RoomUser.
// Note that you need to call RoomUser.Unwrap() before calling this method if this RoomUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ru *RoomUser) Update() *RoomUserUpdateOne {
	return NewRoomUserClient(ru.config).UpdateOne(ru)
}

// Unwrap unwraps the RoomUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ru *RoomUser) Unwrap() *RoomUser {
	_tx, ok := ru.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomUser is not a transactional entity")
	}
	ru.config.driver = _tx.drv
	return ru
}

// String implements the fmt.Stringer.
func (ru *RoomUser) String() string {
	var builder strings.Builder
	builder.WriteString("RoomUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ru.ID))
	if v := ru.RoomID; v != nil {
		builder.WriteString("room_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ru.UserID; v != nil {
		builder.WriteString("user_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ru.LastReadMessageID; v != nil {
		builder.WriteString("last_read_message_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ru *RoomUser) MarshalJSON() ([]byte, error) {
	type Alias RoomUser
	return json.Marshal(&struct {
		*Alias
		RoomUserEdges
	}{
		Alias:         (*Alias)(ru),
		RoomUserEdges: ru.Edges,
	})
}

// RoomUsers is a parsable slice of RoomUser.
type RoomUsers []*RoomUser
