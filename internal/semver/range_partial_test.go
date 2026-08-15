package semver

import "testing"

// Regression: parsePartial must strip build metadata (+) BEFORE
// stripping prerelease (-), otherwise a range like ">=1.0.0-alpha+build"
// will incorrectly include "+build" as part of the prerelease identifier.

func TestRangePartialBuildBeforePre(t *testing.T) {
	// ">=1.0.0-rc.1+build" should parse as prerelease=["rc","1"], build ignored
	// Version 1.0.0-rc.1 should satisfy this range
	v, err := Parse("1.0.0-rc.1")
	if err != nil {
		t.Fatal(err)
	}
	ok, err := SatisfiesRange(v, ">=1.0.0-rc.1+build")
	if err != nil {
		t.Fatalf("SatisfiesRange error: %v", err)
	}
	if !ok {
		t.Errorf("1.0.0-rc.1 should satisfy >=1.0.0-rc.1+build")
	}

	// "^1.2.3-beta+meta" should work the same as "^1.2.3-beta"
	v2, err := Parse("1.2.3-beta")
	if err != nil {
		t.Fatal(err)
	}
	ok, err = SatisfiesRange(v2, "^1.2.3-beta+meta")
	if err != nil {
		t.Fatalf("SatisfiesRange error: %v", err)
	}
	if !ok {
		t.Errorf("1.2.3-beta should satisfy ^1.2.3-beta+meta")
	}

	// Range with build-metadata containing hyphen: ">=1.0.0-alpha+build-info"
	// Must not confuse the hyphen in build as a prerelease separator
	v3, err := Parse("1.0.0-alpha")
	if err != nil {
		t.Fatal(err)
	}
	ok, err = SatisfiesRange(v3, ">=1.0.0-alpha+build-info")
	if err != nil {
		t.Fatalf("SatisfiesRange error: %v", err)
	}
	if !ok {
		t.Errorf("1.0.0-alpha should satisfy >=1.0.0-alpha+build-info")
	}
}
