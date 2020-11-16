package com.testcoders;

import org.junit.Assert;
import org.junit.Test;

import static java.lang.Float.parseFloat;

public class SupermarketTest {

    @Test
    public void addProductTest() {
        Basket basket = new Basket();
        Product product = new Product("name", 12345, 114.44, new ProductCategory("Ham"));
        Assert.assertEquals(basket.getProducts().size(), 0);

        basket.addProduct(product);
        Assert.assertEquals(basket.getProducts().size(), 1);
    }

    @Test
    public void getDiscountPriceTest() {
        Product product = new Product("name", 12345, 114.44, new ProductCategory("Milk"));

        Assert.assertEquals(product.getPrice(), 74.39,0);
    }

    @Test
    public void roundingTest(){
        Assert.assertEquals( "double : " + String.format("%.2f", 64.849999)), 64.85, 64.85, 0);
        Assert.assertEquals( Math.round(74.850000 * 100) / 100, 74.85, 0);
        Assert.assertEquals( Math.round(84.8450000 * 100) / 100, 84.84, 0);
    }

    @Test
    public void totalTest() {
        Basket basket = new Basket();
        Product product = new Product("name", 12345, 114.44,  new ProductCategory("Fiets"));
        Product product2 = new Product("name", 12345, 114.44,  new ProductCategory("Beer"));

        basket.addProduct(product);
        basket.addProduct(product2);
        Assert.assertEquals(basket.getProducts().size(), 2);
        Assert.assertEquals(basket.getTotal(),205.99, 0);
    }
}
