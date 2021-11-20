import 'package:app_resources/src/models/TypeRequest_model.dart';
import 'package:app_resources/src/services/Request_service.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../global/constants.dart';

class ComboBoxCustom extends StatefulWidget {
  final List<TypeRequestModel>? list;
  ComboBoxCustom({
    Key? key,
    @required this.list,
  }) : super(key: key);

  @override
  _ComboBoxCustomState createState() => _ComboBoxCustomState();
}

class _ComboBoxCustomState extends State<ComboBoxCustom> {
  String _dropdownValue = 'Seleccione una Opción';

  RequestService? _requestService;
  @override
  Widget build(BuildContext context) {
    _requestService = Provider.of<RequestService>( context );

    Size size = MediaQuery.of(context).size;

    return Container(
      width: size.width * 0.8,
      // height: 60,
      // decoration: BoxDecoration(
      //   color: kTextfieldColor,
      //   borderRadius: BorderRadius.circular(50.0)
      // ),
      margin: EdgeInsets.only(bottom: 80, top: 40),
      padding: EdgeInsets.symmetric(horizontal: 10, vertical:5),
      child: DropdownButtonFormField<TypeRequestModel>(
        hint: Text(_dropdownValue, style: TextStyle(color: kPrimaryColor),),
        value: _requestService!.typeRequestModel,
        validator: (value) => value == null ? "Seleccione una opción":  null,
        elevation: 16,
        icon: Container(
          margin: EdgeInsets.only(left: 26),
          child: Icon(
            Icons.arrow_drop_down,
            size: 30.0,
          ),
        ),
        style: TextStyle(color: kTextWhiteColor, fontSize: 20,),
        decoration: InputDecoration(
          isDense: true,
          filled: true,
          fillColor: kTextWhiteColor,
          border: OutlineInputBorder(
            borderSide: BorderSide(
              color: Colors.white
            ),
            borderRadius: BorderRadius.circular(50.0)
          ),
          errorStyle: TextStyle(
            color: Colors.red,
          ),
        ),
        onChanged: (TypeRequestModel? newValue) {
          _requestService!.typeRequestModel = newValue!;
        },
        items: widget.list!.length != null ? widget.list!.map((item) => DropdownMenuItem<TypeRequestModel>(
          value: item,
          child: Text(
            item.name!,
            style: TextStyle(
              color: kPrimaryColor
            ),
          ),
        )).toList() : [],
      )
    );
  }
}