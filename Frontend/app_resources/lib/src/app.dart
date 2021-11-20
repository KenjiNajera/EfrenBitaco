import 'package:app_resources/src/global/size_config.dart';
import 'package:app_resources/src/pages/auth/RecoveryPassword_page.dart';
import 'package:app_resources/src/services/Auth_service.dart';
import 'package:app_resources/src/services/Binnacle_service.dart';
import 'package:app_resources/src/services/Checks_service.dart';
import 'package:app_resources/src/services/User_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:provider/provider.dart';

import 'package:app_resources/src/pages/home_page.dart';
import 'package:app_resources/src/pages/auth/login_page.dart';
import 'package:app_resources/src/pages/auth/Landing_page.dart';
import 'package:app_resources/src/services/Request_service.dart';
import 'package:app_resources/src/services/PushNotifi_service.dart';

import 'global/constants.dart';

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

final GlobalKey<NavigatorState> navigatorKey = new GlobalKey<NavigatorState>();

class _MyAppState extends State<MyApp> {
  bool loggedin = false;

  @override
  void initState() {
    super.initState();
    // final pushProvider = new PushNotificationsService();
    // pushProvider.initNotifications();

    // pushProvider.mensajesStream.listen((data) {
    //   navigatorKey.currentState!.pushNamed('mensaje', arguments: data);
    // });
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        return OrientationBuilder(
          builder: (context, orientation) {
            SizeConfig().init(constraints, orientation);
            return MultiProvider(
              providers: [
                ChangeNotifierProvider(create: (_) => RequestService()),
                ChangeNotifierProvider(create: (_) => ChecksService()),
                ChangeNotifierProvider(create: (_) => BinnacleService()),
                ChangeNotifierProvider(create: (_) => UserService()),
                ChangeNotifierProvider(create: (_) => AuthService()),
              ],
              child: MaterialApp(
                title: 'Grupo GIT',
                navigatorKey: navigatorKey,
                theme: ThemeData(
                  primaryColor: kPrimaryColor,
                ),
                debugShowCheckedModeBanner: false,
                routes: {
                  '/': (BuildContext context) => LandingPage(),
                  '/login': (BuildContext context) => LoginPage(),
                  '/home': (BuildContext context) => HomePage(),
                  '/password': (BuildContext context) => RecoveryPassword(),
                  
                },
              ),
            );
          }
        );
      },
    );
  }
}
