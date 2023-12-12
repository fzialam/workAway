import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/config.dart';

String userToJson(User data) => json.encode(data.toJson());

class User {
  final int id;
  final String name;
  final int rank;
  final String email;
  final String password;
  final String token;

  User({
    required this.id,
    required this.name,
    required this.rank,
    required this.email,
    required this.password,
    required this.token,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      rank: json['rank'],
      name: json['name'],
      email: json['email'],
      password: '',
      token: json['token'],
    );
  }

  Future<http.Response> login(User user) async {
    final url = Uri.parse("$URL/login?m=true");

    final response = await http.post(
      url,
      headers: {
        'Content-Type':
            'application/json', // Set header Content-Type ke application/json
      },
      body: jsonEncode(user.toJson()),
    );

    return response;
  }

  Map<String, dynamic> toJson() => {
        "id": id,
        "name": name,
        "rank": rank,
        "email": email,
        "password": password,
        "token": token,
      };
}
