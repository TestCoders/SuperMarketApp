using Service.Interfaces;
using Service.Models;
using System;

namespace Service.Services
{
    public class CalculateCartPrice : ICalculateCartPrice
    {
        private readonly ICalculateProductPrice _calculateProductPrice;

        public CalculateCartPrice(ICalculateProductPrice calculateProductPrice)
        {
            _calculateProductPrice = calculateProductPrice;
        }

        public double Calculate(Cart cart)
        {
            double totalPrice = new double();

            foreach (var product in cart.Products)
            {
                totalPrice += _calculateProductPrice.Calculate(product);
            }

            totalPrice = Math.Round(totalPrice, 2);
            return totalPrice;
        }
    }
}
