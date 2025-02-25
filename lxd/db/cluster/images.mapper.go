//go:build linux && cgo && !agent

package cluster

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
)

var _ = api.ServerEnvironment{}

var imageObjects = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByID = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE images.id = ? ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByProject = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE project = ? ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByProjectAndCached = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE project = ? AND images.cached = ? ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByProjectAndPublic = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE project = ? AND images.public = ? ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByFingerprint = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE images.fingerprint = ? ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByCached = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE images.cached = ? ORDER BY projects.id, images.fingerprint
`)

var imageObjectsByAutoUpdate = RegisterStmt(`
SELECT images.id, projects.name AS project, images.fingerprint, images.type, images.filename, images.size, images.public, images.architecture, images.creation_date, images.expiry_date, images.upload_date, images.cached, images.last_use_date, images.auto_update
  FROM images JOIN projects ON images.project_id = projects.id
  WHERE images.auto_update = ? ORDER BY projects.id, images.fingerprint
`)

// GetImages returns all available images.
// generator: image GetMany
func GetImages(ctx context.Context, tx *sql.Tx, filter ImageFilter) ([]Image, error) {
	var err error

	// Result slice.
	objects := make([]Image, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	var args []any

	if filter.Project != nil && filter.Public != nil && filter.ID == nil && filter.Fingerprint == nil && filter.Cached == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjectsByProjectAndPublic)
		args = []any{
			filter.Project,
			filter.Public,
		}
	} else if filter.Project != nil && filter.Cached != nil && filter.ID == nil && filter.Fingerprint == nil && filter.Public == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjectsByProjectAndCached)
		args = []any{
			filter.Project,
			filter.Cached,
		}
	} else if filter.Project != nil && filter.ID == nil && filter.Fingerprint == nil && filter.Public == nil && filter.Cached == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjectsByProject)
		args = []any{
			filter.Project,
		}
	} else if filter.ID != nil && filter.Project == nil && filter.Fingerprint == nil && filter.Public == nil && filter.Cached == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjectsByID)
		args = []any{
			filter.ID,
		}
	} else if filter.Fingerprint != nil && filter.ID == nil && filter.Project == nil && filter.Public == nil && filter.Cached == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjectsByFingerprint)
		args = []any{
			filter.Fingerprint,
		}
	} else if filter.Cached != nil && filter.ID == nil && filter.Project == nil && filter.Fingerprint == nil && filter.Public == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjectsByCached)
		args = []any{
			filter.Cached,
		}
	} else if filter.AutoUpdate != nil && filter.ID == nil && filter.Project == nil && filter.Fingerprint == nil && filter.Public == nil && filter.Cached == nil {
		sqlStmt = Stmt(tx, imageObjectsByAutoUpdate)
		args = []any{
			filter.AutoUpdate,
		}
	} else if filter.ID == nil && filter.Project == nil && filter.Fingerprint == nil && filter.Public == nil && filter.Cached == nil && filter.AutoUpdate == nil {
		sqlStmt = Stmt(tx, imageObjects)
		args = []any{}
	} else {
		return nil, fmt.Errorf("No statement exists for the given Filter")
	}

	// Dest function for scanning a row.
	dest := func(i int) []any {
		objects = append(objects, Image{})
		return []any{
			&objects[i].ID,
			&objects[i].Project,
			&objects[i].Fingerprint,
			&objects[i].Type,
			&objects[i].Filename,
			&objects[i].Size,
			&objects[i].Public,
			&objects[i].Architecture,
			&objects[i].CreationDate,
			&objects[i].ExpiryDate,
			&objects[i].UploadDate,
			&objects[i].Cached,
			&objects[i].LastUseDate,
			&objects[i].AutoUpdate,
		}
	}

	// Select.
	err = query.SelectObjects(sqlStmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"images\" table: %w", err)
	}

	return objects, nil
}

// GetImage returns the image with the given key.
// generator: image GetOne
func GetImage(ctx context.Context, tx *sql.Tx, project string, fingerprint string) (*Image, error) {
	filter := ImageFilter{}
	filter.Project = &project
	filter.Fingerprint = &fingerprint

	objects, err := GetImages(ctx, tx, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"images\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, api.StatusErrorf(http.StatusNotFound, "Image not found")
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"images\" entry matches")
	}
}
