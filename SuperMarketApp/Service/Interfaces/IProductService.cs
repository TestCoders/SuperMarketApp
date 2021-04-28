using SuperMarketApp.Repositories.Models;

namespace Service.Interfaces
{
    public interface IProductService
    {
        int DecreaseProductAmount(int barcode, int amount);
        int DeleteProduct(int barcode);
        ProductDB GetProduct(int barcode);
        int InsertProduct(ProductDB product);
    }
}
