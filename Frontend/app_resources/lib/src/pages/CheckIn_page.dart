import 'package:app_resources/src/Widgets/ToastCustom.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:app_resources/src/global/constants.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/models/CheckIn_model.dart';
import 'package:app_resources/src/services/Checks_service.dart';
import 'package:app_resources/src/Widgets/ButtonContainer.dart';
import 'package:app_resources/src/Widgets/CalendarStripCustom.dart';
import 'package:app_resources/src/Widgets/TimePicker.dart';
//formato de fecha
import 'package:intl/intl.dart';
import 'package:datetime_picker_formfield/datetime_picker_formfield.dart';

class CheckInPage extends StatefulWidget {
  @override
  _CheckInPageState createState() => _CheckInPageState();
}

class _CheckInPageState extends State<CheckInPage> {
  final _formKey = GlobalKey<FormState>();
  ChecksService? _checksService;
  TextEditingController timeController = TextEditingController();

  DateTime? date;
  DateTime? time;

  onSelect(data) => date = data;

  setTime(data) {
    final now = DateTime.now();
    time = DateTime(now.year, now.month, now.day, data.hour - 6, data.minute);
  }

  sendCheck() async {
    if (!_formKey.currentState!.validate()) return;

    _formKey.currentState!.save();

    final SharedPreference _sharedPreference = SharedPreference();
    int userid = await _sharedPreference.returnValueInt(USERID);

    if (date == null) date = DateTime.now();

    final checkIn = CheckInModel(
      userid: userid,
      daydate: date!.toUtc().toIso8601String(),
      checktime: time!.toUtc().toIso8601String(),
    );

    bool resp = await _checksService!.sendCheckIn(checkIn);

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
    _checksService = Provider.of<ChecksService>(context, listen: false);

    return Form(
      key: _formKey,
      child: Center(
        child: ListView(
          children: <Widget>[
            Column(
              mainAxisAlignment: MainAxisAlignment.center,
              mainAxisSize: MainAxisSize.min,
              children: <Widget>[
                //Titulo
                Container(
                  margin: EdgeInsets.symmetric(vertical: 30),
                  child: Text(
                    "Registro de entrada",
                    style: TextStyle(
                      color: kTextWhiteColor,
                      fontSize: 50.0,
                    ),
                  ),
                ),
                //Calendar
                //  CalendarStripCustom(
                //    onSelect: onSelect,
                //  ),
                // Texto de seleccion

                // Container(
                //   margin: EdgeInsets.only(top: 60, bottom: 25),
                //   child: DateTimeField(
                //     format: DateFormat("yyyy-MM-dd HH:mm"),
                //     onShowPicker: (context, currentValue) async {
                //       final date = await showDatePicker(
                //           context: context,
                //           firstDate: DateTime(1900),
                //           initialDate: currentValue ?? DateTime.now(),
                //           lastDate: DateTime(2100));
                //       if (date != null) {
                //         final time = await showTimePicker(
                //           context: context,
                //           initialTime: TimeOfDay.fromDateTime(
                //               currentValue ?? DateTime.now()),
                //         );
                //         return DateTimeField.combine(date, time);
                //       } else {
                //         return currentValue;
                //       }
                //     },
                //   ),
                // ),

                Container(
                  margin: EdgeInsets.only(top: 60, bottom: 25),
                  child: Text(
                    "Seleccione la hora de entrada",
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
                // Button Check in
                ButtonContainer(
                  text: "Enviar",
                  onPressed: sendCheck(),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
