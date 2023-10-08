package gogeneratorme

func GenerateAll(model interface{}) {
	CreateRepository("client", model)
}
