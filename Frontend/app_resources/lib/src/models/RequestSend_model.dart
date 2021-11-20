import 'dart:convert';

String requestSendModelToJson(RequestSendModel data) => json.encode(data.toJson());

class RequestSendModel {

  final String? description;
  final int? typerequestid;
  final int? userid;

  RequestSendModel({
    this.description,
    this.typerequestid,
    this.userid,
  });

  Map<String, dynamic> toJson() => {
    "description": description,
    "typerequestid": typerequestid,
    "userid": userid,
  };
}