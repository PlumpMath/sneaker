package secman

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/gen/s3"
)

func TestRm(t *testing.T) {
	fakeS3 := &FakeS3{
		DeleteResponses: []s3.DeleteObjectOutput{
			{},
			{},
		},
	}

	man := Manager{
		Objects: fakeS3,
		Bucket:  "bucket",
		Prefix:  "secrets/",
	}

	if err := man.Rm("weeble/wobble.txt"); err != nil {
		t.Fatal(err)
	}

	req := fakeS3.DeleteRequests[0]
	if v, want := *req.Bucket, "bucket"; v != want {
		t.Errorf("Bucket was %q, but expected %q", v, want)
	}

	if v, want := *req.Key, "secrets/weeble/wobble.txt.kms"; v != want {
		t.Errorf("Key was %q, but expected %q", v, want)
	}

	req = fakeS3.DeleteRequests[1]
	if v, want := *req.Bucket, "bucket"; v != want {
		t.Errorf("Bucket was %q, but expected %q", v, want)
	}

	if v, want := *req.Key, "secrets/weeble/wobble.txt.aes"; v != want {
		t.Errorf("Key was %q, but expected %q", v, want)
	}
}
