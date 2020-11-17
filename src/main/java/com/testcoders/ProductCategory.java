package com.testcoders;

public class ProductCategory {
    private static String category;
    private static int discountValue;

    public ProductCategory(String category) {
        if (category.equals("Ham") || category.equals("Beer")) {
            discountValue = 20;
        } else if (category.equals("Milk")) {
            discountValue = 35;
        } else {
            discountValue = 0;
        }
    }

    public int getDiscountValue() {
        return discountValue;
    }
}
