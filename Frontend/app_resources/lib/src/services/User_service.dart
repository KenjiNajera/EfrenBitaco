

import 'package:app_resources/src/global/constants.dart';
import 'package:http/http.dart' as http;
import 'package:app_resources/src/models/User_model.dart';
import 'package:flutter/cupertino.dart';

class UserService extends ChangeNotifier {
  UserModel _user = UserModel(
    resourcedataid: 0,
    rolename: "",
    firstname: "",
    lastname: "",
    image: "",
    about: "",
    cellphone: "",
    personalemail: "",
    datebirth: "",
    address: "",
    country: "",
    city: "",
    postalcode: "",
    hardskills: [],
    softskills: [],
    certificates: [],
    educations: [],
    experiences: [],
    languages: [],
  );

  UserModel get user =>  this._user;
  set user( UserModel value ){
    this._user = value;
    notifyListeners();
  }

  Future<void> getUserData(int id) async {
    try {
      final resp =  await http.get(Uri.parse(url + getDataUser + id.toString()),
        headers: headers,
      );

      // if( resp.statusCode > 299 ) return false;

      final decodedData = userModelFromJson(resp.body);

      this.user = decodedData;
      
    } catch (e) {
    
    }
  }
}