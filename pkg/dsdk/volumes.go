package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Volume struct {
	Path               string             `json:"path,omitempty" mapstructure:"path"`
	ActiveStorageNodes []*StorageNode     `json:"active_storage_nodes,omitempty" mapstructure:"active_storage_nodes"`
	AvailabilityState  string             `json:"availability_state,omitempty" mapstructure:"availability_state"`
	CapacityInUse      int                `json:"capacity_in_use,omitempty" mapstructure:"capacity_in_use"`
	Causes             []string           `json:"causes,omitempty" mapstructure:"causes"`
	DeploymentState    string             `json:"deployment_state,omitempty" mapstructure:"deployment_state"`
	EffectiveSize      int                `json:"effective_size,omitempty" mapstructure:"effective_size"`
	ExclusiveSize      int                `json:"exclusive_size,omitempty" mapstructure:"exclusive_size"`
	Health             string             `json:"health,omitempty" mapstructure:"health"`
	LogicalSize        int                `json:"logical_size,omitempty" mapstructure:"logical_size"`
	Name               string             `json:"name,omitempty" mapstructure:"name"`
	OpState            string             `json:"op_state,omitempty" mapstructure:"op_state"`
	OpStatus           string             `json:"op_status,omitempty" mapstructure:"op_status"`
	PhysicalSize       int                `json:"physical_size,omitempty" mapstructure:"physical_size"`
	PlacementMode      string             `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy    string             `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	RecoveryState      string             `json:"recovery_state,omitempty" mapstructure:"recovery_state"`
	ReplicaCount       int                `json:"replica_count,omitempty" mapstructure:"replica_count"`
	RestorePoint       string             `json:"restore_point,omitempty" mapstructure:"restore_point"`
	Size               int                `json:"size,omitempty" mapstructure:"size"`
	Snapshots          []*Snapshot        `json:"snapshots,omitempty" mapstructure:"snapshots"`
	StoragePool        []*StoragePool     `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
	StorageState       string             `json:"storage_state,omitempty" mapstructure:"storage_state"`
	Uuid               string             `json:"uuid,omitempty" mapstructure:"uuid"`
	SnapshotsEp        *Snapshots         `json:"-"`
	PerformancePolicy  *PerformancePolicy `json:"-"`
}

func RegisterVolumeEndpoints(v *Volume) {
	v.SnapshotsEp = newSnapshots(v.Path)
	v.PerformancePolicy = newPerformancePolicy(v.Path)
}

type Volumes struct {
	Path string
}

type VolumesCreateRequest struct {
	Ctxt            context.Context `json:"-"`
	Name            string          `json:"name,omitempty" mapstructure:"name"`
	ReplicaCount    int             `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size            int             `json:"size,omitempty" mapstructure:"size"`
	PlacementMode   string          `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy string          `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	Force           bool            `json:"force,omitempty" mapstructure:"force"`
}

func newVolumes(path string) *Volumes {
	return &Volumes{
		Path: _path.Join(path, "volumes"),
	}
}

func (e *Volumes) Create(ro *VolumesCreateRequest) (*Volume, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Volume{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type VolumesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *Volumes) List(ro *VolumesListRequest) ([]*Volume, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*Volume{}
	for _, data := range rs.Data {
		elem := &Volume{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type VolumesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"name,omitempty" mapstructure:"name"`
}

func (e *Volumes) Get(ro *VolumesGetRequest) (*Volume, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &Volume{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type VolumeSetRequest struct {
	Ctxt            context.Context `json:"-"`
	ReplicaCount    int             `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size            int             `json:"size,omitempty" mapstructure:"size"`
	PlacementMode   string          `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy string          `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	RestorePoint    string          `json:"restore_point,omitempty" mapstructure:"restore_point"`
	StoragePool     []*StoragePool  `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
}

func (e *Volume) Set(ro *VolumeSetRequest) (*Volume, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Volume{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type VolumeDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *Volume) Delete(ro *VolumeDeleteRequest) (*Volume, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &Volume{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
