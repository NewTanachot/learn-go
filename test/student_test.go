package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/NewTanachot/learn-go/model"
	"github.com/NewTanachot/learn-go/server"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestStudentvalidator(t *testing.T) {
	app := server.Setup()
	tests := getStudentValidateRequestMock()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/student", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
}

func getStudentValidateRequestMock() []FiberTestModel[model.Student] {
	return []FiberTestModel[model.Student]{
		{
			description: "Valid input",
			requestBody: &model.Student{
				Id:    1,
				Name:  "Tanachot Udomsartporn",
				Grade: "A",
				Score: 100,
			},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Nil input",
			requestBody:  nil,
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description: "Invalid fullname",
			requestBody: &model.Student{
				Id:    1,
				Name:  "1234567890",
				Grade: "A",
				Score: 100,
			},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description: "Invalid Max Score",
			requestBody: &model.Student{
				Id:    1,
				Name:  "Tanachot Udomsartporn",
				Grade: "A",
				Score: 1111111,
			},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description: "Invalid Minus Score",
			requestBody: &model.Student{
				Id:    1,
				Name:  "Tanachot Udomsartporn",
				Grade: "A",
				Score: -100,
			},
			expectStatus: fiber.StatusBadRequest,
		},
	}
}
