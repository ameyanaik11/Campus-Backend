// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: tumdev/campus_backend.proto

package api

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Campus_ListNewsAlerts_FullMethodName           = "/api.Campus/ListNewsAlerts"
	Campus_ListNewsSources_FullMethodName          = "/api.Campus/ListNewsSources"
	Campus_ListNews_FullMethodName                 = "/api.Campus/ListNews"
	Campus_ListCanteenRatings_FullMethodName       = "/api.Campus/ListCanteenRatings"
	Campus_GetDishRatings_FullMethodName           = "/api.Campus/GetDishRatings"
	Campus_CreateCanteenRating_FullMethodName      = "/api.Campus/CreateCanteenRating"
	Campus_CreateDishRating_FullMethodName         = "/api.Campus/CreateDishRating"
	Campus_ListAvailableDishTags_FullMethodName    = "/api.Campus/ListAvailableDishTags"
	Campus_ListNameTags_FullMethodName             = "/api.Campus/ListNameTags"
	Campus_ListAvailableCanteenTags_FullMethodName = "/api.Campus/ListAvailableCanteenTags"
	Campus_ListCanteens_FullMethodName             = "/api.Campus/ListCanteens"
	Campus_ListDishes_FullMethodName               = "/api.Campus/ListDishes"
	Campus_GetUpdateNote_FullMethodName            = "/api.Campus/GetUpdateNote"
	Campus_ListMovies_FullMethodName               = "/api.Campus/ListMovies"
	Campus_CreateFeedback_FullMethodName           = "/api.Campus/CreateFeedback"
	Campus_GetUploadStatus_FullMethodName          = "/api.Campus/GetUploadStatus"
	Campus_GetNotification_FullMethodName          = "/api.Campus/GetNotification"
	Campus_GetNotificationConfirm_FullMethodName   = "/api.Campus/GetNotificationConfirm"
	Campus_GetMember_FullMethodName                = "/api.Campus/GetMember"
	Campus_GetCanteenHeadCount_FullMethodName      = "/api.Campus/GetCanteenHeadCount"
	Campus_CreateDevice_FullMethodName             = "/api.Campus/CreateDevice"
	Campus_DeleteDevice_FullMethodName             = "/api.Campus/DeleteDevice"
)

// CampusClient is the client API for Campus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampusClient interface {
	ListNewsAlerts(ctx context.Context, in *ListNewsAlertsRequest, opts ...grpc.CallOption) (*ListNewsAlertsReply, error)
	ListNewsSources(ctx context.Context, in *ListNewsSourcesRequest, opts ...grpc.CallOption) (*ListNewsSourcesReply, error)
	ListNews(ctx context.Context, in *ListNewsRequest, opts ...grpc.CallOption) (*ListNewsReply, error)
	// This endpoint retrieves Canteen Ratings from the Backend.
	ListCanteenRatings(ctx context.Context, in *ListCanteenRatingsRequest, opts ...grpc.CallOption) (*ListCanteenRatingsReply, error)
	// Allows to query ratings for a specific dish in a specific cafeteria.
	GetDishRatings(ctx context.Context, in *GetDishRatingsRequest, opts ...grpc.CallOption) (*GetDishRatingsReply, error)
	CreateCanteenRating(ctx context.Context, in *CreateCanteenRatingRequest, opts ...grpc.CallOption) (*CreateCanteenRatingReply, error)
	CreateDishRating(ctx context.Context, in *CreateDishRatingRequest, opts ...grpc.CallOption) (*CreateDishRatingReply, error)
	ListAvailableDishTags(ctx context.Context, in *ListAvailableDishTagsRequest, opts ...grpc.CallOption) (*ListAvailableDishTagsReply, error)
	ListNameTags(ctx context.Context, in *ListNameTagsRequest, opts ...grpc.CallOption) (*ListNameTagsReply, error)
	ListAvailableCanteenTags(ctx context.Context, in *ListAvailableCanteenTagsRequest, opts ...grpc.CallOption) (*ListAvailableCanteenTagsReply, error)
	ListCanteens(ctx context.Context, in *ListCanteensRequest, opts ...grpc.CallOption) (*ListCanteensReply, error)
	// Returns all dishes for a specific cafeteria, year, week and day
	ListDishes(ctx context.Context, in *ListDishesRequest, opts ...grpc.CallOption) (*ListDishesReply, error)
	GetUpdateNote(ctx context.Context, in *GetUpdateNoteRequest, opts ...grpc.CallOption) (*GetUpdateNoteReply, error)
	ListMovies(ctx context.Context, in *ListMoviesRequest, opts ...grpc.CallOption) (*ListMoviesReply, error)
	CreateFeedback(ctx context.Context, opts ...grpc.CallOption) (Campus_CreateFeedbackClient, error)
	GetUploadStatus(ctx context.Context, in *GetUploadStatusRequest, opts ...grpc.CallOption) (*GetUploadStatusReply, error)
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationReply, error)
	GetNotificationConfirm(ctx context.Context, in *GetNotificationConfirmRequest, opts ...grpc.CallOption) (*GetNotificationConfirmReply, error)
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberReply, error)
	GetCanteenHeadCount(ctx context.Context, in *GetCanteenHeadCountRequest, opts ...grpc.CallOption) (*GetCanteenHeadCountReply, error)
	// Create an device (Android/iOS/Windows) for push notifications
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceReply, error)
	// Delete a device from push notifications
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceReply, error)
}

