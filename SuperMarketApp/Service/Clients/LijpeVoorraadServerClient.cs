using Newtonsoft.Json;
using Service.Interfaces;
using Service.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace Service.Clients
{
    public class LijpeVoorraadServerClient : IProvisioningClient
    {
        private readonly HttpClient _client;

        public LijpeVoorraadServerClient(HttpClient client)
        {
            _client = client;
        }

        public async Task<HttpResponseMessage> SendProvisioningRequest(ProvisioningRequest request)
        {
            var json = JsonConvert.SerializeObject(request);
            var body = new StringContent(json, Encoding.UTF8, "application/json");

            var result = await _client.PostAsync("https://reqbin.com/echo/post/json", body);

            return result;
        }
    }
}
