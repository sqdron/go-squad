package hashers

import (
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_bcrypt_Hash_Algorythm(t *testing.T) {
	Convey("Create application", t, func() {
		somePassword := "DjswkSJ3#%!@$DDJ;'d"
		var hasher = Bcrypt(6)
		So(hasher, should.NotBeEmpty)

		Convey("Hash password", func() {
			hash := hasher.Hash(somePassword)
			So(hash, should.NotBeEmpty)

			Convey("Validate correct password", func() {
				shouldBeTrue := hasher.Check(somePassword, hash)
				So(shouldBeTrue, should.BeTrue)
			})

			Convey("Validate incorrect password", func() {
				shouldBeFalse := hasher.Check("dsdferga34354t63", hash)
				So(shouldBeFalse, should.BeFalse)
			})
		})
	})
}

func BenchmarkSerial(b *testing.B) {
	somePassword := "DjswkSJ3#%!@$DDJ;'d"
	var hasher = Bcrypt(10)
	hash := hasher.Hash(somePassword)
	for i := 0; i < b.N; i++ {
		hasher.Check(somePassword, hash)
	}
}

func BenchmarkAsManyGoroutinesPossible(b *testing.B) {
	var done = make(chan bool)
	somePassword := "DjswkSJ3#%!@$DDJ;'d"
	var hasher = Bcrypt(10)
	hash := hasher.Hash(somePassword)

	for i := 0; i < b.N; i++ {
		go func() {
			hasher.Check(somePassword, hash)
			done <- true
		}()
	}
	for i := 0; i < b.N; i++ {
		<-done
	}
}
