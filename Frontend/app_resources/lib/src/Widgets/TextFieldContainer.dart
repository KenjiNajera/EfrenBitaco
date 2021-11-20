import 'dart:async';

import 'package:app_resources/src/global/size_config.dart';
import 'package:email_validator/email_validator.dart';
import 'package:flutter/material.dart';

import '../global/constants.dart';

class TextFieldContainer extends StatefulWidget {
  final String? text;
  final IconData? icon;
  final TextEditingController? controller;
  final TextInputType? keyboard;
  const TextFieldContainer({
    Key? key,
    this.controller,
    this.text,
    this.keyboard,
    this.icon,
  }) : super(key: key);

  @override
  _TextFieldContainerState createState() => _TextFieldContainerState();
}

class _TextFieldContainerState extends State<TextFieldContainer> {
  String? _valorGeneral;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin:
          EdgeInsets.symmetric(vertical: SizeConfig.isMobilePortrait ? 15 : 0),
      constraints: BoxConstraints(
          maxHeight: 40 *
              (SizeConfig.isMobilePortrait
                  ? SizeConfig.heightMultiplier!
                  : SizeConfig.widthMultiplier!)),
      child: TextFormField(
          keyboardType: widget.keyboard,
          controller: widget.controller,
          validator: (String? value) {
            if (value == null || value.isEmpty) {
              return 'Please enter some text';
            }
            return null;
          },
          // validator: (value) {
          //     if( value!.isEmpty ) return "Llene este campo";
          //     if( !EmailValidator.validate(value) ) return "Ingrese un email valido";
          // },
          onSaved: (value) => _valorGeneral = value,
          style: TextStyle(fontSize: 18.0, color: Colors.black),
          decoration: InputDecoration(
            contentPadding: EdgeInsets.only(left: 23.0, top: 37.0),
            filled: true,
            fillColor: kTextWhiteColor,
            prefixIcon: Icon(
              widget.icon,
              size: 30.0,
              color: kPrimaryColor,
            ),
            hintText: widget.text,
            errorStyle: TextStyle(
              color: Colors.red,
            ),
            border:
                OutlineInputBorder(borderRadius: BorderRadius.circular(5.0)),
          )),
    );
  }
}
