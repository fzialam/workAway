import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:geocoding/geocoding.dart';
import 'package:geolocator/geolocator.dart';
import 'package:http/http.dart' as http;
import 'package:flutter/services.dart';
import 'package:image_picker/image_picker.dart';
import 'package:mobile/config.dart';
import 'package:mobile/list_surat.dart';
// import 'package:mobile/take_photo.dart';
import 'package:path/path.dart';
import './model/surat.dart';
import './model/user.dart';

class DetailSurat extends StatefulWidget {
  final SuratPresensi surat;
  final User user;
  DetailSurat(this.surat, this.user);

  @override
  State<DetailSurat> createState() => _DetailSuratState();
}

class _DetailSuratState extends State<DetailSurat> {
  File? image;
  bool isLoading = false;
  Position? _currentLocation;
  late Position realCurrentLocation;
  late bool servicePermission = false;
  late LocationPermission permission;

  String _currentAddress = "";

  Future pickImage() async {
    try {
      final image = await ImagePicker().pickImage(source: ImageSource.camera);

      if (image == null) return;

      final imageTemporary = File(image.path);
      setState(() {
        this.image = imageTemporary;
      });
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
      String imageExtension = extension(imageFile.path);
      debugPrint('Jenis gambar: $imageExtension');
      List<int> imageBytes = imageFile.readAsBytesSync();

      // Mengonversi byte stream gambar ke base64
      String base64Image = '';
      if (imageExtension.contains('.jpg')) {
        base64Image = 'data:image/jpg;base64,${base64Encode(imageBytes)}';
      } else if (imageExtension.contains('.png')) {
        base64Image = 'data:image/png;base64,${base64Encode(imageBytes)}';
      }

      String laS = _currentLocation?.latitude.toString() as String;
      String loS = _currentLocation?.longitude.toString() as String;

      // Membuat objek JSON dengan data gambar dalam format base64
      Map<String, dynamic> jsonData = {
        'surat_tugas_id': widget.surat.id,
        'gambar': base64Image,
        'lokasi': _currentAddress,
        'koordinat': '$laS, $loS',
      };

      // Membuat permintaan HTTP POST
      var uri = Uri.parse(
          '$URL/${widget.user.id}/mobile'); // Ganti URL server sesuai kebutuhan
      var headers = {'Content-Type': 'application/json'};
      var body = jsonEncode(jsonData);

      // Menjalankan permintaan dan mendapatkan respons
      var response = await http.post(uri, headers: headers, body: body);
      // Cek kode status respons
      if (response.statusCode == 200) {
        debugPrint('Gambar berhasil diunggah');

        Navigator.push(
          context as BuildContext,
          MaterialPageRoute(builder: (_) => GetSurat(widget.user)),
        );
      } else {
        debugPrint(
            'Gagal mengunggah gambar. Kode status: ${response.statusCode}');
      }
    } catch (e) {
      debugPrint('Error: $e');
    } finally {
      setState(() {
        isLoading = false; // Menandai bahwa proses pengiriman telah selesai
      });
    }
  }

  Future<void> _getCurrrentLocation() async {
    servicePermission = await Geolocator.isLocationServiceEnabled();
    if (!servicePermission) {
      debugPrint("Service Disabled");
    }
    checkLocationPermission();

    try {
      _currentLocation = await Geolocator.getCurrentPosition();
      // Update the UI after obtaining the location.
      _getAddressFromCoordinates();
    } catch (e) {
      debugPrint("Error getting location: $e");
    }
  }

  Future<void> checkLocationPermission() async {
    permission = await Geolocator.checkPermission();

    if (permission == LocationPermission.denied) {
      permission = await Geolocator.requestPermission();
    }
  }

  Future<void> _getAddressFromCoordinates() async {
    String laS = _currentLocation?.latitude.toString() as String;
    String loS = _currentLocation?.longitude.toString() as String;

    double latitudeDouble = double.parse(laS);
    double longitudeDouble = double.parse(loS);
    try {
      List<Placemark> placemarks = await placemarkFromCoordinates(
        latitudeDouble,
        longitudeDouble,
      );

      Placemark place = placemarks[0];

      if (placemarks.isNotEmpty) {
        setState(() {
          _currentAddress =
              "${place.street}, ${place.subLocality}, ${place.locality}, ${place.country}, ${place.postalCode}";
        });
      }
    } catch (e) {
      debugPrint('$e');
    }
  }

