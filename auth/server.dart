import 'dart:async';

import 'package:grpc/grpc.dart';
import 'redis.dart';
import 'utils.dart';
import '../proto/done.pb.dart';
import '../proto/done.pbgrpc.dart';

class AuthenticationService extends AuthenticationServiceBase {

  @override
  Future<DeleteTokenResponse> deleteToken(ServiceCall call, DeleteTokenRequest request) async {
    Redis r = Redis();
    var res = DeleteTokenResponse();
    bool deleted = await r.DeleteToken(request.token);
    res.deleted = deleted;
    return res;
  }

  @override
  Future<RegisterTokenResponse> registerToken(ServiceCall call, RegisterTokenRequest request) async {
    String token = CreateCryptoRandomString();
    Redis r = Redis();
    await r.SetToken(token, request.userId);
    RegisterTokenResponse res = RegisterTokenResponse();
    res.token = token;
    return res;
  }

  @override
  Future<ValidateTokenResponse> validateToken(ServiceCall call, ValidateTokenRequest request) async {
    Redis r = Redis();
    String userId = await r.ValidateToken(request.token);
    ValidateTokenResponse res = ValidateTokenResponse();
    bool isValid = userId != null;
    res.valid = isValid;
    if (isValid){
      res.userId = userId;
    }
    return res;
  }
}

Future<void> main(List<String> args) async {
  final server = Server([AuthenticationService()]);
  await server.serve(port: 30052);
  print('Server listening on port ${server.port}...');
}