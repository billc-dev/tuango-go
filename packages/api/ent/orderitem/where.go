// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/billc-dev/tuango-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldID, id))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldOrderID, v))
}

// PostItemID applies equality check predicate on the "post_item_id" field. It's identical to PostItemIDEQ.
func PostItemID(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldPostItemID, v))
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldIdentifier, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldName, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldPrice, v))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldQty, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldLocation, v))
}

// HasName applies equality check predicate on the "has_name" field. It's identical to HasNameEQ.
func HasName(v bool) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldHasName, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldOrderID, v))
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldOrderID, v))
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldOrderID, v))
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldOrderID, v))
}

// OrderIDContains applies the Contains predicate on the "order_id" field.
func OrderIDContains(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContains(FieldOrderID, v))
}

// OrderIDHasPrefix applies the HasPrefix predicate on the "order_id" field.
func OrderIDHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasPrefix(FieldOrderID, v))
}

// OrderIDHasSuffix applies the HasSuffix predicate on the "order_id" field.
func OrderIDHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasSuffix(FieldOrderID, v))
}

// OrderIDEqualFold applies the EqualFold predicate on the "order_id" field.
func OrderIDEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldOrderID, v))
}

// OrderIDContainsFold applies the ContainsFold predicate on the "order_id" field.
func OrderIDContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldOrderID, v))
}

// PostItemIDEQ applies the EQ predicate on the "post_item_id" field.
func PostItemIDEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldPostItemID, v))
}

// PostItemIDNEQ applies the NEQ predicate on the "post_item_id" field.
func PostItemIDNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldPostItemID, v))
}

// PostItemIDIn applies the In predicate on the "post_item_id" field.
func PostItemIDIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldPostItemID, vs...))
}

// PostItemIDNotIn applies the NotIn predicate on the "post_item_id" field.
func PostItemIDNotIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldPostItemID, vs...))
}

// PostItemIDGT applies the GT predicate on the "post_item_id" field.
func PostItemIDGT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldPostItemID, v))
}

// PostItemIDGTE applies the GTE predicate on the "post_item_id" field.
func PostItemIDGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldPostItemID, v))
}

// PostItemIDLT applies the LT predicate on the "post_item_id" field.
func PostItemIDLT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldPostItemID, v))
}

// PostItemIDLTE applies the LTE predicate on the "post_item_id" field.
func PostItemIDLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldPostItemID, v))
}

// PostItemIDContains applies the Contains predicate on the "post_item_id" field.
func PostItemIDContains(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContains(FieldPostItemID, v))
}

// PostItemIDHasPrefix applies the HasPrefix predicate on the "post_item_id" field.
func PostItemIDHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasPrefix(FieldPostItemID, v))
}

// PostItemIDHasSuffix applies the HasSuffix predicate on the "post_item_id" field.
func PostItemIDHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasSuffix(FieldPostItemID, v))
}

// PostItemIDEqualFold applies the EqualFold predicate on the "post_item_id" field.
func PostItemIDEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldPostItemID, v))
}

// PostItemIDContainsFold applies the ContainsFold predicate on the "post_item_id" field.
func PostItemIDContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldPostItemID, v))
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldIdentifier, v))
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContains(FieldIdentifier, v))
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasPrefix(FieldIdentifier, v))
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasSuffix(FieldIdentifier, v))
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldIdentifier, v))
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldIdentifier, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldName, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldPrice, v))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldQty, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldLocation, v))
}

// HasNameEQ applies the EQ predicate on the "has_name" field.
func HasNameEQ(v bool) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldHasName, v))
}

// HasNameNEQ applies the NEQ predicate on the "has_name" field.
func HasNameNEQ(v bool) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldHasName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldStatus, vs...))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPostItem applies the HasEdge predicate on the "post_item" edge.
func HasPostItem() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostItemTable, PostItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostItemWith applies the HasEdge predicate on the "post_item" edge with a given conditions (other predicates).
func HasPostItemWith(preds ...predicate.PostItem) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newPostItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.NotPredicates(p))
}
