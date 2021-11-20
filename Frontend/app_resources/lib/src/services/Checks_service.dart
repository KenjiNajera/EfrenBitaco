import 'package:http/http.dart' as http;
import 'package:flutter/cupertino.dart';

import 'package:app_resources/src/global/constants.dart';
import 'package:app_resources/src/models/CheckIn_model.dart';
import 'package:app_resources/src/models/CheckOut_model.dart';
import 'package:app_resources/src/models/Mealtime_model.dart';

class ChecksService with ChangeNotifier {

  Future<bool> sendCheckIn(CheckInModel checkIn) async{
    try {
      final resp = await http.post(Uri.parse(url+sendCheckInUser),
        body: checkInModelToJson(checkIn),
        headers: headers,
      );

      if( resp.statusCode > 299 ) return false;

      return true;

    } catch( e ) {
      return false;
    }
  }

  Future<bool> sendCheckMeal( MealTimeModel mealtime ) async {
    try {      
      final resp = await http.post(Uri.parse(url+sendCheckMealUser),
        body: mealTimeModelToJson( mealtime ),
        headers: headers,
      );

      if( resp.statusCode > 299 ) return false;

      return true;
    } catch (e) {
      return false;
    }
  }

  Future<bool> sendCheckOut( CheckOutModel checkout ) async {
    try {
      final resp = await http.post(Uri.parse(url+sendCheckOutUser),
        body: checkOutModelToJson( checkout ),
        headers: headers,
      );

      if( resp.statusCode > 299 ) return false;

      return true;
    } catch (e) {
      return false;
    }
  }
}