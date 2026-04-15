import 'package:flutter/material.dart';

class ProductCard extends StatelessWidget {
  const ProductCard({
    super.key,
    required this.productName,
    required this.productPrice,
    required this.isOutOfStock,
  });

  final String productName;
  final double productPrice;
  final bool isOutOfStock;

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 300,
      height: 350,
      decoration: BoxDecoration(border: Border.all()),
      child: Column(
        children: [
          Container(
            width: 250,
            height: 250,
            padding: EdgeInsets.all(5),
            child: Image.network(
              "https://nobero.com/cdn/shop/files/burgundy_e86469e5-d1fb-495c-aa57-e07d8a672b0a.jpg",
              fit: BoxFit.cover,
            ),
          ),
          Text(
            productName,
            style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
          ),
          Text("Price \$$productPrice"),
          Container(
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                ElevatedButton.icon(
                  onPressed: () {},
                  icon: const Icon(Icons.visibility), // Your icon
                  label: const Text('View'), // Your text
                ),
                ElevatedButton.icon(
                  onPressed: () {},
                  icon: const Icon(Icons.shopping_cart), // Your icon
                  label: const Text('Add to cart'), // Your text
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
