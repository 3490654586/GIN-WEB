package jwt

import (
	"fmt"
	"testing"
)



func TestSetToken(t *testing.T) {
    SetToken("admin")
     fmt.Println(  GetToken("dadwa"))
}


func TestGetToken(t *testing.T) {

}

