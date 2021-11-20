import 'package:app_resources/src/global/size_config.dart';
import 'package:flutter/material.dart';

import '../global/constants.dart';


class ButtonContainer extends StatelessWidget {
  final String? text;
  final Function()? onPressed;
  final Color? colorbg;
  final Color? colortext;
  const ButtonContainer({
    Key? key,
    this.text,
    this.onPressed,
    this.colorbg,
    this.colortext
  }) : super(key: key);
  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Container(
      margin: SizeConfig.isMobilePortrait ? EdgeInsets.only(top: 25, bottom: 30) : EdgeInsets.only(bottom: 15),
      width: size.width * 0.8,
      child: ClipRRect(
        borderRadius: BorderRadius.circular(3),
        child: TextButton(
          style: TextButton.styleFrom(
            padding: EdgeInsets.symmetric(vertical: 20, horizontal: 40),
            primary: Colors.white,
            backgroundColor: colorbg == null ? kTextWhiteColor : colorbg,
          ),
          onPressed: onPressed ,
          child: Text(
            text!,
            style: TextStyle(color: colortext == null ? kPrimaryColor : colortext),
          ),
        ),
      ),
    );
  }
}