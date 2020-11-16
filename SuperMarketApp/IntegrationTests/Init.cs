using Microsoft.Extensions.DependencyInjection;
using NUnit.Framework;
using Service.Interfaces;
using Service.Services;

namespace Service.IntegrationTests
{
    public class Init
    {
        protected ICalculateProductPrice CalculateProductPrice;
        protected ICalculateCartPrice CalculateCartPrice;
        protected IRegisterService RegisterService;

        [SetUp]
        public void SetUp()
        {
            var serviceCollection = new ServiceCollection();
            serviceCollection.AddScoped<ICalculateProductPrice, CalculateProductPrice>();
            serviceCollection.AddScoped<ICalculateCartPrice, CalculateCartPrice>();
            serviceCollection.AddScoped<IRegisterService, RegisterService>();
            var serviceProvider = serviceCollection.BuildServiceProvider();

            CalculateProductPrice = serviceProvider.GetService<ICalculateProductPrice>();
            CalculateCartPrice = serviceProvider.GetService<ICalculateCartPrice>();
            RegisterService = serviceProvider.GetService<IRegisterService>();
        }
    }
}
