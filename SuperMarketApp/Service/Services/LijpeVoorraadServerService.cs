using Service.Clients;
using Service.Interfaces;
using Service.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
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

        public async Task PostProvisioning(IProvisioningClient client, ProvisioningRequest request)
        {
            var result = await client.SendProvisioningRequest(request);
        }

        public ProvisioningRequest CreateProvisionRequest(int provisionMax)
        {
            var provisioningRequest = new ProvisioningRequest
            {
                ProvisionProducts = new List<ProvisioningProduct>()
            };

            var provisioningProducts = _productService.GetProvisionProducts(100);

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
