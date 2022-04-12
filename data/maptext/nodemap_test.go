package maptext

import (
	"reflect"
	"testing"
)

var singleNode = NewRoot("SingleNode", BlankData)

var TreeTests = []struct {
	name    string
	want    *Tree
	wantErr bool
}{
	// TODO: Add test cases.
	{"tree", &Tree{singleNode, singleNode}, false},
}

func TestTree_String(t *testing.T) {
	for _, tt := range TreeTests {
		got := makeTree(tt.name)
		want := &Tree{root: got.root, last: got.last}

		TRun(t, "Tree.String", tt.name, got.String(), want.String(), tt.wantErr)
	}
}

func Test_makeTree(t *testing.T) {

	for _, tt := range TreeTests {
		{
			got := makeTree(tt.name)
			want := &Tree{root: got.root, last: got.last}

			TRun(t, "makeTree", tt.name, got, want, tt.wantErr)

			t.Run(tt.name, func(t *testing.T) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("makeTree() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
