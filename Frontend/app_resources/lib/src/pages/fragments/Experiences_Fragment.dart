import 'package:app_resources/src/models/User_model.dart';
import 'package:flutter/material.dart';

import 'package:app_resources/src/pages/fragments/Widgets/ListaExperiencia.dart';
import 'package:app_resources/src/pages/fragments/Widgets/TitlesTextContainer.dart';

class ExperienceFragment extends StatelessWidget {
  final List<Experience>? list;
  ExperienceFragment({
    Key? key,
    this.list
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return ListView(
      padding: EdgeInsets.zero,
      children: <Widget>[
        Container(
          padding: EdgeInsets.only(left: size.width * 0.07, right: size.width * 0.07, bottom:  size.width * 0.07),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              TitlesTextContainer(
                title: "Experiencia",
              ),
              list!.isEmpty 
              ? Text("No hay información.")
              : Column(children: [...list!.map( ( e ) => ListaExperiencia( experiencia: e ) )]),

            ],
          ),
        )
      ],
    );
  }
}