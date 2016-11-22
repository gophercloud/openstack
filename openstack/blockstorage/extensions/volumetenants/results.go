package volumetenants

// An extension to the base Volume object
type VolumeExt struct {
	// TenantID is the id of the project that owns the volume.
	TenantID string `json:"os-vol-tenant-attr:tenant_id"`
}

func (s *VolumeExt) UnmarshalJSON(b []byte) error {
	return nil
}
