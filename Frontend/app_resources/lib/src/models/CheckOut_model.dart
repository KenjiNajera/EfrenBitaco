import 'dart:convert';

String checkOutModelToJson( CheckOutModel data ) => json.encode( data.toJson() );

class CheckOutModel {

  final int? userid;
  final String? daydate;
  final String? departuretime;
  final bool? overtime;
  final String? summary;

  CheckOutModel({
    this.userid,
    this.daydate,
    this.departuretime,
    this.overtime,
    this.summary,
  });

  Map<String, dynamic> toJson() => {
    "userid": userid,
    "daydate": daydate,
    "departuretime": departuretime,
    "overtime": overtime,
    "summary": summary,
  };
}
