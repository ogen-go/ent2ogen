// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ogen-go/ent2ogen/example/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ogen-go/ent2ogen/example/ent/keyboard"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
	"github.com/ogen-go/ent2ogen/example/ent/switchmodel"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Keyboard is the client for interacting with the Keyboard builders.
	Keyboard *KeyboardClient
	// KeycapModel is the client for interacting with the KeycapModel builders.
	KeycapModel *KeycapModelClient
	// SwitchModel is the client for interacting with the SwitchModel builders.
	SwitchModel *SwitchModelClient
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
	c.Keyboard = NewKeyboardClient(c.config)
	c.KeycapModel = NewKeycapModelClient(c.config)
	c.SwitchModel = NewSwitchModelClient(c.config)
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
		ctx:         ctx,
		config:      cfg,
		Keyboard:    NewKeyboardClient(cfg),
		KeycapModel: NewKeycapModelClient(cfg),
		SwitchModel: NewSwitchModelClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Keyboard:    NewKeyboardClient(cfg),
		KeycapModel: NewKeycapModelClient(cfg),
		SwitchModel: NewSwitchModelClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Keyboard.
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
	c.Keyboard.Use(hooks...)
	c.KeycapModel.Use(hooks...)
	c.SwitchModel.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Keyboard.Intercept(interceptors...)
	c.KeycapModel.Intercept(interceptors...)
	c.SwitchModel.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *KeyboardMutation:
		return c.Keyboard.mutate(ctx, m)
	case *KeycapModelMutation:
		return c.KeycapModel.mutate(ctx, m)
	case *SwitchModelMutation:
		return c.SwitchModel.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// KeyboardClient is a client for the Keyboard schema.
type KeyboardClient struct {
	config
}

