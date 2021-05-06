using Service.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace Service.Interfaces
{
    public interface IProvisioningClient
    {
        Task<HttpResponseMessage> SendProvisioningRequest(ProvisioningRequest request);
    }
}
