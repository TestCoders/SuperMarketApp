using Service.Enum;

namespace Service.Models
{
    public class Product
    {
        public string Name { get; set; }
        public int Barcode { get; set; }
        public double Price { get; set; }
        public Discount Discount { get; set; }

        public Product(string name, int barcode, double price, Discount discount = Discount.NoDiscount)
        {
            Name = name;
            Barcode = barcode;
            Price = price;
            Discount = discount;
        }
    }
}
