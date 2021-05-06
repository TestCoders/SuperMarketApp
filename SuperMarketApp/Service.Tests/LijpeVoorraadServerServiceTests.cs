using Moq;
using NUnit.Framework;
using Service.Interfaces;
using Service.Models;
using Service.Services;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;

namespace Service.Tests
{
    public class LijpeVoorraadServerServiceTests
    {
        private ILijpeVoorraadServerService _service;
        private Mock<IProvisioningClient> _mockClient;
        private Mock<IProductService> _mockProductService;
        private ProvisioningRequest _provisioningRequest;

        [SetUp]
        public void Init()
        {
        }

        [Test]
        public async Task PostProvisioning_ShouldPass()
        {
            // Assign
            var mockResponseMessage = new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.OK
            };

            // Assemble
            _provisioningRequest = new ProvisioningRequest
            {
                ProvisionProducts = new List<ProvisioningProduct>()
            };
            _provisioningRequest.ProvisionProducts.Add(new ProvisioningProduct { Amount = 5, Barcode = 123 });
            _provisioningRequest.ProvisionProducts.Add(new ProvisioningProduct { Amount = 12, Barcode = 1834 });

            _mockProductService = new Mock<IProductService>();
            _mockClient = new Mock<IProvisioningClient>();
            _mockClient.Setup(m => m.SendProvisioningRequest(_provisioningRequest)).Returns(Task.FromResult(mockResponseMessage));
            
            _service = new LijpeVoorraadServerService(_mockProductService.Object);

            // Act
            await _service.PostProvisioning(_mockClient.Object, _provisioningRequest);

            // Assert
            _mockClient.Verify(m => m.SendProvisioningRequest(_provisioningRequest), Times.Once);
        }
    }
}
