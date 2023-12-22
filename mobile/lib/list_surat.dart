import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_exit_app/flutter_exit_app.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/detail.dart';
import 'package:mobile/model/surat.dart';

import 'package:mobile/model/user.dart';

class GetSurat extends StatefulWidget {
  final User user;
  const GetSurat(this.user, {super.key});

  @override
  State<GetSurat> createState() => _GetSuratState();
}

class _GetSuratState extends State<GetSurat> {
  File? image;
  bool dataAvailable = false;
  late Future<List<SuratPresensi>> suratPresensiListFuture;

  @override
  void initState() {
    super.initState();

    suratPresensiListFuture = fetchSuratPresensiList(widget.user.id);
  }

  Future<bool?> showExitConfirmationDialog(BuildContext context) async {
    return showDialog<bool>(
      context: context,
      builder: (context) => AlertDialog(
        title: const Align(
          alignment: Alignment.center,
          child: Text(
            'Konfirmasi Keluar',
            style: TextStyle(
              fontWeight: FontWeight.w600,
              fontSize: 24,
              color: Colors.red,
            ),
          ),
        ),
        content: const SizedBox(
          height: 50,
          child: Align(
            alignment: Alignment.center,
            child: Text(
              'Apakah Anda yakin keluar?',
              style: TextStyle(
                fontWeight: FontWeight.normal,
                fontSize: 20,
                color: Colors.grey,
              ),
            ),
          ),
        ),
        actions: <Widget>[
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              TextButton(
                onPressed: () {
                  // Batal keluar, tutup dialog, dan kembalikan nilai false
                  Navigator.of(context).pop(false);
                },
                style: TextButton.styleFrom(
                  backgroundColor: Colors.blue.shade600,
                ),
                child: const Row(
                  children: [
                    Icon(
                      Icons.arrow_back_ios_sharp,
                      color: Colors.white,
                    ),
                    Text(
                      'Batal',
                      style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: 16,
                      ),
                    ),
                  ],
                ),
              ),
              const SizedBox(
                width: 5,
              ),
              TextButton(
                onPressed: () {
                  // Konfirmasi keluar, tutup dialog, dan kembalikan nilai true
                  Navigator.of(context).pop(true);
                },
                style: TextButton.styleFrom(
                  backgroundColor: Colors.red.shade500,
                ),
                child: const Row(
                  children: [
                    Icon(
                      Icons.exit_to_app,
                      color: Colors.white,
                    ),
                    Text(
                      'Keluar',
                      style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: 16,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          )
        ],
      ),
    );
  }

  Future<List<SuratPresensi>> fetchSuratPresensiList(int userId) async {
    SuratPresensi suratPresensi = SuratPresensi(
      id: 0,
      lokasiTujuan: "",
      jenisProgram: "",
      tglAwal: "",
      tglAkhir: "",
      gambarId: 0,
      nameGambar: "",
      gambar: "",
      lokasi: "",
      koordinat: "",
    );

    final response = await suratPresensi.getSuratPresensi(userId);
    if (response.statusCode == 200) {
      return suratPresensiFromJson(response.body);
    } else {
      debugPrint("No Data");
      showErrorDialog(response);
      return [];
    }
  }

  Future<void> _refreshData() async {
    setState(() {
      suratPresensiListFuture = fetchSuratPresensiList(widget.user.id);
    });
    debugPrint('refresh');
  }

  void showErrorDialog(http.Response response) {
    final responseData = json.decode(response.body);
    final String message =
        "ERROR ${response.statusCode}: Terjadi kesalahan ${responseData['message']}";

    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text(
            'ERROR',
            style: TextStyle(
                color: Colors.red,
                fontWeight: FontWeight.w900,
                fontFamily: "HeadlandOne"),
          ),
          icon: const Icon(Icons.error_sharp),
          content: Text(
            message,
            style: const TextStyle(
                color: Color.fromARGB(255, 158, 158, 158),
                fontWeight: FontWeight.bold,
                fontFamily: "HeadlandOne"),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text(
                'OK',
                style: TextStyle(
                  color: Color.fromARGB(186, 244, 67, 54),
                  fontWeight: FontWeight.bold,
                  fontFamily: "HeadlandOne",
                ),
              ),
              onPressed: () {
                Navigator.of(context).pop(); // Tutup dialog
              },
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return WillPopScope(
      onWillPop: () async {
        // Tampilkan dialog konfirmasi sebelum keluar
        bool? exitConfirmed = await showExitConfirmationDialog(context);

        if (exitConfirmed!) {
          FlutterExitApp.exitApp();
          return exitConfirmed;
        } else {
          return false;
        }
      },
      child: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          elevation: 2,
          backgroundColor: Colors.blue.shade50,
          title: const Align(
            alignment: Alignment.center,
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.center,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Image(
                  image: AssetImage('assets/images/logo.png'),
                  height: 40,
                  width: 40,
                ),
                SizedBox(
                  width: 20,
                ),
                Text(
                  'WorkAway',
                  style: TextStyle(
                    color: Colors.black,
                    fontSize: 30,
                    fontWeight: FontWeight.w800,
                  ),
                ),
              ],
            ),
          ),
        ),
        body: RefreshIndicator(
            onRefresh: _refreshData,
            child: FutureBuilder<List<SuratPresensi>>(
              future: suratPresensiListFuture,
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return const CircularProgressIndicator();
                } else if (snapshot.hasError) {
                  return Text('Error: ${snapshot.error}');
                } else {
                  // Assuming your SuratPresensiList is a List<SuratPresensi>
                  List<SuratPresensi>? listSuratPresensi = snapshot.data;
                  return ListView.builder(
                    itemCount: listSuratPresensi?.length ?? 0,
                    itemBuilder: (context, index) {
                      SuratPresensi suratPresensi = listSuratPresensi![index];
                      if (suratPresensi.gambar.isEmpty) {
                        return ListTile(
                            leading: const CircleAvatar(
                              backgroundColor: Colors.red,
                              child: Icon(
                                Icons.close,
                                color: Colors.white,
                              ),
                            ),
                            title: Text(
                              suratPresensi.jenisProgram,
                              style: const TextStyle(
                                fontSize: 20,
                                fontWeight: FontWeight.w600,
                                color: Colors.black,
                              ),
                            ),
                            subtitle: Text(
                              suratPresensi.lokasiTujuan,
                            ),
                            textColor: Colors.black,
                            onTap: () => Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => DetailSurat(
                                      listSuratPresensi[index],
                                      User(
                                        id: widget.user.id,
                                        name: widget.user.name,
                                        rank: widget.user.rank,
                                        email: widget.user.email,
                                        password: widget.user.password,
                                        token: widget.user.token,
                                      ),
                                    ),
                                  ),
                                ));
                      } else {
                        return ListTile(
                          leading: const CircleAvatar(
                            backgroundColor: Colors.green,
                            child:
                                Icon(Icons.check_circle, color: Colors.white),
                          ),
                          title: Text(
                            suratPresensi.jenisProgram,
                            style: const TextStyle(
                              fontSize: 20,
                              fontWeight: FontWeight.w600,
                              color: Colors.black,
                            ),
                          ),
                          subtitle: Text(
                            suratPresensi.lokasiTujuan,
                          ),
                          textColor: Colors.black,
                          onTap: () => Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => DetailSurat(
                                listSuratPresensi[index],
                                User(
                                  id: widget.user.id,
                                  name: widget.user.name,
                                  rank: widget.user.rank,
                                  email: widget.user.email,
                                  password: widget.user.password,
                                  token: widget.user.token,
                                ),
                              ),
                            ),
                          ),
                        );
                      }
                    },
                  );
                }
              },
            )),
      ),
    );
  }
}
