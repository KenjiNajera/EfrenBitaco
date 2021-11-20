import 'dart:convert';

String checkInModelToJson( CheckInModel data ) => json.encode( data.toJson() );

class CheckInModel {

  final int? userid;
  final String? daydate;
  final String? checktime;

  CheckInModel({
    this.userid,
    this.daydate,
    this.checktime
  });

  Map<String, dynamic> toJson() => {
    "userid": userid,
    "daydate": daydate,
    "checktime": checktime,
  };
}