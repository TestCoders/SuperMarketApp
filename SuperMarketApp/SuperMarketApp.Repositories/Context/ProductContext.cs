using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using SuperMarketApp.Repositories.Enum;
using SuperMarketApp.Repositories.Models;

namespace SuperMarketApp.Repositories.Context
{
    public class ProductContext : DbContext 
    {
        public virtual DbSet<ProductDB> Product { get; set; }

        public ProductContext(DbContextOptions options) : base(options)
        {
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<ProductDB>().Property(p => p.Discount).
                HasConversion(new EnumToStringConverter<Discount>());
        }
    }
}
