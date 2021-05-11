using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using SuperMarketApp.Service.Enum;
using SuperMarketApp.Service.Models;

namespace SuperMarketApp.Repositories.Context
{
    public class ProductContext : DbContext 
    {
        public virtual DbSet<Product> Product { get; set; }

        public ProductContext(DbContextOptions options) : base(options)
        {
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Product>().Property(p => p.Discount).
                HasConversion(new EnumToStringConverter<Discount>());
        }
    }
}
