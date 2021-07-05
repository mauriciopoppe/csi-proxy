// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/kubernetes-csi/csi-proxy/client/api/filesystem/v1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

var version = apiversion.NewVersionOrPanic("v1")

type versionedAPI struct {
	apiGroupServer impl.ServerInterface
}

func NewVersionedServer(apiGroupServer impl.ServerInterface) impl.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1.RegisterFilesystemServer(grpcServer, s)
}

func (s *versionedAPI) CreateSymlink(context context.Context, versionedRequest *v1.CreateSymlinkRequest) (*v1.CreateSymlinkResponse, error) {
	request := &impl.CreateSymlinkRequest{}
	if err := Convert_v1_CreateSymlinkRequest_To_impl_CreateSymlinkRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.CreateSymlink(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.CreateSymlinkResponse{}
	if err := Convert_impl_CreateSymlinkResponse_To_v1_CreateSymlinkResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) IsSymlink(context context.Context, versionedRequest *v1.IsSymlinkRequest) (*v1.IsSymlinkResponse, error) {
	request := &impl.IsSymlinkRequest{}
	if err := Convert_v1_IsSymlinkRequest_To_impl_IsSymlinkRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.IsSymlink(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.IsSymlinkResponse{}
	if err := Convert_impl_IsSymlinkResponse_To_v1_IsSymlinkResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) Mkdir(context context.Context, versionedRequest *v1.MkdirRequest) (*v1.MkdirResponse, error) {
	request := &impl.MkdirRequest{}
	if err := Convert_v1_MkdirRequest_To_impl_MkdirRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.Mkdir(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.MkdirResponse{}
	if err := Convert_impl_MkdirResponse_To_v1_MkdirResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) PathExists(context context.Context, versionedRequest *v1.PathExistsRequest) (*v1.PathExistsResponse, error) {
	request := &impl.PathExistsRequest{}
	if err := Convert_v1_PathExistsRequest_To_impl_PathExistsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.PathExists(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.PathExistsResponse{}
	if err := Convert_impl_PathExistsResponse_To_v1_PathExistsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	klog.Infof("versioned response=%v", versionedResponse)
	return versionedResponse, err
}

func (s *versionedAPI) Rmdir(context context.Context, versionedRequest *v1.RmdirRequest) (*v1.RmdirResponse, error) {
	request := &impl.RmdirRequest{}
	if err := Convert_v1_RmdirRequest_To_impl_RmdirRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.Rmdir(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.RmdirResponse{}
	if err := Convert_impl_RmdirResponse_To_v1_RmdirResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
