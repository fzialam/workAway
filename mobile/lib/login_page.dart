import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/listSurat.dart';
import 'package:mobile/model/user.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final TextEditingController email = TextEditingController();
  final TextEditingController password = TextEditingController();

  @override
  void dispose() {
    email.clear();
    password.clear();
    super.dispose();
  }

  final _formKey = GlobalKey<FormState>();

  Future<void> loginToServer(String email, password) async {
    final user = User(
        id: 0, name: "", rank: 0, email: email, password: password, token: "");

    final response = await user.login(user);

    if (response.statusCode == 200) {
      debugPrint('Data berhasil dikirim ke server');
      final userData = jsonDecode(response.body);
      final user = User.fromJson(userData);

      // ignore: use_build_context_synchronously
      Navigator.push(
        context,
        MaterialPageRoute(builder: (_) => GetSurat(user)),
      );
    } else {
      debugPrint(
          'Gagal mengirim data ke server. Status code: ${response.statusCode}');
      showErrorDialog(response);
    }
  }

  void showErrorDialog(http.Response response) {
    final responseData = json.decode(response.body);
    final String message =
        "ERROR ${response.statusCode}: Terjadi kesalahan ${responseData['data']}";
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
      body: Container(
        decoration: const BoxDecoration(
          // color: Colors.red.withOpacity(0.1),
          image: DecorationImage(
            image: AssetImage("assets/images/bg2.jpg"),
            fit: BoxFit.cover,
            colorFilter: ColorFilter.mode(Colors.black, BlendMode.dstATop),
            opacity: 0.35,
          ),
        ),
        child: SafeArea(
          child: Center(
            child: SingleChildScrollView(
              scrollDirection: Axis.vertical,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  // logo here
                  Image.asset(
                    'assets/images/WORKAWAY.png',
                    height: 250,
                    width: 250,
                  ),
                  const Text(
                    'LOG-IN NOW',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 26,
                      fontWeight: FontWeight.bold,
                      fontFamily: "HeadlandOne",
                    ),
                  ),

                  const SizedBox(
                    height: 30,
                  ),
                  Container(
                    height: 240,
                    width: MediaQuery.of(context).size.width / 1.1,
                    decoration: BoxDecoration(
                        color: Colors.white.withOpacity(0.2),
                        borderRadius: BorderRadius.circular(20)),
                    child: Column(
                      children: [
                        Padding(
                          padding: const EdgeInsets.only(
                              left: 20, right: 20, bottom: 20, top: 20),
                          child: TextFormField(
                            controller: email,
                            decoration: const InputDecoration(
                              focusedBorder: UnderlineInputBorder(
                                  borderSide: BorderSide.none,
                                  borderRadius:
                                      BorderRadius.all(Radius.circular(10))),
                              enabledBorder: UnderlineInputBorder(
                                  borderSide: BorderSide.none,
                                  borderRadius:
                                      BorderRadius.all(Radius.circular(10))),
                              prefixIcon: Icon(
                                Icons.person,
                                color: Colors.purple,
                              ),
                              filled: true,
                              fillColor: Colors.white,
                              labelText: "Email",
                              hintText: 'ucup@unesa.ac.id',
                              labelStyle: TextStyle(color: Colors.purple),
                            ),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.only(left: 20, right: 20),
                          child: Form(
                            key: _formKey,
                            child: TextFormField(
                              controller: password,
                              obscuringCharacter: '*',
                              obscureText: true,
                              decoration: const InputDecoration(
                                focusedBorder: UnderlineInputBorder(
                                    borderSide: BorderSide.none,
                                    borderRadius:
                                        BorderRadius.all(Radius.circular(10))),
                                enabledBorder: UnderlineInputBorder(
                                    borderSide: BorderSide.none,
                                    borderRadius:
                                        BorderRadius.all(Radius.circular(10))),
                                prefixIcon: Icon(
                                  Icons.person,
                                  color: Colors.purple,
                                ),
                                filled: true,
                                fillColor: Colors.white,
                                labelText: "Password",
                                hintText: '*********',
                                labelStyle: TextStyle(color: Colors.purple),
                              ),
                              validator: (value) {
                                if (value!.isEmpty && value.length < 5) {
                                  return 'Enter a valid password';
                                }
                                return null;
                              },
                            ),
                          ),
                        ),
                        const SizedBox(
                          height: 20,
                        ),
                        Padding(
                          padding: const EdgeInsets.only(left: 20, right: 20),
                          child: ElevatedButton(
                            style: ElevatedButton.styleFrom(
                                shape: RoundedRectangleBorder(
                                    borderRadius: BorderRadius.circular(80.0)),
                                backgroundColor:
                                    const Color.fromRGBO(0, 74, 173, 1.0),
                                padding: EdgeInsets.symmetric(
                                    horizontal:
                                        MediaQuery.of(context).size.width / 3.3,
                                    vertical: 20)
                                // padding: EdgeInsets.only(
                                //     left: 120, right: 120, top: 20, bottom: 20),
                                ),
                            onPressed: () {
                              loginToServer(email.text, password.text);
                            },
                            child: const Text(
                              'Log In',
                              style: TextStyle(
                                fontSize: 17,
                                fontFamily: "HeadlandOne",
                                fontStyle: FontStyle.normal,
                              ),
                            ),
                          ),
                        )
                      ],
                    ),
                  ),

                  // this is button
                  const SizedBox(
                    height: 30,
                  ),

                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        'You have\'t any account?',
                        style: TextStyle(
                          color: Colors.white.withOpacity(0.6),
                        ),
                      ),
                      TextButton(
                        onPressed: () {},
                        child: const Text(
                          'Sign Up',
                          style: TextStyle(
                              color: Color.fromRGBO(216, 42, 247, 0.9),
                              fontWeight: FontWeight.w500),
                        ),
                      )
                    ],
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
