using NUnit.Framework;
using Service.Interfaces;
using Service.Services;
using SuperMarketApp.Repositories.Context;
using SuperMarketApp.Repositories.Enum;
using SuperMarketApp.Repositories.Models;
using System;

namespace SuperMarketApp.Service.IntegrationTests
{
    public class ProductServiceIntegrationTests
    {
        private IProductService _productService;
        private ProductContext _productContext;

        [SetUp]
        public void Setup()
        {
            _productContext = new ProductContext();
            _productService = new ProductService(_productContext);
        }

        [Test]
        public void GetProductTest_ShouldReturnOneProduct()
        {
            var product = _productService.GetProduct(156734);
            Assert.AreEqual("Kaas", product.ProductName);
        }

        [Test]
        public void InsertProductTest_ShouldThrowIfProductExists()
        {
            // Assign
            Exception ex = new Exception();

            // Assemble
            var product = new ProductDB
            {
                ProductName = "Haribo snoepies",
                Barcode = 81681385,
                Price = 1.49M,
                Discount = Discount.NoDiscount,
                Amount = 100
            };

            // Act
            try
            {
                var rowsAffected = _productService.InsertProduct(product);
            }
            catch (Exception e)
            {
                ex = e;
            }

            // Assert
            Assert.That(ex is ArgumentException, $"No Exception was thrown. Please check if {product.Barcode} was inserted");
        }

        [Test]
        public void InsertProductTest_ShouldReturnOneRowAffectedWhenProductAdded()
        {

        }

        private void CleanUp(ProductDB product)
        {

        }
    }
}