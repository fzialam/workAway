import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      title: 'Flutter MySQL Example',
      home: HomeScreen(),
    );
  }
}

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() {
    // Avoid using private types in public APIs.
    return _HomeScreenState();
  }
}

class _HomeScreenState extends State<HomeScreen> {
  String data = ''; // Store the fetched data

  Future<void> fetchData() async {
    // Replace with your API endpoint
    final apiUrl =
        Uri.parse('http://10.60.5.168:3000/api/data'); // Change to your API URL
    final response = await http.get(apiUrl);

    if (response.statusCode == 200) {
      final jsonData = json.decode(response.body);
      setState(() {
        data = jsonData['message']; // Adjust to your API response structure
      });
    } else {
      setState(() {
        data = 'Failed to fetch data';
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Flutter MySQL Example'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text(
              'Data from MySQL:',
              style: TextStyle(fontSize: 18),
            ),
            const SizedBox(height: 10),
            Text(
              data,
              style: const TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: fetchData,
              child: const Text('Fetch Data'),
            ),
          ],
        ),
      ),
    );
  }
}
