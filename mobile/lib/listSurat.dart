import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/detail.dart';
import 'package:mobile/main.dart';
import 'package:mobile/model/surat.dart';
// import 'package:image_picker/image_picker.dart';
// import 'package:mobile/config.dart';

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

  Future<List<SuratPresensi>> fetchSuratPresensiList(int userId) async {
    // SuratPresensi suratPresensi = SuratPresensi(
    //   id: 0,
    //   lokasiTujuan: "",
    //   jenisProgram: "",
    //   tglAwal: "",
    //   tglAkhir: "",
    //   gambarId: 0,
    //   nameGambar: "",
    //   gambar: "",
    //   lokasi: "",
    //   koordinat: "",
    // );

    // final response = await suratPresensi.getSuratPresensi(userId);

    // if (response.statusCode == 200) {
    //   return suratPresensiFromJson(response.body);
    // } else {
    //   debugPrint("No Data");
    //   showErrorDialog(response);
    //   return []; // or return an empty list or throw an exception based on your requirement
    // }
    var x = """[
  {
    "id": 33,
    "lokasi_tujuan": "a",
    "jenis_program": "0",
    "tgl_awal": "2023-12-12",
    "tgl_akhir": "2023-12-16",
    "gambar_id": 0,
    "name_gambar": "",
    "gambar": "",
    "lokasi": "",
    "koordinat": ""
  },
  {
    "id": 43,
    "lokasi_tujuan": "AMERIKa",
    "jenis_program": "Berenang",
    "tgl_awal": "2023-12-20",
    "tgl_akhir": "2023-12-09",
    "gambar_id": 1,
    "name_gambar": "",
    "gambar": "",
    "lokasi": "",
    "koordinat": ""
  }
]""";
    return suratPresensiFromJson(x);
  }

  Future<void> _refreshData() async {
    setState(() {
      // suratPresensiListFuture = fetchSuratPresensiList(widget.user.id);
      debugPrint("Refresh");
    });
  }

  // Future<void> cameraScreen() async {
  //   final image = await ImagePicker().pickImage(source: ImageSource.camera);
  //
  //   if (image == null) {
  //     showErrorDialog(response);
  //   } else {
  //     var imageTemp = File(image.path);
  //     setState(() {
  //       this.image = imageTemp;
  //     });
  //   }
  //   dataAvailable = false;
  // }

  // @override
  // void initState() {
  //   super.initState();
  //   fetchData(); // Panggil fetchData saat halaman diinisialisasi
  //   if (dataAvailable == true) {
  //     cameraScreen();
  //   }
  // }

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
    return Scaffold(
      appBar: AppBar(
        title: const Text('Surat Presensi List'),
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
                  if (suratPresensi.gambarId == 0) {
                    return ListTile(
                        leading: const CircleAvatar(
                          backgroundColor: Colors.red,
                        ),
                        title: Text(
                            '${suratPresensi.id} SPPD ${suratPresensi.jenisProgram}',
                            style: const TextStyle(
                              fontSize: 20,
                              fontWeight: FontWeight.normal,
                            )),
                        subtitle: Text(
                          'Lokasi:${suratPresensi.lokasiTujuan}\nSurat${suratPresensi.tglAwal}',
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
                            )
                        // title: Text("""ID: ${suratPresensi.id} |
                        //   ${suratPresensi.jenisProgram}
                        //   ${suratPresensi.tglAwal}
                        //   ${suratPresensi.tglAkhir}"""),
                        // onTap: () => Navigator.push(
                        //     context,
                        //     MaterialPageRoute(
                        //         builder: (_) => const GeolocationApp())),
                        // // Add other fields you want to display
                        );
                  } else {
                    return ListTile(
                      leading: const CircleAvatar(
                        backgroundColor: Colors.green,
                      ),
                      title: Text(
                        '${suratPresensi.id} SPPD ${suratPresensi.jenisProgram}',
                        style: const TextStyle(
                          fontSize: 20,
                          fontWeight: FontWeight.normal,
                        ),
                      ),
                      subtitle: Text(
                        'Lokasi:${suratPresensi.lokasiTujuan}\nSurat${suratPresensi.tglAwal}',
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
                    // return ListTile(
                    //   title: Text("""ID: ${suratPresensi.id} |
                    //     ${suratPresensi.jenisProgram}
                    //     ${suratPresensi.tglAwal}
                    //     ${suratPresensi.tglAkhir}"""),
                    //   onTap: () => Navigator.push(
                    //       context,
                    //       MaterialPageRoute(
                    //           builder: (_) => const GeolocationApp())),
                    //   // Add other fields you want to display
                    // );
                  }
                },
              );
            }
          },
        ),
      ),
    );
  }
  // Jika data tidak tersedia, tampilkan "Take Photo" dan tombol refresh
  //   return Scaffold(
  //     appBar: AppBar(
  //       title: const Text('Take Photo'),
  //     ),
  //     body: Center(
  //       child: Column(
  //         mainAxisAlignment: MainAxisAlignment.center,
  //         children: <Widget>[
  //           const Text('Data tidak tersedia'),
  //           ElevatedButton(
  //             onPressed: refresh, // Panggil fetchData saat tombol ditekan
  //             child: const Text('Refresh'),
  //           ),
  //         ],
  //       ),
  //     ),
  //   );
  // }
}
