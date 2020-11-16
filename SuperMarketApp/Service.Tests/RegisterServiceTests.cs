using Moq;
using NUnit.Framework;
using Service.Interfaces;
using Service.Models;
using Service.Services;

namespace Service.Tests
{
    public class RegisterServiceTests
    {
        private Mock<ICalculateCartPrice> _calculatePriceServiceMock;
        private Mock<IRegisterService> _registerServiceMock;

        private Cart _cart;

        [SetUp]
        public void Setup()
        {
            _calculatePriceServiceMock = new Mock<ICalculateCartPrice>();
            _registerServiceMock = new Mock<IRegisterService>();
            
            _cart = new Cart();
            _cart.AddToCart(new Product("Kaas", 156734, 4.99));
            _cart.AddToCart(new Product("Ham", 579843, 1.49));
            _cart.AddToCart(new Product("Melk", 378941, 0.99));
            _cart.AddToCart(new Product("Pizza", 739214, 4.59));
            _cart.AddToCart(new Product("WC papier", 798234, 1.12));
        }

        [Test]
        public void ShouldPass_WhenCheckOut()
        {
            // Assemble
            _calculatePriceServiceMock.Setup(mock => mock.Calculate(_cart));
            var registerService = new RegisterService(_calculatePriceServiceMock.Object);

            // Act
            registerService.CheckOut(_cart);

            // Assert
            _calculatePriceServiceMock.Verify(mock => mock.Calculate(_cart), Times.Once);
        }
    }
}