using Microsoft.Extensions.DependencyInjection;
using NUnit.Framework;
using Service.Interfaces;
using Service.Services;
using SuperMarketApp.Repositories.Context;

namespace Service.IntegrationTests
{
    public class Init
    {
        protected ICalculateProductPrice CalculateProductPrice;
        protected ICalculateCartPrice CalculateCartPrice;
        protected IRegisterService RegisterService;
        protected IProductService ProductService;

        [SetUp]
        public void SetUp()
        {
            var serviceCollection = new ServiceCollection();
            serviceCollection.AddDbContext<ProductContext>();
            serviceCollection.AddScoped<ICalculateProductPrice, CalculateProductPrice>();
            serviceCollection.AddScoped<ICalculateCartPrice, CalculateCartPrice>();
            serviceCollection.AddScoped<IRegisterService, RegisterService>();
            serviceCollection.AddScoped<IProductService, ProductService>();
            var serviceProvider = serviceCollection.BuildServiceProvider();

            CalculateProductPrice = serviceProvider.GetService<ICalculateProductPrice>();
            CalculateCartPrice = serviceProvider.GetService<ICalculateCartPrice>();
            RegisterService = serviceProvider.GetService<IRegisterService>();
            ProductService = serviceProvider.GetService<IProductService>();
        }
    }
}
