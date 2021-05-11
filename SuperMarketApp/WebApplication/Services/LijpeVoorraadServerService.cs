using Service.Interfaces;
using Service.Models;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace Service.Services
{
    public class LijpeVoorraadServerService : ILijpeVoorraadServerService
    {
        private readonly IProductService _productService;

        public LijpeVoorraadServerService(IProductService productService)
        {
            _productService = productService;
        }

        public async Task<HttpResponseMessage> PostProvisioning(IProvisioningClient client, ProvisioningRequest request)
        {
            return await client.SendProvisioningRequest(request);
        }

        public async Task<int> PostSupply(ProvisioningRequest request)
        {
            int rowsAffected = 0;

            foreach (var product in request.ProvisionProducts)
            {
                rowsAffected += await _productService.IncreaseProductAmount(product.Barcode, product.Amount);
            }
            return rowsAffected;
        }

        public async Task<ProvisioningRequest> CreateProvisionRequest(int provisionMax)
        {
            var provisioningRequest = new ProvisioningRequest
            {
                ProvisionProducts = new List<ProvisioningProduct>()
            };

            var provisioningProducts = await _productService.GetProvisionProducts(100);

            foreach (var product in provisioningProducts)
            {
                provisioningRequest.ProvisionProducts.Add(new ProvisioningProduct
                {
                    Barcode = product.Barcode,
                    Amount = provisionMax - product.Amount
                });
            }
            return provisioningRequest;
        }
    }
}
