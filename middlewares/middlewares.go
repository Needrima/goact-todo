package middlewares

import (
	"context"
	"goact-todo/database"
	"goact-todo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasks() ([]models.Todo, error) {
	todosCursor, err := database.TodosCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Todo{}, err
	}
	defer todosCursor.Close(context.TODO())

	var todos []models.Todo

	if err := todosCursor.All(context.TODO(), &todos); err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func AddNewTaskToDB(task models.Todo) error {
	_, err := database.TodosCollection().InsertOne(context.TODO(), task)
	return err
}

func TaskDone(id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = database.TodosCollection().UpdateOne(context.TODO(), bson.M{"_id": idHex}, bson.M{"$set": bson.M{"status": true}})
	return err
}

func TaskUnDone(id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = database.TodosCollection().UpdateOne(context.TODO(), bson.M{"_id": idHex}, bson.M{"$set": bson.M{"status": false}})
	return err
}

func DeleteTaskFromDB(id string) error {
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = database.TodosCollection().DeleteOne(context.TODO(), bson.M{"_id": idHex})
	return err
}

func DeleteAllTaskFromDB() error {
	_, err := database.TodosCollection().DeleteMany(context.TODO(), bson.M{})

	return err
}
