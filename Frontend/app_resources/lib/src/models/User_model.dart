
import 'dart:convert';

import 'package:intl/intl.dart';

UserModel userModelFromJson( String str ) =>  UserModel.fromJson( json.decode( str ) );
class UserModel {
  final int? resourcedataid;
  final String? rolename;
  final String? firstname;
  final String? lastname;
  final String? image;
  final String? about;
  final String? cellphone;
  final String? personalemail;
  final String? datebirth;
  final String? address;
  final String? country;
  final String? city;
  final String? postalcode;
  final List<Hardskill>? hardskills;
  final List<Education>? educations;
  final List<Experience>? experiences;
  final List<Softskill>? softskills;
  final List<Language>? languages;
  final List<Certificate>? certificates;
  
  UserModel({
    this.resourcedataid,
    this.rolename,
    this.firstname,
    this.lastname,
    this.image,
    this.about,
    this.cellphone,
    this.personalemail,
    this.datebirth,
    this.address,
    this.country,
    this.city,
    this.postalcode,
    this.hardskills,
    this.educations,
    this.experiences,
    this.softskills,
    this.languages,
    this.certificates,
  });

  factory UserModel.fromJson( Map<String, dynamic> json ) => UserModel(
    resourcedataid: json['resourcedataid'],
    rolename: json['rolename'],
    firstname: json['firstname'],
    lastname: json['lastname'],
    image: json['Image'],
    about: json['about'],
    cellphone: json['cellphone'],
    personalemail: json['personalemail'],
    datebirth: DateFormat("dd-MM-yyyy").format(DateTime.parse(json['datebirth'])),
    address:json['address'],
    country: json['country'],
    city: json['city'],
    postalcode: json['cp'],
    hardskills: (json['hardskills'] as List ).map( (e) => hardskillFromJson(e) ).toList(),
    educations: (json['educations'] as List ).map( (e) => educationFromJson(e) ).toList(),
    experiences: (json['experiences'] as List ).map( (e) => experienceFromJson(e) ).toList(),
    softskills: (json['softskills'] as List ).map( (e) => softskillFromJson(e) ).toList(),
    languages: (json['languages'] as List ).map( (e) => languageFromJson(e) ).toList(),
    certificates: (json['certificates'] as List ).map( (e) => certificateFromJson(e) ).toList(),
  );
}

Hardskill hardskillFromJson( Map<String, dynamic> json ) => Hardskill.fromJson( json );
class Hardskill {
  final int? hardsresourceid;
  final int? domain;
  final String? hardskillname;
  final String? image;
  
  Hardskill({
    this.hardsresourceid,
    this.domain,
    this.hardskillname,
    this.image,
  });
  
  factory Hardskill.fromJson( Map<String, dynamic> json ) => Hardskill(
    hardsresourceid: json['hardsresourceid'],
    domain: json['domain'],
    hardskillname: json['hardskillname'],
    image: json['image'],
  );
}

Education educationFromJson( Map<String, dynamic> json ) => Education.fromJson( json );
class Education {
  final int? educationid;
  final String? academicname;
  final String? academictitle;
  final int? year;

  Education({
    this.educationid,
    this.academicname,
    this.academictitle,
    this.year
  });

  factory Education.fromJson( Map<String, dynamic> json ) => Education(
    educationid: json['educationid'],
    academicname: json['academicname'],
    academictitle: json['academictitle'],
    year: json['year'],
  );
}

Experience experienceFromJson( Map<String, dynamic> json ) => Experience.fromJson( json );
class Experience {
  final int? experienceid;
  final String? projecttitle;
  final String? description;
  final int? yearelaboration;

  Experience({
    this.experienceid,
    this.projecttitle,
    this.description,
    this.yearelaboration,
  });

  factory Experience.fromJson( Map<String, dynamic> json ) => Experience(
    experienceid: json['experienceid'],
    projecttitle: json['projecttitle'],
    description: json['description'],
    yearelaboration: json['yearelaboration'],
  );
}

Softskill softskillFromJson( Map<String, dynamic> json ) => Softskill.fromJson( json );
class Softskill {
  final int? softsresourceid;
  final String? softskillname;
  
  Softskill({
    this.softsresourceid,
    this.softskillname,
  });
  
  factory Softskill.fromJson( Map<String, dynamic> json ) =>  Softskill(
    softsresourceid: json['softsresourceid'],
    softskillname: json['softskillname'],
  );
}

Language languageFromJson( Map<String, dynamic> json ) => Language.fromJson( json );
class Language {
  final int? languageresourceid;
  final int? domain;
  final String? languagename;

  Language({
    this.languageresourceid,
    this.domain,
    this.languagename,
  });

  factory Language.fromJson( Map<String, dynamic> json ) => new Language(
    languageresourceid: json['languageresourceid'],
    domain: json['domain'],
    languagename: json['languagename'],
  );
}

Certificate certificateFromJson( Map<String, dynamic> json ) => Certificate.fromJson( json );
class Certificate {
  final int? certificateid;
  final String? idcertificate;
  final String? certificatename;
  final String? expedition;
  final String? expiration;

  Certificate({
    this.certificateid,
    this.idcertificate,
    this.certificatename,
    this.expedition,
    this.expiration,
  });

  factory Certificate.fromJson( Map<String, dynamic> json ) => new Certificate(
    certificateid: json['certificateid'],
    idcertificate: json['idcertificate'],
    certificatename: json['certificatename'],
    expedition: json['expedition'],
    expiration: json['expiration'],
  );
}