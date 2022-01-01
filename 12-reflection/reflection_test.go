package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name         string
		Input        interface{}
		ExpectedCals []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Yury"},
			[]string{"Yury"},
		}, {
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Yury", "LA"},
			[]string{"Yury", "LA"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Yury", 30},
			[]string{"Yury"},
		}, {
			"Nested fields",
			Person{
				"Yury",
				Profile{30, "LA"},
			},
			[]string{"Yury", "LA"},
		}, {
			"Pointer  to things",
			&Person{
				"Yury",
				Profile{30, "LA"},
			},
			[]string{"Yury", "LA"},
		}, {
			"Slices",
			[]Profile{
				{30, "LA"},
				{40, "Salt Lake City"},
			},
			[]string{"LA", "Salt Lake City"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCals) {
				t.Errorf("got %v want %v", got, test.ExpectedCals)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "LA"}
			aChannel <- Profile{40, "Denver"}
			close(aChannel)
		}()

		var got []string
		want := []string{"LA", "Denver"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "LA"}, Profile{34, "Denver"}
		}

		var got []string
		want := []string{"LA", "Denver"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
