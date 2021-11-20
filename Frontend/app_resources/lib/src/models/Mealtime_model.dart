import 'dart:convert';

String mealTimeModelToJson( MealTimeModel data ) => json.encode( data.toJson() );

class MealTimeModel {

  final int? userid;
  final String? daydate;
  final String? mealtime;

  MealTimeModel({
    this.userid,
    this.daydate,
    this.mealtime,
  });

  Map<String, dynamic> toJson() => {
    "userid": userid,
    "daydate": daydate,
    "mealtime": mealtime,
  };
}

