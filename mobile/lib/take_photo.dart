import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/config.dart';

import 'package:mobile/user.dart';

class TakePhoto extends StatefulWidget {
  final User user;
  const TakePhoto(this.user, {super.key});

  @override
  State<TakePhoto> createState() => _TakePhotoState();
}

class _TakePhotoState extends State<TakePhoto> {
  File? image;
  bool dataAvailable = false;
  late final http.Response response;

  Future fetchData() async {
    final url = Uri.parse('$URL/wp/${widget.user.id}/mobile');
    response = await http.get(url);

    if (response.statusCode == 200) {
      // Successful response with a status code of 200
      setState(() {
        null;
      });
    } else {
      // Handle errors if the request was not successful
      debugPrint("No Data");
      showErrorDialog(response);
      null;
    }
  }

  Future refresh() async {
    fetchData();
    if (dataAvailable == true) {
      cameraScreen();
    }
  }

  Future<void> cameraScreen() async {
    final image = await ImagePicker().pickImage(source: ImageSource.camera);

    if (image == null) {
      showErrorDialog(response);
    } else {
      var imageTemp = File(image.path);
      setState(() {
        this.image = imageTemp;
      });
    }
    dataAvailable = false;
  }

  @override
  void initState() {
    super.initState();
    fetchData(); // Panggil fetchData saat halaman diinisialisasi
    if (dataAvailable == true) {
      cameraScreen();
    }
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
    // Jika data tidak tersedia, tampilkan "Take Photo" dan tombol refresh
    return Scaffold(
      appBar: AppBar(
        title: const Text('Take Photo'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text('Data tidak tersedia'),
            ElevatedButton(
              onPressed: refresh, // Panggil fetchData saat tombol ditekan
              child: const Text('Refresh'),
            ),
          ],
        ),
      ),
    );
  }
}
