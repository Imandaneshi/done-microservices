///
//  Generated code. Do not modify.
//  source: proto/done.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'done.pb.dart' as $0;
export 'done.pb.dart';

class AuthenticationServiceClient extends $grpc.Client {
  static final _$registerToken =
      $grpc.ClientMethod<$0.RegisterTokenRequest, $0.RegisterTokenResponse>(
          '/done.AuthenticationService/RegisterToken',
          ($0.RegisterTokenRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.RegisterTokenResponse.fromBuffer(value));
  static final _$deleteToken =
      $grpc.ClientMethod<$0.DeleteTokenRequest, $0.DeleteTokenResponse>(
          '/done.AuthenticationService/DeleteToken',
          ($0.DeleteTokenRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.DeleteTokenResponse.fromBuffer(value));
  static final _$validateToken =
      $grpc.ClientMethod<$0.ValidateTokenRequest, $0.ValidateTokenResponse>(
          '/done.AuthenticationService/ValidateToken',
          ($0.ValidateTokenRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.ValidateTokenResponse.fromBuffer(value));

  AuthenticationServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.RegisterTokenResponse> registerToken(
      $0.RegisterTokenRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$registerToken, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.DeleteTokenResponse> deleteToken(
      $0.DeleteTokenRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$deleteToken, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.ValidateTokenResponse> validateToken(
      $0.ValidateTokenRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$validateToken, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class AuthenticationServiceBase extends $grpc.Service {
  $core.String get $name => 'done.AuthenticationService';

  AuthenticationServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.RegisterTokenRequest, $0.RegisterTokenResponse>(
            'RegisterToken',
            registerToken_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.RegisterTokenRequest.fromBuffer(value),
            ($0.RegisterTokenResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.DeleteTokenRequest, $0.DeleteTokenResponse>(
            'DeleteToken',
            deleteToken_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.DeleteTokenRequest.fromBuffer(value),
            ($0.DeleteTokenResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.ValidateTokenRequest, $0.ValidateTokenResponse>(
            'ValidateToken',
            validateToken_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.ValidateTokenRequest.fromBuffer(value),
            ($0.ValidateTokenResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.RegisterTokenResponse> registerToken_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.RegisterTokenRequest> request) async {
    return registerToken(call, await request);
  }

  $async.Future<$0.DeleteTokenResponse> deleteToken_Pre($grpc.ServiceCall call,
      $async.Future<$0.DeleteTokenRequest> request) async {
    return deleteToken(call, await request);
  }

  $async.Future<$0.ValidateTokenResponse> validateToken_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.ValidateTokenRequest> request) async {
    return validateToken(call, await request);
  }

  $async.Future<$0.RegisterTokenResponse> registerToken(
      $grpc.ServiceCall call, $0.RegisterTokenRequest request);
  $async.Future<$0.DeleteTokenResponse> deleteToken(
      $grpc.ServiceCall call, $0.DeleteTokenRequest request);
  $async.Future<$0.ValidateTokenResponse> validateToken(
      $grpc.ServiceCall call, $0.ValidateTokenRequest request);
}
