import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_svg/flutter_svg.dart';

import 'package:app_resources/src/Widgets/SwitchFilter.dart';
import 'package:app_resources/src/global/constants.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/models/Requests_model.dart';
import 'package:app_resources/src/services/Request_service.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';


class ListRequestPage extends StatefulWidget {
  const ListRequestPage({Key? key}) : super(key: key);

  @override
  _ListRequestPageState createState() => _ListRequestPageState();
}

class _ListRequestPageState extends State<ListRequestPage> {
  RequestService? _requestService;
  RefreshController _refreshController = RefreshController(initialRefresh: false);

  bool isAprovadas = false;
  bool isPendientes = false;

  setIsAprobadas(value) => setState((){
    isAprovadas = value;
  });
  setIsPendientes(value) => setState((){
    isPendientes = value;
  });

  @override
  void initState() { 
    super.initState();
    _getListRequestsUser();
  }

  void _onRefresh() async{
    // monitor network fetch
    await Future.delayed(Duration(milliseconds: 1000));
    // if failed,use refreshFailed()
    _getListRequestsUser();
    _refreshController.refreshCompleted();
  }

  void _onLoading() async{
    // monitor network fetch
    await Future.delayed(Duration(milliseconds: 1000));
    // if failed,use loadFailed(),if no data return,use LoadNodata()
    _refreshController.loadComplete();
  }

  @override
  Widget build(BuildContext context) {
    _requestService = Provider.of<RequestService>( context );
    return SmartRefresher(
      controller: _refreshController,
      enablePullDown: true,
      header: WaterDropHeader(),
      onRefresh: _onRefresh,
      onLoading: (){},
      child: Center(
        child: Column(
          children: <Widget>[
            // Filtros
            FractionallySizedBox(
              widthFactor: 0.8,
              child: Container(
                margin: EdgeInsets.symmetric(vertical: 15),
                decoration: BoxDecoration(
                  color: kTextfieldColor,
                  boxShadow: <BoxShadow>[
                    BoxShadow(
                      color: Colors.black12,
                      blurRadius: 5.0,
                      offset: Offset(0.0, 6)
                    )
                  ],
                  borderRadius: BorderRadius.circular(20.0)
                ),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    SwitchFilter(
                      text: "aprobadas",
                      value: isAprovadas,
                      color: kSwitchColorGreen,
                      setStatus: setIsAprobadas,
                    ),
                    SwitchFilter(
                      text: "denegadas",
                      value: isPendientes,
                      setStatus: setIsPendientes,
                      color: kSwitchColorRed,
                    ),
                  ],
                ),
              ),
            ),
            // Lista de cards
            Expanded(
              child: ListView(
                children: <Widget>[
                  ..._requestService!.listRequestsModel
                  .where((element) => 
                    isAprovadas && isPendientes ? element.status == false || element.status == true :
                    isAprovadas ? element.status == true : 
                    isPendientes ? element.status == false : 
                    element.status == false || element.status == true || element.status == null)
                  .map((e)  => cardReqest(e, context)), 
                ],
              ),
            ),
          ],
        ),
      ),
    ); 
    
  }

  Widget cardReqest(RequestsModel item, BuildContext context) {
    final alert = AlertDialog(
      scrollable: true,
      backgroundColor: Colors.transparent,
      contentPadding: EdgeInsets.all(0),
      content: Stack(
          alignment: AlignmentDirectional.centerStart,
          children: <Widget>[
            Card(
              margin: EdgeInsets.only(left: 40),
              elevation: 2.0,
              child: FractionallySizedBox(
                widthFactor: 0.7,
                child: Container(
                  padding: EdgeInsets.only(left: 50, top: 20, bottom: 20, right: 20),
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      // Titulo
                      Text(
                        item.typerequestname!,
                        style: TextStyle(
                          fontSize: 25.0,
                          fontWeight: FontWeight.w600,
                        ),
                      ),
                      //Description
                      Container(
                        margin: EdgeInsets.symmetric(vertical: 12),
                        child: Text("Descripción:"),
                      ),
                      Text(item.description!),
                      // Detalles
                      item.details != "" ?
                        Container(
                        margin: EdgeInsets.symmetric(vertical: 12),
                        child: Text("Detalles:"),
                        )
                      : Text(""),
                        Text(item.details!)
                    ]
                  ),
                ),
              ),
            ),
            SvgPicture.asset(
              item.status == null ? 'assets/images/circuloGris.svg' : item.status! ? 'assets/images/circuloVerde.svg' : 'assets/images/circuloRojo.svg',
              height: 80.0,
              width: 80.0,
            ),
          ],
        ),
    );

    return Center(
      child: Container(
        margin: EdgeInsets.only(bottom: 20),
        child: Stack(
          alignment: AlignmentDirectional.centerStart,
          children: <Widget>[
            Card(
              margin: EdgeInsets.only(left: 40),
              elevation: 3.0,
              child: InkWell(
                onTap: () => item.status != null ? showDialog(
                  context: context,
                  builder: (context) => alert,
                ) : {},
                child: FractionallySizedBox(
                  widthFactor: 0.7,
                  child: Container(
                    padding: EdgeInsets.only(left: 50, top: 20, bottom: 20, right: 20),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.start,
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: <Widget>[
                        Text(
                          item.typerequestname!,
                          style: TextStyle(
                            fontSize: 25.0,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                        Container(
                          margin: EdgeInsets.symmetric(vertical: 12),
                          child: Text("Descripción:"),
                        ),
                        Text(item.description!)
                      ]
                    ),
                  ),
                ),
              ),
            ),
            SvgPicture.asset(
              item.status == null ? 'assets/images/circuloGris.svg' : item.status! ? 'assets/images/circuloVerde.svg' : 'assets/images/circuloRojo.svg',
              height: 80.0,
              width: 80.0,
            ),
          ],
        )
      )
    );   
  }

  _getListRequestsUser() async {
    final SharedPreference _sharedPreference = SharedPreference();
    int id = await _sharedPreference.returnValueInt(USERID);
    _requestService!.getRequests(id);
  }

}