import 'dart:convert';

String loginMovilModelToJson(LoginMovilModel data) => json.encode(data.toJson());

class LoginMovilModel {
  
  final String? email;
  final String? password;
  final String? tokenmovil;

  LoginMovilModel({
    this.email,
    this.password,
    this.tokenmovil,
  });

  Map<String, dynamic> toJson() => {
    "email": email,
    "password": password,
    "tokenmovil": tokenmovil,
  };

}