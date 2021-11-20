import 'package:flutter/material.dart';

class SwitchFilter extends StatefulWidget {
  final String? text;
  final Color? color;
  bool? value;
  Function? setStatus;
  SwitchFilter({
    Key? key,
    @required this.text,
    this.color,
    this.value,
    this.setStatus,
  }) : super(key: key);

  @override
  _SwitchFilterState createState() => _SwitchFilterState();
}

class _SwitchFilterState extends State<SwitchFilter> {

  @override
  Widget build(BuildContext context) {
    return Container(
        child: Row(
          children: <Widget>[
            Switch(
              value: widget.value!,
              onChanged: (value) => widget.setStatus!(value),
              activeColor: widget.color,
            ),
            Text(
              widget.text!,
            )
          ],
        ),
    );
  }
}