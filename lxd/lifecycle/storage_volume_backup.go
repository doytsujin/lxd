package lifecycle

import (
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/version"
)

// StorageVolumeBackupAction represents a lifecycle event action for storage volume backups.
type StorageVolumeBackupAction string

// All supported lifecycle events for storage volume backups.
const (
	StorageVolumeBackupCreated   = StorageVolumeBackupAction(api.EventLifecycleStorageVolumeBackupCreated)
	StorageVolumeBackupDeleted   = StorageVolumeBackupAction(api.EventLifecycleStorageVolumeBackupDeleted)
	StorageVolumeBackupRetrieved = StorageVolumeBackupAction(api.EventLifecycleStorageVolumeBackupRetrieved)
	StorageVolumeBackupRenamed   = StorageVolumeBackupAction(api.EventLifecycleStorageVolumeBackupRenamed)
)

// Event creates the lifecycle event for an action on a storage volume backup.
func (a StorageVolumeBackupAction) Event(poolName string, volumeType string, fullBackupName string, projectName string, requestor *api.EventLifecycleRequestor, ctx map[string]any) api.EventLifecycle {
	volumeName, backupName, _ := shared.InstanceGetParentAndSnapshotName(fullBackupName)

	u := api.NewURL().Path(version.APIVersion, "storage-pools", poolName, "volumes", volumeType, volumeName, "backups", backupName).Project(projectName)

	return api.EventLifecycle{
		Action:    string(a),
		Source:    u.String(),
		Context:   ctx,
		Requestor: requestor,
	}
}
