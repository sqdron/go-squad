package hashers

import (
	"testing"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha1"
	"crypto/sha512"
)

func Test_Md5_Hash_Algorythm(t *testing.T) {
	Convey("Create application", t, func() {
		somePassword := "DjswkSJ3#%!@$DDJ;'d"
		var hasher = DefaultHash(md5.New(), 6);
		So(hasher, should.NotBeEmpty);

		Convey("Hash password", func() {
			hash:= hasher.Hash(somePassword);
			So(hash, should.NotBeEmpty);

			Convey("Validate correct password", func() {
				shouldBeTrue:= hasher.Check(somePassword, hash);
				So(shouldBeTrue, should.BeTrue);
			});

			Convey("Validate incorrect password", func() {
				shouldBeFalse:= hasher.Check("dsdferga34354t63", hash);
				So(shouldBeFalse, should.BeFalse);
			});
		});
	})
}

func Test_sha256_Hash_Algorythm(t *testing.T) {
	Convey("Create application", t, func() {
		somePassword := "DjswkSJ3#%!@$DDJ;'d"
		var hasher = DefaultHash(sha256.New(), 6);
		So(hasher, should.NotBeEmpty);

		Convey("Hash password", func() {
			hash:= hasher.Hash(somePassword);
			So(hash, should.NotBeEmpty);

			Convey("Validate correct password", func() {
				shouldBeTrue:= hasher.Check(somePassword, hash);
				So(shouldBeTrue, should.BeTrue);
			});

			Convey("Validate incorrect password", func() {
				shouldBeFalse:= hasher.Check("dsdferga34354t63", hash);
				So(shouldBeFalse, should.BeFalse);
			});
		});
	})
}

func Test_sha1_Hash_Algorythm(t *testing.T) {
	Convey("Create application", t, func() {
		somePassword := "DjswkSJ3#%!@$DDJ;'d"
		var hasher = DefaultHash(sha1.New(),12);
		So(hasher, should.NotBeEmpty);

		Convey("Hash password", func() {
			hash:= hasher.Hash(somePassword);
			So(hash, should.NotBeEmpty);

			Convey("Validate correct password", func() {
				shouldBeTrue:= hasher.Check(somePassword, hash);
				So(shouldBeTrue, should.BeTrue);
			});

			Convey("Validate incorrect password", func() {
				shouldBeFalse:= hasher.Check("dsdferga34354t63", hash);
				So(shouldBeFalse, should.BeFalse);
			});
		});
	})
}

func Test_sha512_Hash_Algorythm(t *testing.T) {
	Convey("Create application", t, func() {
		somePassword := "d"
		var hasher = DefaultHash(sha512.New(),12);
		So(hasher, should.NotBeEmpty);

		Convey("Hash password", func() {
			hash:= hasher.Hash(somePassword);
			So(hash, should.NotBeEmpty);

			Convey("Validate correct password", func() {
				shouldBeTrue:= hasher.Check(somePassword, hash);
				So(shouldBeTrue, should.BeTrue);
			});

			Convey("Validate incorrect password", func() {
				shouldBeFalse:= hasher.Check("dsdferga34354t63", hash);
				So(shouldBeFalse, should.BeFalse);
			});
		});
	})
}

func Benchmark_sha512_Hash_Algorythm(b *testing.B) {
	somePassword := "DjswkSJ3#%!@$DDJ;"
	var hasher = DefaultHash(sha512.New(),6);
	for i := 0; i < b.N; i++ {

		hash:= hasher.Hash(somePassword);
		hasher.Check(somePassword, hash)
	}
}
