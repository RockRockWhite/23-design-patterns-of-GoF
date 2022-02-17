package AbstractFactory

type IFactory interface {
    GetPhone() IPhone
    GetRouter() IRouter
}
