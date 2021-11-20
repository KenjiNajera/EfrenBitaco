import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:flutter/material.dart';

import '../../global/constants.dart';

class LandingPage extends StatefulWidget {
  LandingPage({Key? key}) : super(key: key);

  @override
  _LandingPageState createState() => _LandingPageState();
}

class _LandingPageState extends State<LandingPage> {
  final SharedPreference _sharedPreference = SharedPreference();

  @override
  void initState() {
    super.initState();
    Future.delayed( Duration( milliseconds: 1500 ),() => _loadUserInfo() );
  }
  _loadUserInfo() async {
    _sharedPreference.returnValueBoolean(LOGGEDIN).then((value) {
      if (value == false) {
        Navigator.pushNamedAndRemoveUntil(context, '/login', ModalRoute.withName('/login'));
      } else {
        Navigator.pushNamedAndRemoveUntil(context, '/home', ModalRoute.withName('/home'));
      }
    });
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: null,
      backgroundColor: kPrimaryColor,
      body: Center(
        child: Padding(
          padding: EdgeInsets.all(12.0),
          child: CircularProgressIndicator(
            backgroundColor: Colors.white38,
            valueColor: AlwaysStoppedAnimation(Colors.white),
          ),
        ),
      ),
    );
  }
}