package Builder

type Worker struct {
    house House
}

func (worker *Worker) BuildWindow(name string) {
    worker.house.Window = name
}

func (worker *Worker) BuildDoor(name string) {
    worker.house.Door = name
}

func (worker *Worker) BuildFloor(name string) {
    worker.house.Floor = name
}

func (worker *Worker) GetHouse() *House {
    return &worker.house
}
