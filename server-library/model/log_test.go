package model_test

import (
	"testing"
	"time"

	"github.com/danielbintar/angel/server-library/factory"
	"github.com/danielbintar/angel/server-library/model"

	"github.com/stretchr/testify/assert"
)

type Dummy struct {
	PreviousData *Dummy
	ID           uint
	Username     string
	Age          uint
	CreatedAt    time.Time
}

func TestLog(t *testing.T) {
	dummy := Dummy{Username: "haha", Age: uint(20)}
	copyiedData := dummy
	dummy.PreviousData = &copyiedData
	dummy.Username = "lala"
	dummy.Age++

	t.Run("no change struct", func(t *testing.T) {
		dummy := Dummy{Username: "haha", Age: uint(20)}
		copyiedData := dummy
		dummy.PreviousData = &copyiedData
		publisher := factory.MockAsyncPublisher()
		assert.NotPanics(t, func() { model.Log("users", dummy, &publisher) })
		assert.Equal(t, uint(0), publisher.PublishCallCount)
	})

	t.Run("success", func(t *testing.T) {
		publisher := factory.MockAsyncPublisher()
		assert.NotPanics(t, func() { model.Log("users", dummy, &publisher) })
		assert.Equal(t, uint(1), publisher.PublishCallCount)
	})
}

func TestGenerateChanges(t *testing.T) {
	t.Run("non struct type", func(t *testing.T) {
		assert.Panics(t, func() { model.GenerateChanges(1) })
	})

	t.Run("not have previous data field", func(t *testing.T) {
		type Broken struct {
			ID uint
		}
		assert.Panics(t, func() { model.GenerateChanges(Broken{}) })
	})

	t.Run("nil previous data", func(t *testing.T) {
		assert.Panics(t, func() { model.GenerateChanges(Dummy{}) })
	})

	t.Run("success", func(t *testing.T) {
		dummy := Dummy{Username: "haha", Age: uint(20)}
		copyiedData := dummy
		dummy.PreviousData = &copyiedData
		dummy.Username = "lala"
		dummy.Age++
		result := model.GenerateChanges(dummy)
		assert.Equal(t, 2, len(result))

		usernameI := 0
		ageI := 1
		if result[0].Key == "Age" {
			usernameI = 1
			ageI = 0
		}

		assert.Equal(t, "Username", result[usernameI].Key)
		assert.Equal(t, "haha", result[usernameI].Previous)
		assert.Equal(t, "lala", result[usernameI].After)
		assert.Equal(t, "Age", result[ageI].Key)
		assert.Equal(t, "20", result[ageI].Previous)
		assert.Equal(t, "21", result[ageI].After)
	})
}
