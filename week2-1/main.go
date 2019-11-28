package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"week2-1/database"

	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Student struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Age       int    `json:"age" bson:"age"`
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	client := database.GetMongoClient()
	collection := client.Database("go-sample").Collection("students")

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error")
	}

	var students []Student

	for cur.Next(context.TODO()) {
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	studentsJSON, err := json.Marshal(students)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(studentsJSON)
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	client := database.GetMongoClient()
	collection := client.Database("go-sample").Collection("students")

	var student Student

	error = json.Unmarshal(body, &student)
	if error != nil {
		log.Fatal(error)
	}

	_, error = collection.InsertOne(context.TODO(), student)
	if error != nil {
		log.Fatal(error)
	}

	bs, error := json.Marshal(student)
	if error != nil {
		log.Fatal(error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func Students(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetStudents(w, r)
	case "POST":
		AddStudent(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{message: "Not Found Method}`))
	}
}

func main() {
	database.MongoConnect()

	http.HandleFunc("/students", Students)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
