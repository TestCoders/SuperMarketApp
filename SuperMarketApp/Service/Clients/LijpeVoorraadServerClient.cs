using Newtonsoft.Json;
using Service.Interfaces;
using Service.Models;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace Service.Clients
{
    public class LijpeVoorraadServerClient : IProvisioningClient
    {
        public async Task<HttpResponseMessage> SendProvisioningRequest(ProvisioningRequest request)
        {
            var json = JsonConvert.SerializeObject(request);
            var body = new StringContent(json, Encoding.UTF8, "application/json");

            var result = await new HttpClient().PostAsync("https://reqbin.com/echo/post/json", body);

            return result;
        }
    }
}
