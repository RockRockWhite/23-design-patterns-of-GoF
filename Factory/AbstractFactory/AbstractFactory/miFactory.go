package AbstractFactory

type MiFactory struct {
}

func (MiFactory) GetPhone() IPhone {
    return &MiPhone{}
}

func (MiFactory) GetRouter() IRouter {
    return &MiRouter{}
}
