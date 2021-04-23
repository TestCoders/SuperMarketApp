using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using SuperMarketApp.Repositories.Enum;
using SuperMarketApp.Repositories.Models;
using System;
using System.Linq;

namespace SuperMarketApp.Repositories.Context
{
    public class ProductContext : DbContext 
    {
        public virtual DbSet<ProductDB> Product { get; set; }

        private static readonly string _connectionString = @"Data Source=PQTE224\SQLEXPRESS;Initial Catalog=SuperMarketApp.Database;Integrated Security=True;Connect Timeout=30;Encrypt=False;TrustServerCertificate=False;ApplicationIntent=ReadWrite;MultiSubnetFailover=False";

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            optionsBuilder.UseSqlServer(_connectionString);
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<ProductDB>().Property(p => p.Discount).
                HasConversion(new EnumToStringConverter<Discount>());
        }
    }
}
