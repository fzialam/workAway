// To parse this JSON data, do
//
//     final suratPresensi = suratPresensiFromJson(jsonString);

import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:mobile/config.dart';

List<SuratPresensi> suratPresensiFromJson(String str) =>
    List<SuratPresensi>.from(
        json.decode(str).map((x) => SuratPresensi.fromJson(x)));

String suratPresensiToJson(List<SuratPresensi> data) =>
    json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class SuratPresensi {
  int id;
  String lokasiTujuan;
  String jenisProgram;
  String tglAwal;
  String tglAkhir;
  int gambarId;
  String nameGambar;
  String gambar;
  String lokasi;
  String koordinat;

  SuratPresensi({
    required this.id,
    required this.lokasiTujuan,
    required this.jenisProgram,
    required this.tglAwal,
    required this.tglAkhir,
    required this.gambarId,
    required this.nameGambar,
    required this.gambar,
    required this.lokasi,
    required this.koordinat,
  });

  factory SuratPresensi.fromJson(Map<String, dynamic> json) => SuratPresensi(
        id: json["id"],
        lokasiTujuan: json["lokasi_tujuan"],
        jenisProgram: json["jenis_program"],
        tglAwal: json["tgl_awal"],
        tglAkhir: json["tgl_akhir"],
        gambarId: json["gambar_id"],
        nameGambar: json["name_gambar"],
        gambar: json["gambar"],
        lokasi: json["lokasi"],
        koordinat: json["koordinat"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "lokasi_tujuan": lokasiTujuan,
        "jenis_program": jenisProgram,
        "tgl_awal": tglAwal,
        "tgl_akhir": tglAkhir,
        "gambar_id": gambarId,
        "name_gambar": nameGambar,
        "gambar": gambar,
        "lokasi": lokasi,
        "koordinat": koordinat,
      };

  Future<http.Response> getSuratPresensi(int userId) async {
    final url = Uri.parse('$URL/$userId/mobile');
    final response = await http.get(url);

    return response;
  }
}
