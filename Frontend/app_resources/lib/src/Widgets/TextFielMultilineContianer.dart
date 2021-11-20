import 'package:app_resources/src/global/constants.dart';
import 'package:flutter/material.dart';

class TextFieldMultilineContainer extends StatefulWidget {
  final TextEditingController? textController;
  final String? text;
  final bool? isvalidator;
  TextFieldMultilineContainer({
    Key? key,
    this.text,
    this.textController,
    this.isvalidator,
  }) : super(key: key);

  @override
  _TextFieldMultilineContainerState createState() => _TextFieldMultilineContainerState();
}

class _TextFieldMultilineContainerState extends State<TextFieldMultilineContainer> {
  String? _description;

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Container(
      padding: EdgeInsets.all(8.0),
      decoration: BoxDecoration(
        color: Colors.white12,
        borderRadius: BorderRadius.circular(8.0),
      ),
      width: size.width * 0.8,
      child: TextFormField(
        controller: widget.textController,
        validator: (value) => widget.isvalidator! ? value!.isEmpty ? "Llene este campo" : null : null,
        onSaved: (newValue) => _description = newValue,
        maxLength: 155,
        decoration: InputDecoration(
          helperStyle: TextStyle(
            color: kTextWhiteColor,
            fontWeight: FontWeight.w400,
          ),
          hintText: widget.text,
          hintStyle: TextStyle(
            color: Colors.white54,
          ),
        ),
        style: TextStyle(
          color: kTextWhiteColor
        ),
        minLines: 4,
        maxLines: 8,
      ),
    );
  }
}