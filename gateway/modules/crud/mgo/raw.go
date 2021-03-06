package mgo

import (
	"context"
	"errors"
)

// RawBatch performs a batch operation for schema creation
// NOTE: not to be exposed externally
func (m *Mongo) RawBatch(ctx context.Context, queries []string) error {
	return errors.New("raw batch operation cannot be performed on mongo")
}

// RawExec performs an operation for schema creation
// NOTE: not to be exposed externally
func (m *Mongo) RawExec(ctx context.Context, query string) error {
	return errors.New("raw exec operation cannot be performed on mongo")
}

// GetConnectionState : function to check connection state
func (m *Mongo) GetConnectionState(ctx context.Context) bool {
	if !m.enabled || m.client == nil {
		return false
	}

	// Ping to check if connection is established
	err := m.client.Ping(ctx, nil)
	return err == nil
}

// CreateDatabaseIfNotExist creates a database if not exist which has same name of project
func (m *Mongo) CreateDatabaseIfNotExist(ctx context.Context, project string) error {
	return errors.New("create project exists cannot be performed over mongo")
}
