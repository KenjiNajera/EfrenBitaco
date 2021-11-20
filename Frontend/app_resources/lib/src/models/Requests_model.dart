import 'dart:convert';

RequestsModel requestsModelFromJson( String str ) => RequestsModel.fromJson(json.decode(str)); 

class RequestsModel {

  final int? requestid;
  final bool? status;
  final String? typerequestname;
  final String? description;
  final String? details;

  RequestsModel({
    this.requestid,
    this.status,
    this.typerequestname,
    this.description,
    this.details,
  });

  factory RequestsModel.fromJson(Map<String, dynamic> json) => new RequestsModel(
    requestid: json['requestid'],
    status: json['status'],
    typerequestname: json['typerequestname'],
    description: json['description'],
    details: json['details'],
  );
}


