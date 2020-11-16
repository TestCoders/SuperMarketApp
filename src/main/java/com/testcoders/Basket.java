package com.testcoders;

import java.util.ArrayList;

public class Basket {
    private final ArrayList<Product> products;

    public Basket() {
        this.products = new ArrayList<Product>();
    }

    public void addProduct(Product product) {
        this.products.add(product);
    }

    public ArrayList<Product> getProducts() {
        return products;
    }

    public double getTotal() {
        return products.stream().mapToDouble(Product::getPrice).sum();
    }

    public double getTotalWithDiscount() {
        return products.stream().mapToDouble(Product::getPrice).sum();
    }
}
