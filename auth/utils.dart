import 'dart:convert';
import 'dart:math';

String CreateCryptoRandomString([int length = 32]) {
  final Random _random = Random.secure();
  var values = List<int>.generate(length, (i) => _random.nextInt(256));
  return base64Url.encode(values);
}
