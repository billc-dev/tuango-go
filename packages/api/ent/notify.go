// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/ent/notify"
)

// Notify is the model entity for the Notify schema.
type Notify struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID *string `json:"user_id,omitempty"`
	// LineToken holds the value of the "line_token" field.
	LineToken *string `json:"line_token,omitempty"`
	// FbToken holds the value of the "fb_token" field.
	FbToken *string `json:"fb_token,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notify) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notify.FieldID, notify.FieldUserID, notify.FieldLineToken, notify.FieldFbToken:
			values[i] = new(sql.NullString)
		case notify.FieldCreatedAt, notify.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notify fields.
func (n *Notify) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notify.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				n.ID = value.String
			}
		case notify.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				n.UserID = new(string)
				*n.UserID = value.String
			}
		case notify.FieldLineToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field line_token", values[i])
			} else if value.Valid {
				n.LineToken = new(string)
				*n.LineToken = value.String
			}
		case notify.FieldFbToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fb_token", values[i])
			} else if value.Valid {
				n.FbToken = new(string)
				*n.FbToken = value.String
			}
		case notify.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = new(time.Time)
				*n.CreatedAt = value.Time
			}
		case notify.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = new(time.Time)
				*n.UpdatedAt = value.Time
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Notify.
// This includes values selected through modifiers, order, etc.
func (n *Notify) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// Update returns a builder for updating this Notify.
// Note that you need to call Notify.Unwrap() before calling this method if this Notify
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notify) Update() *NotifyUpdateOne {
	return NewNotifyClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Notify entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notify) Unwrap() *Notify {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Notify is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notify) String() string {
	var builder strings.Builder
	builder.WriteString("Notify(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	if v := n.UserID; v != nil {
		builder.WriteString("user_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := n.LineToken; v != nil {
		builder.WriteString("line_token=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := n.FbToken; v != nil {
		builder.WriteString("fb_token=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := n.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := n.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Notifies is a parsable slice of Notify.
type Notifies []*Notify