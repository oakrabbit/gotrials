/**
Exploring the context implementation
 */

package mimic

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Context interface {

	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}


	Err() error

	Value(key interface{}) interface{}
}


type emptyCtx int

func (emptyCtx) Deadline() (deadline time.Time, ok bool){ return }
func (emptyCtx)  Done() <-chan struct{} {return nil}
func (emptyCtx)  Err() error {return nil}
func (emptyCtx)  Value(key interface{}) interface{} {return nil}


var (
	background = new(emptyCtx)
	todo = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

type cancelCtx struct{
	Context
	done chan struct{}
	err	error
	mu sync.Mutex
}


// with cancel context
func (ctx *cancelCtx)  Done() <-chan struct{} 				{return ctx.done}
func (ctx *cancelCtx)  Err() error 							{
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}


var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc){
	ctx := &cancelCtx{
		Context: parent,
		done: make(chan struct{}),
	}
	fmt.Sprint(ctx)
	cancel := func(){
		ctx.cancel(Canceled)
	}
	go func(){
		select{
			case <- parent.Done():
					ctx.cancel(parent.Err())
			case <- ctx.Done():

		}
	}()

	return ctx,  cancel
}


func (ctx *cancelCtx) cancel(err error){
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err != nil {
		return
	}
	ctx.err = Canceled
	close(ctx.done)
}


type deadlineCtx struct{
	*cancelCtx
	deadLine time.Time
}
var DeadLineExceeded = errors.New("Deadline exceeded!")


func (ctx *deadlineCtx) Deadline() (deadline time.Time, ok bool){ return ctx.deadLine, true }


func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc){
	/// cancel context after given time
	cctx, cancel := WithCancel(parent)


	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadLine: deadline,
	}
	timer := time.AfterFunc(time.Until(deadline), func(){
		ctx.cancel(DeadLineExceeded)
	})

	stop := func(){
		timer.Stop()
		cancel()
	}

	return ctx, stop
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc){
	// just return with cancel with timout adding the duration to now.
	return WithDeadline(parent, time.Now().Add(timeout))
}

type valueCtx struct{
	Context
	value, key interface{}
}


func (ctx *valueCtx) Value(key interface{}) interface{}{
	if key == ctx.key{
		return ctx.value
	}
	return ctx.Context.Value(key)
}
func WithValue(parent Context, key, val interface{}) Context{

	if key == nil{
		panic("key os nil")
	}

	if reflect.TypeOf(key).Comparable(){
		panic("key is not comparable")
	}

	return &valueCtx{
		Context: parent,
		key: key,
		value: val,
	}
}