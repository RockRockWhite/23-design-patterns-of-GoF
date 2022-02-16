using FactoryMethod;

var car1 = new BenzFactory().GetCar();
var car2 = new TeslaFactory().GetCar();

car1.Name();
car2.Name();