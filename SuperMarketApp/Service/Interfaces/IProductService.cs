using SuperMarketApp.Repositories.Models;
using System.Collections.Generic;

namespace Service.Interfaces
{
    public interface IProductService
    {
        int DecreaseProductAmount(int barcode, int amount);
        int DeleteProduct(int barcode);
        ProductDB GetProduct(int barcode);
        IEnumerable<ProductDB> GetProvisionProducts(int provisionMax);
        int InsertProduct(ProductDB product);
        int IncreaseProductAmount(int barcode, int amount);
    }
}
