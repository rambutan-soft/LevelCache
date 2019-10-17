package levelcache

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	numCounters = 256
	bufferItems = numCounters / 4
	sampleItems = numCounters * 8
)

const (
	_  = iota
	Ki = 1 << (10 * iota)
	Mi
	Gi
)

func TestAdd(t *testing.T) {

	Convey("Given an initialized Cache", t, func() {
		a := NewCache()

		Convey("When a Key Value pair is added to the cache", func() {
			a.Add("KEY_TRAMS", "VALUE_TRAMS")
			Convey("Then size should be 1", func() {
				So(a.Size, ShouldEqual, 1)
			})

			Convey("Then retrieved key should be 'KEY_TRAMS'", func() {
				So(a.Key[0], ShouldEqual, "KEY_TRAMS")
			})

			Convey("Then retrieved value should be 'VALUE_TRAMS'", func() {
				So(a.KeyValue[a.Key[0]], ShouldEqual, "VALUE_TRAMS")
			})
		})
		Convey("When a duplicate Key Value pair is added to the cache", func() {
			a.Add("KEY_TRAMS", "VALUE_TRAMS")
			Convey("Then size should remain 1", func() {
				So(a.Size, ShouldEqual, 1)
			})
		})
		Convey("When a updated Key Value pair is added to the cache", func() {
			a.Add("KEY_TRAMS", "UPDATED_VALUE_TRAMS")
			Convey("Then size should remain 1", func() {
				So(a.Size, ShouldEqual, 1)
			})

			Convey("Then retrieved key should be 'KEY_TRAMS'", func() {
				So(a.Key[0], ShouldEqual, "KEY_TRAMS")
			})

			Convey("Then retrieved value should be 'UPDATED_VALUE_TRAMS'", func() {
				So(a.KeyValue[a.Key[0]], ShouldEqual, "UPDATED_VALUE_TRAMS")
			})
		})
	})
}

func TestDelete(t *testing.T) {

	Convey("Given an initialized Cache", t, func() {
		a := NewCache()

		Convey("When 3 Key Value pairs are added to the cache\n"+
			"And the 2nd key is deleted from the cache", func() {
			a.Add("KEY_TRAMS1", "VALUE_TRAMS1")
			a.Add("KEY_TRAMS2", "VALUE_TRAMS2")
			a.Add("KEY_TRAMS3", "VALUE_TRAMS3")
			So(a.Size, ShouldEqual, 3)
			a.Delete("KEY_TRAMS2")

			Convey("Then size should be 2", func() {
				So(a.Size, ShouldEqual, 2)
			})

			Convey("Then the last retrieved key should be 'KEY_TRAMS3'", func() {
				So(a.Key[1], ShouldEqual, "KEY_TRAMS3")
			})

			Convey("Then the last retrieved value should be 'VALUE_TRAMS3'", func() {
				So(a.KeyValue[a.Key[1]], ShouldEqual, "VALUE_TRAMS3")
			})
		})
	})
}

func TestGetValueByKey(t *testing.T) {

	Convey("Given a Cache with Key Value pairs", t, func() {
		a := NewCache()
		a.Add("KEY_TRAMS", "VALUE_TRAMS")

		Convey("When value is retrieved using the key", func() {
			b := a.GetValueByKey("KEY_TRAMS")
			Convey("Then Value should be 'VALUE_TRAMS'", func() {
				So(b, ShouldEqual, "VALUE_TRAMS")

			})

		})

	})

}

func TestLen(t *testing.T) {

	Convey("Given a Cache with 2 Key Value pairs", t, func() {
		a := NewCache()
		a.Add("KEY_TRAMS", "VALUE_TRAMS")
		a.Add("KEY_TRAMS1", "VALUE_TRAMS1")

		Convey("When length is retrieved using the function LEN", func() {
			b := a.Len()
			Convey("Then Length should be 2", func() {
				So(b, ShouldEqual, 2)

			})

		})

	})

	/*Convey("Adding 2 KV pairs to check length", func() {

		a := NewCache()
		b := a.Len()
		So(b, ShouldEqual, 0)

	})*/

}

func TestGetKeyByIndex(t *testing.T) {

	Convey("Given a Cache with 2 Key Value pairs", t, func() {
		a := NewCache()
		a.Add("KEY_TRAMS", "VALUE_TRAMS")
		a.Add("KEY_TRAMS1", "VALUE_TRAMS1")

		Convey("When key at index 1 is retreived", func() {
			b := a.GetKeyByIndex(1)
			Convey("Then Key should equal 'KEY_TRAMS1'", func() {

				So(b, ShouldEqual, "KEY_TRAMS1")

			})

		})

	})

}

func TestGetValueByIndex(t *testing.T) {

	Convey("Given a Cache with 2 Key Value pairs", t, func() {
		a := NewCache()
		a.Add("KEY_TRAMS", "VALUE_TRAMS")
		a.Add("KEY_TRAMS1", "VALUE_TRAMS1")

		Convey("When value at index 1 is retreived", func() {
			b := a.GetValueByIndex(1)
			Convey("Then Key should equal 'VALUE_TRAMS1'", func() {
				So(b, ShouldEqual, "VALUE_TRAMS1")

			})

		})

	})

}

func TestGetRandomKey(t *testing.T) {

	Convey("Given a Cache with 2 Key Value pairs", t, func() {
		a := NewCache()
		a.Add("KEY_TRAMS", "VALUE_TRAMS")
		a.Add("KEY_TRAMS1", "VALUE_TRAMS1")

		Convey("When a random key is accessed", func() {
			b := a.GetRandomKey()

			Convey("Then the random key should be from the added keys", func() {
				So(b, ShouldBeIn, []string{"KEY_TRAMS", "KEY_TRAMS1"})
			})

			Convey("Then the random key should not be from the keys not added", func() {
				So(b, ShouldNotBeIn, []string{"KEY_TRAMS2", "KEY_TRAMS#"})
			})

		})

	})

}

func TestGetRandomKV(t *testing.T) {

	Convey("Given a Cache with 2 Key Value pairs", t, func() {
		a := NewCache()
		a.Add("KEY_TRAMS", "VALUE_TRAMS")
		a.Add("KEY_TRAMS1", "VALUE_TRAMS1")

		Convey("When a random key value pair is accessed", func() {

			b, c := a.GetRandomKV()
			b += c

			Convey("Then the random key value pair should be from the added pairs", func() {
				So(b, ShouldBeIn, []string{"KEY_TRAMSVALUE_TRAMS", "KEY_TRAMS1VALUE_TRAMS1"})
			})

			Convey("Then the random key value pair should not be from the key value pair not added", func() {
				So(b, ShouldNotBeIn, []string{"KEY_TRAMSVALUETRAMS1", "KEY_TRAMS1VALUETRAMS"})

			})

		})

	})

}

func TestNewCache(t *testing.T) {

	Convey("Given two Initialized Caches", t, func() {
		a := NewCache()
		b := NewCache()

		Convey("When the returned pointer for the caches is compared", func() {
			Convey("Then the pointer should not be Nil", func() {
				So(a, ShouldNotBeNil)
			})

			Convey("Then both pointers should not be equal", func() {
				So(b, ShouldNotEqual, a)
			})

		})

	})

}
