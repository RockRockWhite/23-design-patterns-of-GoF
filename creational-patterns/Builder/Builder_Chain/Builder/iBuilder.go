package Builder

type IBuilder interface {
    BuildWindow(string) *Worker
    BuildDoor(string) *Worker
    BuildFloor(string) *Worker
    GetHouse() *House
}
