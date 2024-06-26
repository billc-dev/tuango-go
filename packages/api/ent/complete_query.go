// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/billc-dev/tuango-go/ent/complete"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/user"
)

// CompleteQuery is the builder for querying Complete entities.
type CompleteQuery struct {
	config
	ctx        *QueryContext
	order      []complete.OrderOption
	inters     []Interceptor
	predicates []predicate.Complete
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompleteQuery builder.
func (cq *CompleteQuery) Where(ps ...predicate.Complete) *CompleteQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CompleteQuery) Limit(limit int) *CompleteQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CompleteQuery) Offset(offset int) *CompleteQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CompleteQuery) Unique(unique bool) *CompleteQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CompleteQuery) Order(o ...complete.OrderOption) *CompleteQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryUser chains the current query on the "user" edge.
func (cq *CompleteQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(complete.Table, complete.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, complete.UserTable, complete.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Complete entity from the query.
// Returns a *NotFoundError when no Complete was found.
func (cq *CompleteQuery) First(ctx context.Context) (*Complete, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{complete.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CompleteQuery) FirstX(ctx context.Context) *Complete {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Complete ID from the query.
// Returns a *NotFoundError when no Complete ID was found.
func (cq *CompleteQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{complete.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CompleteQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Complete entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Complete entity is found.
// Returns a *NotFoundError when no Complete entities are found.
func (cq *CompleteQuery) Only(ctx context.Context) (*Complete, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{complete.Label}
	default:
		return nil, &NotSingularError{complete.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CompleteQuery) OnlyX(ctx context.Context) *Complete {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Complete ID in the query.
// Returns a *NotSingularError when more than one Complete ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CompleteQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{complete.Label}
	default:
		err = &NotSingularError{complete.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CompleteQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Completes.
func (cq *CompleteQuery) All(ctx context.Context) ([]*Complete, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Complete, *CompleteQuery]()
	return withInterceptors[[]*Complete](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CompleteQuery) AllX(ctx context.Context) []*Complete {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Complete IDs.
func (cq *CompleteQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(complete.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CompleteQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CompleteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CompleteQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CompleteQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CompleteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CompleteQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompleteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CompleteQuery) Clone() *CompleteQuery {
	if cq == nil {
		return nil
	}
	return &CompleteQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]complete.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Complete{}, cq.predicates...),
		withUser:   cq.withUser.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompleteQuery) WithUser(opts ...func(*UserQuery)) *CompleteQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withUser = query
	return cq
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
//	client.Complete.Query().
//		GroupBy(complete.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CompleteQuery) GroupBy(field string, fields ...string) *CompleteGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompleteGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = complete.Label
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
//	client.Complete.Query().
//		Select(complete.FieldUserID).
//		Scan(ctx, &v)
func (cq *CompleteQuery) Select(fields ...string) *CompleteSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CompleteSelect{CompleteQuery: cq}
	sbuild.label = complete.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompleteSelect configured with the given aggregations.
func (cq *CompleteQuery) Aggregate(fns ...AggregateFunc) *CompleteSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CompleteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !complete.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CompleteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Complete, error) {
	var (
		nodes       = []*Complete{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Complete).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Complete{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withUser; query != nil {
		if err := cq.loadUser(ctx, query, nodes, nil,
			func(n *Complete, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CompleteQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Complete, init func(*Complete), assign func(*Complete, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Complete)
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

func (cq *CompleteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CompleteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(complete.Table, complete.Columns, sqlgraph.NewFieldSpec(complete.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, complete.FieldID)
		for i := range fields {
			if fields[i] != complete.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withUser != nil {
			_spec.Node.AddColumnOnce(complete.FieldUserID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CompleteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(complete.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = complete.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CompleteQuery) Modify(modifiers ...func(s *sql.Selector)) *CompleteSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CompleteGroupBy is the group-by builder for Complete entities.
type CompleteGroupBy struct {
	selector
	build *CompleteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CompleteGroupBy) Aggregate(fns ...AggregateFunc) *CompleteGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CompleteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompleteQuery, *CompleteGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CompleteGroupBy) sqlScan(ctx context.Context, root *CompleteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompleteSelect is the builder for selecting fields of Complete entities.
type CompleteSelect struct {
	*CompleteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CompleteSelect) Aggregate(fns ...AggregateFunc) *CompleteSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CompleteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompleteQuery, *CompleteSelect](ctx, cs.CompleteQuery, cs, cs.inters, v)
}

func (cs *CompleteSelect) sqlScan(ctx context.Context, root *CompleteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CompleteSelect) Modify(modifiers ...func(s *sql.Selector)) *CompleteSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