type campusClient struct {
	cc grpc.ClientConnInterface
}

func NewCampusClient(cc grpc.ClientConnInterface) CampusClient {
	return &campusClient{cc}
}

func (c *campusClient) ListNewsAlerts(ctx context.Context, in *ListNewsAlertsRequest, opts ...grpc.CallOption) (*ListNewsAlertsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNewsAlertsReply)
	err := c.cc.Invoke(ctx, Campus_ListNewsAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListNewsSources(ctx context.Context, in *ListNewsSourcesRequest, opts ...grpc.CallOption) (*ListNewsSourcesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNewsSourcesReply)
	err := c.cc.Invoke(ctx, Campus_ListNewsSources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListNews(ctx context.Context, in *ListNewsRequest, opts ...grpc.CallOption) (*ListNewsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNewsReply)
	err := c.cc.Invoke(ctx, Campus_ListNews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListCanteenRatings(ctx context.Context, in *ListCanteenRatingsRequest, opts ...grpc.CallOption) (*ListCanteenRatingsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCanteenRatingsReply)
	err := c.cc.Invoke(ctx, Campus_ListCanteenRatings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) GetDishRatings(ctx context.Context, in *GetDishRatingsRequest, opts ...grpc.CallOption) (*GetDishRatingsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDishRatingsReply)
	err := c.cc.Invoke(ctx, Campus_GetDishRatings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) CreateCanteenRating(ctx context.Context, in *CreateCanteenRatingRequest, opts ...grpc.CallOption) (*CreateCanteenRatingReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCanteenRatingReply)
	err := c.cc.Invoke(ctx, Campus_CreateCanteenRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) CreateDishRating(ctx context.Context, in *CreateDishRatingRequest, opts ...grpc.CallOption) (*CreateDishRatingReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDishRatingReply)
	err := c.cc.Invoke(ctx, Campus_CreateDishRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListAvailableDishTags(ctx context.Context, in *ListAvailableDishTagsRequest, opts ...grpc.CallOption) (*ListAvailableDishTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAvailableDishTagsReply)
	err := c.cc.Invoke(ctx, Campus_ListAvailableDishTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListNameTags(ctx context.Context, in *ListNameTagsRequest, opts ...grpc.CallOption) (*ListNameTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNameTagsReply)
	err := c.cc.Invoke(ctx, Campus_ListNameTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListAvailableCanteenTags(ctx context.Context, in *ListAvailableCanteenTagsRequest, opts ...grpc.CallOption) (*ListAvailableCanteenTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAvailableCanteenTagsReply)
	err := c.cc.Invoke(ctx, Campus_ListAvailableCanteenTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListCanteens(ctx context.Context, in *ListCanteensRequest, opts ...grpc.CallOption) (*ListCanteensReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCanteensReply)
	err := c.cc.Invoke(ctx, Campus_ListCanteens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListDishes(ctx context.Context, in *ListDishesRequest, opts ...grpc.CallOption) (*ListDishesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDishesReply)
	err := c.cc.Invoke(ctx, Campus_ListDishes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) GetUpdateNote(ctx context.Context, in *GetUpdateNoteRequest, opts ...grpc.CallOption) (*GetUpdateNoteReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUpdateNoteReply)
	err := c.cc.Invoke(ctx, Campus_GetUpdateNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) ListMovies(ctx context.Context, in *ListMoviesRequest, opts ...grpc.CallOption) (*ListMoviesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMoviesReply)
	err := c.cc.Invoke(ctx, Campus_ListMovies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) CreateFeedback(ctx context.Context, opts ...grpc.CallOption) (Campus_CreateFeedbackClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Campus_ServiceDesc.Streams[0], Campus_CreateFeedback_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &campusCreateFeedbackClient{ClientStream: stream}
	return x, nil
}

type Campus_CreateFeedbackClient interface {
	Send(*CreateFeedbackRequest) error
	CloseAndRecv() (*CreateFeedbackReply, error)
	grpc.ClientStream
}

type campusCreateFeedbackClient struct {
	grpc.ClientStream
}

func (x *campusCreateFeedbackClient) Send(m *CreateFeedbackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *campusCreateFeedbackClient) CloseAndRecv() (*CreateFeedbackReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateFeedbackReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *campusClient) GetUploadStatus(ctx context.Context, in *GetUploadStatusRequest, opts ...grpc.CallOption) (*GetUploadStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUploadStatusReply)
	err := c.cc.Invoke(ctx, Campus_GetUploadStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationReply)
	err := c.cc.Invoke(ctx, Campus_GetNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) GetNotificationConfirm(ctx context.Context, in *GetNotificationConfirmRequest, opts ...grpc.CallOption) (*GetNotificationConfirmReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationConfirmReply)
	err := c.cc.Invoke(ctx, Campus_GetNotificationConfirm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMemberReply)
	err := c.cc.Invoke(ctx, Campus_GetMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) GetCanteenHeadCount(ctx context.Context, in *GetCanteenHeadCountRequest, opts ...grpc.CallOption) (*GetCanteenHeadCountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCanteenHeadCountReply)
	err := c.cc.Invoke(ctx, Campus_GetCanteenHeadCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDeviceReply)
	err := c.cc.Invoke(ctx, Campus_CreateDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campusClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDeviceReply)
	err := c.cc.Invoke(ctx, Campus_DeleteDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampusServer is the server API for Campus service.
// All implementations must embed UnimplementedCampusServer
// for forward compatibility
type CampusServer interface {
	ListNewsAlerts(context.Context, *ListNewsAlertsRequest) (*ListNewsAlertsReply, error)
	ListNewsSources(context.Context, *ListNewsSourcesRequest) (*ListNewsSourcesReply, error)
	ListNews(context.Context, *ListNewsRequest) (*ListNewsReply, error)
	// This endpoint retrieves Canteen Ratings from the Backend.
	ListCanteenRatings(context.Context, *ListCanteenRatingsRequest) (*ListCanteenRatingsReply, error)
	// Allows to query ratings for a specific dish in a specific cafeteria.
	GetDishRatings(context.Context, *GetDishRatingsRequest) (*GetDishRatingsReply, error)
	CreateCanteenRating(context.Context, *CreateCanteenRatingRequest) (*CreateCanteenRatingReply, error)
	CreateDishRating(context.Context, *CreateDishRatingRequest) (*CreateDishRatingReply, error)
	ListAvailableDishTags(context.Context, *ListAvailableDishTagsRequest) (*ListAvailableDishTagsReply, error)
	ListNameTags(context.Context, *ListNameTagsRequest) (*ListNameTagsReply, error)
	ListAvailableCanteenTags(context.Context, *ListAvailableCanteenTagsRequest) (*ListAvailableCanteenTagsReply, error)
	ListCanteens(context.Context, *ListCanteensRequest) (*ListCanteensReply, error)
	// Returns all dishes for a specific cafeteria, year, week and day
	ListDishes(context.Context, *ListDishesRequest) (*ListDishesReply, error)
	GetUpdateNote(context.Context, *GetUpdateNoteRequest) (*GetUpdateNoteReply, error)
	ListMovies(context.Context, *ListMoviesRequest) (*ListMoviesReply, error)
	CreateFeedback(Campus_CreateFeedbackServer) error
	GetUploadStatus(context.Context, *GetUploadStatusRequest) (*GetUploadStatusReply, error)
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationReply, error)
	GetNotificationConfirm(context.Context, *GetNotificationConfirmRequest) (*GetNotificationConfirmReply, error)
	GetMember(context.Context, *GetMemberRequest) (*GetMemberReply, error)
	GetCanteenHeadCount(context.Context, *GetCanteenHeadCountRequest) (*GetCanteenHeadCountReply, error)
	// Create an device (Android/iOS/Windows) for push notifications
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceReply, error)
	// Delete a device from push notifications
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceReply, error)
	mustEmbedUnimplementedCampusServer()
}

// UnimplementedCampusServer must be embedded to have forward compatible implementations.
type UnimplementedCampusServer struct {
}

func (UnimplementedCampusServer) ListNewsAlerts(context.Context, *ListNewsAlertsRequest) (*ListNewsAlertsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewsAlerts not implemented")
}
func (UnimplementedCampusServer) ListNewsSources(context.Context, *ListNewsSourcesRequest) (*ListNewsSourcesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewsSources not implemented")
}
func (UnimplementedCampusServer) ListNews(context.Context, *ListNewsRequest) (*ListNewsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNews not implemented")
}
func (UnimplementedCampusServer) ListCanteenRatings(context.Context, *ListCanteenRatingsRequest) (*ListCanteenRatingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCanteenRatings not implemented")
}
func (UnimplementedCampusServer) GetDishRatings(context.Context, *GetDishRatingsRequest) (*GetDishRatingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDishRatings not implemented")
}
func (UnimplementedCampusServer) CreateCanteenRating(context.Context, *CreateCanteenRatingRequest) (*CreateCanteenRatingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCanteenRating not implemented")
}
func (UnimplementedCampusServer) CreateDishRating(context.Context, *CreateDishRatingRequest) (*CreateDishRatingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDishRating not implemented")
}
func (UnimplementedCampusServer) ListAvailableDishTags(context.Context, *ListAvailableDishTagsRequest) (*ListAvailableDishTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableDishTags not implemented")
}
func (UnimplementedCampusServer) ListNameTags(context.Context, *ListNameTagsRequest) (*ListNameTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNameTags not implemented")
}
func (UnimplementedCampusServer) ListAvailableCanteenTags(context.Context, *ListAvailableCanteenTagsRequest) (*ListAvailableCanteenTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableCanteenTags not implemented")
}
func (UnimplementedCampusServer) ListCanteens(context.Context, *ListCanteensRequest) (*ListCanteensReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCanteens not implemented")
}
func (UnimplementedCampusServer) ListDishes(context.Context, *ListDishesRequest) (*ListDishesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDishes not implemented")
}
func (UnimplementedCampusServer) GetUpdateNote(context.Context, *GetUpdateNoteRequest) (*GetUpdateNoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdateNote not implemented")
}
func (UnimplementedCampusServer) ListMovies(context.Context, *ListMoviesRequest) (*ListMoviesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMovies not implemented")
}
func (UnimplementedCampusServer) CreateFeedback(Campus_CreateFeedbackServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateFeedback not implemented")
}
func (UnimplementedCampusServer) GetUploadStatus(context.Context, *GetUploadStatusRequest) (*GetUploadStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadStatus not implemented")
}
func (UnimplementedCampusServer) GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedCampusServer) GetNotificationConfirm(context.Context, *GetNotificationConfirmRequest) (*GetNotificationConfirmReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationConfirm not implemented")
}
func (UnimplementedCampusServer) GetMember(context.Context, *GetMemberRequest) (*GetMemberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedCampusServer) GetCanteenHeadCount(context.Context, *GetCanteenHeadCountRequest) (*GetCanteenHeadCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCanteenHeadCount not implemented")
}
func (UnimplementedCampusServer) CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedCampusServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedCampusServer) mustEmbedUnimplementedCampusServer() {}

// UnsafeCampusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampusServer will
// result in compilation errors.
type UnsafeCampusServer interface {
	mustEmbedUnimplementedCampusServer()
}

func RegisterCampusServer(s grpc.ServiceRegistrar, srv CampusServer) {
	s.RegisterService(&Campus_ServiceDesc, srv)
}

func _Campus_ListNewsAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewsAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListNewsAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListNewsAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListNewsAlerts(ctx, req.(*ListNewsAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListNewsSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewsSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListNewsSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListNewsSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListNewsSources(ctx, req.(*ListNewsSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListNews(ctx, req.(*ListNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListCanteenRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCanteenRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListCanteenRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListCanteenRatings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListCanteenRatings(ctx, req.(*ListCanteenRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_GetDishRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDishRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetDishRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetDishRatings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetDishRatings(ctx, req.(*GetDishRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_CreateCanteenRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCanteenRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).CreateCanteenRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_CreateCanteenRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).CreateCanteenRating(ctx, req.(*CreateCanteenRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_CreateDishRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDishRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).CreateDishRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_CreateDishRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).CreateDishRating(ctx, req.(*CreateDishRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListAvailableDishTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableDishTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListAvailableDishTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListAvailableDishTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListAvailableDishTags(ctx, req.(*ListAvailableDishTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListNameTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNameTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListNameTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListNameTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListNameTags(ctx, req.(*ListNameTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListAvailableCanteenTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableCanteenTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListAvailableCanteenTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListAvailableCanteenTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListAvailableCanteenTags(ctx, req.(*ListAvailableCanteenTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListCanteens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCanteensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListCanteens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListCanteens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListCanteens(ctx, req.(*ListCanteensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListDishes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDishesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListDishes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListDishes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListDishes(ctx, req.(*ListDishesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_GetUpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetUpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetUpdateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetUpdateNote(ctx, req.(*GetUpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_ListMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).ListMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_ListMovies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).ListMovies(ctx, req.(*ListMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_CreateFeedback_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CampusServer).CreateFeedback(&campusCreateFeedbackServer{ServerStream: stream})
}

type Campus_CreateFeedbackServer interface {
	SendAndClose(*CreateFeedbackReply) error
	Recv() (*CreateFeedbackRequest, error)
	grpc.ServerStream
}

type campusCreateFeedbackServer struct {
	grpc.ServerStream
}

func (x *campusCreateFeedbackServer) SendAndClose(m *CreateFeedbackReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *campusCreateFeedbackServer) Recv() (*CreateFeedbackRequest, error) {
	m := new(CreateFeedbackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Campus_GetUploadStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetUploadStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetUploadStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetUploadStatus(ctx, req.(*GetUploadStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_GetNotificationConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetNotificationConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetNotificationConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetNotificationConfirm(ctx, req.(*GetNotificationConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_GetCanteenHeadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCanteenHeadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).GetCanteenHeadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_GetCanteenHeadCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).GetCanteenHeadCount(ctx, req.(*GetCanteenHeadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_CreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campus_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Campus_DeleteDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Campus_ServiceDesc is the grpc.ServiceDesc for Campus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Campus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Campus",
	HandlerType: (*CampusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNewsAlerts",
			Handler:    _Campus_ListNewsAlerts_Handler,
		},
		{
			MethodName: "ListNewsSources",
			Handler:    _Campus_ListNewsSources_Handler,
		},
		{
			MethodName: "ListNews",
			Handler:    _Campus_ListNews_Handler,
		},
		{
			MethodName: "ListCanteenRatings",
			Handler:    _Campus_ListCanteenRatings_Handler,
		},
		{
			MethodName: "GetDishRatings",
			Handler:    _Campus_GetDishRatings_Handler,
		},
		{
			MethodName: "CreateCanteenRating",
			Handler:    _Campus_CreateCanteenRating_Handler,
		},
		{
			MethodName: "CreateDishRating",
			Handler:    _Campus_CreateDishRating_Handler,
		},
		{
			MethodName: "ListAvailableDishTags",
			Handler:    _Campus_ListAvailableDishTags_Handler,
		},
		{
			MethodName: "ListNameTags",
			Handler:    _Campus_ListNameTags_Handler,
		},
		{
			MethodName: "ListAvailableCanteenTags",
			Handler:    _Campus_ListAvailableCanteenTags_Handler,
		},
		{
			MethodName: "ListCanteens",
			Handler:    _Campus_ListCanteens_Handler,
		},
		{
			MethodName: "ListDishes",
			Handler:    _Campus_ListDishes_Handler,
		},
		{
			MethodName: "GetUpdateNote",
			Handler:    _Campus_GetUpdateNote_Handler,
		},
		{
			MethodName: "ListMovies",
			Handler:    _Campus_ListMovies_Handler,
		},
		{
			MethodName: "GetUploadStatus",
			Handler:    _Campus_GetUploadStatus_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _Campus_GetNotification_Handler,
		},
		{
			MethodName: "GetNotificationConfirm",
			Handler:    _Campus_GetNotificationConfirm_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _Campus_GetMember_Handler,
		},
		{
			MethodName: "GetCanteenHeadCount",
			Handler:    _Campus_GetCanteenHeadCount_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _Campus_CreateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Campus_DeleteDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateFeedback",
			Handler:       _Campus_CreateFeedback_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tumdev/campus_backend.proto",
}
