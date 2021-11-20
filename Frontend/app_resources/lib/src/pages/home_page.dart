import 'package:app_resources/src/app.dart';
import 'package:app_resources/src/global/size_config.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/pages/Binnacle_page.dart';
import 'package:app_resources/src/pages/CheckIn_page.dart';
import 'package:app_resources/src/pages/CheckOut_page.dart';
import 'package:app_resources/src/pages/ListRequest_page.dart';
import 'package:app_resources/src/pages/Mealtime_page.dart';
import 'package:app_resources/src/pages/Profile_page.dart';
import 'package:app_resources/src/pages/Request_page.dart';
import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';

import 'package:app_resources/src/Widgets/ButtonContainer.dart';
import '../global/constants.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _page = 0;

  SharedPreference _sharedPreference = SharedPreference();

  GlobalKey<ScaffoldState> mScaffoldState = GlobalKey<ScaffoldState>();
  GlobalKey bottomNavigationBar = GlobalKey();

  Widget _callPage(int actuallyPage) {
    switch (actuallyPage) {
      case 0:
        return _homeWidget();
      case 1:
        return ListRequestPage();
      case 2:
        return BinnaclePage();
      case 3:
        return RequestPage();
      case 4:
        return ProfilePage();
      case 5:
        return Container();
      case 6:
        return CheckInPage();
      case 7:
        return MealtimePage();
      case 8:
        return CheckOutPage();
      default:
        return _homeWidget();
    }
  }

  setStatePage(int index) => setState(() => _page = index);

  Widget _homeWidget() => Center(
        child: SingleChildScrollView(
          child: Column(
            children: <Widget>[
              // Check in
              ButtonContainer(
                text: "Entrada",
                onPressed: setStatePage(6),
              ),
              RaisedButton(
                disabledColor: Colors.amber,
                child: Text("Raised Button"),
                splashColor: Colors.amber,
                color: Colors.blueAccent,
                onPressed: () {
                  print("Hola que hace?");
                  
                },
              ),
              // Meal time
              ButtonContainer(
                text: "Comida",
                onPressed: setStatePage(7),
              ),
              // Check out
              ButtonContainer(
                text: "Salida",
                onPressed: setStatePage(8),
              ),
            ],
          ),
        ),
      );

  @override
  Widget build(BuildContext context) {
    final alert = AlertDialog(
      scrollable: true,
      elevation: 2.0,
      actionsOverflowButtonSpacing: 10.0,
      actions: [
        Container(
          margin: EdgeInsets.only(left: 10, right: 5),
          child: FlatButton(
            child: Text(
              "Cancelar",
              style: TextStyle(color: kGrisColor),
            ),
            onPressed: () => Navigator.of(context).pop(),
          ),
        ),
        Container(
            margin: EdgeInsets.only(left: 5, right: 10),
            child: FlatButton(
              child: Text(
                "Cerrar",
                style: TextStyle(color: kRojoColor),
              ),
              onPressed: () {
                _sharedPreference.removeAll();
                Navigator.pushReplacementNamed(context, '/');
              },
            ))
      ],
      contentPadding: EdgeInsets.zero,
      content: FractionallySizedBox(
          widthFactor: 1,
          child: Column(
            children: <Widget>[
              Container(
                width: SizeConfig.widthMultiplier! * 100,
                padding: EdgeInsets.only(top: 10, bottom: 10),
                decoration: BoxDecoration(
                    color: kPrimaryColor,
                    borderRadius: BorderRadius.only(
                        topLeft: Radius.circular(4.0),
                        topRight: Radius.circular(4.0))),
                child: Text(
                  "Cerrar sesion",
                  textAlign: TextAlign.center,
                  style: TextStyle(
                      fontWeight: FontWeight.w500,
                      color: kTextWhiteColor,
                      fontSize: 28.0),
                ),
              ),
              Container(
                margin: EdgeInsets.only(top: 20, right: 15, left: 15),
                child: Text(
                  "¿Esta seguro que desea cerrar la sesion?",
                  textAlign: TextAlign.justify,
                  style: TextStyle(fontWeight: FontWeight.w300, fontSize: 18.0),
                ),
              ),
            ],
          )),
    );

    return WillPopScope(
      onWillPop: () async {
        if (_page != 0) setStatePage(0);
        if (_page == 0) {
          await navigatorKey.currentState!.maybePop();
          return true;
        }
        return false;
      },
      child: Scaffold(
        key: mScaffoldState,
        backgroundColor: kPrimaryColor,
        appBar: null,
        body: SafeArea(child: _callPage(_page)),
        bottomNavigationBar: CurvedNavigationBar(
          key: bottomNavigationBar,
          index: _page > 5 ? 0 : _page,
          letIndexChange: (value) => true,
          height: 50.0,
          backgroundColor: kPrimaryColor,
          items: <Widget>[
            Icon(Icons.home),
            Icon(Icons.list),
            Icon(Icons.view_day),
            Icon(Icons.sticky_note_2),
            Icon(Icons.account_circle),
            Icon(Icons.exit_to_app),
          ],
          onTap: (index) {
            if (index != 5) setStatePage(index);
            if (index == 5)
              showDialog(context: context, builder: (context) => alert);
          },
        ),
      ),
    );
  }
}
