import 'package:flutter/material.dart';
import 'package:geocoding/geocoding.dart';
import 'package:geolocator/geolocator.dart';
import 'package:mobile/splash_screen.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      title: "WORKAWAY",
      // debugShowCheckedModeBanner: false,
      home: SplashScreen(),
      // home: GeolocationApp(),
    );
  }
}

class GeolocationApp extends StatefulWidget {
  const GeolocationApp(int idUser, int idSurat, {super.key});

  @override
  State<GeolocationApp> createState() => _GeolocationAppState();
}

class _GeolocationAppState extends State<GeolocationApp> {
  Position? _currentLocation;
  late Position realCurrentLocation;
  late bool servicePermission = false;
  late LocationPermission permission;

  String _currentAddress = "";

  @override
  void initState() {
    super.initState();

    // _getCurrrentLocation();
    checkLocationPermission();
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
              "${place.street}, ${place.subLocality}, ${place.locality},  ${place.country}, ${place.postalCode}";
        });
      }
    } catch (e) {
      debugPrint('$e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Get Current Location"),
        centerTitle: true,
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            const Text(
              "Location Coordinates",
              style: TextStyle(
                fontSize: 32,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(
              height: 20.0,
            ),
            Text(
              "Latitude: ${_currentLocation?.latitude ?? 'Loading...'}, Longitude: ${_currentLocation?.longitude ?? 'Loading...'}",
              style: const TextStyle(
                fontSize: 18,
              ),
            ),
            const SizedBox(
              height: 50.0,
            ),
            const Text(
              "Location Address",
              style: TextStyle(
                fontSize: 32,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(
              height: 20.0,
            ),
            Text(
              _currentAddress,
              style: const TextStyle(
                fontSize: 18,
              ),
            ),
            const SizedBox(
              height: 50.0,
            ),
            ElevatedButton(
              onPressed: () async {
                await _getCurrrentLocation();
                await _getAddressFromCoordinates();
                debugPrint("$_currentLocation");
              },
              child: const Text("Press"),
            )
          ],
        ),
      ),
    );
  }
}
