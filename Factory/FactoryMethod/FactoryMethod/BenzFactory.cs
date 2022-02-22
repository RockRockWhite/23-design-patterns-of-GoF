namespace FactoryMethod
{
    public class BenzFactory : ICarFactory
    {
        public ICar GetCar()
        {
            return new Benz();
        }
    }
}