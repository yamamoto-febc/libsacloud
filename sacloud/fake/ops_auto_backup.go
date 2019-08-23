package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *AutoBackupOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.AutoBackupFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.AutoBackup
	for _, res := range results {
		dest := &sacloud.AutoBackup{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.AutoBackupFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		AutoBackups: values,
	}, nil
}

// Create is fake implementation
func (o *AutoBackupOp) Create(ctx context.Context, zone string, param *sacloud.AutoBackupCreateRequest) (*sacloud.AutoBackup, error) {
	result := &sacloud.AutoBackup{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	result.AccountID = accountID
	result.ZoneID = zoneIDs[zone]
	result.ZoneName = zone

	putAutoBackup(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *AutoBackupOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.AutoBackup, error) {
	value := getAutoBackupByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.AutoBackup{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *AutoBackupOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.AutoBackupUpdateRequest) (*sacloud.AutoBackup, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putAutoBackup(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *AutoBackupOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.AutoBackupPatchRequest) (*sacloud.AutoBackup, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	patchParam := make(map[string]interface{})
	if err := mergo.Map(&patchParam, value); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(&patchParam, param); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(param, &patchParam); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	copySameNameField(param, value)

	if param.PatchEmptyToDescription {
		value.Description = ""
	}
	if param.PatchEmptyToTags {
		value.Tags = nil
	}
	if param.PatchEmptyToIconID {
		value.IconID = types.ID(int64(0))
	}
	if param.PatchEmptyToBackupSpanWeekdays {
		value.BackupSpanWeekdays = nil
	}
	if param.PatchEmptyToMaximumNumberOfArchives {
		value.MaximumNumberOfArchives = 0
	}

	putAutoBackup(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *AutoBackupOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}