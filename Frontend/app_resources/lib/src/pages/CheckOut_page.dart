
import 'package:app_resources/src/Widgets/TextFielMultilineContianer.dart';
import 'package:app_resources/src/Widgets/ToastCustom.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/models/CheckOut_model.dart';
import 'package:app_resources/src/services/Checks_service.dart';
import 'package:flutter/material.dart';

import '../global/constants.dart';
import 'package:app_resources/src/Widgets/ButtonContainer.dart';
import 'package:app_resources/src/Widgets/CalendarStripCustom.dart';
import 'package:app_resources/src/Widgets/TimePicker.dart';

class CheckOutPage extends StatefulWidget {
  const CheckOutPage({
    Key? key
  }) : super(key: key);

  @override
  _CheckOutPageState createState() => _CheckOutPageState();
}

class _CheckOutPageState extends State<CheckOutPage> {

  final _formKey = GlobalKey<FormState>();
  ChecksService? _checksService;
  TextEditingController timeController = TextEditingController();
  TextEditingController textController = TextEditingController();

  DateTime? date;
  DateTime? time;
  bool overtime =  false;

  onSelect(data) => date = data;

  setTime(data){
    final now = DateTime.now();
    time = DateTime(now.year, now.month, now.day, data.hour - 6, data.minute);
  }
  
  setOvertime(value) => overtime = value;

  sendCheck() async {
    if(!_formKey.currentState!.validate()) return;
    
    _formKey.currentState!.save();

    final SharedPreference _sharedPreference = SharedPreference();
    int userid = await _sharedPreference.returnValueInt(USERID);
    
    if(date == null) date = DateTime.now();

    final checkout = CheckOutModel(
      userid: userid,
      daydate: date!.toUtc().toIso8601String(),
      departuretime: time!.toUtc().toIso8601String(),
      overtime: overtime,
      summary: textController.text,
    );

    bool resp =  await _checksService!.sendCheckOut(checkout);

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
    Size size = MediaQuery.of(context).size;
    return Form(
      key: _formKey,
      child: Center(
        child: ListView(
          children: <Widget>[
            Column(
              children: <Widget>[
                //Titulo
                Container(
                  margin: EdgeInsets.symmetric(vertical:30),
                  child: Text(
                    "Registro de salida",
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
                  margin: EdgeInsets.only(top: 25, bottom: 20),
                  child: Text(
                    "Seleccione la hora de salida",
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
                // Summary
                TextFieldMultilineContainer(
                  textController: textController,
                    text: "descripcción...",
                    isvalidator: false,
                ),
                // Horas extra?
                Container(
                  width: size.width * 0.8,
                  height: size.height *0.1,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: <Widget>[
                      Text(
                        "Horas Extra?",
                        style: TextStyle(
                          color: kTextWhiteColor,
                          fontSize: 18.0
                        ),
                      ),
                      Flexible(
                        child: Switch(
                          value: overtime,
                          onChanged: (value) {
                            setState((){

                            });
                            overtime = value;
                          },
                          activeTrackColor: kSwitchTrackColor,
                          activeColor: kSwitchColor
                        ),
                      )
                    ],
                  ),
                ),
                // Button Meal check
                ButtonContainer(
                  text: "Enviar",
                  onPressed: sendCheck(),
                ),
              ],
              mainAxisAlignment: MainAxisAlignment.center,
            ),
          ],
        ),
      ),
    );
  }
}