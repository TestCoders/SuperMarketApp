using Moq;
using NUnit.Framework;
using Service.Interfaces;
using Service.Models;
using Service.Services;

namespace Service.Tests
{
    public class CalculateCartPriceTests
    {
        private Mock<ICalculateProductPrice> _calculateProductPriceMock;
        private Cart _cart;

        [SetUp]
        public void Setup()
        {
            _calculateProductPriceMock = new Mock<ICalculateProductPrice>();

            _cart = new Cart();
            _cart.AddToCart(new Product("Kaas", 156734, 4.99));
            _cart.AddToCart(new Product("Ham", 579843, 1.49));
            _cart.AddToCart(new Product("Melk", 378941, 0.99));
            _cart.AddToCart(new Product("Pizza", 739214, 4.59));
            _cart.AddToCart(new Product("WC papier", 798234, 1.12));
        }

        [Test]
        public void CalculatePrice_ShouldReturnCorrectDouble_WhenGivenCartNoDiscount()
        {
            // Assemble
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[0])).Returns(_cart.Products[0].Price);
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[1])).Returns(_cart.Products[1].Price);
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[2])).Returns(_cart.Products[2].Price);
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[3])).Returns(_cart.Products[3].Price);
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[4])).Returns(_cart.Products[4].Price);

            var calculatePriceService = new CalculateCartPrice(_calculateProductPriceMock.Object);

            // Assign
            var expectedPrice = 13.18;

            // Act
            var price = calculatePriceService.Calculate(_cart);

            // Assert
            Assert.AreEqual(expectedPrice, price);
        }

        [Test]
        public void CalculatePrice_ShouldRoundUp_WhenGivenThreeDecimals()
        {
            // Assemble
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[0])).Returns(5.495);
            var calculatePriceService = new CalculateCartPrice(_calculateProductPriceMock.Object);

            // Assign
            var expectedPrice = 5.50;

            // Act
            var price = calculatePriceService.Calculate(_cart);

            // Assert
            Assert.AreEqual(expectedPrice, price);
        }

        [Test]
        public void CalculatePrice_ShouldRoundDown_WhenGivenThreeDecimals()
        {
            // Assemble
            _calculateProductPriceMock.Setup(mock => mock.Calculate(_cart.Products[0])).Returns(5.494);
            var calculatePriceService = new CalculateCartPrice(_calculateProductPriceMock.Object);

            // Assign
            var expectedPrice = 5.49;

            // Act
            var price = calculatePriceService.Calculate(_cart);

            // Assert
            Assert.AreEqual(expectedPrice, price);
        }
    }
}
