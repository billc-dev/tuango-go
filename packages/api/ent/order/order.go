// Code generated by ent, DO NOT EDIT.

package order

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPostID holds the string denoting the post_id field in the database.
	FieldPostID = "post_id"
	// FieldOrderNum holds the string denoting the order_num field in the database.
	FieldOrderNum = "order_num"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldSellerComment holds the string denoting the seller_comment field in the database.
	FieldSellerComment = "seller_comment"
	// FieldHasName holds the string denoting the has_name field in the database.
	FieldHasName = "has_name"
	// FieldIsExtra holds the string denoting the is_extra field in the database.
	FieldIsExtra = "is_extra"
	// FieldFb holds the string denoting the fb field in the database.
	FieldFb = "fb"
	// FieldIsInStock holds the string denoting the is_in_stock field in the database.
	FieldIsInStock = "is_in_stock"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeOrderItems holds the string denoting the order_items edge name in mutations.
	EdgeOrderItems = "order_items"
	// EdgeOrderHistories holds the string denoting the order_histories edge name in mutations.
	EdgeOrderHistories = "order_histories"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "orders"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_id"
	// OrderItemsTable is the table that holds the order_items relation/edge.
	OrderItemsTable = "order_items"
	// OrderItemsInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemsInverseTable = "order_items"
	// OrderItemsColumn is the table column denoting the order_items relation/edge.
	OrderItemsColumn = "order_id"
	// OrderHistoriesTable is the table that holds the order_histories relation/edge.
	OrderHistoriesTable = "order_histories"
	// OrderHistoriesInverseTable is the table name for the OrderHistory entity.
	// It exists in this package in order to avoid circular dependency with the "orderhistory" package.
	OrderHistoriesInverseTable = "order_histories"
	// OrderHistoriesColumn is the table column denoting the order_histories relation/edge.
	OrderHistoriesColumn = "order_id"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldPostID,
	FieldOrderNum,
	FieldComment,
	FieldSellerComment,
	FieldHasName,
	FieldIsExtra,
	FieldFb,
	FieldIsInStock,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultComment holds the default value on creation for the "comment" field.
	DefaultComment string
	// DefaultSellerComment holds the default value on creation for the "seller_comment" field.
	DefaultSellerComment string
	// DefaultHasName holds the default value on creation for the "has_name" field.
	DefaultHasName bool
	// DefaultIsExtra holds the default value on creation for the "is_extra" field.
	DefaultIsExtra bool
	// DefaultFb holds the default value on creation for the "fb" field.
	DefaultFb bool
	// DefaultIsInStock holds the default value on creation for the "is_in_stock" field.
	DefaultIsInStock bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
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
		return fmt.Errorf("order: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPostID orders the results by the post_id field.
func ByPostID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostID, opts...).ToFunc()
}

// ByOrderNum orders the results by the order_num field.
func ByOrderNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderNum, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// BySellerComment orders the results by the seller_comment field.
func BySellerComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSellerComment, opts...).ToFunc()
}

// ByHasName orders the results by the has_name field.
func ByHasName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasName, opts...).ToFunc()
}

// ByIsExtra orders the results by the is_extra field.
func ByIsExtra(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsExtra, opts...).ToFunc()
}

// ByFb orders the results by the fb field.
func ByFb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFb, opts...).ToFunc()
}

// ByIsInStock orders the results by the is_in_stock field.
func ByIsInStock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsInStock, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostField orders the results by post field.
func ByPostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrderItemsCount orders the results by order_items count.
func ByOrderItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderItemsStep(), opts...)
	}
}

// ByOrderItems orders the results by order_items terms.
func ByOrderItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrderHistoriesCount orders the results by order_histories count.
func ByOrderHistoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderHistoriesStep(), opts...)
	}
}

// ByOrderHistories orders the results by order_histories terms.
func ByOrderHistories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderHistoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
	)
}
func newOrderItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrderItemsTable, OrderItemsColumn),
	)
}
func newOrderHistoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderHistoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrderHistoriesTable, OrderHistoriesColumn),
	)
}