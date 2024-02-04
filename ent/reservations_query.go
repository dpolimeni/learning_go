// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dpolimeni/fiber_app/ent/events"
	"github.com/dpolimeni/fiber_app/ent/predicate"
	"github.com/dpolimeni/fiber_app/ent/reservations"
	"github.com/dpolimeni/fiber_app/ent/user"
)

// ReservationsQuery is the builder for querying Reservations entities.
type ReservationsQuery struct {
	config
	ctx        *QueryContext
	order      []reservations.OrderOption
	inters     []Interceptor
	predicates []predicate.Reservations
	withUser   *UserQuery
	withEvent  *EventsQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReservationsQuery builder.
func (rq *ReservationsQuery) Where(ps ...predicate.Reservations) *ReservationsQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReservationsQuery) Limit(limit int) *ReservationsQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReservationsQuery) Offset(offset int) *ReservationsQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReservationsQuery) Unique(unique bool) *ReservationsQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReservationsQuery) Order(o ...reservations.OrderOption) *ReservationsQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUser chains the current query on the "user" edge.
func (rq *ReservationsQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reservations.Table, reservations.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reservations.UserTable, reservations.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvent chains the current query on the "event" edge.
func (rq *ReservationsQuery) QueryEvent() *EventsQuery {
	query := (&EventsClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reservations.Table, reservations.FieldID, selector),
			sqlgraph.To(events.Table, events.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reservations.EventTable, reservations.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Reservations entity from the query.
// Returns a *NotFoundError when no Reservations was found.
func (rq *ReservationsQuery) First(ctx context.Context) (*Reservations, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reservations.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReservationsQuery) FirstX(ctx context.Context) *Reservations {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Reservations ID from the query.
// Returns a *NotFoundError when no Reservations ID was found.
func (rq *ReservationsQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reservations.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReservationsQuery) FirstIDX(ctx context.Context) string {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Reservations entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Reservations entity is found.
// Returns a *NotFoundError when no Reservations entities are found.
func (rq *ReservationsQuery) Only(ctx context.Context) (*Reservations, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reservations.Label}
	default:
		return nil, &NotSingularError{reservations.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReservationsQuery) OnlyX(ctx context.Context) *Reservations {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Reservations ID in the query.
// Returns a *NotSingularError when more than one Reservations ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReservationsQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reservations.Label}
	default:
		err = &NotSingularError{reservations.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReservationsQuery) OnlyIDX(ctx context.Context) string {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReservationsSlice.
func (rq *ReservationsQuery) All(ctx context.Context) ([]*Reservations, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Reservations, *ReservationsQuery]()
	return withInterceptors[[]*Reservations](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReservationsQuery) AllX(ctx context.Context) []*Reservations {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Reservations IDs.
func (rq *ReservationsQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(reservations.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReservationsQuery) IDsX(ctx context.Context) []string {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReservationsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReservationsQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReservationsQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReservationsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReservationsQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReservationsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReservationsQuery) Clone() *ReservationsQuery {
	if rq == nil {
		return nil
	}
	return &ReservationsQuery{
		config:     rq.config,
		ctx:        rq.ctx.Clone(),
		order:      append([]reservations.OrderOption{}, rq.order...),
		inters:     append([]Interceptor{}, rq.inters...),
		predicates: append([]predicate.Reservations{}, rq.predicates...),
		withUser:   rq.withUser.Clone(),
		withEvent:  rq.withEvent.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReservationsQuery) WithUser(opts ...func(*UserQuery)) *ReservationsQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReservationsQuery) WithEvent(opts ...func(*EventsQuery)) *ReservationsQuery {
	query := (&EventsClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withEvent = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Reservations.Query().
//		GroupBy(reservations.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReservationsQuery) GroupBy(field string, fields ...string) *ReservationsGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReservationsGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = reservations.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Reservations.Query().
//		Select(reservations.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *ReservationsQuery) Select(fields ...string) *ReservationsSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReservationsSelect{ReservationsQuery: rq}
	sbuild.label = reservations.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReservationsSelect configured with the given aggregations.
func (rq *ReservationsQuery) Aggregate(fns ...AggregateFunc) *ReservationsSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReservationsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !reservations.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReservationsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Reservations, error) {
	var (
		nodes       = []*Reservations{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withUser != nil,
			rq.withEvent != nil,
		}
	)
	if rq.withUser != nil || rq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, reservations.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Reservations).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Reservations{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Reservations, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withEvent; query != nil {
		if err := rq.loadEvent(ctx, query, nodes, nil,
			func(n *Reservations, e *Events) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReservationsQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Reservations, init func(*Reservations), assign func(*Reservations, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Reservations)
	for i := range nodes {
		if nodes[i].user_reservations == nil {
			continue
		}
		fk := *nodes[i].user_reservations
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
			return fmt.Errorf(`unexpected foreign-key "user_reservations" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReservationsQuery) loadEvent(ctx context.Context, query *EventsQuery, nodes []*Reservations, init func(*Reservations), assign func(*Reservations, *Events)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*Reservations)
	for i := range nodes {
		if nodes[i].events_reservations == nil {
			continue
		}
		fk := *nodes[i].events_reservations
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(events.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "events_reservations" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *ReservationsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReservationsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(reservations.Table, reservations.Columns, sqlgraph.NewFieldSpec(reservations.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reservations.FieldID)
		for i := range fields {
			if fields[i] != reservations.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReservationsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(reservations.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = reservations.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReservationsGroupBy is the group-by builder for Reservations entities.
type ReservationsGroupBy struct {
	selector
	build *ReservationsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReservationsGroupBy) Aggregate(fns ...AggregateFunc) *ReservationsGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReservationsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReservationsQuery, *ReservationsGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReservationsGroupBy) sqlScan(ctx context.Context, root *ReservationsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReservationsSelect is the builder for selecting fields of Reservations entities.
type ReservationsSelect struct {
	*ReservationsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReservationsSelect) Aggregate(fns ...AggregateFunc) *ReservationsSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReservationsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReservationsQuery, *ReservationsSelect](ctx, rs.ReservationsQuery, rs, rs.inters, v)
}

func (rs *ReservationsSelect) sqlScan(ctx context.Context, root *ReservationsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}