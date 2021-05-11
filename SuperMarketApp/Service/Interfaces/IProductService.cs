using SuperMarketApp.Repositories.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Service.Interfaces
{
    public interface IProductService
    {
        Task<int> DecreaseProductAmount(int barcode, int amount);
        Task<int> DeleteProduct(int barcode);
        Task<ProductDB> GetProduct(int barcode);
        Task<IEnumerable<ProductDB>> GetProvisionProducts(int provisionMax);
        Task<int> InsertProduct(ProductDB product);
        Task<int> IncreaseProductAmount(int barcode, int amount);
    }
}
