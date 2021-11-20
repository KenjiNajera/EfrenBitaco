import 'package:app_resources/src/global/size_config.dart';
import 'package:app_resources/src/models/Binnacle_model.dart';
import 'package:app_resources/src/services/Binnacle_service.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';

import 'package:app_resources/src/global/constants.dart';
import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/services/Request_service.dart';

class BinnaclePage extends StatefulWidget {
  const BinnaclePage({Key? key}) : super(key: key);

  @override
  _BinnaclePageState createState() => _BinnaclePageState();
}

class _BinnaclePageState extends State<BinnaclePage> {
  BinnacleService? _binnacleService;
  RefreshController _refreshController = RefreshController(initialRefresh: false);

  @override
  void initState() { 
    super.initState();
    _getBinnacleView();
  }

  void _onRefresh() async{
    // monitor network fetch
    await Future.delayed(Duration(milliseconds: 1000));
    // if failed,use refreshFailed()
    _getBinnacleView();
    _refreshController.refreshCompleted();
  }

  void _onLoading() async{
    // monitor network fetch
    await Future.delayed(Duration(milliseconds: 1000));
    // if failed,use loadFailed(),if no data return,use LoadNodata()
    _refreshController.loadComplete();
  }

  _getBinnacleView() async {
    final SharedPreference _sharedPreference = SharedPreference();
    int id = await _sharedPreference.returnValueInt(USERID);
    bool resp = await _binnacleService!.getBinnacleView( id );
  }

  @override
  Widget build(BuildContext context) {
    _binnacleService = Provider.of<BinnacleService>( context );
    return SmartRefresher(
      controller: _refreshController,
      enablePullDown: true,
      header: WaterDropHeader(),
      onLoading: (){},
      onRefresh: _onRefresh,
      child: SizeConfig.isMobilePortrait ? ContainerPortrait(binnacle: _binnacleService!.binnacle,) : ContainerLanscape(binnacle:_binnacleService!.binnacle)
    );
  }
}

class ContainerPortrait extends StatefulWidget {
  final BinnacleViewModel? binnacle;
  const ContainerPortrait({
    Key? key,
    this.binnacle,
  }) : super(key: key);

  @override
  _ContainerPortraitState createState() => _ContainerPortraitState();
}

class _ContainerPortraitState extends State<ContainerPortrait> {
  @override
  Widget build(BuildContext context) {
    return FractionallySizedBox(
      widthFactor: 0.92,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          //Titulo
          Container(
            margin: SizeConfig.isMobilePortrait ? EdgeInsets.only(top:40, bottom: 25) : EdgeInsets.only(top:15, bottom:10),
            child: Text(
              "Bitácora",
              style: TextStyle(
                color: kTextWhiteColor,
                fontSize: SizeConfig.isMobilePortrait ? 50.0 : 30.0,
              ),
            ),
          ),
          // Datos
          Container(
            margin: EdgeInsets.only(bottom: 20),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: <Widget>[
                Container(
                  margin: EdgeInsets.only(right: 20),
                  child: Text(
                    "Proyecto: ${ widget.binnacle!.projects!.isEmpty ? "" : widget.binnacle!.projects!.map((e) => e) }",
                    style: TextStyle(
                      fontSize: 18,
                      color: kTextWhiteColor,
                    ),
                  ),
                ),
              ],
            ),
          ),
          Container(
            margin: EdgeInsets.only(bottom: 10),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: <Widget>[
                Text(
                  "Mes: ${ widget.binnacle!.monthname == null ? "" : widget.binnacle!.monthname!.toUpperCase() }",
                  style: TextStyle(
                    fontSize: 16,
                    color: kTextWhiteColor,
                  ),
                ),
                Text(
                  "Año: ${ widget.binnacle!.year == null ? "" : widget.binnacle!.year }",
                  style: TextStyle(
                    fontSize: 16,
                    color: kTextWhiteColor,
                  ),
                ),
                Text(
                  "Horas: ${ widget.binnacle!.totalhours == null ? "" : widget.binnacle!.totalhours }",
                  style: TextStyle(
                    fontSize: 16,
                    color: kTextWhiteColor,
                  ),
                )
              ],
            ),
          ),
          // Divisor
          Divider(
            color: kTextWhiteColor,
            height: SizeConfig.isMobilePortrait ? 15.0 : 5.0,
            thickness: 2.0,
            indent: 10.0,
            endIndent: 10.0,
          ),
          FractionallySizedBox(
            widthFactor: 1,
            child: Container(
              height: SizeConfig.isMobilePortrait ? SizeConfig.heightMultiplier! * 6  : SizeConfig.heightMultiplier! * 6,
              decoration: BoxDecoration(
                border: Border(bottom: BorderSide(color: Colors.black, width: 1.0))
              ),
              child: ListTile(
                contentPadding: EdgeInsets.only(top: 0, bottom: 0, left: 0, right: 0),
                onTap: null,
                title: Row(
                  children: <Widget>[
                    Expanded(
                      child: Text("Fecha", 
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hr. Entrada", 
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hr. Comida",
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hr. Salida",
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hrs. Totales",
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                  ]
                ),
              )
            ),
          ),
          widget.binnacle!.days!.isEmpty ? 
          Container(
            margin: EdgeInsets.symmetric(vertical: 180),
            child: CircularProgressIndicator(
              backgroundColor: Colors.white38,
              valueColor: AlwaysStoppedAnimation(Colors.white),
            )
          ) 
          : Expanded(
            child:ListView.builder(
              padding: EdgeInsets.zero,
              itemCount: widget.binnacle!.days!.length,
              itemBuilder: (BuildContext context, int index) {
                return  Rows(item:widget.binnacle!.days![index]);
              },
            ),
          )
        ],
      ),
    );
  }
}

