import 'package:app_resources/src/global/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_rounded_progress_bar/flutter_rounded_progress_bar.dart';
import 'package:flutter_rounded_progress_bar/rounded_progress_bar_style.dart';

class ListaPorcentajes extends StatelessWidget {
  final String? text;
  final double? percent;
  const ListaPorcentajes({
    Key? key,
    this.text,
    this.percent,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Container(
      width: size.width *0.8,
      margin: EdgeInsets.only(top:15),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Container(
            child: Text(
              text!,
              style: TextStyle(
                fontSize: 16,
              ),
            ),
          ),
          Container(
            width: size.width * 0.6,
            child: RoundedProgressBar(
              height: 21.0,
              childCenter: Text(
                "${ percent!.toInt() }%",
                style: TextStyle(
                  color: kTextWhiteColor,
                  fontSize: 16,
                ),
              ),
              style: RoundedProgressBarStyle(
                colorProgress: Colors.green,
                borderWidth: 0, 
                widthShadow: 0
              ),
              percent: percent!,
            ),
          )
        ],
      ),
    );
  }
}