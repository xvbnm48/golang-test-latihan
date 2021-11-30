package integer
import (
	"fmt"
	"testing"
)
func TestAdd(t *testing.T){
	sum := Add(2,2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%v' but got '%d' ", expected, sum)
	}
}

func TestExampleAdd(t *testing.T){
	sum := Add(1,5)
	fmt.Println(sum)
}