  Image imageFromBase64() {
    String data = widget.surat.gambar;

    // Memisahkan string menggunakan metode split
    List<String> parts = data.split(',');

    // Mengambil bagian kedua (base64)
    String base64String = parts.length > 1 ? parts[1] : '';

    Uint8List bytes = base64Decode(base64String);

    // Membuat objek Image dari byte array
    return Image.memory(
      bytes,
      height: 160,
      width: 160,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('${widget.surat.id} | ${widget.surat.jenisProgram}'),
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Align(
                alignment: Alignment.centerLeft,
                child: Title(
                  color: Colors.black,
                  child: Text(
                    "Program: ${widget.surat.jenisProgram}",
                    style: const TextStyle(
                      fontSize: 22,
                      fontWeight: FontWeight.normal,
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 10.0),
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Lokasi Tujuan: ${widget.surat.lokasiTujuan}",
                  style: const TextStyle(
                    fontSize: 22,
                    fontWeight: FontWeight.normal,
                  ),
                ),
              ),
              const SizedBox(height: 10.0),
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Tanggal Awal: ${widget.surat.tglAwal}",
                  style: const TextStyle(
                    fontSize: 22,
                    fontWeight: FontWeight.normal,
                  ),
                ),
              ),
              const SizedBox(height: 10.0),
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Tanggal Akhir: ${widget.surat.tglAkhir}",
                  style: const TextStyle(
                    fontSize: 22,
                    fontWeight: FontWeight.normal,
                  ),
                ),
              ),
              const SizedBox(height: 50.0),
              if (widget.surat.gambar.isNotEmpty)
                Align(
                  alignment: Alignment.center,
                  child: Column(
                    children: [
                      imageFromBase64(),
                      const SizedBox(height: 16),
                      const Text(
                        "Sudah Melakukan presensi",
                        style: TextStyle(
                          color: Colors.grey,
                          fontSize: 30,
                          fontWeight: FontWeight.w700,
                        ),
                      ),
                    ],
                  ),
                ),
              if (widget.surat.gambar.isEmpty)
                Align(
                  alignment: Alignment.center,
                  child: image != null
                      ? Image.file(
                          image!,
                          width: 160,
                          height: 160,
                          fit: BoxFit.cover,
                        )
                      : const Icon(Icons.image_not_supported_outlined,
                          size: 160),
                ),
              const SizedBox(height: 50.0),
              const Spacer(),
              if (widget.surat.gambar.isEmpty)
                ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    foregroundColor: Colors.black,
                    backgroundColor: Colors.white,
                    textStyle: const TextStyle(fontSize: 20),
                    minimumSize: const Size.fromHeight(56),
                  ),
                  onPressed: () {
                    pickImage();
                    _getCurrrentLocation();
                  },
                  child: const Row(
                    children: [
                      Icon(Icons.camera_alt_outlined, size: 28),
                      SizedBox(width: 16),
                      Text('Pick Camera'),
                    ],
                  ),
                ),
              const Spacer(),
              if (widget.surat.gambar.isEmpty)
                ElevatedButton(
                  onPressed: isLoading || image == null
                      ? () {
                          showDialog(
                              context: context,
                              builder: (BuildContext context) {
                                return AlertDialog(
                                  title: const Align(
                                    alignment: Alignment.center,
                                    child: Text(
                                      "Ambil Gambar Terlebih Dahulu!!",
                                      style: TextStyle(
                                        color: Colors.red,
                                        fontWeight: FontWeight.w900,
                                        fontFamily: "HeadlandOne",
                                      ),
                                    ),
                                  ),
                                  actions: <Widget>[
                                    TextButton(
                                      onPressed: () {
                                        Navigator.of(context).pop();
                                      },
                                      child: const Text(
                                        'OK',
                                        style: TextStyle(
                                          color:
                                              Color.fromARGB(186, 244, 67, 54),
                                          fontWeight: FontWeight.bold,
                                          fontFamily: "HeadlandOne",
                                        ),
                                      ),
                                    ),
                                  ],
                                );
                              });
                          null;
                        }
                      : () {
                          postImageToServer(image!);
                        },
                  child: const Row(
                    children: [
                      Icon(
                        Icons.check,
                        size: 28,
                      ),
                      SizedBox(width: 16),
                      Text('Send'),
                    ],
                  ),
                ),
              const SizedBox(height: 16),
              if (isLoading) const CircularProgressIndicator(),
            ],
          ),
        ),
      ),
    );
  }
}
