package com.testcoders;

import java.util.Date;
import java.util.List;

public class Invoice {
    private Product product;
    private double total;
    private Date date;

    public Invoice(Product product, double total, Date date) {
        this.product = product;
        this.total = total;
        this.date = date;
    }

//    public void generateInvoice(List<Product>) {
//
//    }
}
