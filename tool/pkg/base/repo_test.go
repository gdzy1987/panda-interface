package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("https://github.com/PandaGoAdmin/panda-template-go.git", "")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), "/tmp/test_repo", "github.com/PandaGoAdmin/panda-template-go", nil); err != nil {
		t.Fatal(err)
	}
}
