using Service.Models;
using System.Net.Http;
using System.Threading.Tasks;

namespace Service.Interfaces
{
    public interface ILijpeVoorraadServerService
    {
        Task<HttpResponseMessage> PostProvisioning(IProvisioningClient client, ProvisioningRequest request);
        Task <int> PostSupply(ProvisioningRequest request);
        Task<ProvisioningRequest> CreateProvisionRequest(int provisionMax);
    }
}
