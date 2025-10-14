package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ademayowa/deploy-go-demo/db"
	"github.com/Ademayowa/deploy-go-demo/routes"
	"github.com/gin-gonic/gin"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestCreateProperty tests the CreateProperty endpoint
func TestCreateProperty(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	mock.ExpectPrepare("INSERT INTO properties").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/", bytes.NewBufferString(`{"title":"2 bedroom for rent","location":"UK"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	routes.CreateProperty(c)

	if w.Code != http.StatusCreated {
		t.Errorf("got %d, want 201", w.Code)
	}
}

// TestGetProperties tests the GetProperties endpoint
func TestGetProperties(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	rows := sqlmock.NewRows([]string{"id", "title", "location"}).
		AddRow("1", "Test", "Loc")
	mock.ExpectQuery("SELECT (.+) FROM properties").WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)

	routes.GetProperties(c)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, want 200", w.Code)
	}
}

// TestGetProperty tests the GetProperty endpoint
func TestGetProperty(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	rows := sqlmock.NewRows([]string{"id", "title", "location"}).
		AddRow("1", "Test", "Loc")
	mock.ExpectQuery("SELECT (.+) FROM properties WHERE id").WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("GET", "/1", nil)

	routes.GetProperty(c)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, want 200", w.Code)
	}
}

// TestDeleteProperty tests the DeleteProperty endpoint
func TestDeleteProperty(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	rows := sqlmock.NewRows([]string{"id", "title", "location"}).
		AddRow("1", "Test", "Loc")
	mock.ExpectQuery("SELECT (.+) FROM properties WHERE id").WillReturnRows(rows)
	mock.ExpectPrepare("DELETE FROM properties").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("DELETE", "/1", nil)

	routes.DeleteProperty(c)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, want 200", w.Code)
	}
}

// TestUpdateProperty tests the UpdateProperty endpoint
func TestUpdateProperty(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	mock.ExpectExec("UPDATE properties SET title").WillReturnResult(sqlmock.NewResult(0, 1))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("PUT", "/1", bytes.NewBufferString(`{"title":"Updated","location":"Loc"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	routes.UpdateProperty(c)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, want 200", w.Code)
	}
	if mock.ExpectationsWereMet() != nil {
		t.Error("mock expectations not met")
	}
}
