using Service.Models;

namespace Service.Interfaces
{
    public interface ICalculateProductPrice
    {
        double Calculate(Product product);
    }
}
