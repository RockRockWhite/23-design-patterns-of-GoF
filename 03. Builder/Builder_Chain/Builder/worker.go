package Builder

type Worker struct {
    house House
}

func (worker *Worker) BuildWindow(name string) *Worker {
    worker.house.Window = name
    return worker
}

func (worker *Worker) BuildDoor(name string) *Worker {
    worker.house.Door = name
    return worker
}

func (worker *Worker) BuildFloor(name string) *Worker {
    worker.house.Floor = name
    return worker
}

func (worker *Worker) GetHouse() *House {
    return &worker.house
}
