import 'package:app_resources/src/models/Binnacle_model.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;

import 'package:app_resources/src/global/constants.dart';

class BinnacleService with ChangeNotifier {

  BinnacleViewModel _binnacle = BinnacleViewModel(projects: [], days: []);

  BinnacleViewModel get binnacle => this._binnacle;
  set binnacle( BinnacleViewModel value ) {
    this._binnacle = value;
    notifyListeners();
  }

  Future<bool> getBinnacleView(int id) async {
    try{
      final resp =  await http.get(Uri.parse(url + getBinnacleUser + id.toString()),
        headers: headers,
      );

      if( resp.statusCode > 299 ) return false;

      final decodedData = binnacleViewFromJson(resp.body);

      this.binnacle = decodedData;

      return true;
    }
    catch( e ) {
      return false;
    }
  }
}