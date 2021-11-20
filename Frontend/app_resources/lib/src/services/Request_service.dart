import 'dart:convert';

import 'package:app_resources/src/global/constants.dart';
import 'package:http/http.dart' as http;
import 'package:flutter/cupertino.dart';

import 'package:app_resources/src/models/RequestSend_model.dart';
import 'package:app_resources/src/models/Requests_model.dart';
import 'package:app_resources/src/models/TypeRequest_model.dart';

class RequestService with ChangeNotifier {
  TypeRequestModel? _typeRequestModel = null;
  List<TypeRequestModel> _listTypeRequest = [];
  List<RequestsModel> _listRequestsModel = [];

  List<TypeRequestModel> get listTypeRequest => this._listTypeRequest;
  set listTypeRequest(List<TypeRequestModel> value) {
    this._listTypeRequest = value;
    notifyListeners();
  }

  TypeRequestModel get typeRequestModel => this._typeRequestModel!;

  set typeRequestModel(TypeRequestModel value) {
    this._typeRequestModel = value;
    notifyListeners();
  }

  List<RequestsModel> get listRequestsModel => this._listRequestsModel;
  set listRequestsModel(List<RequestsModel> value) {
    this._listRequestsModel = value;
    notifyListeners();
  }

  void getTypeRequest() async {
    try {
      final resp = await http.get(
        Uri.parse(url + typeRequests),
        headers: headers,
      );

      final List<dynamic> decodedData = json.decode(resp.body);
      final List<TypeRequestModel> typeRequestsList = [];

      if (decodedData == null) return;

      decodedData.forEach((value) {
        final typerequestTemp = TypeRequestModel.fronJson(value);
        typeRequestsList.add(typerequestTemp);
      });

      this.listTypeRequest = typeRequestsList;
      notifyListeners();
      return;
    } catch (e) {
      return;
    }
  }

  Future<bool> sendRequest(RequestSendModel request) async {
    final resp = await http.post(
      Uri.parse(url + sendRequests),
      body: requestSendModelToJson(request),
      headers: headers,
    );

    if (resp.statusCode > 299) return false;

    return true;
  }

  void getRequests(int id) async {
    try {
      final resp = await http.get(
        Uri.parse(url + getRequestsUser + id.toString()),
        headers: headers,
      );

      final List<dynamic> decodedData = json.decode(resp.body);
      final List<RequestsModel> listTemp = [];

      decodedData.forEach((element) {
        final newValue = RequestsModel.fromJson(element);
        listTemp.add(newValue);
      });

      this.listRequestsModel = listTemp;
      notifyListeners();

      return;
    } catch (e) {
      return;
    }
  }
}
