import 'package:flutter/material.dart';

class TitlesTextContainer extends StatelessWidget {
  final String? title;
  const TitlesTextContainer({
    Key? key,
    this.title,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(top: 20),
      child: Text(
        title!,
        style: TextStyle(
          fontSize: 25,
          fontWeight: FontWeight.w600
        ),
      ),
    );
  }
}