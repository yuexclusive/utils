package db

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Config Config
type Config struct {
	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool
	// NamingStrategy tables, columns naming strategy
	NamingStrategy schema.Namer
	// FullSaveAssociations full save associations
	FullSaveAssociations bool
	// Logger
	Logger logger.Interface
	// NowFunc the function to be used when creating a new timestamp
	NowFunc func() time.Time
	// DryRun generate sql without execute
	DryRun bool
	// PrepareStmt executes the given query in cached statement
	PrepareStmt bool
	// DisableAutomaticPing
	DisableAutomaticPing bool
	// DisableForeignKeyConstraintWhenMigrating
	DisableForeignKeyConstraintWhenMigrating bool
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool
	// CreateBatchSize default create batch size
	CreateBatchSize int

	// ClauseBuilders clause builder
	ClauseBuilders map[string]clause.ClauseBuilder
	// ConnPool db conn pool
	ConnPool gorm.ConnPool
	// Dialector database dialector
	gorm.Dialector
	// Plugins registered plugins
	Plugins map[string]gorm.Plugin

	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
	// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	//
	// If n <= 0, no idle connections are retained.
	//
	// The default max idle connections is currently 2. This may change in
	// a future release.
	MaxIdleConns *int

	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
	// MaxIdleConns, then MaxIdleConns will be reduced to match the new
	// MaxOpenConns limit.
	//
	// If n <= 0, then there is no limit on the number of open connections.
	// The default is 0 (unlimited).
	MaxOpenConns *int

	// Name name
	Name string
}

func toGormConfig(cfg *Config) *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction:                   cfg.SkipDefaultTransaction,
		NamingStrategy:                           cfg.NamingStrategy,
		FullSaveAssociations:                     cfg.FullSaveAssociations,
		Logger:                                   cfg.Logger,
		NowFunc:                                  cfg.NowFunc,
		DryRun:                                   cfg.DryRun,
		PrepareStmt:                              cfg.PrepareStmt,
		DisableAutomaticPing:                     cfg.DisableAutomaticPing,
		DisableForeignKeyConstraintWhenMigrating: cfg.DisableForeignKeyConstraintWhenMigrating,
		DisableNestedTransaction:                 cfg.DisableNestedTransaction,
		AllowGlobalUpdate:                        cfg.AllowGlobalUpdate,
		QueryFields:                              cfg.QueryFields,
		CreateBatchSize:                          cfg.CreateBatchSize,
		ClauseBuilders:                           cfg.ClauseBuilders,
		ConnPool:                                 cfg.ConnPool,
		Dialector:                                cfg.Dialector,
		Plugins:                                  cfg.Plugins,
	}
}

func toConfig(cfg *gorm.Config) *Config {
	return &Config{
		SkipDefaultTransaction:                   cfg.SkipDefaultTransaction,
		NamingStrategy:                           cfg.NamingStrategy,
		FullSaveAssociations:                     cfg.FullSaveAssociations,
		Logger:                                   cfg.Logger,
		NowFunc:                                  cfg.NowFunc,
		DryRun:                                   cfg.DryRun,
		PrepareStmt:                              cfg.PrepareStmt,
		DisableAutomaticPing:                     cfg.DisableAutomaticPing,
		DisableForeignKeyConstraintWhenMigrating: cfg.DisableForeignKeyConstraintWhenMigrating,
		DisableNestedTransaction:                 cfg.DisableNestedTransaction,
		AllowGlobalUpdate:                        cfg.AllowGlobalUpdate,
		QueryFields:                              cfg.QueryFields,
		CreateBatchSize:                          cfg.CreateBatchSize,
		ClauseBuilders:                           cfg.ClauseBuilders,
		ConnPool:                                 cfg.ConnPool,
		Dialector:                                cfg.Dialector,
		Plugins:                                  cfg.Plugins,
	}
}

func (c *Config) Apply(config *gorm.Config) error {
	*config = *toGormConfig(c)
	return nil
}

func (c *Config) AfterInitialize(db *gorm.DB) error {
	if db != nil {
		for _, plugin := range c.Plugins {
			if err := plugin.Initialize(db); err != nil {
				return err
			}
		}
	}
	return nil
}
