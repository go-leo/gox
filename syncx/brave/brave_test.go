package brave

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/syncx/chanx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	Do(func() {
		panic("this is a panic")
	}, func(p any) {
		t.Log(p)
	})
}

func TestGo(t *testing.T) {
	Go(func() {
		panic("this is a panic")
	}, func(p any) {
		t.Log(p)
	})
	time.Sleep(time.Second)
}

func TestDoE(t *testing.T) {
	err := DoE(func() error {
		return errors.New("this is do error")
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(err)
}

func TestDoEWithPanic(t *testing.T) {
	err := DoE(func() error {
		panic("this is a panic")
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(err)
}

func TestGoE(t *testing.T) {
	errC := GoE(func() error {
		return errors.New("this is do error")
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(<-errC)
}

func TestGoEWithPanic(t *testing.T) {
	errC := GoE(func() error {
		panic("this is a panic")
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(<-errC)
}

func TestGoRE(t *testing.T) {
	r1C, e1C := GoRE(func() (int, error) {
		<-time.After(time.Second)
		return 10, nil
	})
	r2C, e2C := GoRE(func() (int, error) {
		<-time.After(2 * time.Second)
		return 20, nil
	})
	r3C, e3C := GoRE(func() (int, error) {
		<-time.After(3 * time.Second)
		return 30, nil
	})

	ch := chanx.FanIn(context.Background(), e1C, e2C, e3C)
	err, ok := <-ch
	if ok {
		panic(err)
	}

	assert.Equal(t, 10, <-r1C)
	assert.Equal(t, 20, <-r2C)
	assert.Equal(t, 30, <-r3C)

}
