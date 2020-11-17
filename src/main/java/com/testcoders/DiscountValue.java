package com.testcoders;

public class DiscountValue {
    private Discount discount;
    private int discountPercentage;

    public DiscountValue() {
    }

    public int getDiscountValue(Discount discount) {
        if (discount.equals(Discount.OutOfDate)) {
            discountPercentage = 35;
        } if (discount.equals(Discount.BonusDiscount))
            discountPercentage = 20;
        else discountPercentage = 0;

        return discountPercentage;
    }
}
