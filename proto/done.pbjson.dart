///
//  Generated code. Do not modify.
//  source: proto/done.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const RegisterTokenRequest$json = const {
  '1': 'RegisterTokenRequest',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 9, '10': 'userId'},
  ],
};

const RegisterTokenResponse$json = const {
  '1': 'RegisterTokenResponse',
  '2': const [
    const {'1': 'Token', '3': 1, '4': 1, '5': 9, '10': 'Token'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.done.Error', '10': 'error'},
  ],
};

const DeleteTokenRequest$json = const {
  '1': 'DeleteTokenRequest',
  '2': const [
    const {'1': 'Token', '3': 1, '4': 1, '5': 9, '10': 'Token'},
  ],
};

const DeleteTokenResponse$json = const {
  '1': 'DeleteTokenResponse',
  '2': const [
    const {'1': 'deleted', '3': 1, '4': 1, '5': 8, '10': 'deleted'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.done.Error', '10': 'error'},
  ],
};

const ValidateTokenRequest$json = const {
  '1': 'ValidateTokenRequest',
  '2': const [
    const {'1': 'Token', '3': 1, '4': 1, '5': 9, '10': 'Token'},
  ],
};

const ValidateTokenResponse$json = const {
  '1': 'ValidateTokenResponse',
  '2': const [
    const {'1': 'valid', '3': 1, '4': 1, '5': 8, '10': 'valid'},
    const {'1': 'userId', '3': 2, '4': 1, '5': 9, '10': 'userId'},
    const {'1': 'error', '3': 3, '4': 1, '5': 11, '6': '.done.Error', '10': 'error'},
  ],
};

const Error$json = const {
  '1': 'Error',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 9, '10': 'code'},
    const {'1': 'json', '3': 2, '4': 1, '5': 9, '10': 'json'},
    const {'1': 'message', '3': 3, '4': 1, '5': 9, '10': 'message'},
  ],
};

