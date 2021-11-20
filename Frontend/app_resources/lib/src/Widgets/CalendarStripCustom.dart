// import 'package:app_resources/src/global/constants.dart';
// import 'package:flutter/material.dart';
// import 'package:calendar_strip/calendar_strip.dart';

// class CalendarStripCustom extends StatefulWidget {
  
//   final Function? onSelect;
//   CalendarStripCustom({
//     Key? key,
//     this.onSelect,
//   }) : super(key: key);

//   @override
//   _CalendarStripCustomState createState() => _CalendarStripCustomState();
// }

// class _CalendarStripCustomState extends State<CalendarStripCustom> {
//   DateTime startDate = DateTime.now().subtract(Duration(days: 15));
//   DateTime endDate = DateTime.now().add(Duration(days: 1));
//   DateTime selectedDate = DateTime.now();
//   List<DateTime> markedDates = [
//     DateTime.now().subtract(Duration(days: 1)),
//     DateTime.now().subtract(Duration(days: 2)),
//     DateTime.now().add(Duration(days: 4))
//   ];

//   _monthNameWidget(monthName) {
//     return Container(
//       child: Text(
//         monthName,
//         style: TextStyle(
//           fontSize: 17,
//           fontWeight: FontWeight.w600,
//           color: Colors.black87,
//           fontStyle: FontStyle.italic,
//         ),
//       ),
//       padding: EdgeInsets.only(top: 8, bottom: 4),
//     );
//   }

//   getMarkedIndicatorWidget() {
//     return Row(mainAxisAlignment: MainAxisAlignment.center, children: [
//       Container(
//         margin: EdgeInsets.only(left: 1, right: 1),
//         width: 7,
//         height: 7,
//         decoration: BoxDecoration(shape: BoxShape.circle, color: Colors.red),
//       ),
//       Container(
//         width: 7,
//         height: 7,
//         decoration: BoxDecoration(shape: BoxShape.circle, color: Colors.blue),
//       )
//     ]);
//   }

//   dateTileBuilder(date, selectedDate, rowIndex, dayName, isDateMarked, isDateOutOfRange) {
//     bool isSelectedDate = date.compareTo(selectedDate) == 0;
//     Color fontColor = isDateOutOfRange ? Colors.black26 : Colors.black87;
//     TextStyle normalStyle = TextStyle(fontSize: 17, fontWeight: FontWeight.w800, color: fontColor);
//     TextStyle selectedStyle = TextStyle(fontSize: 17, fontWeight: FontWeight.w800, color: Colors.black87);
//     TextStyle dayNameStyle = TextStyle(fontSize: 14.5, color: fontColor);
//     List<Widget> _children = [
//       Text(dayName, style: dayNameStyle),
//       Text(date.day.toString(), style: !isSelectedDate ? normalStyle : selectedStyle),
//     ];

//     if (isDateMarked == true) {
//       _children.add(getMarkedIndicatorWidget());
//     }

//     return AnimatedContainer(
//       duration: Duration(milliseconds: 100),
//       alignment: Alignment.center,
//       padding: EdgeInsets.only(top: 8, left: 5, right: 5, bottom: 5),
//       decoration: BoxDecoration(
//         color: !isSelectedDate ? Colors.transparent : kSelectorColor,
//         borderRadius: BorderRadius.all(Radius.circular(60)),
//       ),
//       child: Column(
//         children: _children,
//       ),
//     );
//   }
//   onWeekSelect(data) => null;

//   onSelect(data) {
//     widget.onSelect(data);
//     selectedDate = data;
//   } 

//   @override
//   Widget build(BuildContext context) {
//     Size size = MediaQuery.of(context).size;
//     return Container(
//       width: size.width * 0.9,
//       height: 120,
//       child: CalendarStrip(
//         startDate: startDate,
//         endDate: endDate,
//         selectedDate: selectedDate,
//         onDateSelected: onSelect,
//         dateTileBuilder: dateTileBuilder,
//         onWeekSelected: onWeekSelect,
//         iconColor: Colors.black87,
//         monthNameWidget: _monthNameWidget,
//         markedDates: markedDates,
//         containerDecoration: BoxDecoration(
//           color: Colors.white,
//           boxShadow: <BoxShadow>[
//             BoxShadow(
//               color: Colors.black26,
//               blurRadius: 5.0,
//               offset: Offset(0.0, 6)
//             )
//           ],
//           borderRadius: BorderRadius.circular(15),
//         ),
//       ),
//     );
//   }
// }

