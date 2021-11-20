import 'package:app_resources/src/Widgets/ButtonContainer.dart';
import 'package:app_resources/src/Widgets/CalendarStripCustom.dart';
import 'package:app_resources/src/Widgets/TimePicker.dart';
import 'package:app_resources/src/Widgets/ToastCustom.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/models/Mealtime_model.dart';
import 'package:app_resources/src/services/Checks_service.dart';
import 'package:flutter/material.dart';

import '../global/constants.dart';

class MealtimePage extends StatefulWidget {
  const MealtimePage({
    Key? key
  }) : super(key: key);

  @override
  _MealtimePageState createState() => _MealtimePageState();
}

class _MealtimePageState extends State<MealtimePage> {

  final _formKey = GlobalKey<FormState>();
  ChecksService? _checksService;
  TextEditingController timeController = TextEditingController();

  DateTime? date;
  DateTime? time;

  onSelect(data) => date = data;

  setTime(data){
    final now = DateTime.now();
    time = DateTime(now.year, now.month, now.day, data.hour - 6, data.minute);
  }
  
  sendCheck() async {
    if(!_formKey.currentState!.validate()) return;
    
    _formKey.currentState! .save();

    final SharedPreference _sharedPreference = SharedPreference();
    int userid = await _sharedPreference.returnValueInt(USERID);
    
    if(date == null) date = DateTime.now();

    final mealtime = MealTimeModel(
      userid: userid,
      daydate: date!.toUtc().toIso8601String(),
      mealtime: time!.toUtc().toIso8601String(),
    );

    bool resp =  await _checksService!.sendCheckMeal(mealtime);

    if (resp) {
      // TODO: volver a home y mandar snackbar
      Navigator.pushReplacementNamed(context, '/home');
    } else {
      // TODO: Mandar snackbar
      // ToastCustom().getToastError(context);
    }
  }
  
  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Center(
        child: ListView(
          children: <Widget>[
            Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                //Titulo
                Container(
                  margin: EdgeInsets.symmetric(vertical:30),
                  child: Text(
                    "Registro de comida",
                    style: TextStyle(
                      color: kTextWhiteColor,
                      fontSize: 50.0,
                    ),
                  ),
                ),
                //Calendar
                // CalendarStripCustom(
                //   onSelect: onSelect,
                // ),
                // Texto de seleccion
                Container(
                  margin: EdgeInsets.only(top: 60, bottom: 25),
                  child: Text(
                    "Seleccione la hora de comida",
                    style: TextStyle(
                      color: kTextWhiteColor,
                      fontSize: 23.0,
                    ),
                  ),
                ),
                // Selector Time
                TimePicker(
                  timeController: timeController,
                  setTime: setTime,
                ),
                // Button Meal check
                ButtonContainer(
                  text: "Enviar",
                  onPressed: sendCheck(),
                ),
              ],
            ),
          ],
        ),
      )
    );
  }
}