// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/casbin-ent-adapter/ent/casbinrule"
	"github.com/woocoos/casbin-ent-adapter/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCasbinRule = "CasbinRule"
)

// CasbinRuleMutation represents an operation that mutates the CasbinRule nodes in the graph.
type CasbinRuleMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_Ptype        *string
	_V0           *string
	_V1           *string
	_V2           *string
	_V3           *string
	_V4           *string
	_V5           *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*CasbinRule, error)
	predicates    []predicate.CasbinRule
}

var _ ent.Mutation = (*CasbinRuleMutation)(nil)

// casbinruleOption allows management of the mutation configuration using functional options.
type casbinruleOption func(*CasbinRuleMutation)

// newCasbinRuleMutation creates new mutation for the CasbinRule entity.
func newCasbinRuleMutation(c config, op Op, opts ...casbinruleOption) *CasbinRuleMutation {
	m := &CasbinRuleMutation{
		config:        c,
		op:            op,
		typ:           TypeCasbinRule,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCasbinRuleID sets the ID field of the mutation.
func withCasbinRuleID(id int) casbinruleOption {
	return func(m *CasbinRuleMutation) {
		var (
			err   error
			once  sync.Once
			value *CasbinRule
		)
		m.oldValue = func(ctx context.Context) (*CasbinRule, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().CasbinRule.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCasbinRule sets the old CasbinRule of the mutation.
func withCasbinRule(node *CasbinRule) casbinruleOption {
	return func(m *CasbinRuleMutation) {
		m.oldValue = func(context.Context) (*CasbinRule, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CasbinRuleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CasbinRuleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CasbinRuleMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CasbinRuleMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().CasbinRule.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPtype sets the "Ptype" field.
func (m *CasbinRuleMutation) SetPtype(s string) {
	m._Ptype = &s
}

// Ptype returns the value of the "Ptype" field in the mutation.
func (m *CasbinRuleMutation) Ptype() (r string, exists bool) {
	v := m._Ptype
	if v == nil {
		return
	}
	return *v, true
}

// OldPtype returns the old "Ptype" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldPtype(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPtype is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPtype requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPtype: %w", err)
	}
	return oldValue.Ptype, nil
}

// ResetPtype resets all changes to the "Ptype" field.
func (m *CasbinRuleMutation) ResetPtype() {
	m._Ptype = nil
}

// SetV0 sets the "V0" field.
func (m *CasbinRuleMutation) SetV0(s string) {
	m._V0 = &s
}

// V0 returns the value of the "V0" field in the mutation.
func (m *CasbinRuleMutation) V0() (r string, exists bool) {
	v := m._V0
	if v == nil {
		return
	}
	return *v, true
}

// OldV0 returns the old "V0" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldV0(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldV0 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldV0 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldV0: %w", err)
	}
	return oldValue.V0, nil
}

// ResetV0 resets all changes to the "V0" field.
func (m *CasbinRuleMutation) ResetV0() {
	m._V0 = nil
}

// SetV1 sets the "V1" field.
func (m *CasbinRuleMutation) SetV1(s string) {
	m._V1 = &s
}

// V1 returns the value of the "V1" field in the mutation.
func (m *CasbinRuleMutation) V1() (r string, exists bool) {
	v := m._V1
	if v == nil {
		return
	}
	return *v, true
}

// OldV1 returns the old "V1" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldV1(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldV1 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldV1 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldV1: %w", err)
	}
	return oldValue.V1, nil
}

// ResetV1 resets all changes to the "V1" field.
func (m *CasbinRuleMutation) ResetV1() {
	m._V1 = nil
}

// SetV2 sets the "V2" field.
func (m *CasbinRuleMutation) SetV2(s string) {
	m._V2 = &s
}

// V2 returns the value of the "V2" field in the mutation.
func (m *CasbinRuleMutation) V2() (r string, exists bool) {
	v := m._V2
	if v == nil {
		return
	}
	return *v, true
}

// OldV2 returns the old "V2" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldV2(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldV2 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldV2 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldV2: %w", err)
	}
	return oldValue.V2, nil
}

// ResetV2 resets all changes to the "V2" field.
func (m *CasbinRuleMutation) ResetV2() {
	m._V2 = nil
}

// SetV3 sets the "V3" field.
func (m *CasbinRuleMutation) SetV3(s string) {
	m._V3 = &s
}

// V3 returns the value of the "V3" field in the mutation.
func (m *CasbinRuleMutation) V3() (r string, exists bool) {
	v := m._V3
	if v == nil {
		return
	}
	return *v, true
}

// OldV3 returns the old "V3" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldV3(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldV3 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldV3 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldV3: %w", err)
	}
	return oldValue.V3, nil
}

// ResetV3 resets all changes to the "V3" field.
func (m *CasbinRuleMutation) ResetV3() {
	m._V3 = nil
}

