using Service.Models;
using System.Net.Http;
using System.Threading.Tasks;

namespace Service.Interfaces
{
    public interface IProvisioningClient
    {
        Task<HttpResponseMessage> SendProvisioningRequest(ProvisioningRequest request);
    }
}
