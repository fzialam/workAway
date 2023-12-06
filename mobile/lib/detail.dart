import 'package:flutter/material.dart';
import 'package:mobile/take_photo.dart';
import './model/surat.dart';
import './model/user.dart';

class DetailSurat extends StatelessWidget {
  final SuratPresensi surat;
  final User user;
  DetailSurat(this.surat, this.user);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          '${surat.id} | ${surat.jenisProgram}',
        ),
      ),
      body: Container(
        alignment: Alignment.center,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Title(
              color: Colors.black,
              child: Text(
                surat.jenisProgram,
              ),
            ),
            SizedBox(
              height: 50.0,
            ),
            Text(
              '${surat.lokasiTujuan}',
              style: TextStyle(
                fontSize: 30,
                fontWeight: FontWeight.normal,
              ),
            ),
            SizedBox(
              height: 50.0,
            ),
            ElevatedButton(
                onPressed: () => Navigator.push(
                      context,
                      MaterialPageRoute(
                          builder: (context) => TakePhoto(
                                user.id,
                                surat.id,
                              )),
                    ),
                child: Text(
                  'Ambil Foto',
                  style: TextStyle(
                    fontSize: 20,
                  ),
                ))
          ],
        ),
      ),
    );
  }
}
