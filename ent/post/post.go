// Code generated by ent, DO NOT EDIT.

package post

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSellerID holds the string denoting the seller_id field in the database.
	FieldSellerID = "seller_id"
	// FieldPostNum holds the string denoting the post_num field in the database.
	FieldPostNum = "post_num"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldDeliveryDate holds the string denoting the delivery_date field in the database.
	FieldDeliveryDate = "delivery_date"
	// FieldLikeCount holds the string denoting the like_count field in the database.
	FieldLikeCount = "like_count"
	// FieldCommentCount holds the string denoting the comment_count field in the database.
	FieldCommentCount = "comment_count"
	// FieldOrderCount holds the string denoting the order_count field in the database.
	FieldOrderCount = "order_count"
	// FieldImages holds the string denoting the images field in the database.
	FieldImages = "images"
	// FieldStorageType holds the string denoting the storage_type field in the database.
	FieldStorageType = "storage_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldDelivered holds the string denoting the delivered field in the database.
	FieldDelivered = "delivered"
	// FieldIsInStock holds the string denoting the is_in_stock field in the database.
	FieldIsInStock = "is_in_stock"
	// FieldNormalTotal holds the string denoting the normal_total field in the database.
	FieldNormalTotal = "normal_total"
	// FieldNormalFee holds the string denoting the normal_fee field in the database.
	FieldNormalFee = "normal_fee"
	// FieldExtraTotal holds the string denoting the extra_total field in the database.
	FieldExtraTotal = "extra_total"
	// FieldExtraFee holds the string denoting the extra_fee field in the database.
	FieldExtraFee = "extra_fee"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeSeller holds the string denoting the seller edge name in mutations.
	EdgeSeller = "seller"
	// EdgePostComments holds the string denoting the post_comments edge name in mutations.
	EdgePostComments = "post_comments"
	// EdgePostDelivers holds the string denoting the post_delivers edge name in mutations.
	EdgePostDelivers = "post_delivers"
	// EdgePostItems holds the string denoting the post_items edge name in mutations.
	EdgePostItems = "post_items"
	// EdgePostLikes holds the string denoting the post_likes edge name in mutations.
	EdgePostLikes = "post_likes"
	// EdgePostOrders holds the string denoting the post_orders edge name in mutations.
	EdgePostOrders = "post_orders"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// SellerTable is the table that holds the seller relation/edge.
	SellerTable = "posts"
	// SellerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SellerInverseTable = "users"
	// SellerColumn is the table column denoting the seller relation/edge.
	SellerColumn = "seller_id"
	// PostCommentsTable is the table that holds the post_comments relation/edge.
	PostCommentsTable = "comments"
	// PostCommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	PostCommentsInverseTable = "comments"
	// PostCommentsColumn is the table column denoting the post_comments relation/edge.
	PostCommentsColumn = "post_id"
	// PostDeliversTable is the table that holds the post_delivers relation/edge.
	PostDeliversTable = "delivers"
	// PostDeliversInverseTable is the table name for the Deliver entity.
	// It exists in this package in order to avoid circular dependency with the "deliver" package.
	PostDeliversInverseTable = "delivers"
	// PostDeliversColumn is the table column denoting the post_delivers relation/edge.
	PostDeliversColumn = "post_id"
	// PostItemsTable is the table that holds the post_items relation/edge.
	PostItemsTable = "post_items"
	// PostItemsInverseTable is the table name for the PostItem entity.
	// It exists in this package in order to avoid circular dependency with the "postitem" package.
	PostItemsInverseTable = "post_items"
	// PostItemsColumn is the table column denoting the post_items relation/edge.
	PostItemsColumn = "post_id"
	// PostLikesTable is the table that holds the post_likes relation/edge.
	PostLikesTable = "likes"
	// PostLikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	PostLikesInverseTable = "likes"
	// PostLikesColumn is the table column denoting the post_likes relation/edge.
	PostLikesColumn = "post_id"
	// PostOrdersTable is the table that holds the post_orders relation/edge.
	PostOrdersTable = "orders"
	// PostOrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	PostOrdersInverseTable = "orders"
	// PostOrdersColumn is the table column denoting the post_orders relation/edge.
	PostOrdersColumn = "post_id"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldSellerID,
	FieldPostNum,
	FieldTitle,
	FieldBody,
	FieldDeadline,
	FieldDeliveryDate,
	FieldLikeCount,
	FieldCommentCount,
	FieldOrderCount,
	FieldImages,
	FieldStorageType,
	FieldStatus,
	FieldComment,
	FieldDelivered,
	FieldIsInStock,
	FieldNormalTotal,
	FieldNormalFee,
	FieldExtraTotal,
	FieldExtraFee,
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
	// DefaultDeadline holds the default value on creation for the "deadline" field.
	DefaultDeadline string
	// DeadlineValidator is a validator for the "deadline" field. It is called by the builders before save.
	DeadlineValidator func(string) error
	// DefaultDeliveryDate holds the default value on creation for the "delivery_date" field.
	DefaultDeliveryDate string
	// DeliveryDateValidator is a validator for the "delivery_date" field. It is called by the builders before save.
	DeliveryDateValidator func(string) error
	// DefaultLikeCount holds the default value on creation for the "like_count" field.
	DefaultLikeCount int
	// DefaultCommentCount holds the default value on creation for the "comment_count" field.
	DefaultCommentCount int
	// DefaultOrderCount holds the default value on creation for the "order_count" field.
	DefaultOrderCount int
	// DefaultComment holds the default value on creation for the "comment" field.
	DefaultComment string
	// DefaultDelivered holds the default value on creation for the "delivered" field.
	DefaultDelivered bool
	// DefaultIsInStock holds the default value on creation for the "is_in_stock" field.
	DefaultIsInStock bool
	// DefaultNormalTotal holds the default value on creation for the "normal_total" field.
	DefaultNormalTotal float64
	// DefaultNormalFee holds the default value on creation for the "normal_fee" field.
	DefaultNormalFee float64
	// DefaultExtraTotal holds the default value on creation for the "extra_total" field.
	DefaultExtraTotal float64
	// DefaultExtraFee holds the default value on creation for the "extra_fee" field.
	DefaultExtraFee float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// StorageType defines the type for the "storage_type" enum field.
