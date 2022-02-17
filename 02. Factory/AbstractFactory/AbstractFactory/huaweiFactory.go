package AbstractFactory

type HuaweiFactory struct {
}

func (HuaweiFactory) GetPhone() IPhone {
    return &HuaweiPhone{}

}
func (HuaweiFactory) GetRouter() IRouter {
    return &HuaweiRouter{}
}
