package email

import (
	"testing"
)

func Test_Send(t *testing.T) {
	dialer := NewDialer("smtp.gmail.com", 587, "evolve.publisher123@gmail.com", "")
	err := dialer.Send("haha", "", "this is a fucking body", "xxx@icloud.com")
	if err != nil {
		t.Error(err)
	}
}
