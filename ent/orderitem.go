// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/billc-dev/tuango-go/ent/postitem"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID *string `json:"order_id,omitempty"`
	// PostItemID holds the value of the "post_item_id" field.
	PostItemID *string `json:"post_item_id,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier *string `json:"identifier,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Price holds the value of the "price" field.
	Price *float64 `json:"price,omitempty"`
	// Qty holds the value of the "qty" field.
	Qty *float64 `json:"qty,omitempty"`
	// Location holds the value of the "location" field.
	Location *string `json:"location,omitempty"`
	// HasName holds the value of the "has_name" field.
	HasName *bool `json:"has_name,omitempty"`
	// Status holds the value of the "status" field.
	Status *orderitem.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges        OrderItemEdges `json:"-"`
	selectValues sql.SelectValues
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// PostItem holds the value of the post_item edge.
	PostItem *PostItem `json:"post_item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// PostItemOrErr returns the PostItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) PostItemOrErr() (*PostItem, error) {
	if e.loadedTypes[1] {
		if e.PostItem == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: postitem.Label}
		}
		return e.PostItem, nil
	}
	return nil, &NotLoadedError{edge: "post_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldHasName:
			values[i] = new(sql.NullBool)
		case orderitem.FieldPrice, orderitem.FieldQty:
			values[i] = new(sql.NullFloat64)
		case orderitem.FieldID, orderitem.FieldOrderID, orderitem.FieldPostItemID, orderitem.FieldIdentifier, orderitem.FieldName, orderitem.FieldLocation, orderitem.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oi.ID = value.String
			}
		case orderitem.FieldOrderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oi.OrderID = new(string)
				*oi.OrderID = value.String
			}
		case orderitem.FieldPostItemID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_item_id", values[i])
			} else if value.Valid {
				oi.PostItemID = new(string)
				*oi.PostItemID = value.String
			}
		case orderitem.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				oi.Identifier = new(string)
				*oi.Identifier = value.String
			}
		case orderitem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				oi.Name = new(string)
				*oi.Name = value.String
			}
		case orderitem.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				oi.Price = new(float64)
				*oi.Price = value.Float64
			}
		case orderitem.FieldQty:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value.Valid {
				oi.Qty = new(float64)
				*oi.Qty = value.Float64
			}
		case orderitem.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				oi.Location = new(string)
				*oi.Location = value.String
			}
		case orderitem.FieldHasName:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_name", values[i])
			} else if value.Valid {
				oi.HasName = new(bool)
				*oi.HasName = value.Bool
			}
		case orderitem.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				oi.Status = new(orderitem.Status)
				*oi.Status = orderitem.Status(value.String)
			}
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderItem.
// This includes values selected through modifiers, order, etc.
func (oi *OrderItem) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrder() *OrderQuery {
	return NewOrderItemClient(oi.config).QueryOrder(oi)
}

// QueryPostItem queries the "post_item" edge of the OrderItem entity.
func (oi *OrderItem) QueryPostItem() *PostItemQuery {
	return NewOrderItemClient(oi.config).QueryPostItem(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return NewOrderItemClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	if v := oi.OrderID; v != nil {
		builder.WriteString("order_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.PostItemID; v != nil {
		builder.WriteString("post_item_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.Identifier; v != nil {
		builder.WriteString("identifier=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.Price; v != nil {
		builder.WriteString("price=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.Qty; v != nil {
		builder.WriteString("qty=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.Location; v != nil {
		builder.WriteString("location=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oi.HasName; v != nil {
		builder.WriteString("has_name=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := oi.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (oi *OrderItem) MarshalJSON() ([]byte, error) {
	type Alias OrderItem
	return json.Marshal(&struct {
		*Alias
		OrderItemEdges
	}{
		Alias:          (*Alias)(oi),
		OrderItemEdges: oi.Edges,
	})
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem
