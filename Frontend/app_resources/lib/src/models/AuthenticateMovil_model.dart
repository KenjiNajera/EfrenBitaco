import 'dart:convert';

AuthenticateMovilModel authenticateMovilModelFromJson(String str) =>
    AuthenticateMovilModel.fromJson(json.decode(str));

class AuthenticateMovilModel {
  final int? userid;
  final int? resourcedataid;
  final int? roleid;
  final bool? status;

  AuthenticateMovilModel({
    this.userid,
    this.resourcedataid,
    this.roleid,
    this.status,
  });

  factory AuthenticateMovilModel.fromJson(Map<String, dynamic> json) =>
      new AuthenticateMovilModel(
        userid: json['id'],
        resourcedataid: json['resourcedataid'],
        roleid: json['roleid'],
        status: json['Status'],
      );
}
