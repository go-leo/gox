package syncx

import (
	"sync"
	"sync/atomic"
)

// MapInterface is the interface map implements.
// copy from sync/map_reference_test.go
type MapInterface interface {
	Load(any) (any, bool)
	Store(key, value any)
	LoadOrStore(key, value any) (actual any, loaded bool)
	LoadAndDelete(key any) (value any, loaded bool)
	Delete(any)
	Swap(key, value any) (previous any, loaded bool)
	CompareAndSwap(key, old, new any) (swapped bool)
	CompareAndDelete(key, old any) (deleted bool)
	Range(func(key, value any) (shouldContinue bool))
}

// RWMutexMap is an implementation of mapInterface using a sync.RWMutex.
type RWMutexMap struct {
	mu    sync.RWMutex
	dirty map[any]any
}

func (m *RWMutexMap) Load(key any) (value any, ok bool) {
	m.mu.RLock()
	value, ok = m.dirty[key]
	m.mu.RUnlock()
	return
}

func (m *RWMutexMap) Store(key, value any) {
	m.mu.Lock()
	if m.dirty == nil {
		m.dirty = make(map[any]any)
	}
	m.dirty[key] = value
	m.mu.Unlock()
}

func (m *RWMutexMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	m.mu.Lock()
	actual, loaded = m.dirty[key]
	if !loaded {
		actual = value
		if m.dirty == nil {
			m.dirty = make(map[any]any)
		}
		m.dirty[key] = value
	}
	m.mu.Unlock()
	return actual, loaded
}

func (m *RWMutexMap) Swap(key, value any) (previous any, loaded bool) {
	m.mu.Lock()
	if m.dirty == nil {
		m.dirty = make(map[any]any)
	}

	previous, loaded = m.dirty[key]
	m.dirty[key] = value
	m.mu.Unlock()
	return
}

func (m *RWMutexMap) LoadAndDelete(key any) (value any, loaded bool) {
	m.mu.Lock()
	value, loaded = m.dirty[key]
	if !loaded {
		m.mu.Unlock()
		return nil, false
	}
	delete(m.dirty, key)
	m.mu.Unlock()
	return value, loaded
}

func (m *RWMutexMap) Delete(key any) {
	m.mu.Lock()
	delete(m.dirty, key)
	m.mu.Unlock()
}

func (m *RWMutexMap) CompareAndSwap(key, old, new any) (swapped bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.dirty == nil {
		return false
	}

	value, loaded := m.dirty[key]
	if loaded && value == old {
		m.dirty[key] = new
		return true
	}
	return false
}

func (m *RWMutexMap) CompareAndDelete(key, old any) (deleted bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.dirty == nil {
		return false
	}

	value, loaded := m.dirty[key]
	if loaded && value == old {
		delete(m.dirty, key)
		return true
	}
	return false
}

