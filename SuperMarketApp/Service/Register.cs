using Microsoft.Extensions.DependencyInjection;
using Service.Interfaces;
using Service.Services;

namespace Service
{
    public class Register
    {
        public void ConfigureServices(IServiceCollection serviceColletion)
        {
            serviceColletion.AddScoped<ICalculateProductPrice, CalculateProductPrice>();
            serviceColletion.AddScoped<ICalculateCartPrice, CalculateCartPrice>();
            serviceColletion.AddScoped<IRegisterService, RegisterService>();
            serviceColletion.AddScoped<IProductService, ProductService>();
            serviceColletion.AddScoped<ILijpeVoorraadServerService, LijpeVoorraadServerService>();
        }
    }
}
