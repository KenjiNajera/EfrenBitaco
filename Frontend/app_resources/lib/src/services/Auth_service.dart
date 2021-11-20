import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;

import '../global/constants.dart';
import 'package:app_resources/src/models/LoginMovil_model.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/models/AuthenticateMovil_model.dart';

class AuthService with ChangeNotifier {
  final SharedPreference _sharedPreference = SharedPreference();

  Future<bool> login(LoginMovilModel login) async {
    final resp = await http.post(
      Uri.parse(url + authenticatemovil),
      body: loginMovilModelToJson(login),
      headers: headers,
    );

    final decodedData = authenticateMovilModelFromJson(resp.body);

    if (resp.statusCode != 200) {
      return false;
    }
    print(decodedData.userid);
    print(decodedData.resourcedataid);
    print(decodedData.roleid);
    print(decodedData.status);
    _sharedPreference.saveValueInt(decodedData.userid!, USERID);
    _sharedPreference.saveValueInt(decodedData.resourcedataid!, RESOURCEDATAID);
    _sharedPreference.saveValueInt(decodedData.roleid!, ROLEID);
    _sharedPreference.saveValueBoolean(decodedData.status!, USERSTATUS);
    _sharedPreference.saveValueBoolean(true, LOGGEDIN);

    notifyListeners();
    return true;
  }
}
