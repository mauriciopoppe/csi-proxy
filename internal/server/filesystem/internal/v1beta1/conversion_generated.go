// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubernetes-csi/csi-proxy/client/api/filesystem/v1beta1"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/filesystem/internal"
)

func autoConvert_v1beta1_IsMountPointRequest_To_internal_IsMountPointRequest(in *v1beta1.IsMountPointRequest, out *internal.IsMountPointRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1beta1_IsMountPointRequest_To_internal_IsMountPointRequest is an autogenerated conversion function.
func Convert_v1beta1_IsMountPointRequest_To_internal_IsMountPointRequest(in *v1beta1.IsMountPointRequest, out *internal.IsMountPointRequest) error {
	return autoConvert_v1beta1_IsMountPointRequest_To_internal_IsMountPointRequest(in, out)
}

func autoConvert_internal_IsMountPointRequest_To_v1beta1_IsMountPointRequest(in *internal.IsMountPointRequest, out *v1beta1.IsMountPointRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_internal_IsMountPointRequest_To_v1beta1_IsMountPointRequest is an autogenerated conversion function.
func Convert_internal_IsMountPointRequest_To_v1beta1_IsMountPointRequest(in *internal.IsMountPointRequest, out *v1beta1.IsMountPointRequest) error {
	return autoConvert_internal_IsMountPointRequest_To_v1beta1_IsMountPointRequest(in, out)
}

func autoConvert_v1beta1_IsMountPointResponse_To_internal_IsMountPointResponse(in *v1beta1.IsMountPointResponse, out *internal.IsMountPointResponse) error {
	out.Error = in.Error
	out.IsMountPoint = in.IsMountPoint
	return nil
}

// Convert_v1beta1_IsMountPointResponse_To_internal_IsMountPointResponse is an autogenerated conversion function.
func Convert_v1beta1_IsMountPointResponse_To_internal_IsMountPointResponse(in *v1beta1.IsMountPointResponse, out *internal.IsMountPointResponse) error {
	return autoConvert_v1beta1_IsMountPointResponse_To_internal_IsMountPointResponse(in, out)
}

func autoConvert_internal_IsMountPointResponse_To_v1beta1_IsMountPointResponse(in *internal.IsMountPointResponse, out *v1beta1.IsMountPointResponse) error {
	out.Error = in.Error
	out.IsMountPoint = in.IsMountPoint
	return nil
}

// Convert_internal_IsMountPointResponse_To_v1beta1_IsMountPointResponse is an autogenerated conversion function.
func Convert_internal_IsMountPointResponse_To_v1beta1_IsMountPointResponse(in *internal.IsMountPointResponse, out *v1beta1.IsMountPointResponse) error {
	return autoConvert_internal_IsMountPointResponse_To_v1beta1_IsMountPointResponse(in, out)
}

func autoConvert_v1beta1_LinkPathRequest_To_internal_LinkPathRequest(in *v1beta1.LinkPathRequest, out *internal.LinkPathRequest) error {
	out.SourcePath = in.SourcePath
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_v1beta1_LinkPathRequest_To_internal_LinkPathRequest is an autogenerated conversion function.
func Convert_v1beta1_LinkPathRequest_To_internal_LinkPathRequest(in *v1beta1.LinkPathRequest, out *internal.LinkPathRequest) error {
	return autoConvert_v1beta1_LinkPathRequest_To_internal_LinkPathRequest(in, out)
}

func autoConvert_internal_LinkPathRequest_To_v1beta1_LinkPathRequest(in *internal.LinkPathRequest, out *v1beta1.LinkPathRequest) error {
	out.SourcePath = in.SourcePath
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_internal_LinkPathRequest_To_v1beta1_LinkPathRequest is an autogenerated conversion function.
func Convert_internal_LinkPathRequest_To_v1beta1_LinkPathRequest(in *internal.LinkPathRequest, out *v1beta1.LinkPathRequest) error {
	return autoConvert_internal_LinkPathRequest_To_v1beta1_LinkPathRequest(in, out)
}

func autoConvert_v1beta1_LinkPathResponse_To_internal_LinkPathResponse(in *v1beta1.LinkPathResponse, out *internal.LinkPathResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_v1beta1_LinkPathResponse_To_internal_LinkPathResponse is an autogenerated conversion function.
func Convert_v1beta1_LinkPathResponse_To_internal_LinkPathResponse(in *v1beta1.LinkPathResponse, out *internal.LinkPathResponse) error {
	return autoConvert_v1beta1_LinkPathResponse_To_internal_LinkPathResponse(in, out)
}

func autoConvert_internal_LinkPathResponse_To_v1beta1_LinkPathResponse(in *internal.LinkPathResponse, out *v1beta1.LinkPathResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_internal_LinkPathResponse_To_v1beta1_LinkPathResponse is an autogenerated conversion function.
func Convert_internal_LinkPathResponse_To_v1beta1_LinkPathResponse(in *internal.LinkPathResponse, out *v1beta1.LinkPathResponse) error {
	return autoConvert_internal_LinkPathResponse_To_v1beta1_LinkPathResponse(in, out)
}

func autoConvert_v1beta1_MkdirRequest_To_internal_MkdirRequest(in *v1beta1.MkdirRequest, out *internal.MkdirRequest) error {
	out.Path = in.Path
	out.Context = internal.PathContext(in.Context)
	return nil
}

// Convert_v1beta1_MkdirRequest_To_internal_MkdirRequest is an autogenerated conversion function.
func Convert_v1beta1_MkdirRequest_To_internal_MkdirRequest(in *v1beta1.MkdirRequest, out *internal.MkdirRequest) error {
	return autoConvert_v1beta1_MkdirRequest_To_internal_MkdirRequest(in, out)
}

func autoConvert_internal_MkdirRequest_To_v1beta1_MkdirRequest(in *internal.MkdirRequest, out *v1beta1.MkdirRequest) error {
	out.Path = in.Path
	out.Context = v1beta1.PathContext(in.Context)
	return nil
}

// Convert_internal_MkdirRequest_To_v1beta1_MkdirRequest is an autogenerated conversion function.
func Convert_internal_MkdirRequest_To_v1beta1_MkdirRequest(in *internal.MkdirRequest, out *v1beta1.MkdirRequest) error {
	return autoConvert_internal_MkdirRequest_To_v1beta1_MkdirRequest(in, out)
}

func autoConvert_v1beta1_MkdirResponse_To_internal_MkdirResponse(in *v1beta1.MkdirResponse, out *internal.MkdirResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_v1beta1_MkdirResponse_To_internal_MkdirResponse is an autogenerated conversion function.
func Convert_v1beta1_MkdirResponse_To_internal_MkdirResponse(in *v1beta1.MkdirResponse, out *internal.MkdirResponse) error {
	return autoConvert_v1beta1_MkdirResponse_To_internal_MkdirResponse(in, out)
}

func autoConvert_internal_MkdirResponse_To_v1beta1_MkdirResponse(in *internal.MkdirResponse, out *v1beta1.MkdirResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_internal_MkdirResponse_To_v1beta1_MkdirResponse is an autogenerated conversion function.
func Convert_internal_MkdirResponse_To_v1beta1_MkdirResponse(in *internal.MkdirResponse, out *v1beta1.MkdirResponse) error {
	return autoConvert_internal_MkdirResponse_To_v1beta1_MkdirResponse(in, out)
}

func autoConvert_v1beta1_PathExistsRequest_To_internal_PathExistsRequest(in *v1beta1.PathExistsRequest, out *internal.PathExistsRequest) error {
	out.Path = in.Path
	out.Context = internal.PathContext(in.Context)
	return nil
}

// Convert_v1beta1_PathExistsRequest_To_internal_PathExistsRequest is an autogenerated conversion function.
func Convert_v1beta1_PathExistsRequest_To_internal_PathExistsRequest(in *v1beta1.PathExistsRequest, out *internal.PathExistsRequest) error {
	return autoConvert_v1beta1_PathExistsRequest_To_internal_PathExistsRequest(in, out)
}

func autoConvert_internal_PathExistsRequest_To_v1beta1_PathExistsRequest(in *internal.PathExistsRequest, out *v1beta1.PathExistsRequest) error {
	out.Path = in.Path
	out.Context = v1beta1.PathContext(in.Context)
	return nil
}

// Convert_internal_PathExistsRequest_To_v1beta1_PathExistsRequest is an autogenerated conversion function.
func Convert_internal_PathExistsRequest_To_v1beta1_PathExistsRequest(in *internal.PathExistsRequest, out *v1beta1.PathExistsRequest) error {
	return autoConvert_internal_PathExistsRequest_To_v1beta1_PathExistsRequest(in, out)
}

func autoConvert_v1beta1_PathExistsResponse_To_internal_PathExistsResponse(in *v1beta1.PathExistsResponse, out *internal.PathExistsResponse) error {
	out.Error = in.Error
	out.Exists = in.Exists
	return nil
}

// Convert_v1beta1_PathExistsResponse_To_internal_PathExistsResponse is an autogenerated conversion function.
func Convert_v1beta1_PathExistsResponse_To_internal_PathExistsResponse(in *v1beta1.PathExistsResponse, out *internal.PathExistsResponse) error {
	return autoConvert_v1beta1_PathExistsResponse_To_internal_PathExistsResponse(in, out)
}

func autoConvert_internal_PathExistsResponse_To_v1beta1_PathExistsResponse(in *internal.PathExistsResponse, out *v1beta1.PathExistsResponse) error {
	out.Error = in.Error
	out.Exists = in.Exists
	return nil
}

// Convert_internal_PathExistsResponse_To_v1beta1_PathExistsResponse is an autogenerated conversion function.
func Convert_internal_PathExistsResponse_To_v1beta1_PathExistsResponse(in *internal.PathExistsResponse, out *v1beta1.PathExistsResponse) error {
	return autoConvert_internal_PathExistsResponse_To_v1beta1_PathExistsResponse(in, out)
}

func autoConvert_v1beta1_RmdirRequest_To_internal_RmdirRequest(in *v1beta1.RmdirRequest, out *internal.RmdirRequest) error {
	out.Path = in.Path
	out.Context = internal.PathContext(in.Context)
	out.Force = in.Force
	return nil
}

// Convert_v1beta1_RmdirRequest_To_internal_RmdirRequest is an autogenerated conversion function.
func Convert_v1beta1_RmdirRequest_To_internal_RmdirRequest(in *v1beta1.RmdirRequest, out *internal.RmdirRequest) error {
	return autoConvert_v1beta1_RmdirRequest_To_internal_RmdirRequest(in, out)
}

func autoConvert_internal_RmdirRequest_To_v1beta1_RmdirRequest(in *internal.RmdirRequest, out *v1beta1.RmdirRequest) error {
	out.Path = in.Path
	out.Context = v1beta1.PathContext(in.Context)
	out.Force = in.Force
	return nil
}

// Convert_internal_RmdirRequest_To_v1beta1_RmdirRequest is an autogenerated conversion function.
func Convert_internal_RmdirRequest_To_v1beta1_RmdirRequest(in *internal.RmdirRequest, out *v1beta1.RmdirRequest) error {
	return autoConvert_internal_RmdirRequest_To_v1beta1_RmdirRequest(in, out)
}

func autoConvert_v1beta1_RmdirResponse_To_internal_RmdirResponse(in *v1beta1.RmdirResponse, out *internal.RmdirResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_v1beta1_RmdirResponse_To_internal_RmdirResponse is an autogenerated conversion function.
func Convert_v1beta1_RmdirResponse_To_internal_RmdirResponse(in *v1beta1.RmdirResponse, out *internal.RmdirResponse) error {
	return autoConvert_v1beta1_RmdirResponse_To_internal_RmdirResponse(in, out)
}

func autoConvert_internal_RmdirResponse_To_v1beta1_RmdirResponse(in *internal.RmdirResponse, out *v1beta1.RmdirResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_internal_RmdirResponse_To_v1beta1_RmdirResponse is an autogenerated conversion function.
func Convert_internal_RmdirResponse_To_v1beta1_RmdirResponse(in *internal.RmdirResponse, out *v1beta1.RmdirResponse) error {
	return autoConvert_internal_RmdirResponse_To_v1beta1_RmdirResponse(in, out)
}
