///
//  Generated code. Do not modify.
//  source: proto/done.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class RegisterTokenRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegisterTokenRequest', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOS(1, 'userId', protoName: 'userId')
    ..hasRequiredFields = false
  ;

  RegisterTokenRequest._() : super();
  factory RegisterTokenRequest() => create();
  factory RegisterTokenRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegisterTokenRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegisterTokenRequest clone() => RegisterTokenRequest()..mergeFromMessage(this);
  RegisterTokenRequest copyWith(void Function(RegisterTokenRequest) updates) => super.copyWith((message) => updates(message as RegisterTokenRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegisterTokenRequest create() => RegisterTokenRequest._();
  RegisterTokenRequest createEmptyInstance() => create();
  static $pb.PbList<RegisterTokenRequest> createRepeated() => $pb.PbList<RegisterTokenRequest>();
  static RegisterTokenRequest getDefault() => _defaultInstance ??= create()..freeze();
  static RegisterTokenRequest _defaultInstance;

  $core.String get userId => $_getS(0, '');
  set userId($core.String v) { $_setString(0, v); }
  $core.bool hasUserId() => $_has(0);
  void clearUserId() => clearField(1);
}

class RegisterTokenResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegisterTokenResponse', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOS(1, 'Token', protoName: 'Token')
    ..a<Error>(2, 'error', $pb.PbFieldType.OM, defaultOrMaker: Error.getDefault, subBuilder: Error.create)
    ..hasRequiredFields = false
  ;

  RegisterTokenResponse._() : super();
  factory RegisterTokenResponse() => create();
  factory RegisterTokenResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegisterTokenResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegisterTokenResponse clone() => RegisterTokenResponse()..mergeFromMessage(this);
  RegisterTokenResponse copyWith(void Function(RegisterTokenResponse) updates) => super.copyWith((message) => updates(message as RegisterTokenResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegisterTokenResponse create() => RegisterTokenResponse._();
  RegisterTokenResponse createEmptyInstance() => create();
  static $pb.PbList<RegisterTokenResponse> createRepeated() => $pb.PbList<RegisterTokenResponse>();
  static RegisterTokenResponse getDefault() => _defaultInstance ??= create()..freeze();
  static RegisterTokenResponse _defaultInstance;

  $core.String get token => $_getS(0, '');
  set token($core.String v) { $_setString(0, v); }
  $core.bool hasToken() => $_has(0);
  void clearToken() => clearField(1);

  Error get error => $_getN(1);
  set error(Error v) { setField(2, v); }
  $core.bool hasError() => $_has(1);
  void clearError() => clearField(2);
}

class DeleteTokenRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DeleteTokenRequest', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOS(1, 'Token', protoName: 'Token')
    ..hasRequiredFields = false
  ;

  DeleteTokenRequest._() : super();
  factory DeleteTokenRequest() => create();
  factory DeleteTokenRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteTokenRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DeleteTokenRequest clone() => DeleteTokenRequest()..mergeFromMessage(this);
  DeleteTokenRequest copyWith(void Function(DeleteTokenRequest) updates) => super.copyWith((message) => updates(message as DeleteTokenRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteTokenRequest create() => DeleteTokenRequest._();
  DeleteTokenRequest createEmptyInstance() => create();
  static $pb.PbList<DeleteTokenRequest> createRepeated() => $pb.PbList<DeleteTokenRequest>();
  static DeleteTokenRequest getDefault() => _defaultInstance ??= create()..freeze();
  static DeleteTokenRequest _defaultInstance;

  $core.String get token => $_getS(0, '');
  set token($core.String v) { $_setString(0, v); }
  $core.bool hasToken() => $_has(0);
  void clearToken() => clearField(1);
}

class DeleteTokenResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DeleteTokenResponse', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOB(1, 'deleted')
    ..a<Error>(2, 'error', $pb.PbFieldType.OM, defaultOrMaker: Error.getDefault, subBuilder: Error.create)
    ..hasRequiredFields = false
  ;

  DeleteTokenResponse._() : super();
  factory DeleteTokenResponse() => create();
  factory DeleteTokenResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteTokenResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DeleteTokenResponse clone() => DeleteTokenResponse()..mergeFromMessage(this);
  DeleteTokenResponse copyWith(void Function(DeleteTokenResponse) updates) => super.copyWith((message) => updates(message as DeleteTokenResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteTokenResponse create() => DeleteTokenResponse._();
  DeleteTokenResponse createEmptyInstance() => create();
  static $pb.PbList<DeleteTokenResponse> createRepeated() => $pb.PbList<DeleteTokenResponse>();
  static DeleteTokenResponse getDefault() => _defaultInstance ??= create()..freeze();
  static DeleteTokenResponse _defaultInstance;

  $core.bool get deleted => $_get(0, false);
  set deleted($core.bool v) { $_setBool(0, v); }
  $core.bool hasDeleted() => $_has(0);
  void clearDeleted() => clearField(1);

  Error get error => $_getN(1);
  set error(Error v) { setField(2, v); }
  $core.bool hasError() => $_has(1);
  void clearError() => clearField(2);
}

class ValidateTokenRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ValidateTokenRequest', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOS(1, 'Token', protoName: 'Token')
    ..hasRequiredFields = false
  ;

  ValidateTokenRequest._() : super();
  factory ValidateTokenRequest() => create();
  factory ValidateTokenRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ValidateTokenRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ValidateTokenRequest clone() => ValidateTokenRequest()..mergeFromMessage(this);
  ValidateTokenRequest copyWith(void Function(ValidateTokenRequest) updates) => super.copyWith((message) => updates(message as ValidateTokenRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ValidateTokenRequest create() => ValidateTokenRequest._();
  ValidateTokenRequest createEmptyInstance() => create();
  static $pb.PbList<ValidateTokenRequest> createRepeated() => $pb.PbList<ValidateTokenRequest>();
  static ValidateTokenRequest getDefault() => _defaultInstance ??= create()..freeze();
  static ValidateTokenRequest _defaultInstance;

  $core.String get token => $_getS(0, '');
  set token($core.String v) { $_setString(0, v); }
  $core.bool hasToken() => $_has(0);
  void clearToken() => clearField(1);
}

class ValidateTokenResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ValidateTokenResponse', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOB(1, 'valid')
    ..aOS(2, 'userId', protoName: 'userId')
    ..a<Error>(3, 'error', $pb.PbFieldType.OM, defaultOrMaker: Error.getDefault, subBuilder: Error.create)
    ..hasRequiredFields = false
  ;

  ValidateTokenResponse._() : super();
  factory ValidateTokenResponse() => create();
  factory ValidateTokenResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ValidateTokenResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ValidateTokenResponse clone() => ValidateTokenResponse()..mergeFromMessage(this);
  ValidateTokenResponse copyWith(void Function(ValidateTokenResponse) updates) => super.copyWith((message) => updates(message as ValidateTokenResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ValidateTokenResponse create() => ValidateTokenResponse._();
  ValidateTokenResponse createEmptyInstance() => create();
  static $pb.PbList<ValidateTokenResponse> createRepeated() => $pb.PbList<ValidateTokenResponse>();
  static ValidateTokenResponse getDefault() => _defaultInstance ??= create()..freeze();
  static ValidateTokenResponse _defaultInstance;

  $core.bool get valid => $_get(0, false);
  set valid($core.bool v) { $_setBool(0, v); }
  $core.bool hasValid() => $_has(0);
  void clearValid() => clearField(1);

  $core.String get userId => $_getS(1, '');
  set userId($core.String v) { $_setString(1, v); }
  $core.bool hasUserId() => $_has(1);
  void clearUserId() => clearField(2);

  Error get error => $_getN(2);
  set error(Error v) { setField(3, v); }
  $core.bool hasError() => $_has(2);
  void clearError() => clearField(3);
}

class Error extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Error', package: const $pb.PackageName('done'), createEmptyInstance: create)
    ..aOS(1, 'code')
    ..aOS(2, 'json')
    ..aOS(3, 'message')
    ..hasRequiredFields = false
  ;

  Error._() : super();
  factory Error() => create();
  factory Error.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Error.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Error clone() => Error()..mergeFromMessage(this);
  Error copyWith(void Function(Error) updates) => super.copyWith((message) => updates(message as Error));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Error create() => Error._();
  Error createEmptyInstance() => create();
  static $pb.PbList<Error> createRepeated() => $pb.PbList<Error>();
  static Error getDefault() => _defaultInstance ??= create()..freeze();
  static Error _defaultInstance;

  $core.String get code => $_getS(0, '');
  set code($core.String v) { $_setString(0, v); }
  $core.bool hasCode() => $_has(0);
  void clearCode() => clearField(1);

  $core.String get json => $_getS(1, '');
  set json($core.String v) { $_setString(1, v); }
  $core.bool hasJson() => $_has(1);
  void clearJson() => clearField(2);

  $core.String get message => $_getS(2, '');
  set message($core.String v) { $_setString(2, v); }
  $core.bool hasMessage() => $_has(2);
  void clearMessage() => clearField(3);
}

