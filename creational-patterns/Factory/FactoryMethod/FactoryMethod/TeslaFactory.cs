namespace FactoryMethod
{
    public class TeslaFactory : ICarFactory
    {
        public ICar GetCar()
        {
            return new Tesla();
        }
    }
}