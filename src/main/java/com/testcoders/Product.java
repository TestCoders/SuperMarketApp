package com.testcoders;

public class Product {
    private String name;
    private int barCode;
    private double price;
    private int discount;

    public Product(String name, int barCode, double price, int discount) {
        this.name = name;
        this.barCode = barCode;
        this.price = price;
        this.discount = discount;
    }

    public int getBarCode() {
        return barCode;
    }

    public void setBarCode(int barCode) {
        this.barCode = barCode;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }

    public int getDiscount() {
        return discount;
    }

    public void setDiscount(int discount) {
        this.discount = discount;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
