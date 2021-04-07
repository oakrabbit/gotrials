package mimic

import (
	"fmt"
	"testing"
	"time"
)

func TestBackground(t *testing.T) {
	todo:= fmt.Sprint(TODO())
	background := fmt.Sprint(Background())

	if todo == background {
		t.Errorf("TODO and BG are equal: %q vs %q", todo, background)
	}
}

func TestWithCancel(t *testing.T){
	ctx , cancel := WithCancel(Background())
	if err := ctx.Err(); err != nil{
		t.Errorf("error should be nill first, got %v",err)
	}
	cancel()

	<-ctx.Done()

	if err := ctx.Err(); err != Canceled{
		t.Errorf("error Should be canceled now, got %v", err)
	}
}
func TestWithCancelConcurrent(t *testing.T){
	ctx , cancel := WithCancel(Background())
	time.AfterFunc(1*time.Second, cancel)

	if err := ctx.Err(); err != nil{
		t.Errorf("error should be nill first, got %v",err)
	}
	<-ctx.Done()
	if err := ctx.Err(); err != Canceled{
		t.Errorf("error Should be canceled now, got %v", err)
	}
}

func TestWithCancelPropagation(t *testing.T){
	ctxA, cancelA := WithCancel(Background())
	ctxB, cancelB := WithCancel(ctxA)

	t.Log("value: ",ctxB)
	defer cancelB()
	cancelA()
	select {
		case <-ctxB.Done():
		case <- time.After(1* time.Second):
				t.Errorf("timeout")
	}
	if err := ctxB.Err(); err != Canceled{
		t.Errorf("error should be canceled now, got %v", err)
	}
}
