package Builder

type Director struct {
}

func (Director) Build(builder IBuilder) *House {
    builder.BuildFloor("Floor from Command")
    builder.BuildWindow("Window from Command")
    builder.BuildDoor("Door from Command")

    return builder.GetHouse()
}
