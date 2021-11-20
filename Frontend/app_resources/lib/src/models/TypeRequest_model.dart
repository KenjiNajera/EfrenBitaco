
class TypeRequestModel {
  final String? name;
  final int? id;
  
  TypeRequestModel({
    this.name,
    this.id,
  });

  factory TypeRequestModel.fronJson(Map<String, dynamic> json) => new TypeRequestModel(
    name: json['typerequestname'],
    id: json['id'],
  );

}