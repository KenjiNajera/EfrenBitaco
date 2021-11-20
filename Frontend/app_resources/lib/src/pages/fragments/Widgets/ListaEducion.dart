import 'package:app_resources/src/models/User_model.dart';
import 'package:flutter/material.dart';

class ListaEducacion extends StatelessWidget {
  final Education? education;

  const ListaEducacion({
    Key? key,
    this.education,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.symmetric(vertical: 10),
      child:Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              Text(
                "${education!.academicname}",
                style: TextStyle(
                  fontSize: 14,
                ),
              ),
              Text(
                "${education!.year}",
                style: TextStyle(
                  fontSize: 14,
                ),
              ),
            ],
          ),
          Text(
            "${education!.academictitle}",
            style: TextStyle(
              fontSize: 14,
            ),
          ),
        ],
      ),
    );
  }
}