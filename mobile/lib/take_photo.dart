import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:flutter/services.dart';
import 'package:image_picker/image_picker.dart';
import 'package:mobile/config.dart';
import 'package:path/path.dart';

class TakePhoto extends StatefulWidget {
  TakePhoto(int userId, int suratId);

  @override
  State<TakePhoto> createState() => _TakePhotoState();
}

class _TakePhotoState extends State<TakePhoto> {
  File? image;
  bool isLoading = false;
  Future pickImage() async {
    try {
      final image = await ImagePicker().pickImage(source: ImageSource.camera);

      if (image == null) return;

      final imageTemporary = File(image.path);
      setState(() {
        this.image = imageTemporary;
      });
      String imageExtension = extension(image.path);
      print('Jenis gambar: $imageExtension');
      int imageSizeInBytes = this.image!.lengthSync();

      double imageSizeInMB = imageSizeInBytes / (1024 * 1024);
      print('Ukuran gambar: $imageSizeInMB MB');
    } on PlatformException catch (e) {
      debugPrint('failed pick image: $e');
    }
  }

  Future<void> postImageToServer(File imageFile) async {
    try {
      setState(() {
        isLoading = true; // Menandai bahwa proses pengiriman dimulai
      });
      // Membuka file gambar sebagai byte stream
      List<int> imageBytes = imageFile.readAsBytesSync();

      // Mengonversi byte stream gambar ke base64
      String base64Image = base64Encode(imageBytes);

      // Membuat objek JSON dengan data gambar dalam format base64
      Map<String, dynamic> jsonData = {
        'gambar': base64Image,
        'lokasi': 'disini',
        'surat_tugas_id': 1,
      };

      // Membuat permintaan HTTP POST
      var uri = Uri.parse('$URL/1/mobile'); // Ganti URL server sesuai kebutuhan
      var headers = {'Content-Type': 'application/json'};
      var body = jsonEncode(jsonData);

      // Menjalankan permintaan dan mendapatkan respons
      var response = await http.post(uri, headers: headers, body: body);
      // Cek kode status respons
      if (response.statusCode == 200) {
        print('Gambar berhasil diunggah');
      } else {
        print('Gagal mengunggah gambar. Kode status: ${response.statusCode}');
      }
    } catch (e) {
      print('Error: $e');
    } finally {
      setState(() {
        isLoading = false; // Menandai bahwa proses pengiriman telah selesai
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.amber.shade300,
      body: Container(
        padding: EdgeInsets.all(32),
        child: Column(
          children: [
            const Spacer(),
            image != null
                ? Image.file(
                    image!,
                    width: 160,
                    height: 160,
                    fit: BoxFit.cover,
                  )
                : const FlutterLogo(size: 160),
            const SizedBox(height: 24),
            const Text(
              'Image Picker',
              style: TextStyle(
                fontSize: 48,
                fontWeight: FontWeight.bold,
              ),
            ),
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                foregroundColor: Colors.black,
                backgroundColor: Colors.white,
                textStyle: const TextStyle(fontSize: 20),
                minimumSize: const Size.fromHeight(56),
              ),
              child: const Row(
                children: [
                  Icon(Icons.camera_alt_outlined, size: 28),
                  SizedBox(width: 16),
                  Text('Pick Camera'),
                ],
              ),
              onPressed: () {
                pickImage();
              },
            ),
            const Spacer(),
            ElevatedButton(
              onPressed: isLoading
                  ? null
                  : () {
                      debugPrint('clicked');
                      postImageToServer(image!);
                    },
              child: const Row(
                children: [
                  Icon(
                    Icons.check,
                    size: 28,
                  ),
                  SizedBox(
                    width: 16,
                  ),
                  Text('Send')
                ],
              ),
            ),
            isLoading ? const CircularProgressIndicator() : Container(),
          ],
        ),
      ),
    );
  }
}
