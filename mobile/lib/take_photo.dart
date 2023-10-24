import 'dart:io';

import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
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
  String data = '';

  Future fetchData() async {
    // Di sini Anda dapat mengambil data dari server (misalnya, melalui HTTP request)
    // Contoh sederhana:
    data = 'Data dari server untuk user ${widget.user.id}';
    dataAvailable = true;
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
      // ignore: use_build_context_synchronously
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return const AlertDialog(
            title: Text(
              'ERROR',
              style: TextStyle(
                  color: Colors.red,
                  fontWeight: FontWeight.w900,
                  fontFamily: "HeadlandOne"),
            ),
            icon: Icon(Icons.error_sharp),
            content: Text(
              "No Camera",
              style: TextStyle(
                  color: Color.fromARGB(255, 158, 158, 158),
                  fontWeight: FontWeight.bold,
                  fontFamily: "HeadlandOne"),
            ),
          );
        },
      );
    }

    var imageTemp = File(image!.path);
    setState(() {
      this.image = imageTemp;
    });
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
