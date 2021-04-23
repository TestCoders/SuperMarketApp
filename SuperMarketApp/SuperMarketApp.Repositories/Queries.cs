namespace SuperMarketApp.Repositories
{
    public static class Queries
    {
        public static string GetAllProducts() => "SELECT * FROM dbo.Product";
        public static string GetProduct(int barcode) => $"SELECT * FROM dbo.Product WHERE Barcode = {barcode}";
        public static string Update(int barcode, int newAmount) => $"UPDATE dbo.Product SET Amount = {newAmount} WHERE Barcode = {barcode}";
    }
}
