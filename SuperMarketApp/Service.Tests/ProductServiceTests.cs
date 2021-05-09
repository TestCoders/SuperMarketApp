﻿using Microsoft.EntityFrameworkCore;
using Moq;
using NUnit.Framework;
using Service.Interfaces;
using Service.Services;
using SuperMarketApp.Repositories.Context;
using SuperMarketApp.Repositories.Enum;
using SuperMarketApp.Repositories.Models;

namespace Service.Tests
{
    public class ProductServiceTests
    {
        private Mock<DbSet<ProductDB>> _mockSet;
        private Mock<ProductContext> _mockContext;
        private IProductService _productService;

        [SetUp]
        public void Init()
        {
            _mockSet = new Mock<DbSet<ProductDB>>();
            _mockContext = new Mock<ProductContext>();
        }

        //[Test]
        //public void AddProductTest_ShouldReturnOneRowAffected()
        //{
        //    // Assemble
        //    var product = new ProductDB
        //    {
        //        ProductName = "Test product",
        //        Barcode = 12345,
        //        Price = 1.99M,
        //        Discount = Discount.NoDiscount,
        //        Amount = 1
        //    };

        //    _mockContext.Setup(m => m.Product).Returns(() =>_mockSet.Object);
        //    _productService = new ProductService(_mockContext.Object);

        //    // Act
        //    _productService.InsertProduct(product);

        //    // Assert
        //    _mockSet.Verify(m => m.Add(It.IsAny<ProductDB>()), Times.Once);
        //    _mockContext.Verify(m => m.SaveChanges(), Times.Once);
        //}
    }
}
