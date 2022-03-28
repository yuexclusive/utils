package validator

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func Example() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	w := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/bookable?check_in=2118-04-16&check_out=2118-04-17", nil)

	route.ServeHTTP(w, r)

	w2 := httptest.NewRecorder()

	r2 := httptest.NewRequest("GET", "/bookable?/check_in=2118-03-10&check_out=2118-03-09", nil)

	route.ServeHTTP(w2, r2)

	//Output:{"message":"Booking dates are valid!"}
	//{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'required' tag"}
	fmt.Println(w.Body.String())
	fmt.Println(w2.Body.String())

}
