import 'dart:io';

import 'package:redis/redis.dart';

class Redis {

  get redisHost {
    Map<String, String> env = Platform.environment;
    if (env.containsKey("REDIS_HOST")){
      return env['REDIS_HOST'];
    } else {
      return 'localhost';
    }
  }

  get redisPort {
    Map<String, String> env = Platform.environment;
    if (env.containsKey("REDIS_PORT")){
      return int.parse(env['REDIS_PORT']);
    } else {
      return 6379;
    }
  }

  Future<bool> SetToken(String token, String userId) async {
    RedisConnection conn = new RedisConnection();
    bool tokenSet = false;
    await conn.connect(this.redisHost, this.redisPort).then((Command command) async {
      await command.send_object(["SET", token, userId]).then((var response) async {
        tokenSet = response == 'OK';
        await command.send_object(["SADD", "user_tokens:" + userId, token]);
      });
    });
    return tokenSet;
  }

  Future<bool> DeleteToken(String token) async {
    RedisConnection conn = new RedisConnection();
    bool tokenDeleted = false;
    await conn.connect(this.redisHost, this.redisPort).then((Command command) async {
      await command.send_object(["GET", token]).then((var userId) async {
        if (userId != null){
          await command.send_object(["DEL", token]).then((var response) async {
            tokenDeleted = response == 1;
            await command.send_object(["SREM", "user_tokens:" + userId, token]);
          });
        }
      });
    });
    return tokenDeleted;
  }

  Future<String> ValidateToken(String token) async {
    RedisConnection conn = new RedisConnection();
    String userId;
    await conn.connect(this.redisHost, this.redisPort).then((Command command) async {
      await command.send_object(["GET", token]).then((var response) {
        userId = response != null? response:null;
      });
    });
    return userId;
  }

}
