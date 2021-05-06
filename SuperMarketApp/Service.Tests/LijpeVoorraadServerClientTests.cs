using Moq;
using Moq.Protected;
using NUnit.Framework;
using Service.Clients;
using Service.Models;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Threading;
using System.Threading.Tasks;

namespace Service.Tests
{
    public class LijpeVoorraadServerClientTests
    {
        [Test]
        public async Task SendProvisioningRequestTest_ShouldReturnStatus200()
        {
            // Assign
            var expectedStatusCode = HttpStatusCode.OK;

            // Assemble
            var provisioningRequest = new ProvisioningRequest
            {
                ProvisionProducts = new List<ProvisioningProduct>()
            };
            provisioningRequest.ProvisionProducts.Add(new ProvisioningProduct { Amount = 5, Barcode = 123 });
            provisioningRequest.ProvisionProducts.Add(new ProvisioningProduct { Amount = 12, Barcode = 1834 });

            var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
            mockHttpMessageHandler.Protected().Setup<Task<HttpResponseMessage>>("SendAsync", ItExpr.IsAny<HttpRequestMessage>(), ItExpr.IsAny<CancellationToken>()).
                ReturnsAsync(new HttpResponseMessage { StatusCode = HttpStatusCode.OK });

            var httpClient = new HttpClient(mockHttpMessageHandler.Object);
            var lijpeVoorraadServerClient = new LijpeVoorraadServerClient(httpClient);

            // Act
            var result = await lijpeVoorraadServerClient.SendProvisioningRequest(provisioningRequest);

            // Assert
            Assert.AreEqual(expectedStatusCode, result.StatusCode);
            mockHttpMessageHandler.Protected().Verify("SendAsync", Times.Exactly(1), ItExpr.IsAny<HttpRequestMessage>(), ItExpr.IsAny<CancellationToken>());
        }
    }
}
