using Service.Models;

namespace Service.Interfaces
{
    public interface ICalculateCartPrice
    {
        double Calculate(Cart cart);
    }
}
