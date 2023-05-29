package services

import (
	"reflect"
	"testing"
)

func TestEncodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   EncodingTable
		want DecodingTree
	}{
		{
			name: "test #1",
			et: EncodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: DecodingTree{
				Left: &DecodingTree{
					Right: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: "z",
							},
						},
					},
				},
				Right: &DecodingTree{
					Left: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: "b",
							},
						},
					},
					Right: &DecodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodingTable.DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
