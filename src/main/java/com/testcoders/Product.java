package com.testcoders;

public class Product {
    private final String name;
    private int barCode;
    private final int discount;
    private double price;

    public Product(String name, int barCode, double price, ProductCategory productCategory) {
        this.name = name;
        this.barCode = barCode;
        this.price = price;
        this.discount = productCategory.getDiscountValue();
    }

    public int getBarCode() {
        return barCode;
    }

    public void setBarCode(int barCode) {
        this.barCode = barCode;
    }

    public double getPrice() {
        double discountPrice = price * (100-this.discount) / 100;
        return  Math.rint(discountPrice * 100) / 100;
    }

    public void setPrice(double price) {
        this.price = price;
    }
}
