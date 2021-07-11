package clerk

import "go.mongodb.org/mongo-driver/bson"

type mongoUpdateCommand struct {
	Command // Interface

	collection Collection
	filter bson.D
	entity interface{}
}

func NewMongoUpdateCommand(collection Collection, entity interface{}) *mongoUpdateCommand {
	command := new(mongoUpdateCommand)

	command.collection = collection
	command.entity = entity

	return command
}

func (command mongoUpdateCommand) Where(key string, value interface{}) mongoUpdateCommand {
	command.filter = append(command.filter, bson.E{Key: key, Value: value})
	return command
}

func (command mongoUpdateCommand) GetCollection() Collection {
	return command.collection
}

func (command mongoUpdateCommand) Handle(handler CommandHandler) error {
	return handler.Update(command.filter, command.entity)
}
