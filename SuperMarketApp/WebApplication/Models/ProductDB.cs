using SuperMarketApp.Service.Enum;

namespace SuperMarketApp.Service.Models
{
    public class ProductDB
    {
        public int Id { get; set; }
        public string ProductName { get; set; }
        public int Barcode { get; set; }
        public decimal Price { get; set; }
        public Discount Discount { get; set; }
        public int Amount { get; set; }
    }
}
