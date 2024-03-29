// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ogen-go/ent2ogen/internal/test/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemaa"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemab"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// SchemaA is the client for interacting with the SchemaA builders.
	SchemaA *SchemaAClient
	// SchemaB is the client for interacting with the SchemaB builders.
	SchemaB *SchemaBClient
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
	c.SchemaA = NewSchemaAClient(c.config)
	c.SchemaB = NewSchemaBClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		SchemaA: NewSchemaAClient(cfg),
		SchemaB: NewSchemaBClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		SchemaA: NewSchemaAClient(cfg),
		SchemaB: NewSchemaBClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		SchemaA.
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
	c.SchemaA.Use(hooks...)
	c.SchemaB.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.SchemaA.Intercept(interceptors...)
	c.SchemaB.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *SchemaAMutation:
		return c.SchemaA.mutate(ctx, m)
	case *SchemaBMutation:
		return c.SchemaB.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// SchemaAClient is a client for the SchemaA schema.
type SchemaAClient struct {
	config
}

// NewSchemaAClient returns a client for the SchemaA from the given config.
func NewSchemaAClient(c config) *SchemaAClient {
	return &SchemaAClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schemaa.Hooks(f(g(h())))`.
func (c *SchemaAClient) Use(hooks ...Hook) {
	c.hooks.SchemaA = append(c.hooks.SchemaA, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `schemaa.Intercept(f(g(h())))`.
func (c *SchemaAClient) Intercept(interceptors ...Interceptor) {
	c.inters.SchemaA = append(c.inters.SchemaA, interceptors...)
}

// Create returns a builder for creating a SchemaA entity.
func (c *SchemaAClient) Create() *SchemaACreate {
	mutation := newSchemaAMutation(c.config, OpCreate)
	return &SchemaACreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SchemaA entities.
func (c *SchemaAClient) CreateBulk(builders ...*SchemaACreate) *SchemaACreateBulk {
	return &SchemaACreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SchemaA.
func (c *SchemaAClient) Update() *SchemaAUpdate {
	mutation := newSchemaAMutation(c.config, OpUpdate)
	return &SchemaAUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SchemaAClient) UpdateOne(s *SchemaA) *SchemaAUpdateOne {
	mutation := newSchemaAMutation(c.config, OpUpdateOne, withSchemaA(s))
	return &SchemaAUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SchemaAClient) UpdateOneID(id int) *SchemaAUpdateOne {
	mutation := newSchemaAMutation(c.config, OpUpdateOne, withSchemaAID(id))
	return &SchemaAUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SchemaA.
func (c *SchemaAClient) Delete() *SchemaADelete {
	mutation := newSchemaAMutation(c.config, OpDelete)
	return &SchemaADelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SchemaAClient) DeleteOne(s *SchemaA) *SchemaADeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SchemaAClient) DeleteOneID(id int) *SchemaADeleteOne {
	builder := c.Delete().Where(schemaa.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SchemaADeleteOne{builder}
}

// Query returns a query builder for SchemaA.
func (c *SchemaAClient) Query() *SchemaAQuery {
	return &SchemaAQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchemaA},
		inters: c.Interceptors(),
	}
}

// Get returns a SchemaA entity by its id.
func (c *SchemaAClient) Get(ctx context.Context, id int) (*SchemaA, error) {
	return c.Query().Where(schemaa.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SchemaAClient) GetX(ctx context.Context, id int) *SchemaA {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEdgeSchemabUniqueRequired queries the edge_schemab_unique_required edge of a SchemaA.
func (c *SchemaAClient) QueryEdgeSchemabUniqueRequired(s *SchemaA) *SchemaBQuery {
	query := (&SchemaBClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, id),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, schemaa.EdgeSchemabUniqueRequiredTable, schemaa.EdgeSchemabUniqueRequiredColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEdgeSchemabUniqueRequiredBindtoBs queries the edge_schemab_unique_required_bindto_bs edge of a SchemaA.
func (c *SchemaAClient) QueryEdgeSchemabUniqueRequiredBindtoBs(s *SchemaA) *SchemaBQuery {
	query := (&SchemaBClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, id),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, schemaa.EdgeSchemabUniqueRequiredBindtoBsTable, schemaa.EdgeSchemabUniqueRequiredBindtoBsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEdgeSchemabUniqueOptional queries the edge_schemab_unique_optional edge of a SchemaA.
func (c *SchemaAClient) QueryEdgeSchemabUniqueOptional(s *SchemaA) *SchemaBQuery {
	query := (&SchemaBClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, id),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, schemaa.EdgeSchemabUniqueOptionalTable, schemaa.EdgeSchemabUniqueOptionalColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEdgeSchemab queries the edge_schemab edge of a SchemaA.
func (c *SchemaAClient) QueryEdgeSchemab(s *SchemaA) *SchemaBQuery {
	query := (&SchemaBClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, id),
			sqlgraph.To(schemab.Table, schemab.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, schemaa.EdgeSchemabTable, schemaa.EdgeSchemabColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEdgeSchemaaRecursive queries the edge_schemaa_recursive edge of a SchemaA.
func (c *SchemaAClient) QueryEdgeSchemaaRecursive(s *SchemaA) *SchemaAQuery {
	query := (&SchemaAClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schemaa.Table, schemaa.FieldID, id),
			sqlgraph.To(schemaa.Table, schemaa.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, schemaa.EdgeSchemaaRecursiveTable, schemaa.EdgeSchemaaRecursivePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SchemaAClient) Hooks() []Hook {
	return c.hooks.SchemaA
}

// Interceptors returns the client interceptors.
func (c *SchemaAClient) Interceptors() []Interceptor {
	return c.inters.SchemaA
}

func (c *SchemaAClient) mutate(ctx context.Context, m *SchemaAMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SchemaACreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SchemaAUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SchemaAUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SchemaADelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SchemaA mutation op: %q", m.Op())
	}
}

// SchemaBClient is a client for the SchemaB schema.
type SchemaBClient struct {
	config
}

// NewSchemaBClient returns a client for the SchemaB from the given config.
func NewSchemaBClient(c config) *SchemaBClient {
	return &SchemaBClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schemab.Hooks(f(g(h())))`.
func (c *SchemaBClient) Use(hooks ...Hook) {
	c.hooks.SchemaB = append(c.hooks.SchemaB, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `schemab.Intercept(f(g(h())))`.
func (c *SchemaBClient) Intercept(interceptors ...Interceptor) {
	c.inters.SchemaB = append(c.inters.SchemaB, interceptors...)
}

// Create returns a builder for creating a SchemaB entity.
func (c *SchemaBClient) Create() *SchemaBCreate {
	mutation := newSchemaBMutation(c.config, OpCreate)
	return &SchemaBCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SchemaB entities.
func (c *SchemaBClient) CreateBulk(builders ...*SchemaBCreate) *SchemaBCreateBulk {
	return &SchemaBCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SchemaB.
func (c *SchemaBClient) Update() *SchemaBUpdate {
	mutation := newSchemaBMutation(c.config, OpUpdate)
	return &SchemaBUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SchemaBClient) UpdateOne(s *SchemaB) *SchemaBUpdateOne {
	mutation := newSchemaBMutation(c.config, OpUpdateOne, withSchemaB(s))
	return &SchemaBUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SchemaBClient) UpdateOneID(id int64) *SchemaBUpdateOne {
	mutation := newSchemaBMutation(c.config, OpUpdateOne, withSchemaBID(id))
	return &SchemaBUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SchemaB.
func (c *SchemaBClient) Delete() *SchemaBDelete {
	mutation := newSchemaBMutation(c.config, OpDelete)
	return &SchemaBDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SchemaBClient) DeleteOne(s *SchemaB) *SchemaBDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SchemaBClient) DeleteOneID(id int64) *SchemaBDeleteOne {
	builder := c.Delete().Where(schemab.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SchemaBDeleteOne{builder}
}

// Query returns a query builder for SchemaB.
func (c *SchemaBClient) Query() *SchemaBQuery {
	return &SchemaBQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchemaB},
		inters: c.Interceptors(),
	}
}

// Get returns a SchemaB entity by its id.
func (c *SchemaBClient) Get(ctx context.Context, id int64) (*SchemaB, error) {
	return c.Query().Where(schemab.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SchemaBClient) GetX(ctx context.Context, id int64) *SchemaB {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SchemaBClient) Hooks() []Hook {
	return c.hooks.SchemaB
}

// Interceptors returns the client interceptors.
func (c *SchemaBClient) Interceptors() []Interceptor {
	return c.inters.SchemaB
}

func (c *SchemaBClient) mutate(ctx context.Context, m *SchemaBMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SchemaBCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SchemaBUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SchemaBUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SchemaBDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SchemaB mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		SchemaA, SchemaB []ent.Hook
	}
	inters struct {
		SchemaA, SchemaB []ent.Interceptor
	}
)
