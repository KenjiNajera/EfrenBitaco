import 'package:app_resources/src/models/User_model.dart';
import 'package:flutter/material.dart';

class ListCertificates extends StatelessWidget {
  final Certificate? certificate;
  const ListCertificates({
    Key? key,
    this.certificate,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Container(
      width: size.width * 0.8,
      margin: EdgeInsets.symmetric(vertical: 10),
      child: Column(
        // mainAxisAlignment: MainAxisAlignment.end,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Container(
            margin: EdgeInsets.only(),
            child: Text(
              "● ${certificate!.certificatename}",
              style: TextStyle(
                fontSize: 14
              ),
            ),
          ),
          Container(
            margin: EdgeInsets.only(top: 8, left: 20),
            child: Text(
              "○ Expedition: ${certificate!.expedition} ${certificate!.expiration}",
              style: TextStyle(
                fontSize: 14
              ),
            ),
          ),
          Container(
            margin: EdgeInsets.only(top: 8, left: 20),
            child: Text(
              "○ ID: ${certificate!.idcertificate}",
              style: TextStyle(
                fontSize: 14
              ),
            ),
          )
        ],
      )
    );
  }
}