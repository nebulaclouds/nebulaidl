# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from nebulaidl.service import dataproxy_pb2 as nebulaidl_dot_service_dot_dataproxy__pb2


class DataProxyServiceStub(object):
    """DataProxyService defines an RPC Service that allows access to user-data in a controlled manner.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateUploadLocation = channel.unary_unary(
                '/nebulaidl.service.DataProxyService/CreateUploadLocation',
                request_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateUploadLocationRequest.SerializeToString,
                response_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateUploadLocationResponse.FromString,
                )
        self.CreateDownloadLocation = channel.unary_unary(
                '/nebulaidl.service.DataProxyService/CreateDownloadLocation',
                request_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLocationRequest.SerializeToString,
                response_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLocationResponse.FromString,
                )
        self.CreateDownloadLink = channel.unary_unary(
                '/nebulaidl.service.DataProxyService/CreateDownloadLink',
                request_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLinkRequest.SerializeToString,
                response_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLinkResponse.FromString,
                )
        self.GetData = channel.unary_unary(
                '/nebulaidl.service.DataProxyService/GetData',
                request_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.GetDataRequest.SerializeToString,
                response_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.GetDataResponse.FromString,
                )


class DataProxyServiceServicer(object):
    """DataProxyService defines an RPC Service that allows access to user-data in a controlled manner.
    """

    def CreateUploadLocation(self, request, context):
        """CreateUploadLocation creates a signed url to upload artifacts to for a given project/domain.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateDownloadLocation(self, request, context):
        """CreateDownloadLocation creates a signed url to download artifacts.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateDownloadLink(self, request, context):
        """CreateDownloadLocation creates a signed url to download artifacts.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetData(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DataProxyServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateUploadLocation': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateUploadLocation,
                    request_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateUploadLocationRequest.FromString,
                    response_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateUploadLocationResponse.SerializeToString,
            ),
            'CreateDownloadLocation': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateDownloadLocation,
                    request_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLocationRequest.FromString,
                    response_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLocationResponse.SerializeToString,
            ),
            'CreateDownloadLink': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateDownloadLink,
                    request_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLinkRequest.FromString,
                    response_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLinkResponse.SerializeToString,
            ),
            'GetData': grpc.unary_unary_rpc_method_handler(
                    servicer.GetData,
                    request_deserializer=nebulaidl_dot_service_dot_dataproxy__pb2.GetDataRequest.FromString,
                    response_serializer=nebulaidl_dot_service_dot_dataproxy__pb2.GetDataResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'nebulaidl.service.DataProxyService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DataProxyService(object):
    """DataProxyService defines an RPC Service that allows access to user-data in a controlled manner.
    """

    @staticmethod
    def CreateUploadLocation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/nebulaidl.service.DataProxyService/CreateUploadLocation',
            nebulaidl_dot_service_dot_dataproxy__pb2.CreateUploadLocationRequest.SerializeToString,
            nebulaidl_dot_service_dot_dataproxy__pb2.CreateUploadLocationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateDownloadLocation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/nebulaidl.service.DataProxyService/CreateDownloadLocation',
            nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLocationRequest.SerializeToString,
            nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLocationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateDownloadLink(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/nebulaidl.service.DataProxyService/CreateDownloadLink',
            nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLinkRequest.SerializeToString,
            nebulaidl_dot_service_dot_dataproxy__pb2.CreateDownloadLinkResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetData(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/nebulaidl.service.DataProxyService/GetData',
            nebulaidl_dot_service_dot_dataproxy__pb2.GetDataRequest.SerializeToString,
            nebulaidl_dot_service_dot_dataproxy__pb2.GetDataResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
