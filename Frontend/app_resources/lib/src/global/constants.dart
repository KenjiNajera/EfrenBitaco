import 'package:flutter/material.dart';

// Consumos
  const headers = { "Content-Type": "application/json", "Accept" : "application/json, text/plain" };
  const url = "https://resource-grupogit.herokuapp.com/";
  const authenticatemovil ="auth/authenticatemovil/";
  const typeRequests = "typerequest/gettyperequests/";
  const sendRequests = "request/create/";
  const getRequestsUser = "request/getrequestuser/";
  const getBinnacleUser = "binnacle/getbinnacleuser/";
  const sendCheckInUser = "binnacle/checkin/";
  const sendCheckMealUser = "binnacle/mealtime/";
  const sendCheckOutUser = "binnacle/checkout/"; 
  const getDataUser = "user/getmydatamovil/";

// TAG'S Shared Preferences
  const TOKENMOVBILE = "TokenMobile";
  const USERID = "UserId";
  const RESOURCEDATAID = "ResourcedataId";
  const ROLEID = "RoleId";
  const USERSTATUS = "Status";
  const LOGGEDIN = "LoggedIn";
  
//Colors
  const kPrimaryColor = Color.fromRGBO(3, 103, 166, 1.0);
  const kSelectorColor = Color.fromRGBO(3, 139, 225, 0.8);
  const kTextfieldColor = Color.fromRGBO(255, 255, 255, 1.0);
  const kButtonColor = Color.fromRGBO(255, 255, 255, 1.0);
  const kDividerColor = Color.fromRGBO(255, 255, 255, 1.0);
  const kTextWhiteColor = Color.fromRGBO(255, 255, 255, 1.0);
  
  const kRojoColor = Color.fromRGBO(220, 53, 69, 1.0);
  const kGrisColor = Color.fromRGBO(192, 192, 192, 1.0);

  // Switch Colors
    const kSwitchColor = Color.fromRGBO(255, 255, 255, 1.0);
    const kSwitchTrackColor = Color.fromRGBO(174, 174, 174, 1.0);

    const kSwitchColorGreen = Color.fromRGBO(40, 167, 69, 1.0);
    const kSwitchColorRed = Color.fromRGBO(220, 53, 69, 1.0);
  
  // Table Colors
  const kTableColor = Color.fromRGBO(3, 139, 255, 0.2);
