package mymodproj

import "testing"

func TestReturnString(t *testing.T){
        w := "Arun"
        if x :=ReturnString(w); x == " " {
                t.Errorf(x)
        }
}
