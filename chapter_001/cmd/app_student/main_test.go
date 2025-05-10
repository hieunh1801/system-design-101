package main

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-resty/resty/v2"
)

func TestPing(t *testing.T) {
	client := resty.New()
	count := 1000000
	for i := 0; i < count; i++ {
		res, err := client.R().
			Get("http://localhost:8080/ping")

		if err != nil {
			panic(err)
		}
		t.Log("DONE", res)
	}
}

func TestGetStudents(t *testing.T) {
	client := resty.New()
	res, err := client.R().
		Get("http://localhost:8080/students")

	if err != nil {
		panic(err)
	}

	t.Log("DONE", res)

}

type StudentFake struct {
	Name  string `faker:"name"`
	Email string `faker:"email"`
}

func TestCreateStudent(t *testing.T) {
	client := resty.New()
	count := 100000
	for i := 0; i < count; i++ {
		var studentData StudentFake
		err := faker.FakeData(&studentData)
		if err != nil {
			panic(err)
		}
		_, err = client.R().
			SetBody(studentData).
			Post("http://localhost:8080/students")

		if err != nil {
			panic(err)
		}
		t.Log("DONE", i)
	}

}

func TestCountStudents(t *testing.T) {
	client := resty.New()
	res, err := client.R().
		Get("http://localhost:8080/students/count")
	if err != nil {
		panic(err)
	}
	t.Log("DONE", res)
}