// SetV4 sets the "V4" field.
func (m *CasbinRuleMutation) SetV4(s string) {
	m._V4 = &s
}

// V4 returns the value of the "V4" field in the mutation.
func (m *CasbinRuleMutation) V4() (r string, exists bool) {
	v := m._V4
	if v == nil {
		return
	}
	return *v, true
}

// OldV4 returns the old "V4" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldV4(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldV4 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldV4 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldV4: %w", err)
	}
	return oldValue.V4, nil
}

// ResetV4 resets all changes to the "V4" field.
func (m *CasbinRuleMutation) ResetV4() {
	m._V4 = nil
}

// SetV5 sets the "V5" field.
func (m *CasbinRuleMutation) SetV5(s string) {
	m._V5 = &s
}

// V5 returns the value of the "V5" field in the mutation.
func (m *CasbinRuleMutation) V5() (r string, exists bool) {
	v := m._V5
	if v == nil {
		return
	}
	return *v, true
}

// OldV5 returns the old "V5" field's value of the CasbinRule entity.
// If the CasbinRule object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CasbinRuleMutation) OldV5(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldV5 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldV5 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldV5: %w", err)
	}
	return oldValue.V5, nil
}

// ResetV5 resets all changes to the "V5" field.
func (m *CasbinRuleMutation) ResetV5() {
	m._V5 = nil
}

// Where appends a list predicates to the CasbinRuleMutation builder.
func (m *CasbinRuleMutation) Where(ps ...predicate.CasbinRule) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CasbinRuleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CasbinRuleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.CasbinRule, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CasbinRuleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CasbinRuleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (CasbinRule).
func (m *CasbinRuleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CasbinRuleMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m._Ptype != nil {
		fields = append(fields, casbinrule.FieldPtype)
	}
	if m._V0 != nil {
		fields = append(fields, casbinrule.FieldV0)
	}
	if m._V1 != nil {
		fields = append(fields, casbinrule.FieldV1)
	}
	if m._V2 != nil {
		fields = append(fields, casbinrule.FieldV2)
	}
	if m._V3 != nil {
		fields = append(fields, casbinrule.FieldV3)
	}
	if m._V4 != nil {
		fields = append(fields, casbinrule.FieldV4)
	}
	if m._V5 != nil {
		fields = append(fields, casbinrule.FieldV5)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CasbinRuleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case casbinrule.FieldPtype:
		return m.Ptype()
	case casbinrule.FieldV0:
		return m.V0()
	case casbinrule.FieldV1:
		return m.V1()
	case casbinrule.FieldV2:
		return m.V2()
	case casbinrule.FieldV3:
		return m.V3()
	case casbinrule.FieldV4:
		return m.V4()
	case casbinrule.FieldV5:
		return m.V5()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CasbinRuleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case casbinrule.FieldPtype:
		return m.OldPtype(ctx)
	case casbinrule.FieldV0:
		return m.OldV0(ctx)
	case casbinrule.FieldV1:
		return m.OldV1(ctx)
	case casbinrule.FieldV2:
		return m.OldV2(ctx)
	case casbinrule.FieldV3:
		return m.OldV3(ctx)
	case casbinrule.FieldV4:
		return m.OldV4(ctx)
	case casbinrule.FieldV5:
		return m.OldV5(ctx)
	}
	return nil, fmt.Errorf("unknown CasbinRule field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CasbinRuleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case casbinrule.FieldPtype:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPtype(v)
		return nil
	case casbinrule.FieldV0:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetV0(v)
		return nil
	case casbinrule.FieldV1:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetV1(v)
		return nil
	case casbinrule.FieldV2:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetV2(v)
		return nil
	case casbinrule.FieldV3:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetV3(v)
		return nil
	case casbinrule.FieldV4:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetV4(v)
		return nil
	case casbinrule.FieldV5:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetV5(v)
		return nil
	}
	return fmt.Errorf("unknown CasbinRule field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CasbinRuleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CasbinRuleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CasbinRuleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown CasbinRule numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CasbinRuleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CasbinRuleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CasbinRuleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown CasbinRule nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CasbinRuleMutation) ResetField(name string) error {
	switch name {
	case casbinrule.FieldPtype:
		m.ResetPtype()
		return nil
	case casbinrule.FieldV0:
		m.ResetV0()
		return nil
	case casbinrule.FieldV1:
		m.ResetV1()
		return nil
	case casbinrule.FieldV2:
		m.ResetV2()
		return nil
	case casbinrule.FieldV3:
		m.ResetV3()
		return nil
	case casbinrule.FieldV4:
		m.ResetV4()
		return nil
	case casbinrule.FieldV5:
		m.ResetV5()
		return nil
	}
	return fmt.Errorf("unknown CasbinRule field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CasbinRuleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CasbinRuleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CasbinRuleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CasbinRuleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CasbinRuleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CasbinRuleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CasbinRuleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown CasbinRule unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CasbinRuleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown CasbinRule edge %s", name)
}
