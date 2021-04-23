using NUnit.Framework;
using Service.Interfaces;
using Service.Models;
using Service.Services;
using SuperMarketApp.Repositories.Enum;
using System;

namespace Service.Tests
{
    public class CalculateProductPriceTests
    {
        private ICalculateProductPrice _calculateProductPrice;

        [SetUp]
        public void SetUp()
        {
            _calculateProductPrice = new CalculateProductPrice();
        }

        [Test]
        public void CalculateProductPrice_ShouldReturnProductPrice_WhenNoDiscount()
        {
            // Assemble
            var product = new Product("Kaaaas", 123, 5.49, Discount.NoDiscount);

            // Act
            var price = _calculateProductPrice.Calculate(product);

            // Assert
            Assert.AreEqual(product.Price, price);
        }

        [Test]
        public void CalculateProductPrice_ShouldReturnPriceWithBonusDiscount_WhenGivenBonusDiscount()
        {
            // Assemble 
            var product = new Product("Kaaaas", 123, 5.49, Discount.Bonus);

            // Assign
            var expectedPrice = product.Price * Constants.BonusDiscount;

            // Act
            var price = _calculateProductPrice.Calculate(product);

            // Assert
            Assert.AreEqual(expectedPrice, price);
        }

        [Test]
        public void CalculateProductPrice_ShouldReturnPriceWithExpiryDiscount_WhenGivenExpiryDiscount()
        {
            // Assemble 
            var product = new Product("Kaaaas", 123, 5.49, Discount.Expiry);

            // Assign
            var expectedPrice = product.Price * Constants.ExpiryDiscount;

            // Act
            var price = _calculateProductPrice.Calculate(product);

            // Assert
            Assert.AreEqual(expectedPrice, price);
        }

        [Test]
        public void CalculateProductPrice_ShouldThrowNullReferenceExpcetion_WhenGivenNullProduct()
        {
            Exception ex = Assert.Throws<NullReferenceException>(delegate
            {
                _calculateProductPrice.Calculate(null);
            });
            Assert.AreEqual("given product is null!", ex.Message.ToLower());
        }
    }
}
