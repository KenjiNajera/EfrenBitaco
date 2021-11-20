import 'package:app_resources/src/global/constants.dart';
import 'package:flutter/material.dart';


class TimePicker extends StatefulWidget {
  TextEditingController? timeController;
  Function? setTime;

  TimePicker({
    Key? key,
    this.timeController,
    this.setTime,
  }) : super(key: key);

  @override
  _TimePickerState createState() => _TimePickerState();
}

class _TimePickerState extends State<TimePicker> {
  DateTime? _data;
  bool isButtonEnabled = false;
  String value = "00:00";

  void _setButtonState() {
    setState(() {
      if (_data == null) {
        isButtonEnabled = true;
      } else {
        isButtonEnabled = false;
      }
    });
  }

  Future<Null> _showStartTimeDialog() async {
    final TimeOfDay? picked = await showTimePicker(
      context: context, 
      initialTime: TimeOfDay.now(),
      builder: (context, child) {
        return Directionality(
          textDirection: TextDirection.ltr,
          child: child!,
        );
      }, 
    );
    
    if (picked != null) {
      setState(() {
        value ="${picked.hour > 9 ? picked.hour : "0"+picked.hour.toString()}:${picked.minute > 9 ? picked.minute : "0"+picked.minute.toString()}";
        widget.timeController = TextEditingController(text: value);
      });
      _setButtonState();
      widget.setTime!(picked);
    }
  }


  @override
  Widget build(BuildContext context) {
    widget.timeController!.text = value;
    Size size = MediaQuery.of(context).size;
    return Container(
      width: size.width * 0.8,
      margin: EdgeInsets.only(bottom: 45),
      child: GestureDetector(
        onTap: () {
          _showStartTimeDialog();
        },
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            Flexible(
              child: Container(
                child: TextFormField(
                  controller: widget.timeController,
                  validator: (value) => value!.isEmpty ? "Seleccione una hora" : null,
                  enabled: false,
                  decoration: InputDecoration(
                    contentPadding: EdgeInsets.only(left:23.0, top:32.0),
                    isDense: true,
                    filled: true,
                    fillColor: kTextWhiteColor,
                    hintText: "00:00",
                    errorStyle: TextStyle(
                      color: Colors.red,
                    ),
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(5.0)
                    ),
                  ),
                  style: TextStyle(
                    fontSize: 18.0, 
                    color: Colors.black
                  ),
                  maxLines: 1,
                ),
              ),
            ),
            Container(
              margin: EdgeInsets.only(left:20),
              padding: EdgeInsets.all(10.0),
              decoration: BoxDecoration(
                color: kTextWhiteColor,
                borderRadius: BorderRadius.circular(50.0)
              ),
              child: Icon(
                Icons.access_time,
                color: kPrimaryColor,
                size: 35,
              ),
            ),  
          ],
        ),
      ),
    );
  }
}

