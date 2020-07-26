package mymodproj

import "testing"

func TestReturnString(t *testing.T){
        w := "ArunavaSahoo"
        if x :=ReturnString(w); x == " " {
                t.Errorf(x)
        }
}
