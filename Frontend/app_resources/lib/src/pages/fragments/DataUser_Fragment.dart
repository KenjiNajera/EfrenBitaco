import 'package:app_resources/src/models/User_model.dart';
import 'package:app_resources/src/pages/fragments/Widgets/ListCetificates.dart';
import 'package:flutter/material.dart';

import 'package:app_resources/src/pages/fragments/Widgets/ListaEducion.dart';
import 'package:app_resources/src/pages/fragments/Widgets/TitlesTextContainer.dart';
import 'package:app_resources/src/pages/fragments/Widgets/TextContainer.dart';

class DataUserFragment extends StatelessWidget {
  final UserModel? user;
  const DataUserFragment({
    Key? key,
    this.user,
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
            children: [
              
              // Datos
              TitlesTextContainer(
                title: "Datos",
              ),
              TextContainer(
                title: "About",
                info: "${user!.about != "" ? user!.about : "No hay información."}" ,
              ),
              TextContainer(
                title: "Email",
                info: "${user!.personalemail != "" ? user!.personalemail : "No hay información."}",
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: <Widget>[
                  TextContainer(
                    title: "Celular",
                    info: "${user!.cellphone != "" ? user!.cellphone : "No hay información."}",
                  ),
                  TextContainer(
                    title: "Fecha de nacimiento",
                    info: "${user!.datebirth != "" ? user!.datebirth : "No hay información."}",
                  ),
                ],
              ),
              TextContainer(
                title: "Dirección",
                info: "${user!.address != "" ? user!.address : "No hay informacion."}",
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: <Widget>[
                  TextContainer(
                    title: "Estado",
                    info: "${user!.country != "" ? user!.country : "No hay información."}",
                  ),
                  TextContainer(
                    title: "Ciudad",
                    info: "${user!.city != "" ? user!.city : "No hay información."}",
                  ),
                  TextContainer(
                    title: "Código postal",
                    info: "${user!.postalcode != "" ? user!.postalcode : "No hay información."}",
                  ),
                ],
              ),
              // Education
              TitlesTextContainer(
                title: "Educación",
              ),
              user!.educations!.isEmpty 
                ? Text("No hay información.")
                : Column(children: [...user!.educations!.map( ( e ) => ListaEducacion( education: e ) )]),
              
              // Certificates
              TitlesTextContainer(
                title: "Certificados",
              ),
              user!.certificates!.isEmpty 
              ? Text("No hay información.")
              : Column(children: [...user!.certificates!.map( ( e ) => ListCertificates( certificate: e ) )]),
            ],
          ),
        )
      ],
    );
  }
}