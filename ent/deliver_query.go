// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/billc-dev/tuango-go/ent/deliver"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/user"
)

// DeliverQuery is the builder for querying Deliver entities.
type DeliverQuery struct {
	config
	ctx        *QueryContext
	order      []deliver.OrderOption
	inters     []Interceptor
	predicates []predicate.Deliver
	withUser   *UserQuery
	withPost   *PostQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeliverQuery builder.
func (dq *DeliverQuery) Where(ps ...predicate.Deliver) *DeliverQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DeliverQuery) Limit(limit int) *DeliverQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DeliverQuery) Offset(offset int) *DeliverQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeliverQuery) Unique(unique bool) *DeliverQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DeliverQuery) Order(o ...deliver.OrderOption) *DeliverQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryUser chains the current query on the "user" edge.
func (dq *DeliverQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deliver.Table, deliver.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deliver.UserTable, deliver.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPost chains the current query on the "post" edge.
func (dq *DeliverQuery) QueryPost() *PostQuery {
	query := (&PostClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deliver.Table, deliver.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deliver.PostTable, deliver.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Deliver entity from the query.
// Returns a *NotFoundError when no Deliver was found.
func (dq *DeliverQuery) First(ctx context.Context) (*Deliver, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deliver.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeliverQuery) FirstX(ctx context.Context) *Deliver {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Deliver ID from the query.
// Returns a *NotFoundError when no Deliver ID was found.
func (dq *DeliverQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deliver.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeliverQuery) FirstIDX(ctx context.Context) string {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Deliver entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Deliver entity is found.
// Returns a *NotFoundError when no Deliver entities are found.
func (dq *DeliverQuery) Only(ctx context.Context) (*Deliver, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deliver.Label}
	default:
		return nil, &NotSingularError{deliver.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeliverQuery) OnlyX(ctx context.Context) *Deliver {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Deliver ID in the query.
// Returns a *NotSingularError when more than one Deliver ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeliverQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deliver.Label}
	default:
		err = &NotSingularError{deliver.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeliverQuery) OnlyIDX(ctx context.Context) string {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Delivers.
func (dq *DeliverQuery) All(ctx context.Context) ([]*Deliver, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Deliver, *DeliverQuery]()
	return withInterceptors[[]*Deliver](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeliverQuery) AllX(ctx context.Context) []*Deliver {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Deliver IDs.
func (dq *DeliverQuery) IDs(ctx context.Context) (ids []string, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(deliver.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeliverQuery) IDsX(ctx context.Context) []string {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeliverQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DeliverQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeliverQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeliverQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeliverQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeliverQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeliverQuery) Clone() *DeliverQuery {
	if dq == nil {
		return nil
	}
	return &DeliverQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]deliver.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Deliver{}, dq.predicates...),
		withUser:   dq.withUser.Clone(),
		withPost:   dq.withPost.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeliverQuery) WithUser(opts ...func(*UserQuery)) *DeliverQuery {
	query := (&UserClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withUser = query
	return dq
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeliverQuery) WithPost(opts ...func(*PostQuery)) *DeliverQuery {
	query := (&PostClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withPost = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Deliver.Query().
//		GroupBy(deliver.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DeliverQuery) GroupBy(field string, fields ...string) *DeliverGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeliverGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = deliver.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.Deliver.Query().
//		Select(deliver.FieldUserID).
//		Scan(ctx, &v)
func (dq *DeliverQuery) Select(fields ...string) *DeliverSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DeliverSelect{DeliverQuery: dq}
	sbuild.label = deliver.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeliverSelect configured with the given aggregations.
func (dq *DeliverQuery) Aggregate(fns ...AggregateFunc) *DeliverSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DeliverQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !deliver.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeliverQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Deliver, error) {
	var (
		nodes       = []*Deliver{}
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withUser != nil,
			dq.withPost != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Deliver).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Deliver{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withUser; query != nil {
		if err := dq.loadUser(ctx, query, nodes, nil,
			func(n *Deliver, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withPost; query != nil {
		if err := dq.loadPost(ctx, query, nodes, nil,
			func(n *Deliver, e *Post) { n.Edges.Post = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeliverQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Deliver, init func(*Deliver), assign func(*Deliver, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Deliver)
	for i := range nodes {
		if nodes[i].UserID == nil {
			continue
		}
		fk := *nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DeliverQuery) loadPost(ctx context.Context, query *PostQuery, nodes []*Deliver, init func(*Deliver), assign func(*Deliver, *Post)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Deliver)
	for i := range nodes {
		if nodes[i].PostID == nil {
			continue
		}
		fk := *nodes[i].PostID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DeliverQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeliverQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(deliver.Table, deliver.Columns, sqlgraph.NewFieldSpec(deliver.FieldID, field.TypeString))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deliver.FieldID)
		for i := range fields {
			if fields[i] != deliver.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dq.withUser != nil {
			_spec.Node.AddColumnOnce(deliver.FieldUserID)
		}
		if dq.withPost != nil {
			_spec.Node.AddColumnOnce(deliver.FieldPostID)
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeliverQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(deliver.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = deliver.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dq.modifiers {
		m(selector)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dq *DeliverQuery) Modify(modifiers ...func(s *sql.Selector)) *DeliverSelect {
	dq.modifiers = append(dq.modifiers, modifiers...)
	return dq.Select()
}

// DeliverGroupBy is the group-by builder for Deliver entities.
type DeliverGroupBy struct {
	selector
	build *DeliverQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeliverGroupBy) Aggregate(fns ...AggregateFunc) *DeliverGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DeliverGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeliverQuery, *DeliverGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DeliverGroupBy) sqlScan(ctx context.Context, root *DeliverQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeliverSelect is the builder for selecting fields of Deliver entities.
type DeliverSelect struct {
	*DeliverQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DeliverSelect) Aggregate(fns ...AggregateFunc) *DeliverSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeliverSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeliverQuery, *DeliverSelect](ctx, ds.DeliverQuery, ds, ds.inters, v)
}

func (ds *DeliverSelect) sqlScan(ctx context.Context, root *DeliverQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ds *DeliverSelect) Modify(modifiers ...func(s *sql.Selector)) *DeliverSelect {
	ds.modifiers = append(ds.modifiers, modifiers...)
	return ds
}