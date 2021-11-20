import 'package:app_resources/src/Widgets/ToastCustom.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/models/RequestSend_model.dart';
import 'package:app_resources/src/pages/home_page.dart';
import 'package:flutter/material.dart';

import 'package:provider/provider.dart';
import 'package:app_resources/src/Widgets/ButtonContainer.dart';
import 'package:app_resources/src/Widgets/ComboBoxCustom.dart';
import 'package:app_resources/src/Widgets/TextFielMultilineContianer.dart';
import 'package:app_resources/src/services/Request_service.dart';

import '../global/constants.dart';

class RequestPage extends StatefulWidget {
  @override
  _RequestPage createState() => _RequestPage();
}

class _RequestPage extends State<RequestPage> {
  final formKey = GlobalKey<FormState>();

  RequestService? _requestService;

  final textController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    _requestService = Provider.of<RequestService>(context);

    _getTypeRequest();
    return Container(
      child: Center(
        child: Form(
          key: formKey,
          child: ListView(
            children: <Widget>[
              Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  //Titulo
                  Container(
                    margin: EdgeInsets.symmetric(vertical: 30),
                    child: Text(
                      "Solicitudes",
                      style: TextStyle(
                        color: kTextWhiteColor,
                        fontSize: 50.0,
                      ),
                    ),
                  ),
                  // Select type request
                  ComboBoxCustom(
                    list: _requestService!.listTypeRequest,
                  ),
                  // Summary
                  TextFieldMultilineContainer(
                    textController: textController,
                    text: "Descripción...",
                    isvalidator: true,
                  ),
                  // Button Meal check
                  ButtonContainer(
                    text: "Enviar",
                    onPressed: sendRequest(),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }

  sendRequest() async {
    //if(!formKey.currentState!.validate()) return;
    if (formKey.currentState == null) {
      print("a");
    } else {
      formKey.currentState!.save();

      final SharedPreference _sharedPreference = SharedPreference();
      int userid = await _sharedPreference.returnValueInt(USERID);

      final request = RequestSendModel(
        typerequestid: _requestService!.typeRequestModel.id,
        description: textController.text,
        userid: userid,
      );
      bool resp = await _requestService!.sendRequest(request);

      if (resp) {
        // TODO: limpiar componentes y mostrar dialogo de que se envio correctamente.
        textController.clear();
        // ToastCustom().getToastSuccesful(context);
      } else {
        // TODO: mandar dialogo de que hubo un error durante el envio que lo intente mas tarde.
        // ToastCustom().getToastError(context);
      }
    }
  }

  void _getTypeRequest() {
    if (this._requestService!.listTypeRequest.isEmpty) {
      this._requestService!.getTypeRequest();
    }
  }
}
