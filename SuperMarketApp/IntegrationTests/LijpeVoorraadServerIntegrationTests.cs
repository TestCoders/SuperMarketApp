using NUnit.Framework;
using System;
using System.Collections.Generic;
using System.Text;

namespace Service.IntegrationTests
{
    public class LijpeVoorraadServerIntegrationTests : Init
    {
        [Test]
        public void GetProvisionProductsTest_ShouldReturnAtLeastTwoProducts()
        {
            // Act
            var provisionProducts = LijpeVoorraadServerService.CreateProvisionRequest(100);

            // Assert
            Assert.That(provisionProducts.ProvisionProducts.Count >= 2);
        }
    }
}
