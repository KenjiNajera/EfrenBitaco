import 'package:app_resources/src/global/constants.dart';
import 'package:app_resources/src/models/User_model.dart';
import 'package:app_resources/src/pages/fragments/Widgets/TextContainer.dart';
import 'package:flutter/material.dart';

class ListaExperiencia extends StatelessWidget {
  final Experience? experiencia;
  const ListaExperiencia({
    Key? key,
    this.experiencia,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(top: 15),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Divider(
            thickness: 2.0,
            color: kDividerColor,
          ),
          TextContainer(
            title: "Nombre del Proyecto",
            info: "${experiencia!.projecttitle}",
          ),
          TextContainer(
            title: "Descripción",
            info: "${experiencia!.description}"
          ),
          TextContainer(
            title: "Año de elaboración",
            info: "${experiencia!.yearelaboration}"
          ),
        ],
      ),
    );
  }
}