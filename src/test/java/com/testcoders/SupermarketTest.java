package com.testcoders;

import org.junit.Assert;
import org.junit.Test;

import java.util.List;

public class SupermarketTest {
    List<Product> products;

    @Test
    public void addProductTest() {
        Basket basket = new Basket(products);
        products.add(new Product("name", 12345, 114.44, 20));
        Assert.assertEquals(products.size(), 1);
    }
    //addProductsTest

    //sumProductsInBasketTest

    //productCategoriesTest
    //Name, discount

    //
}