class ContainerLanscape extends StatefulWidget {
  final BinnacleViewModel? binnacle;
  const ContainerLanscape({
    Key? key,
    this.binnacle,
  }) : super(key: key);

  @override
  _ContainerLanscapeState createState() => _ContainerLanscapeState();
}

class _ContainerLanscapeState extends State<ContainerLanscape> {
  @override
  Widget build(BuildContext context) {
    return NestedScrollView(
      body: Builder(
        builder: (BuildContext context) {
          return CustomScrollView(
            slivers: [
              SliverOverlapInjector(handle: NestedScrollView.sliverOverlapAbsorberHandleFor(context)),
              SliverToBoxAdapter(
                child: widget.binnacle!.days!.isEmpty ? 
                Container(
                  margin: EdgeInsets.symmetric(vertical: 180),
                  child: CircularProgressIndicator(
                    backgroundColor: Colors.white38,
                    valueColor: AlwaysStoppedAnimation(Colors.white),
                  )
                ) : 
                Column(children: [
                  ...widget.binnacle!.days!.map((e) => Rows(item: e,))
                  ],
                )
              ),
            ],
          );
        }
      ),
      headerSliverBuilder: (BuildContext context, bool innerBoxIsScrolled) => <Widget>[
        SliverAppBar(
          floating: true,
          pinned: false,
          expandedHeight: 143,
          flexibleSpace: FractionallySizedBox(
            widthFactor: 0.92,
            child: Column(
              children: <Widget>[
                //Titulo
                Container(
                  margin: SizeConfig.isMobilePortrait ? EdgeInsets.only(top:40, bottom: 25) : EdgeInsets.only(top:15, bottom:10),
                  child: Text(
                    "Bitácora",
                    style: TextStyle(
                      color: kTextWhiteColor,
                      fontSize: SizeConfig.isMobilePortrait ? 50.0 : 30.0,
                    ),
                  ),
                ),
                // Datos
                Container(
                  margin: EdgeInsets.only(bottom: 20),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: <Widget>[
                      Container(
                        margin: EdgeInsets.only(right: 20),
                        child: Text(
                          "Proyecto: ${ widget.binnacle!.projects!.isEmpty ? "" : widget.binnacle!.projects!.map((e) => e) }",
                          style: TextStyle(
                            fontSize: 18,
                            color: kTextWhiteColor,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                Container(
                  margin: EdgeInsets.only(bottom: 10),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: <Widget>[
                      Text(
                        "Mes: ${ widget.binnacle!.monthname == null ? "" : widget.binnacle!.monthname!.toUpperCase() }",
                        style: TextStyle(
                          fontSize: 16,
                          color: kTextWhiteColor,
                        ),
                      ),
                      Text(
                        "Año: ${ widget.binnacle!.year == null ? "" : widget.binnacle!.year }",
                        style: TextStyle(
                          fontSize: 16,
                          color: kTextWhiteColor,
                        ),
                      ),
                      Text(
                        "Horas: ${ widget.binnacle!.totalhours == null ? "" : widget.binnacle!.totalhours }",
                        style: TextStyle(
                          fontSize: 16,
                          color: kTextWhiteColor,
                        ),
                      )
                    ],
                  ),
                ),
                // Divisor
                Divider(
                  color: kTextWhiteColor,
                  height: SizeConfig.isMobilePortrait ? 15.0 : 5.0,
                  thickness: 2.0,
                  indent: 10.0,
                  endIndent: 10.0,
                ),
              ],
            ),
          ),
        ),
    
        SliverOverlapAbsorber(
          handle: NestedScrollView.sliverOverlapAbsorberHandleFor(context),
          sliver: SliverAppBar(
            pinned: true,
            flexibleSpace: FractionallySizedBox(
            widthFactor: 1,
            child: Container(
              height: SizeConfig.isMobilePortrait ? SizeConfig.heightMultiplier! * 6  : SizeConfig.heightMultiplier! * 6,
              decoration: BoxDecoration(
                border: Border(bottom: BorderSide(color: Colors.black, width: 1.0))
              ),
              child: ListTile(
                contentPadding: EdgeInsets.only(top: 0, bottom: 0, left: 0, right: 0),
                onTap: null,
                title: Row(
                  children: <Widget>[
                    Expanded(
                      child: Text("Fecha", 
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hr. Entrada", 
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hr. Comida",
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hr. Salida",
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                    Expanded(
                      child: Text("Hrs. Totales",
                        style: TextStyle(fontStyle: FontStyle.italic, color: kTextWhiteColor), 
                        textAlign: TextAlign.center,
                      )
                    ),
                  ]
                ),
              )
            ),
          ),
          )
        ),
      ] 
    );
  }
}

class Header extends SliverPersistentHeaderDelegate {
  
  Header({
    @required this.child,
  });

  final PreferredSize? child;

  @override
  Widget build(BuildContext context, double shrinkOffset, bool overlapsContent) {
    return this.child!;
  }

  @override
  // TODO: implement maxExtent
  double get maxExtent => this.child!.preferredSize.height;

  @override
  // TODO: implement minExtent
  double get minExtent => this.child!.preferredSize.height;

  @override
  bool shouldRebuild(covariant SliverPersistentHeaderDelegate oldDelegate){
    return true;
  }
}

class Rows extends StatelessWidget {
  final DayModel? item;
  const Rows({
    Key? key,
    this.item,
  }) : super(key: key);

  // final AlertDialog viewMore = AlertDialog(
  //   scrollable: true,
  //   backgroundColor: Colors.white,
  //   contentPadding: EdgeInsets.all(0),
  //   content: Container(
  //     decoration: BoxDecoration(
  //       borderRadius: BorderRadius.circular(5.0)
  //     ),
  //     child: Column(
  //       children: <Widget>[
  //         Text("Holas")
  //       ],
  //     ),
  //   ),
  // );

    

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onLongPress: null,
      onTap: () =>showDialog(
        context: context,
        builder: (context) => Text(""),
      ),
      child:  Container(
        decoration: BoxDecoration(
          border: Border(bottom: BorderSide(color: Colors.black, width: 1.0))
        ),
        child: Row(
          children: <Widget>[
            Expanded(
              child: Container(
                padding: EdgeInsets.symmetric(vertical:19.0),
                child: Text(item!.daydate!, 
                  textAlign: TextAlign.center,
                  style: TextStyle(fontStyle: FontStyle.italic, color: Colors.grey[400], fontSize: 16.0),
                )
              )
            ),
            Expanded(
              child: Container(
                padding: EdgeInsets.symmetric(vertical:19.0),
                color: kTableColor,
                child: Text(
                  item!.checktime!, 
                  textAlign: TextAlign.center,
                  style: TextStyle(fontStyle: FontStyle.italic, color: Colors.grey[400], fontSize: 16.0),
                ),
              )
            ),
            Expanded(
              child: Container(
                padding: EdgeInsets.symmetric(vertical:19.0),
                child: Text(item!.mealtime!, 
                  textAlign: TextAlign.center,
                  style: TextStyle(fontStyle: FontStyle.italic, color: Colors.grey[400], fontSize: 16.0),
                )
              ),
            ),
            Expanded(
              child: Container(
                padding: EdgeInsets.symmetric(vertical:19.0),
                color: kTableColor,
                child: Text(item!.departuretime!, 
                  textAlign: TextAlign.center,
                  style: TextStyle(fontStyle: FontStyle.italic, color: Colors.grey[400], fontSize: 16.0),
                )
              ),
            ),
            Expanded(
              child: Container(
                padding: EdgeInsets.symmetric(vertical:19.0),
                child: Text(item!.totalhoursday!, 
                  textAlign: TextAlign.center,
                  style: TextStyle(fontStyle: FontStyle.italic, color: Colors.grey[400], fontSize: 16.0),
                )
              )
            ),
          ]
        ),
      )
    );
  }
}