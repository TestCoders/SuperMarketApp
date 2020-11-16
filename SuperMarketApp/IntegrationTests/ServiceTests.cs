using NUnit.Framework;
using Service.Enum;
using Service.Models;

namespace Service.IntegrationTests
{
    public class ServiceTests : Init
    {
        private Cart _cart;

        [Test]
        public void CalculateCart_NoDiscount()
        {
            // Assemble
            _cart = new Cart();
            _cart.AddToCart(new Product("Kaas", 156734, 4.99));
            _cart.AddToCart(new Product("Ham", 579843, 1.49));
            _cart.AddToCart(new Product("Melk", 378941, 0.99));
            _cart.AddToCart(new Product("Pizza", 739214, 4.59));
            _cart.AddToCart(new Product("WC papier", 798234, 1.12));

            // Act
            var price = CalculateCartPrice.Calculate(_cart);

            // Assert
            Assert.AreEqual(13.18, price);
        }

        [Test]
        public void CalculateCart_WithBonusAndExpiryDiscount()
        {
            _cart = new Cart();
            _cart.AddToCart(new Product("Kaas", 156734, 4.99, Discount.Bonus));
            _cart.AddToCart(new Product("Ham", 579843, 1.49));
            _cart.AddToCart(new Product("Melk", 378941, 0.99, Discount.Expiry));
            _cart.AddToCart(new Product("Pizza", 739214, 4.59));
            _cart.AddToCart(new Product("WC papier", 798234, 1.12, Discount.Bonus));

            // Act
            var price = CalculateCartPrice.Calculate(_cart);

            // Assert
            Assert.AreEqual(11.61, price);
        }

        [Test]
        public void CheckOut_ShouldPrintCorrectReceipt()
        {
            // Assemble
            _cart = new Cart();
            _cart.AddToCart(new Product("Kaas", 156734, 4.99));
            _cart.AddToCart(new Product("Ham", 579843, 1.49));
            _cart.AddToCart(new Product("Melk", 378941, 0.99));
            _cart.AddToCart(new Product("Pizza", 739214, 4.59));
            _cart.AddToCart(new Product("WC papier", 798234, 1.12));

            // Assign
            var expectedPrice = 13.18;

            // Act
            var receipt = RegisterService.CheckOut(_cart);

            // Assert
            Assert.That(receipt.Contains(expectedPrice.ToString()));
        }
    }
}