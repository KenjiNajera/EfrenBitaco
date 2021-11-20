

import 'dart:convert';

import 'package:intl/intl.dart';


BinnacleViewModel binnacleViewFromJson( String str ) => BinnacleViewModel.fromJson(json.decode(str));
class BinnacleViewModel {
  final int? binnacleid;
	final String? monthname;
  final int? year;
	final double? totalhours;
	final bool? status;       
	List<dynamic>? projects;   
  List<DayModel>? days;
  
  BinnacleViewModel({
    this.binnacleid,
    this.monthname,
    this.year,
    this.totalhours,
    this.status,
    this.projects,
    this.days,
  });

  factory BinnacleViewModel.fromJson( Map<String, dynamic> json ) => new BinnacleViewModel(
    binnacleid: json['binnacleid'],
    monthname: json['monthname'],
    year: json['year'],
    totalhours: json['totalhours'],
    status: json['status'],
    projects: (json['projectnames'] as List),
    days: (json['days'] as List).map((e) => dayDataFromJson(e)).toList(),
  );

  Map<String, dynamic> toJson() => {
    "binnacleid": binnacleid,
    "monthname": monthname,
    "year": year,
    "totalhours": totalhours,
    "status": status,
    "projects": projects,
    "days": days,
  };
}

DayModel dayDataFromJson( Map<String, dynamic> json ) => DayModel.fromJson(json);
class DayModel {
  final int? dayid;
	final String? dayname;
	final String? daydate;
	final String? totalhoursday;
	final String? checktime;
	final String? departuretime;
	final String? overtime;
	final String? mealtime;
	final String? summary;
	final String? invoice;

  DayModel({
    this.dayid,
    this.dayname,
    this.daydate,
    this.totalhoursday,
    this.checktime,
    this.departuretime,
    this.overtime,
    this.mealtime,
    this.summary,
    this.invoice,
  });

  factory DayModel.fromJson( Map<String, dynamic> json ) => new DayModel(
    dayid: json['dayid'],
    dayname: json['dayname'],
    daydate: DateFormat("dd-MM-yy").format(DateTime.parse(json['daydate'])),
    totalhoursday: DateFormat("HH:mm::ss").format(DateTime.parse(json['totalhoursday'])),
    checktime: DateFormat("HH:mm:ss").format(DateTime.parse(json['checktime'])),
    departuretime: DateFormat("HH:mm:ss").format(DateTime.parse(json['departuretime'])),
    overtime: DateFormat("HH:mm:ss").format(DateTime.parse(json['overtime'])),
    mealtime: DateFormat("HH:mm:ss").format(DateTime.parse(json['mealtime'])),
    summary: json['summary'],
    invoice: json['invoice'],
  );
}


