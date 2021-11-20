import 'package:app_resources/src/models/User_model.dart';
import 'package:app_resources/src/pages/fragments/Widgets/ListaPorcentajes.dart';
import 'package:app_resources/src/pages/fragments/Widgets/TitlesTextContainer.dart';
import 'package:flutter/material.dart';

class SkillsFragment extends StatelessWidget {
  final List<Hardskill>? listHardskill;
  final List<Softskill>? listSoftskill;
  final List<Language>? listLanguage;
  const SkillsFragment({
    Key? key,
    this.listHardskill,
    this.listSoftskill,
    this.listLanguage,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Container(
      padding: EdgeInsets.only(left: size.width * 0.07, right:  size.width * 0.07, bottom:  size.width * 0.07),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          
          // Softskills
          TitlesTextContainer(
            title: "Soft Skills" 
          ),
          listSoftskill!.isEmpty 
            ? Text("No hay información.")
            : Column(children: [...listSoftskill!.map( ( e ) => Container(
            margin: EdgeInsets.only(top: 12),
            child: Text(
              "• ${e.softskillname}",
              style: TextStyle(
                fontSize: 16
              ),
            ),
          ), )]),

          // Hardskills
          TitlesTextContainer(
            title: "Habilidades Técnicas" 
          ),
          listHardskill!.isEmpty 
            ? Text("No hay información.")
            : Column(children: [...listHardskill!.map( ( e ) => ListaPorcentajes( text: e.hardskillname, percent: e.domain!.toDouble(), ) )]),         

          // Languages
          TitlesTextContainer(
            title: "Idiomas" 
          ),
          listLanguage!.isEmpty 
            ? Text("No hay información.")
            : Column(children: [...listLanguage!.map( ( e ) => ListaPorcentajes( text: e.languagename, percent: e.domain!.toDouble(), ) )]),

        ],
      )
    );
  }
}