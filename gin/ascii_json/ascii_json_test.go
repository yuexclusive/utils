package ascii_json

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func ExampleFoo() {
	engine := gin.Default()
	Foo(engine)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/someJSON", nil)
	if err != nil {
		panic(err)
	}
	engine.ServeHTTP(w, r)

	//Output: {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	fmt.Println(w.Body.String())
}
