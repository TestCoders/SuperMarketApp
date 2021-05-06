using NUnit.Framework;
using Service.IntegrationTests;
using SuperMarketApp.Repositories.Enum;
using SuperMarketApp.Repositories.Models;
using System;

namespace SuperMarketApp.Service.IntegrationTests
{
    public class ProductServiceIntegrationTests : Init
    {
        [Test]
        public void DecreaseProductAmountByOne_ShouldReturnOneRowAffected()
        {
            // Assign
            var expectedRowsAffected = 1;

            // Assemble
            var barcode = 156734;

            // Act
            var rowsAffected = ProductService.DecreaseProductAmount(barcode, 1);

            // Assert
            Assert.AreEqual(expectedRowsAffected, rowsAffected);
        }

        [Test]
        public void DeleteProductTest_ShouldReturnOneRowAffected()
        {
            // Assign
            var barcode = 54752848;

            // Assemble
            var product = new ProductDB
            {
                Amount = 100,
                Barcode = barcode,
                Discount = Discount.NoDiscount,
                Price = 1.59M,
                ProductName = "Should not be visible in DB"
            };

            ProductService.InsertProduct(product);

            // Act
            var rowsAffected = ProductService.DeleteProduct(barcode);

            // Assert
            Assert.AreEqual(1, rowsAffected);
        }

        [Test]
        public void GetProductTest_ShouldReturnOneProduct()
        {
            var product = ProductService.GetProduct(156734);
            Assert.AreEqual("Kaas", product.ProductName);
        }

        [Test]
        public void InsertProductTest_ShouldReturnOneRowAffectedWhenProductAdded()
        {
            // Assign 
            var expectedRowsAffected = 1;

            // Assemble
            var barcode = 23564234;

            var product = new ProductDB
            {
                ProductName = "Bananen",
                Barcode = barcode,
                Price = 1.99M,
                Discount = Discount.Bonus,
                Amount = 100
            };

            var actualRowsAffected = ProductService.InsertProduct(product);

            Assert.AreEqual(expectedRowsAffected, actualRowsAffected);

            // Clean up
            CleanUp(barcode);
        }

        private void CleanUp(int barcode)
        {
            try
            {
                ProductService.DeleteProduct(barcode);
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error during clean up! Product with barcode {barcode} is NOT deleted");
                Console.WriteLine(ex.Message);
                Console.WriteLine(ex.StackTrace);
                Console.WriteLine(ex.InnerException);
            }
        }
    }
}