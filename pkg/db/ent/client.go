// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/api-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ServiceAPI is the client for interacting with the ServiceAPI builders.
	ServiceAPI *ServiceAPIClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ServiceAPI = NewServiceAPIClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ServiceAPI: NewServiceAPIClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:     cfg,
		ServiceAPI: NewServiceAPIClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ServiceAPI.
//		Query().
//		Count(ctx)
//
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
	c.ServiceAPI.Use(hooks...)
}

// ServiceAPIClient is a client for the ServiceAPI schema.
type ServiceAPIClient struct {
	config
}

// NewServiceAPIClient returns a client for the ServiceAPI from the given config.
func NewServiceAPIClient(c config) *ServiceAPIClient {
	return &ServiceAPIClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `serviceapi.Hooks(f(g(h())))`.
func (c *ServiceAPIClient) Use(hooks ...Hook) {
	c.hooks.ServiceAPI = append(c.hooks.ServiceAPI, hooks...)
}

// Create returns a create builder for ServiceAPI.
func (c *ServiceAPIClient) Create() *ServiceAPICreate {
	mutation := newServiceAPIMutation(c.config, OpCreate)
	return &ServiceAPICreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ServiceAPI entities.
func (c *ServiceAPIClient) CreateBulk(builders ...*ServiceAPICreate) *ServiceAPICreateBulk {
	return &ServiceAPICreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ServiceAPI.
func (c *ServiceAPIClient) Update() *ServiceAPIUpdate {
	mutation := newServiceAPIMutation(c.config, OpUpdate)
	return &ServiceAPIUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServiceAPIClient) UpdateOne(sa *ServiceAPI) *ServiceAPIUpdateOne {
	mutation := newServiceAPIMutation(c.config, OpUpdateOne, withServiceAPI(sa))
	return &ServiceAPIUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServiceAPIClient) UpdateOneID(id uuid.UUID) *ServiceAPIUpdateOne {
	mutation := newServiceAPIMutation(c.config, OpUpdateOne, withServiceAPIID(id))
	return &ServiceAPIUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ServiceAPI.
func (c *ServiceAPIClient) Delete() *ServiceAPIDelete {
	mutation := newServiceAPIMutation(c.config, OpDelete)
	return &ServiceAPIDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ServiceAPIClient) DeleteOne(sa *ServiceAPI) *ServiceAPIDeleteOne {
	return c.DeleteOneID(sa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ServiceAPIClient) DeleteOneID(id uuid.UUID) *ServiceAPIDeleteOne {
	builder := c.Delete().Where(serviceapi.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServiceAPIDeleteOne{builder}
}

// Query returns a query builder for ServiceAPI.
func (c *ServiceAPIClient) Query() *ServiceAPIQuery {
	return &ServiceAPIQuery{
		config: c.config,
	}
}

// Get returns a ServiceAPI entity by its id.
func (c *ServiceAPIClient) Get(ctx context.Context, id uuid.UUID) (*ServiceAPI, error) {
	return c.Query().Where(serviceapi.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServiceAPIClient) GetX(ctx context.Context, id uuid.UUID) *ServiceAPI {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ServiceAPIClient) Hooks() []Hook {
	return c.hooks.ServiceAPI
}
