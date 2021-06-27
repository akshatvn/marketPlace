package controller

import (
	"reflect"
	"testing"
)

func TestSplitIntoArgs(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{
			text: `CREATE_LISTING user1 'Phone model 8' 'Black color, brand new' 1000 'Electronics'`,
			want: []string{"CREATE_LISTING", "user1", "Phone model 8", "Black color, brand new", "1000", "Electronics"},
		},
		{
			text: `CREATE_LISTING user2 'BTWIN 900' 'Akshats old bike' 1000 'Bicycles'`,
			want: []string{"CREATE_LISTING", "user2", "BTWIN 900", "Akshats old bike", "1000", "Bicycles"},
		},
	}

	for i, test := range tests  {
		got := splitIntoArgs(test.text)
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("Failed test: %d\n Wanted: %+v\n Got: %+v", i, test.want, got)
		}
	}

}
