using Service.Models;
using System.Threading.Tasks;

namespace Service.Interfaces
{
    public interface ILijpeVoorraadServerService
    {
        Task PostProvisioning(IProvisioningClient client, ProvisioningRequest request);
        ProvisioningRequest CreateProvisionRequest(int provisionMax);
    }
}
