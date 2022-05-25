package cache

import (
	"reflect"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	testCases := []struct {
		name string
		want Cache
	}{
		{
			name: "Constructor",
			want: Cache{cacheItems: make(map[string]Item)},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := NewCache(); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name  string
		items map[string]Item
		input string
		want  string
		ok    bool
	}{
		{
			name: "Existed key",
			items: map[string]Item{
				"ABC": {
					itemValue: "DEF",
				},
			},
			input: "ABC",
			want:  "DEF",
			ok:    true,
		},
		{
			name: "Expired key",
			items: map[string]Item{
				"GHI": {
					itemValue:  "JKL",
					itemExpire: time.Now().Add(time.Duration(-60) * time.Second),
				},
			},
			input: "GHI",
			want:  "",
			ok:    false,
		},
		{
			name:  "Not existed key",
			items: nil,
			input: "XYZ",
			want:  "",
			ok:    false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newCache := &Cache{
				cacheItems: testCase.items,
			}

			got, ok := newCache.Get(testCase.input)

			if got != testCase.want {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}

			if ok != testCase.ok {
				t.Errorf("ok: %v, want: %v", ok, testCase.ok)
			}
		})
	}
}

func TestPut(t *testing.T) {
	type inp struct {
		key   string
		value string
	}

	testCases := []struct {
		name  string
		items map[string]Item
		input inp
		want  string
		ok    bool
	}{
		{
			name:  "Put key",
			items: make(map[string]Item),
			input: inp{key: "ABC", value: "DEF"},
			want:  "DEF",
			ok:    true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newCache := &Cache{
				cacheItems: testCase.items,
			}

			newCache.Put(testCase.input.key, testCase.input.value)

			got, ok := newCache.Get(testCase.input.key)

			if got != testCase.want {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}

			if ok != testCase.ok {
				t.Errorf("ok: %v, want: %v", ok, testCase.ok)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type fields struct {
		items map[string]Item
	}

	testCases := []struct {
		name  string
		items map[string]Item
		want  []string
	}{
		{
			name: "Get keys",
			items: map[string]Item{
				"ABC": {
					itemValue: "DEF",
				},
			},
			want: []string{
				"ABC",
			},
		},
		{
			name: "Get keys without expired",
			items: map[string]Item{
				"ABC": {
					itemValue: "DEF",
				},
				"GHI": {
					itemValue:  "KLM",
					itemExpire: time.Now().Add(time.Duration(-60) * time.Second),
				},
			},
			want: []string{
				"ABC",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newCache := &Cache{
				cacheItems: testCase.items,
			}

			if got := newCache.Keys(); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}
		})
	}
}

func TestPutTill(t *testing.T) {
	type inp struct {
		key      string
		value    string
		deadline time.Time
	}

	testCases := []struct {
		name  string
		items map[string]Item
		input inp
		want  string
		ok    bool
	}{
		{
			name:  "Put key till",
			items: make(map[string]Item),
			input: inp{key: "ABC", value: "DEF", deadline: time.Now().Add(36000 * time.Minute)},
			want:  "DEF",
			ok:    true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newCache := &Cache{
				cacheItems: testCase.items,
			}

			newCache.PutTill(testCase.input.key, testCase.input.value, testCase.input.deadline)

			got, ok := newCache.Get(testCase.input.key)

			if got != testCase.want {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}

			if ok != testCase.ok {
				t.Errorf("ok: %v, want: %v", ok, testCase.ok)
			}
		})
	}
}
