using Service.Interfaces;
using SuperMarketApp.Repositories.Context;
using SuperMarketApp.Repositories.Models;
using System;
using System.Linq;

namespace Service.Services
{
    public class ProductService : IProductService
    {
        private ProductContext _context;

        public ProductService(ProductContext context)
        {
            _context = context;
        }

        public int DecreaseProductAmount(int barcode, int amount)
        {
            throw new NotImplementedException("Is yet to be build");
        }

        public ProductDB GetProduct(int barcode)
        {
            return _context.Product.First(x => x.Barcode == barcode);
        }

        public int InsertProduct(ProductDB product)
        {
            _context.Product.Add(product);
            var rowsAffected = _context.SaveChanges();
            return rowsAffected;
        }
    }
}
