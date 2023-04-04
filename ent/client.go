// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/woocoos/casbin-ent-adapter/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/casbin-ent-adapter/ent/casbinrule"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CasbinRule is the client for interacting with the CasbinRule builders.
	CasbinRule *CasbinRuleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CasbinRule = NewCasbinRuleClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		CasbinRule: NewCasbinRuleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		CasbinRule: NewCasbinRuleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CasbinRule.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CasbinRule.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.CasbinRule.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CasbinRuleMutation:
		return c.CasbinRule.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CasbinRuleClient is a client for the CasbinRule schema.
type CasbinRuleClient struct {
	config
}

// NewCasbinRuleClient returns a client for the CasbinRule from the given config.
func NewCasbinRuleClient(c config) *CasbinRuleClient {
	return &CasbinRuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `casbinrule.Hooks(f(g(h())))`.
func (c *CasbinRuleClient) Use(hooks ...Hook) {
	c.hooks.CasbinRule = append(c.hooks.CasbinRule, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `casbinrule.Intercept(f(g(h())))`.
func (c *CasbinRuleClient) Intercept(interceptors ...Interceptor) {
	c.inters.CasbinRule = append(c.inters.CasbinRule, interceptors...)
}

// Create returns a builder for creating a CasbinRule entity.
func (c *CasbinRuleClient) Create() *CasbinRuleCreate {
	mutation := newCasbinRuleMutation(c.config, OpCreate)
	return &CasbinRuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CasbinRule entities.
func (c *CasbinRuleClient) CreateBulk(builders ...*CasbinRuleCreate) *CasbinRuleCreateBulk {
	return &CasbinRuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CasbinRule.
func (c *CasbinRuleClient) Update() *CasbinRuleUpdate {
	mutation := newCasbinRuleMutation(c.config, OpUpdate)
	return &CasbinRuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CasbinRuleClient) UpdateOne(cr *CasbinRule) *CasbinRuleUpdateOne {
	mutation := newCasbinRuleMutation(c.config, OpUpdateOne, withCasbinRule(cr))
	return &CasbinRuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CasbinRuleClient) UpdateOneID(id int) *CasbinRuleUpdateOne {
	mutation := newCasbinRuleMutation(c.config, OpUpdateOne, withCasbinRuleID(id))
	return &CasbinRuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CasbinRule.
func (c *CasbinRuleClient) Delete() *CasbinRuleDelete {
	mutation := newCasbinRuleMutation(c.config, OpDelete)
	return &CasbinRuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CasbinRuleClient) DeleteOne(cr *CasbinRule) *CasbinRuleDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CasbinRuleClient) DeleteOneID(id int) *CasbinRuleDeleteOne {
	builder := c.Delete().Where(casbinrule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CasbinRuleDeleteOne{builder}
}

// Query returns a query builder for CasbinRule.
func (c *CasbinRuleClient) Query() *CasbinRuleQuery {
	return &CasbinRuleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCasbinRule},
		inters: c.Interceptors(),
	}
}

// Get returns a CasbinRule entity by its id.
func (c *CasbinRuleClient) Get(ctx context.Context, id int) (*CasbinRule, error) {
	return c.Query().Where(casbinrule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CasbinRuleClient) GetX(ctx context.Context, id int) *CasbinRule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CasbinRuleClient) Hooks() []Hook {
	return c.hooks.CasbinRule
}

// Interceptors returns the client interceptors.
func (c *CasbinRuleClient) Interceptors() []Interceptor {
	return c.inters.CasbinRule
}

func (c *CasbinRuleClient) mutate(ctx context.Context, m *CasbinRuleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CasbinRuleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CasbinRuleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CasbinRuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CasbinRuleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CasbinRule mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		CasbinRule []ent.Hook
	}
	inters struct {
		CasbinRule []ent.Interceptor
	}
)
