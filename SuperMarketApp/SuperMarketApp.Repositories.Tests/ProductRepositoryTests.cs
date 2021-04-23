using Microsoft.EntityFrameworkCore;
using Moq;
using NUnit.Framework;
using SuperMarketApp.Repositories.Context;
using SuperMarketApp.Repositories.Models;

namespace SuperMarketApp.Repositories.Tests
{
    public class ProductRepositoryTests
    {
        [Test]
        public void InsertProduct_ShouldReturnOneRowAffected()
        {
            var mockSet = new Mock<DbSet<ProductDB>>();
            var mockContext = new Mock<ProductContext>();

            mockContext.Setup(m => m.Product).Returns(mockSet.Object);
            //mockContext.InsertProduct

            // https://docs.microsoft.com/en-us/ef/ef6/fundamentals/testing/mocking
        }
    }
}