type StorageType string

// StorageType values.
const (
	StorageTypeRoomTemp     StorageType = "roomTemp"
	StorageTypeFarmGoods    StorageType = "farmGoods"
	StorageTypeRefrigerated StorageType = "refrigerated"
	StorageTypeFrozen       StorageType = "frozen"
)

func (st StorageType) String() string {
	return string(st)
}

// StorageTypeValidator is a validator for the "storage_type" field enum values. It is called by the builders before save.
func StorageTypeValidator(st StorageType) error {
	switch st {
	case StorageTypeRoomTemp, StorageTypeFarmGoods, StorageTypeRefrigerated, StorageTypeFrozen:
		return nil
	default:
		return fmt.Errorf("post: invalid enum value for storage_type field: %q", st)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusOpen is the default value of the Status enum.
const DefaultStatus = StatusOpen

// Status values.
const (
	StatusOpen      Status = "open"
	StatusClosed    Status = "closed"
	StatusCompleted Status = "completed"
	StatusCanceled  Status = "canceled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOpen, StatusClosed, StatusCompleted, StatusCanceled:
		return nil
	default:
		return fmt.Errorf("post: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Post queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySellerID orders the results by the seller_id field.
func BySellerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSellerID, opts...).ToFunc()
}

// ByPostNum orders the results by the post_num field.
func ByPostNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostNum, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByDeadline orders the results by the deadline field.
func ByDeadline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeadline, opts...).ToFunc()
}

// ByDeliveryDate orders the results by the delivery_date field.
func ByDeliveryDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeliveryDate, opts...).ToFunc()
}

// ByLikeCount orders the results by the like_count field.
func ByLikeCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikeCount, opts...).ToFunc()
}

// ByCommentCount orders the results by the comment_count field.
func ByCommentCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentCount, opts...).ToFunc()
}

// ByOrderCount orders the results by the order_count field.
func ByOrderCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderCount, opts...).ToFunc()
}

// ByStorageType orders the results by the storage_type field.
func ByStorageType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStorageType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByDelivered orders the results by the delivered field.
func ByDelivered(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelivered, opts...).ToFunc()
}

// ByIsInStock orders the results by the is_in_stock field.
func ByIsInStock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsInStock, opts...).ToFunc()
}

// ByNormalTotal orders the results by the normal_total field.
func ByNormalTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalTotal, opts...).ToFunc()
}

// ByNormalFee orders the results by the normal_fee field.
func ByNormalFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalFee, opts...).ToFunc()
}

// ByExtraTotal orders the results by the extra_total field.
func ByExtraTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraTotal, opts...).ToFunc()
}

// ByExtraFee orders the results by the extra_fee field.
func ByExtraFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraFee, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySellerField orders the results by seller field.
func BySellerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSellerStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostCommentsCount orders the results by post_comments count.
func ByPostCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostCommentsStep(), opts...)
	}
}

// ByPostComments orders the results by post_comments terms.
func ByPostComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostDeliversCount orders the results by post_delivers count.
func ByPostDeliversCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostDeliversStep(), opts...)
	}
}

// ByPostDelivers orders the results by post_delivers terms.
func ByPostDelivers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostDeliversStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostItemsCount orders the results by post_items count.
func ByPostItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostItemsStep(), opts...)
	}
}

// ByPostItems orders the results by post_items terms.
func ByPostItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostLikesCount orders the results by post_likes count.
func ByPostLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostLikesStep(), opts...)
	}
}

// ByPostLikes orders the results by post_likes terms.
func ByPostLikes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostLikesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostOrdersCount orders the results by post_orders count.
func ByPostOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostOrdersStep(), opts...)
	}
}

// ByPostOrders orders the results by post_orders terms.
func ByPostOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSellerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SellerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SellerTable, SellerColumn),
	)
}
func newPostCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostCommentsTable, PostCommentsColumn),
	)
}
func newPostDeliversStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostDeliversInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostDeliversTable, PostDeliversColumn),
	)
}
func newPostItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostItemsTable, PostItemsColumn),
	)
}
func newPostLikesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostLikesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostLikesTable, PostLikesColumn),
	)
}
func newPostOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostOrdersTable, PostOrdersColumn),
	)
}