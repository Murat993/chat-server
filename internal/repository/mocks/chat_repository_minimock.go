package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/Murat993/chat-server/internal/repository.ChatRepository -o ./mocks/chat_repository_minimock.go -n ChatRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Murat993/chat-server/internal/dto"
	"github.com/gojuno/minimock/v3"
)

// ChatRepositoryMock implements repository.ChatRepository
type ChatRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, chat *dto.Chat) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, chat *dto.Chat)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatRepositoryMockCreate
}

// NewChatRepositoryMock returns a mock for repository.ChatRepository
func NewChatRepositoryMock(t minimock.Tester) *ChatRepositoryMock {
	m := &ChatRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatRepositoryMockCreateParams{}

	return m
}

type mChatRepositoryMockCreate struct {
	mock               *ChatRepositoryMock
	defaultExpectation *ChatRepositoryMockCreateExpectation
	expectations       []*ChatRepositoryMockCreateExpectation

	callArgs []*ChatRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// ChatRepositoryMockCreateExpectation specifies expectation struct of the ChatRepository.Create
type ChatRepositoryMockCreateExpectation struct {
	mock    *ChatRepositoryMock
	params  *ChatRepositoryMockCreateParams
	results *ChatRepositoryMockCreateResults
	Counter uint64
}

// ChatRepositoryMockCreateParams contains parameters of the ChatRepository.Create
type ChatRepositoryMockCreateParams struct {
	ctx  context.Context
	chat *dto.Chat
}

// ChatRepositoryMockCreateResults contains results of the ChatRepository.Create
type ChatRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Expect(ctx context.Context, chat *dto.Chat) *mChatRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &ChatRepositoryMockCreateParams{ctx, chat}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Inspect(f func(ctx context.Context, chat *dto.Chat)) *mChatRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Return(i1 int64, err error) *ChatRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the ChatRepository.Create method
func (mmCreate *mChatRepositoryMockCreate) Set(f func(ctx context.Context, chat *dto.Chat) (i1 int64, err error)) *ChatRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the ChatRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the ChatRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the ChatRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatRepositoryMockCreate) When(ctx context.Context, chat *dto.Chat) *ChatRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	expectation := &ChatRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatRepositoryMockCreateParams{ctx, chat},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up ChatRepository.Create return parameters for the expectation previously defined by the When method
func (e *ChatRepositoryMockCreateExpectation) Then(i1 int64, err error) *ChatRepositoryMock {
	e.results = &ChatRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.ChatRepository
func (mmCreate *ChatRepositoryMock) Create(ctx context.Context, chat *dto.Chat) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, chat)
	}

	mm_params := &ChatRepositoryMockCreateParams{ctx, chat}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := ChatRepositoryMockCreateParams{ctx, chat}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, chat)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatRepositoryMock.Create. %v %v", ctx, chat)
	return
}

// CreateAfterCounter returns a count of finished ChatRepositoryMock.Create invocations
func (mmCreate *ChatRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatRepositoryMock.Create invocations
func (mmCreate *ChatRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatRepositoryMockCreate) Calls() []*ChatRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *ChatRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to ChatRepositoryMock.Create")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ChatRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone()
}