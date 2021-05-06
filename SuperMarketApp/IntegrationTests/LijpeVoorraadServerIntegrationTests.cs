using NUnit.Framework;
using Service.Clients;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;

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

        [Test]
        public async Task PostProvisionRequest_ShouldPass()
        {
            // Assign
            var expectedStatusCode = HttpStatusCode.OK;

            // Assemble
            var provisionRequest = LijpeVoorraadServerService.CreateProvisionRequest(100);
            var provisioningClient = new LijpeVoorraadServerClient(new HttpClient());

            // Act
            var result = await LijpeVoorraadServerService.PostProvisioning(provisioningClient, provisionRequest);

            // Assert
            Assert.AreEqual(expectedStatusCode, result.StatusCode);
        }
    }
}
