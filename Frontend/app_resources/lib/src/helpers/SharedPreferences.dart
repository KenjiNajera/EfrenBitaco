
import 'package:shared_preferences/shared_preferences.dart';

class SharedPreference {
  
  Future<void> saveValueString(String data, String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString(tag, data);
  }

  Future<void> saveValueBoolean(bool data, String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setBool(tag, data);
  }

  Future<void> saveValueInt(int data, String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setInt(tag, data);
  }

  Future<String> returnValueString(String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    final result = prefs.getString(tag) ?? '';
    return result;
  }

  Future<bool> returnValueBoolean(String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    final result = prefs.getBool(tag) ?? false;
    return result;
  }
  
  Future<int> returnValueInt(String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    final result = prefs.getInt(tag) ?? 0;
    return result;
  }

  Future<void> removeAll() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.clear();
  }

  Future<void> removeOne(String tag) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.remove(tag);
  }
}