// NewKeyboardClient returns a client for the Keyboard from the given config.
func NewKeyboardClient(c config) *KeyboardClient {
	return &KeyboardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `keyboard.Hooks(f(g(h())))`.
func (c *KeyboardClient) Use(hooks ...Hook) {
	c.hooks.Keyboard = append(c.hooks.Keyboard, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `keyboard.Intercept(f(g(h())))`.
func (c *KeyboardClient) Intercept(interceptors ...Interceptor) {
	c.inters.Keyboard = append(c.inters.Keyboard, interceptors...)
}

// Create returns a builder for creating a Keyboard entity.
func (c *KeyboardClient) Create() *KeyboardCreate {
	mutation := newKeyboardMutation(c.config, OpCreate)
	return &KeyboardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Keyboard entities.
func (c *KeyboardClient) CreateBulk(builders ...*KeyboardCreate) *KeyboardCreateBulk {
	return &KeyboardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Keyboard.
func (c *KeyboardClient) Update() *KeyboardUpdate {
	mutation := newKeyboardMutation(c.config, OpUpdate)
	return &KeyboardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *KeyboardClient) UpdateOne(k *Keyboard) *KeyboardUpdateOne {
	mutation := newKeyboardMutation(c.config, OpUpdateOne, withKeyboard(k))
	return &KeyboardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *KeyboardClient) UpdateOneID(id int64) *KeyboardUpdateOne {
	mutation := newKeyboardMutation(c.config, OpUpdateOne, withKeyboardID(id))
	return &KeyboardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Keyboard.
func (c *KeyboardClient) Delete() *KeyboardDelete {
	mutation := newKeyboardMutation(c.config, OpDelete)
	return &KeyboardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *KeyboardClient) DeleteOne(k *Keyboard) *KeyboardDeleteOne {
	return c.DeleteOneID(k.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *KeyboardClient) DeleteOneID(id int64) *KeyboardDeleteOne {
	builder := c.Delete().Where(keyboard.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &KeyboardDeleteOne{builder}
}

// Query returns a query builder for Keyboard.
func (c *KeyboardClient) Query() *KeyboardQuery {
	return &KeyboardQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeKeyboard},
		inters: c.Interceptors(),
	}
}

// Get returns a Keyboard entity by its id.
func (c *KeyboardClient) Get(ctx context.Context, id int64) (*Keyboard, error) {
	return c.Query().Where(keyboard.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *KeyboardClient) GetX(ctx context.Context, id int64) *Keyboard {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySwitches queries the switches edge of a Keyboard.
func (c *KeyboardClient) QuerySwitches(k *Keyboard) *SwitchModelQuery {
	query := (&SwitchModelClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := k.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(keyboard.Table, keyboard.FieldID, id),
			sqlgraph.To(switchmodel.Table, switchmodel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, keyboard.SwitchesTable, keyboard.SwitchesColumn),
		)
		fromV = sqlgraph.Neighbors(k.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryKeycaps queries the keycaps edge of a Keyboard.
func (c *KeyboardClient) QueryKeycaps(k *Keyboard) *KeycapModelQuery {
	query := (&KeycapModelClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := k.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(keyboard.Table, keyboard.FieldID, id),
			sqlgraph.To(keycapmodel.Table, keycapmodel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, keyboard.KeycapsTable, keyboard.KeycapsColumn),
		)
		fromV = sqlgraph.Neighbors(k.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *KeyboardClient) Hooks() []Hook {
	return c.hooks.Keyboard
}

// Interceptors returns the client interceptors.
func (c *KeyboardClient) Interceptors() []Interceptor {
	return c.inters.Keyboard
}

func (c *KeyboardClient) mutate(ctx context.Context, m *KeyboardMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&KeyboardCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&KeyboardUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&KeyboardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&KeyboardDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Keyboard mutation op: %q", m.Op())
	}
}

// KeycapModelClient is a client for the KeycapModel schema.
type KeycapModelClient struct {
	config
}

// NewKeycapModelClient returns a client for the KeycapModel from the given config.
func NewKeycapModelClient(c config) *KeycapModelClient {
	return &KeycapModelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `keycapmodel.Hooks(f(g(h())))`.
func (c *KeycapModelClient) Use(hooks ...Hook) {
	c.hooks.KeycapModel = append(c.hooks.KeycapModel, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `keycapmodel.Intercept(f(g(h())))`.
func (c *KeycapModelClient) Intercept(interceptors ...Interceptor) {
	c.inters.KeycapModel = append(c.inters.KeycapModel, interceptors...)
}

// Create returns a builder for creating a KeycapModel entity.
func (c *KeycapModelClient) Create() *KeycapModelCreate {
	mutation := newKeycapModelMutation(c.config, OpCreate)
	return &KeycapModelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of KeycapModel entities.
func (c *KeycapModelClient) CreateBulk(builders ...*KeycapModelCreate) *KeycapModelCreateBulk {
	return &KeycapModelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for KeycapModel.
func (c *KeycapModelClient) Update() *KeycapModelUpdate {
	mutation := newKeycapModelMutation(c.config, OpUpdate)
	return &KeycapModelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *KeycapModelClient) UpdateOne(km *KeycapModel) *KeycapModelUpdateOne {
	mutation := newKeycapModelMutation(c.config, OpUpdateOne, withKeycapModel(km))
	return &KeycapModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *KeycapModelClient) UpdateOneID(id int64) *KeycapModelUpdateOne {
	mutation := newKeycapModelMutation(c.config, OpUpdateOne, withKeycapModelID(id))
	return &KeycapModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for KeycapModel.
func (c *KeycapModelClient) Delete() *KeycapModelDelete {
	mutation := newKeycapModelMutation(c.config, OpDelete)
	return &KeycapModelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *KeycapModelClient) DeleteOne(km *KeycapModel) *KeycapModelDeleteOne {
	return c.DeleteOneID(km.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *KeycapModelClient) DeleteOneID(id int64) *KeycapModelDeleteOne {
	builder := c.Delete().Where(keycapmodel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &KeycapModelDeleteOne{builder}
}

// Query returns a query builder for KeycapModel.
func (c *KeycapModelClient) Query() *KeycapModelQuery {
	return &KeycapModelQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeKeycapModel},
		inters: c.Interceptors(),
	}
}

// Get returns a KeycapModel entity by its id.
func (c *KeycapModelClient) Get(ctx context.Context, id int64) (*KeycapModel, error) {
	return c.Query().Where(keycapmodel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *KeycapModelClient) GetX(ctx context.Context, id int64) *KeycapModel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *KeycapModelClient) Hooks() []Hook {
	return c.hooks.KeycapModel
}

// Interceptors returns the client interceptors.
func (c *KeycapModelClient) Interceptors() []Interceptor {
	return c.inters.KeycapModel
}

func (c *KeycapModelClient) mutate(ctx context.Context, m *KeycapModelMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&KeycapModelCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&KeycapModelUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&KeycapModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&KeycapModelDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown KeycapModel mutation op: %q", m.Op())
	}
}

// SwitchModelClient is a client for the SwitchModel schema.
type SwitchModelClient struct {
	config
}

// NewSwitchModelClient returns a client for the SwitchModel from the given config.
func NewSwitchModelClient(c config) *SwitchModelClient {
	return &SwitchModelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `switchmodel.Hooks(f(g(h())))`.
func (c *SwitchModelClient) Use(hooks ...Hook) {
	c.hooks.SwitchModel = append(c.hooks.SwitchModel, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `switchmodel.Intercept(f(g(h())))`.
func (c *SwitchModelClient) Intercept(interceptors ...Interceptor) {
	c.inters.SwitchModel = append(c.inters.SwitchModel, interceptors...)
}

// Create returns a builder for creating a SwitchModel entity.
func (c *SwitchModelClient) Create() *SwitchModelCreate {
	mutation := newSwitchModelMutation(c.config, OpCreate)
	return &SwitchModelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SwitchModel entities.
func (c *SwitchModelClient) CreateBulk(builders ...*SwitchModelCreate) *SwitchModelCreateBulk {
	return &SwitchModelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SwitchModel.
func (c *SwitchModelClient) Update() *SwitchModelUpdate {
	mutation := newSwitchModelMutation(c.config, OpUpdate)
	return &SwitchModelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SwitchModelClient) UpdateOne(sm *SwitchModel) *SwitchModelUpdateOne {
	mutation := newSwitchModelMutation(c.config, OpUpdateOne, withSwitchModel(sm))
	return &SwitchModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SwitchModelClient) UpdateOneID(id int64) *SwitchModelUpdateOne {
	mutation := newSwitchModelMutation(c.config, OpUpdateOne, withSwitchModelID(id))
	return &SwitchModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SwitchModel.
func (c *SwitchModelClient) Delete() *SwitchModelDelete {
	mutation := newSwitchModelMutation(c.config, OpDelete)
	return &SwitchModelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SwitchModelClient) DeleteOne(sm *SwitchModel) *SwitchModelDeleteOne {
	return c.DeleteOneID(sm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SwitchModelClient) DeleteOneID(id int64) *SwitchModelDeleteOne {
	builder := c.Delete().Where(switchmodel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SwitchModelDeleteOne{builder}
}

// Query returns a query builder for SwitchModel.
func (c *SwitchModelClient) Query() *SwitchModelQuery {
	return &SwitchModelQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSwitchModel},
		inters: c.Interceptors(),
	}
}

// Get returns a SwitchModel entity by its id.
func (c *SwitchModelClient) Get(ctx context.Context, id int64) (*SwitchModel, error) {
	return c.Query().Where(switchmodel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SwitchModelClient) GetX(ctx context.Context, id int64) *SwitchModel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SwitchModelClient) Hooks() []Hook {
	return c.hooks.SwitchModel
}

// Interceptors returns the client interceptors.
func (c *SwitchModelClient) Interceptors() []Interceptor {
	return c.inters.SwitchModel
}

func (c *SwitchModelClient) mutate(ctx context.Context, m *SwitchModelMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SwitchModelCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SwitchModelUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SwitchModelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SwitchModelDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SwitchModel mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Keyboard, KeycapModel, SwitchModel []ent.Hook
	}
	inters struct {
		Keyboard, KeycapModel, SwitchModel []ent.Interceptor
	}
)
