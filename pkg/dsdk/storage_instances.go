package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageInstance struct {
	Path                 string               `json:"path,omitempty" mapstructure:"path"`
	Access               *Access              `json:"access,omitempty" mapstructure:"access"`
	AccessControlMode    string               `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AclPolicy            *AclPolicy           `json:"acl_policy,omitempty" mapstructure:"acl_policy"`
	ActiveInitiators     []*Initiator         `json:"active_initiators,omitempty" mapstructure:"active_initiators"`
	ActiveStorageNodes   []*StorageNode       `json:"active_storage_nodes,omitempty" mapstructure:"active_storage_nodes"`
	AdminState           string               `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Auth                 *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	Causes               []string             `json:"causes,omitempty" mapstructure:"causes"`
	DeploymentState      string               `json:"deployment_state,omitempty" mapstructure:"deployment_state"`
	Health               string               `json:"health,omitempty" mapstructure:"health"`
	IpPool               *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	Name                 string               `json:"name,omitempty" mapstructure:"name"`
	OpState              string               `json:"op_state,omitempty" mapstructure:"op_state"`
	ServiceConfiguration string               `json:"service_configuration,omitempty" mapstructure:"service_configuration"`
	Uuid                 string               `json:"uuid,omitempty" mapstructure:"uuid"`
	Volumes              []*Volume            `json:"volumes,omitempty" mapstructure:"volumes"`
	VolumesEp            *Volumes             `json:"-"`
}

type StorageInstances struct {
	Path string
}

type StorageInstancesCreateRequest struct {
	Ctxt                 context.Context      `json:"-"`
	AccessControlMode    string               `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AclPolicy            *AclPolicy           `json:"acl_policy,omitempty" mapstructure:"acl_policy"`
	AdminState           string               `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Auth                 *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	IpPool               *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	Name                 string               `json:"name,omitempty" mapstructure:"name"`
	ServiceConfiguration string               `json:"service_configuration,omitempty" mapstructure:"service_configuration"`
	Volumes              []*Volume            `json:"volumes,omitempty" mapstructure:"volumes"`
}

type StorageInstancesCreateResponse StorageInstance

func newStorageInstances(path string) *StorageInstances {
	return &StorageInstances{
		Path: _path.Join(path, "storage_instances"),
	}
}

func (e *StorageInstances) Create(ro *StorageInstancesCreateRequest) (*StorageInstancesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstancesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.VolumesEp = newVolumes(e.Path)
	return resp, nil
}

type StorageInstancesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type StorageInstancesListResponse []StorageInstance

func (e *StorageInstances) List(ro *StorageInstancesListRequest) (*StorageInstancesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := StorageInstancesListResponse{}
	for _, data := range rs.Data {
		elem := &StorageInstance{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, r := range resp {
		r.VolumesEp = newVolumes(e.Path)
	}
	return &resp, nil
}

type StorageInstancesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string
}

type StorageInstancesGetResponse StorageInstance

func (e *StorageInstances) Get(ro *StorageInstancesGetRequest) (*StorageInstancesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstancesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.VolumesEp = newVolumes(e.Path)
	return resp, nil
}

type StorageInstanceSetRequest struct {
	Ctxt              context.Context      `json:"-"`
	AccessControlMode string               `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AclPolicy         *AclPolicy           `json:"acl_policy,omitempty" mapstructure:"acl_policy"`
	AdminState        string               `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Auth              *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	Force             bool                 `json:"force,omitempty" mapstructure:"force"`
	IpPool            *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	Volumes           []*Volume            `json:"volumes,omitempty" mapstructure:"volumes"`
}

type StorageInstanceSetResponse StorageInstance

func (e *StorageInstance) Set(ro *StorageInstanceSetRequest) (*StorageInstanceSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstanceSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.VolumesEp = newVolumes(e.Path)
	return resp, nil

}

type StorageInstanceDeleteRequest struct {
	Ctxt  context.Context `json:"-"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

type StorageInstanceDeleteResponse StorageInstance

func (e *StorageInstance) Delete(ro *StorageInstanceDeleteRequest) (*StorageInstanceDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstanceDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.VolumesEp = newVolumes(e.Path)
	return resp, nil
}
