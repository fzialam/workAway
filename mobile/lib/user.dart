class User {
  final int id;
  final String nip;
  final String email;

  User({required this.id, required this.nip, required this.email});

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      nip: json['nip'],
      email: json['email'],
    );
  }
}