func (m *RWMutexMap) Range(f func(key, value any) (shouldContinue bool)) {
	m.mu.RLock()
	keys := make([]any, 0, len(m.dirty))
	for k := range m.dirty {
		keys = append(keys, k)
	}
	m.mu.RUnlock()

	for _, k := range keys {
		v, ok := m.Load(k)
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

// DeepCopyMap is an implementation of mapInterface using a Mutex and
// atomic.Value.  It makes deep copies of the map on every write to avoid
// acquiring the Mutex in Load.
type DeepCopyMap struct {
	mu    sync.Mutex
	clean atomic.Value
}

func (m *DeepCopyMap) Load(key any) (value any, ok bool) {
	clean, _ := m.clean.Load().(map[any]any)
	value, ok = clean[key]
	return value, ok
}

func (m *DeepCopyMap) Store(key, value any) {
	m.mu.Lock()
	dirty := m.dirty()
	dirty[key] = value
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *DeepCopyMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	clean, _ := m.clean.Load().(map[any]any)
	actual, loaded = clean[key]
	if loaded {
		return actual, loaded
	}

	m.mu.Lock()
	// Reload clean in case it changed while we were waiting on MapInterface.mu.
	clean, _ = m.clean.Load().(map[any]any)
	actual, loaded = clean[key]
	if !loaded {
		dirty := m.dirty()
		dirty[key] = value
		actual = value
		m.clean.Store(dirty)
	}
	m.mu.Unlock()
	return actual, loaded
}

func (m *DeepCopyMap) Swap(key, value any) (previous any, loaded bool) {
	m.mu.Lock()
	dirty := m.dirty()
	previous, loaded = dirty[key]
	dirty[key] = value
	m.clean.Store(dirty)
	m.mu.Unlock()
	return
}

func (m *DeepCopyMap) LoadAndDelete(key any) (value any, loaded bool) {
	m.mu.Lock()
	dirty := m.dirty()
	value, loaded = dirty[key]
	delete(dirty, key)
	m.clean.Store(dirty)
	m.mu.Unlock()
	return
}

func (m *DeepCopyMap) Delete(key any) {
	m.mu.Lock()
	dirty := m.dirty()
	delete(dirty, key)
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *DeepCopyMap) CompareAndSwap(key, old, new any) (swapped bool) {
	clean, _ := m.clean.Load().(map[any]any)
	if previous, ok := clean[key]; !ok || previous != old {
		return false
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	dirty := m.dirty()
	value, loaded := dirty[key]
	if loaded && value == old {
		dirty[key] = new
		m.clean.Store(dirty)
		return true
	}
	return false
}

func (m *DeepCopyMap) CompareAndDelete(key, old any) (deleted bool) {
	clean, _ := m.clean.Load().(map[any]any)
	if previous, ok := clean[key]; !ok || previous != old {
		return false
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	dirty := m.dirty()
	value, loaded := dirty[key]
	if loaded && value == old {
		delete(dirty, key)
		m.clean.Store(dirty)
		return true
	}
	return false
}

func (m *DeepCopyMap) Range(f func(key, value any) (shouldContinue bool)) {
	clean, _ := m.clean.Load().(map[any]any)
	for k, v := range clean {
		if !f(k, v) {
			break
		}
	}
}

func (m *DeepCopyMap) dirty() map[any]any {
	clean, _ := m.clean.Load().(map[any]any)
	dirty := make(map[any]any, len(clean)+1)
	for k, v := range clean {
		dirty[k] = v
	}
	return dirty
}

type GenericMap[K comparable, V any] struct {
	MapInterface MapInterface
}

func (m *GenericMap[K, V]) Load(key K) (V, bool) {
	var v V
	load, ok := m.MapInterface.Load(key)
	if !ok {
		return v, false
	}
	v = load.(V)
	return v, true
}

func (m *GenericMap[K, V]) Store(key K, value V) {
	m.MapInterface.Swap(key, value)
}

func (m *GenericMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	var v V
	actual, loaded := m.MapInterface.LoadOrStore(key, value)
	if !loaded {
		return v, false
	}
	v = actual.(V)
	return v, true
}

func (m *GenericMap[K, V]) LoadAndDelete(key K) (V, bool) {
	var v V
	value, loaded := m.MapInterface.LoadAndDelete(key)
	if !loaded {
		return v, false
	}
	v = value.(V)
	return v, true
}

func (m *GenericMap[K, V]) Delete(key K) {
	m.MapInterface.Delete(key)
}

func (m *GenericMap[K, V]) Swap(key K, value V) (V, bool) {
	var v V
	previous, loaded := m.MapInterface.Swap(key, value)
	if !loaded {
		return v, false
	}
	v = previous.(V)
	return v, true
}

func (m *GenericMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.MapInterface.CompareAndSwap(key, old, new)
}

func (m *GenericMap[K, V]) CompareAndDelete(key K, old V) bool {
	return m.MapInterface.CompareAndDelete(key, old)
}

func (m *GenericMap[K, V]) Range(f func(key K, value V) bool) {
	m.MapInterface.Range(func(key, value any) bool {
		k := key.(K)
		v := value.(V)
		return f(k, v)
	})
}
