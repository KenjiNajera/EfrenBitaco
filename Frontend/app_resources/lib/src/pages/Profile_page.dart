import 'package:app_resources/src/helpers/SharedPreferences.dart';
import 'package:app_resources/src/services/User_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

import 'package:app_resources/src/global/constants.dart';
import 'package:app_resources/src/pages/fragments/DataUser_Fragment.dart';
import 'package:app_resources/src/pages/fragments/Experiences_Fragment.dart';
import 'package:app_resources/src/pages/fragments/SkillsFragment.dart';
import 'package:provider/provider.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';

class ProfilePage extends StatefulWidget {
  const ProfilePage({Key? key}) : super(key: key);

  @override
  _ProfilePageState createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {

  UserService? _userService;
  RefreshController _refreshController = RefreshController(initialRefresh: false);

  @override
  void initState() { 
    super.initState();
    _getDataUser();
  }

  void _onRefresh() async{
    // monitor network fetch
    await Future.delayed(Duration(milliseconds: 1000));
    // if failed,use refreshFailed()
    _getDataUser();
    _refreshController.refreshCompleted();
  }

  void _onLoading() async{
    // monitor network fetch
    await Future.delayed(Duration(milliseconds: 1000));
    // if failed,use loadFailed(),if no data return,use LoadNodata()
    _refreshController.loadComplete();
  }

  void _getDataUser() async{
    final SharedPreference _sharedPreference = SharedPreference();
    int id = await _sharedPreference.returnValueInt(USERID);
    await _userService!.getUserData(id);
  }
  @override
  Widget build(BuildContext context) {
    _userService = Provider.of<UserService>( context );
    Size size = MediaQuery.of(context).size;
    return SmartRefresher(
      controller: _refreshController,
      enablePullDown: true,
      header: WaterDropHeader(),
      onLoading: (){},
      onRefresh: _onRefresh,
      child: _data(size),
    );
  }

  Widget _data(Size size) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: <Widget>[
        Container(
          width: 130,
          height: 130,
          margin: EdgeInsets.only(top: 40),
          padding: EdgeInsets.all(20.0),
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(100.0),
            color: Colors.brown
          ),
          child: SvgPicture.asset(
            '${_userService!.user.image != "" ? _userService!.user.image : "assets/images/usuario.svg"}',
            color: kTextWhiteColor,
          ) 
        ),
        // Name user
        Container(
          margin: EdgeInsets.only(top:15),
          child: Text(
            "${_userService!.user.firstname} ${_userService!.user.lastname}",
            style: TextStyle(
              color: kTextWhiteColor,
              fontSize: 20.0
            ),
          ),
        ),
        // Puesto
        Container(
          margin: EdgeInsets.symmetric(vertical:10,),
          child: Text(
            "${_userService!.user.rolename}",
            style: TextStyle(
              color: kTextWhiteColor,
              fontSize: 17.0
            ),
          ),
        ),
        // Tabs
        Container(
          height: size.height - 305,
          child: DefaultTabController(
            length: 3,
            child: Column(
              children: <Widget>[
                Container(
                  decoration: BoxDecoration(
                    color: kPrimaryColor,
                    boxShadow: <BoxShadow>[
                      BoxShadow(
                        color: Colors.black12,
                        blurRadius: 3.0,
                        offset: Offset(0.0, 3.0)
                      )
                    ],
                    border: Border.all(color: kTableColor, width: 2.2)
                  ),
                  child: TabBar(
                    tabs: [
                      Tab(text: "Datos"),
                      Tab(text: "Habilidades"),
                      Tab(text: "Experiencia"),
                    ]
                  ),
                ),
                Expanded(
                  child: Container(
                    child: TabBarView(
                      children: [
                        DataUserFragment(
                          user: _userService!.user,
                        ),
                        SkillsFragment(
                          listHardskill: _userService!.user.hardskills,
                          listLanguage: _userService!.user.languages,
                          listSoftskill: _userService!.user.softskills,
                        ),
                        ExperienceFragment(
                          list: _userService!.user.experiences,
                        )
                      ]
                    ),
                  ), 
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}