// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/casbin-ent-adapter/ent/casbinrule"
	"github.com/woocoos/casbin-ent-adapter/ent/predicate"
)

// CasbinRuleQuery is the builder for querying CasbinRule entities.
type CasbinRuleQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.CasbinRule
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CasbinRuleQuery builder.
func (crq *CasbinRuleQuery) Where(ps ...predicate.CasbinRule) *CasbinRuleQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit the number of records to be returned by this query.
func (crq *CasbinRuleQuery) Limit(limit int) *CasbinRuleQuery {
	crq.ctx.Limit = &limit
	return crq
}

// Offset to start from.
func (crq *CasbinRuleQuery) Offset(offset int) *CasbinRuleQuery {
	crq.ctx.Offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *CasbinRuleQuery) Unique(unique bool) *CasbinRuleQuery {
	crq.ctx.Unique = &unique
	return crq
}

// Order specifies how the records should be ordered.
func (crq *CasbinRuleQuery) Order(o ...OrderFunc) *CasbinRuleQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// First returns the first CasbinRule entity from the query.
// Returns a *NotFoundError when no CasbinRule was found.
func (crq *CasbinRuleQuery) First(ctx context.Context) (*CasbinRule, error) {
	nodes, err := crq.Limit(1).All(setContextOp(ctx, crq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{casbinrule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *CasbinRuleQuery) FirstX(ctx context.Context) *CasbinRule {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CasbinRule ID from the query.
// Returns a *NotFoundError when no CasbinRule ID was found.
func (crq *CasbinRuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(setContextOp(ctx, crq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{casbinrule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *CasbinRuleQuery) FirstIDX(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CasbinRule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CasbinRule entity is found.
// Returns a *NotFoundError when no CasbinRule entities are found.
func (crq *CasbinRuleQuery) Only(ctx context.Context) (*CasbinRule, error) {
	nodes, err := crq.Limit(2).All(setContextOp(ctx, crq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{casbinrule.Label}
	default:
		return nil, &NotSingularError{casbinrule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *CasbinRuleQuery) OnlyX(ctx context.Context) *CasbinRule {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CasbinRule ID in the query.
// Returns a *NotSingularError when more than one CasbinRule ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *CasbinRuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(setContextOp(ctx, crq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{casbinrule.Label}
	default:
		err = &NotSingularError{casbinrule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *CasbinRuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CasbinRules.
func (crq *CasbinRuleQuery) All(ctx context.Context) ([]*CasbinRule, error) {
	ctx = setContextOp(ctx, crq.ctx, "All")
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CasbinRule, *CasbinRuleQuery]()
	return withInterceptors[[]*CasbinRule](ctx, crq, qr, crq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crq *CasbinRuleQuery) AllX(ctx context.Context) []*CasbinRule {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CasbinRule IDs.
func (crq *CasbinRuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, crq.ctx, "IDs")
	if err := crq.Select(casbinrule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *CasbinRuleQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *CasbinRuleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crq.ctx, "Count")
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crq, querierCount[*CasbinRuleQuery](), crq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crq *CasbinRuleQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *CasbinRuleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crq.ctx, "Exist")
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *CasbinRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CasbinRuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *CasbinRuleQuery) Clone() *CasbinRuleQuery {
	if crq == nil {
		return nil
	}
	return &CasbinRuleQuery{
		config:     crq.config,
		ctx:        crq.ctx.Clone(),
		order:      append([]OrderFunc{}, crq.order...),
		inters:     append([]Interceptor{}, crq.inters...),
		predicates: append([]predicate.CasbinRule{}, crq.predicates...),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Ptype string `json:"Ptype,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CasbinRule.Query().
//		GroupBy(casbinrule.FieldPtype).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *CasbinRuleQuery) GroupBy(field string, fields ...string) *CasbinRuleGroupBy {
	crq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CasbinRuleGroupBy{build: crq}
	grbuild.flds = &crq.ctx.Fields
	grbuild.label = casbinrule.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Ptype string `json:"Ptype,omitempty"`
//	}
//
//	client.CasbinRule.Query().
//		Select(casbinrule.FieldPtype).
//		Scan(ctx, &v)
func (crq *CasbinRuleQuery) Select(fields ...string) *CasbinRuleSelect {
	crq.ctx.Fields = append(crq.ctx.Fields, fields...)
	sbuild := &CasbinRuleSelect{CasbinRuleQuery: crq}
	sbuild.label = casbinrule.Label
	sbuild.flds, sbuild.scan = &crq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CasbinRuleSelect configured with the given aggregations.
func (crq *CasbinRuleQuery) Aggregate(fns ...AggregateFunc) *CasbinRuleSelect {
	return crq.Select().Aggregate(fns...)
}

func (crq *CasbinRuleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crq); err != nil {
				return err
			}
		}
	}
	for _, f := range crq.ctx.Fields {
		if !casbinrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *CasbinRuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CasbinRule, error) {
	var (
		nodes = []*CasbinRule{}
		_spec = crq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CasbinRule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CasbinRule{config: crq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (crq *CasbinRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	_spec.Node.Columns = crq.ctx.Fields
	if len(crq.ctx.Fields) > 0 {
		_spec.Unique = crq.ctx.Unique != nil && *crq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *CasbinRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   casbinrule.Table,
			Columns: casbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: casbinrule.FieldID,
			},
		},
		From:   crq.sql,
		Unique: true,
	}
	if unique := crq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := crq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casbinrule.FieldID)
		for i := range fields {
			if fields[i] != casbinrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *CasbinRuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(casbinrule.Table)
	columns := crq.ctx.Fields
	if len(columns) == 0 {
		columns = casbinrule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.ctx.Unique != nil && *crq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CasbinRuleGroupBy is the group-by builder for CasbinRule entities.
type CasbinRuleGroupBy struct {
	selector
	build *CasbinRuleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *CasbinRuleGroupBy) Aggregate(fns ...AggregateFunc) *CasbinRuleGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the selector query and scans the result into the given value.
func (crgb *CasbinRuleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crgb.build.ctx, "GroupBy")
	if err := crgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CasbinRuleQuery, *CasbinRuleGroupBy](ctx, crgb.build, crgb, crgb.build.inters, v)
}

func (crgb *CasbinRuleGroupBy) sqlScan(ctx context.Context, root *CasbinRuleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crgb.flds)+len(crgb.fns))
		for _, f := range *crgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CasbinRuleSelect is the builder for selecting fields of CasbinRule entities.
type CasbinRuleSelect struct {
	*CasbinRuleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crs *CasbinRuleSelect) Aggregate(fns ...AggregateFunc) *CasbinRuleSelect {
	crs.fns = append(crs.fns, fns...)
	return crs
}

// Scan applies the selector query and scans the result into the given value.
func (crs *CasbinRuleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crs.ctx, "Select")
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CasbinRuleQuery, *CasbinRuleSelect](ctx, crs.CasbinRuleQuery, crs, crs.inters, v)
}

func (crs *CasbinRuleSelect) sqlScan(ctx context.Context, root *CasbinRuleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crs.fns))
	for _, fn := range crs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
