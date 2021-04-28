using Service.Interfaces;
using SuperMarketApp.Repositories.Context;
using SuperMarketApp.Repositories.Models;
using System;
using System.Linq;

namespace Service.Services
{
    public class ProductService : IProductService
    {
        private readonly ProductContext _context;

        public ProductService(ProductContext context)
        {
            _context = context;
        }

        public int DecreaseProductAmount(int barcode, int amount)
        {
            var product = GetProduct(barcode);
            product.Amount -= amount;
            var rowsAffected = _context.SaveChanges();
            return rowsAffected;
        }

        public int DeleteProduct(int barcode)
        {
            var productToDelete = GetProduct(barcode);
            _context.Product.Remove(productToDelete);
            var rowsAffected = _context.SaveChanges();
            return rowsAffected;
        }

        public ProductDB GetProduct(int barcode)
        {
            return _context.Product.FirstOrDefault(x => x.Barcode == barcode);
        }

        public int InsertProduct(ProductDB product)
        {
            _context.Product.Add(product);
            var rowsAffected = _context.SaveChanges();
            return rowsAffected;
        }
    }
}
