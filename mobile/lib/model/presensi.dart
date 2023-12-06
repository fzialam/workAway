// To parse this JSON data, do
//
//     final presensi = presensiFromJson(jsonString);

import 'dart:convert';
import 'package:mobile/config.dart';
import 'package:http/http.dart' as http;

Presensi presensiFromJson(String str) => Presensi.fromJson(json.decode(str));

String presensiToJson(Presensi data) => json.encode(data.toJson());

Future<http.Response> send(int userId, Presensi presensi) async {
  final url = Uri.parse("$URL/wp/$userId/mobile");

  final response = await http.post(
    url,
    headers: {
      'Content-Type':
          'application/json', // Set header Content-Type ke application/json
    },
    body: jsonEncode(presensi.toJson()),
  );

  return response;
}

class Presensi {
  int id;
  int userId;
  int suratId;
  String name;
  String gambar;
  String lokasi;
  String koordinat;

  Presensi({
    required this.id,
    required this.userId,
    required this.suratId,
    required this.name,
    required this.gambar,
    required this.lokasi,
    required this.koordinat,
  });

  factory Presensi.fromJson(Map<String, dynamic> json) => Presensi(
        id: json["id"],
        userId: json["user_id"],
        suratId: json["surat_id"],
        name: json["name"],
        gambar: json["gambar"],
        lokasi: json["lokasi"],
        koordinat: json["koordinat"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "user_id": userId,
        "surat_id": suratId,
        "name": name,
        "gambar": gambar,
        "lokasi": lokasi,
        "koordinat": koordinat,
      };
}
