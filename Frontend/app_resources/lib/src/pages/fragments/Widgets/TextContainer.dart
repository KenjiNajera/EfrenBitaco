import 'package:flutter/material.dart';

class TextContainer extends StatelessWidget {
  final String? title;
  final String? info;
  const TextContainer({
    Key? key,
    this.title,
    this.info,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Container(
            margin: EdgeInsets.only(top: 10),
            child: Text(
              title!,
              style: TextStyle(
                fontSize: 18,
              ),
            ),
          ),
          Container(
            margin: EdgeInsets.only(top: 3),
            child: Text(
              info!,
              style: TextStyle(
                fontSize: 14,
              ),
            ),
          )
        ],
      ),
    );
  }
}