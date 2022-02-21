package Builder

type IBuilder interface {
    BuildWindow(string)
    BuildDoor(string)
    BuildFloor(string)
    GetHouse() *House
}
