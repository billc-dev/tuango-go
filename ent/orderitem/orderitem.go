// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orderitem type in the database.
	Label = "order_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldPostItemID holds the string denoting the post_item_id field in the database.
	FieldPostItemID = "post_item_id"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldQty holds the string denoting the qty field in the database.
	FieldQty = "qty"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldHasName holds the string denoting the has_name field in the database.
	FieldHasName = "has_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// EdgePostItem holds the string denoting the post_item edge name in mutations.
	EdgePostItem = "post_item"
	// Table holds the table name of the orderitem in the database.
	Table = "order_items"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "order_items"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_id"
	// PostItemTable is the table that holds the post_item relation/edge.
	PostItemTable = "order_items"
	// PostItemInverseTable is the table name for the PostItem entity.
	// It exists in this package in order to avoid circular dependency with the "postitem" package.
	PostItemInverseTable = "post_items"
	// PostItemColumn is the table column denoting the post_item relation/edge.
	PostItemColumn = "post_item_id"
)

// Columns holds all SQL columns for orderitem fields.
var Columns = []string{
	FieldID,
	FieldOrderID,
	FieldPostItemID,
	FieldIdentifier,
	FieldName,
	FieldPrice,
	FieldQty,
	FieldLocation,
	FieldHasName,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLocation holds the default value on creation for the "location" field.
	DefaultLocation string
	// DefaultHasName holds the default value on creation for the "has_name" field.
	DefaultHasName bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusOrdered   Status = "ordered"
	StatusConfirmed Status = "confirmed"
	StatusDelivered Status = "delivered"
	StatusCompleted Status = "completed"
	StatusMissing   Status = "missing"
	StatusCanceled  Status = "canceled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOrdered, StatusConfirmed, StatusDelivered, StatusCompleted, StatusMissing, StatusCanceled:
		return nil
	default:
		return fmt.Errorf("orderitem: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the OrderItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByPostItemID orders the results by the post_item_id field.
func ByPostItemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostItemID, opts...).ToFunc()
}

// ByIdentifier orders the results by the identifier field.
func ByIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentifier, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByQty orders the results by the qty field.
func ByQty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQty, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByHasName orders the results by the has_name field.
func ByHasName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostItemField orders the results by post_item field.
func ByPostItemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostItemStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
	)
}
func newPostItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostItemTable, PostItemColumn),
	)
}