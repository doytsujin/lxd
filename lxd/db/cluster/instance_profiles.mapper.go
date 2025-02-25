//go:build linux && cgo && !agent

package cluster

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
)

var _ = api.ServerEnvironment{}

var instanceProfileObjectsByProfileID = RegisterStmt(`
SELECT instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order
  FROM instances_profiles
  WHERE instances_profiles.profile_id = ? ORDER BY instances_profiles.instance_id, instances_profiles.apply_order
`)

var instanceProfileObjectsByInstanceID = RegisterStmt(`
SELECT instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order
  FROM instances_profiles
  WHERE instances_profiles.instance_id = ? ORDER BY instances_profiles.instance_id, instances_profiles.apply_order
`)

var instanceProfileCreate = RegisterStmt(`
INSERT INTO instances_profiles (instance_id, profile_id, apply_order)
  VALUES (?, ?, ?)
`)

var instanceProfileDeleteByInstanceID = RegisterStmt(`
DELETE FROM instances_profiles WHERE instance_id = ?
`)

// GetProfileInstances returns all available Instances for the Profile.
// generator: instance_profile GetMany
func GetProfileInstances(ctx context.Context, tx *sql.Tx, profileID int) ([]Instance, error) {
	var err error

	// Result slice.
	objects := make([]InstanceProfile, 0)

	sqlStmt := Stmt(tx, instanceProfileObjectsByProfileID)
	args := []any{profileID}

	// Dest function for scanning a row.
	dest := func(i int) []any {
		objects = append(objects, InstanceProfile{})
		return []any{
			&objects[i].InstanceID,
			&objects[i].ProfileID,
			&objects[i].ApplyOrder,
		}
	}

	// Select.
	err = query.SelectObjects(sqlStmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	result := make([]Instance, len(objects))
	for i, object := range objects {
		instance, err := GetInstances(ctx, tx, InstanceFilter{ID: &object.InstanceID})
		if err != nil {
			return nil, err
		}

		result[i] = instance[0]
	}

	return result, nil
}

// GetInstanceProfiles returns all available Profiles for the Instance.
// generator: instance_profile GetMany
func GetInstanceProfiles(ctx context.Context, tx *sql.Tx, instanceID int) ([]Profile, error) {
	var err error

	// Result slice.
	objects := make([]InstanceProfile, 0)

	sqlStmt := Stmt(tx, instanceProfileObjectsByInstanceID)
	args := []any{instanceID}

	// Dest function for scanning a row.
	dest := func(i int) []any {
		objects = append(objects, InstanceProfile{})
		return []any{
			&objects[i].InstanceID,
			&objects[i].ProfileID,
			&objects[i].ApplyOrder,
		}
	}

	// Select.
	err = query.SelectObjects(sqlStmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	result := make([]Profile, len(objects))
	for i, object := range objects {
		profile, err := GetProfiles(ctx, tx, ProfileFilter{ID: &object.ProfileID})
		if err != nil {
			return nil, err
		}

		result[i] = profile[0]
	}

	return result, nil
}

// CreateInstanceProfiles adds a new instance_profile to the database.
// generator: instance_profile Create
func CreateInstanceProfiles(ctx context.Context, tx *sql.Tx, objects []InstanceProfile) error {
	for _, object := range objects {
		args := make([]any, 3)

		// Populate the statement arguments.
		args[0] = object.InstanceID
		args[1] = object.ProfileID
		args[2] = object.ApplyOrder

		// Prepared statement to use.
		stmt := Stmt(tx, instanceProfileCreate)

		// Execute the statement.
		_, err := stmt.Exec(args...)
		if err != nil {
			return fmt.Errorf("Failed to create \"instances_profiles\" entry: %w", err)
		}

	}

	return nil
}

// DeleteInstanceProfiles deletes the instance_profile matching the given key parameters.
// generator: instance_profile DeleteMany
func DeleteInstanceProfiles(ctx context.Context, tx *sql.Tx, instanceID int) error {
	stmt := Stmt(tx, instanceProfileDeleteByInstanceID)
	result, err := stmt.Exec(int(instanceID))
	if err != nil {
		return fmt.Errorf("Delete \"instances_profiles\" entry failed: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	return nil
}
