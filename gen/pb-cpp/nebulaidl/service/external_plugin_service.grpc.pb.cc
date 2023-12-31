// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: nebulaidl/service/external_plugin_service.proto

#include "nebulaidl/service/external_plugin_service.pb.h"
#include "nebulaidl/service/external_plugin_service.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace nebulaidl {
namespace service {

static const char* ExternalPluginService_method_names[] = {
  "/nebulaidl.service.ExternalPluginService/CreateTask",
  "/nebulaidl.service.ExternalPluginService/GetTask",
  "/nebulaidl.service.ExternalPluginService/DeleteTask",
};

std::unique_ptr< ExternalPluginService::Stub> ExternalPluginService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< ExternalPluginService::Stub> stub(new ExternalPluginService::Stub(channel));
  return stub;
}

ExternalPluginService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_CreateTask_(ExternalPluginService_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetTask_(ExternalPluginService_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_DeleteTask_(ExternalPluginService_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status ExternalPluginService::Stub::CreateTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskCreateRequest& request, ::nebulaidl::service::TaskCreateResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_CreateTask_, context, request, response);
}

void ExternalPluginService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskCreateRequest* request, ::nebulaidl::service::TaskCreateResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, std::move(f));
}

void ExternalPluginService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::TaskCreateResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, std::move(f));
}

void ExternalPluginService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskCreateRequest* request, ::nebulaidl::service::TaskCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, reactor);
}

void ExternalPluginService::Stub::experimental_async::CreateTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::TaskCreateResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateTask_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::nebulaidl::service::TaskCreateResponse>* ExternalPluginService::Stub::AsyncCreateTaskRaw(::grpc::ClientContext* context, const ::nebulaidl::service::TaskCreateRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::nebulaidl::service::TaskCreateResponse>::Create(channel_.get(), cq, rpcmethod_CreateTask_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::nebulaidl::service::TaskCreateResponse>* ExternalPluginService::Stub::PrepareAsyncCreateTaskRaw(::grpc::ClientContext* context, const ::nebulaidl::service::TaskCreateRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::nebulaidl::service::TaskCreateResponse>::Create(channel_.get(), cq, rpcmethod_CreateTask_, context, request, false);
}

::grpc::Status ExternalPluginService::Stub::GetTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskGetRequest& request, ::nebulaidl::service::TaskGetResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetTask_, context, request, response);
}

void ExternalPluginService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskGetRequest* request, ::nebulaidl::service::TaskGetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, std::move(f));
}

void ExternalPluginService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::TaskGetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, std::move(f));
}

void ExternalPluginService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskGetRequest* request, ::nebulaidl::service::TaskGetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, reactor);
}

void ExternalPluginService::Stub::experimental_async::GetTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::TaskGetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetTask_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::nebulaidl::service::TaskGetResponse>* ExternalPluginService::Stub::AsyncGetTaskRaw(::grpc::ClientContext* context, const ::nebulaidl::service::TaskGetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::nebulaidl::service::TaskGetResponse>::Create(channel_.get(), cq, rpcmethod_GetTask_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::nebulaidl::service::TaskGetResponse>* ExternalPluginService::Stub::PrepareAsyncGetTaskRaw(::grpc::ClientContext* context, const ::nebulaidl::service::TaskGetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::nebulaidl::service::TaskGetResponse>::Create(channel_.get(), cq, rpcmethod_GetTask_, context, request, false);
}

::grpc::Status ExternalPluginService::Stub::DeleteTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskDeleteRequest& request, ::nebulaidl::service::TaskDeleteResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_DeleteTask_, context, request, response);
}

void ExternalPluginService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskDeleteRequest* request, ::nebulaidl::service::TaskDeleteResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, std::move(f));
}

void ExternalPluginService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::TaskDeleteResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, std::move(f));
}

void ExternalPluginService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::nebulaidl::service::TaskDeleteRequest* request, ::nebulaidl::service::TaskDeleteResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, reactor);
}

void ExternalPluginService::Stub::experimental_async::DeleteTask(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::TaskDeleteResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_DeleteTask_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::nebulaidl::service::TaskDeleteResponse>* ExternalPluginService::Stub::AsyncDeleteTaskRaw(::grpc::ClientContext* context, const ::nebulaidl::service::TaskDeleteRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::nebulaidl::service::TaskDeleteResponse>::Create(channel_.get(), cq, rpcmethod_DeleteTask_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::nebulaidl::service::TaskDeleteResponse>* ExternalPluginService::Stub::PrepareAsyncDeleteTaskRaw(::grpc::ClientContext* context, const ::nebulaidl::service::TaskDeleteRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::nebulaidl::service::TaskDeleteResponse>::Create(channel_.get(), cq, rpcmethod_DeleteTask_, context, request, false);
}

ExternalPluginService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      ExternalPluginService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< ExternalPluginService::Service, ::nebulaidl::service::TaskCreateRequest, ::nebulaidl::service::TaskCreateResponse>(
          std::mem_fn(&ExternalPluginService::Service::CreateTask), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      ExternalPluginService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< ExternalPluginService::Service, ::nebulaidl::service::TaskGetRequest, ::nebulaidl::service::TaskGetResponse>(
          std::mem_fn(&ExternalPluginService::Service::GetTask), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      ExternalPluginService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< ExternalPluginService::Service, ::nebulaidl::service::TaskDeleteRequest, ::nebulaidl::service::TaskDeleteResponse>(
          std::mem_fn(&ExternalPluginService::Service::DeleteTask), this)));
}

ExternalPluginService::Service::~Service() {
}

::grpc::Status ExternalPluginService::Service::CreateTask(::grpc::ServerContext* context, const ::nebulaidl::service::TaskCreateRequest* request, ::nebulaidl::service::TaskCreateResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status ExternalPluginService::Service::GetTask(::grpc::ServerContext* context, const ::nebulaidl::service::TaskGetRequest* request, ::nebulaidl::service::TaskGetResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status ExternalPluginService::Service::DeleteTask(::grpc::ServerContext* context, const ::nebulaidl::service::TaskDeleteRequest* request, ::nebulaidl::service::TaskDeleteResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace nebulaidl
}  // namespace service

