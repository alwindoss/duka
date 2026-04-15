import 'package:flutter/material.dart';
import 'package:ui/components/products/product_card.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key, required this.title});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: Center(
        child: Container(
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              ProductCard(
                productName: "Polo T-shirt",
                productPrice: 10.39,
                isOutOfStock: false,
                productThumbnailURL:
                    "https://nobero.com/cdn/shop/files/burgundy_e86469e5-d1fb-495c-aa57-e07d8a672b0a.jpg",
              ),
              ProductCard(
                productName: "Cotton Formal shirt",
                productPrice: 25.99,
                isOutOfStock: false,
                productThumbnailURL:
                    "https://hamercop.com/cdn/shop/files/3303_3_800x.jpg",
              ),
              ProductCard(
                productName: "Cotton floral dress",
                productPrice: 89.99,
                isOutOfStock: false,
                productThumbnailURL:
                    "https://kasya.in/cdn/shop/products/DRE131-01_ae6c36e6-55e5-4684-9f28-9c7a21c84015.jpg",
              ),
              ProductCard(
                productName: "Men's Black leather shoes",
                productPrice: 190.79,
                isOutOfStock: false,
                productThumbnailURL:
                    "https://tiimg.tistatic.com/fp/1/007/742/comfortable-foot-friendly-easy-to-use-plain-black-light-heel-formal-shoes-for-men-779.jpg",
              ),
            ],
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {},
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}
