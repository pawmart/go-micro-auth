package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	for _, test := range initTestCases {
		Source = test.source
		// TODO: stub open file operation
		FilePath = "../service.cfg"
		actualResult, err := Init()

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected error to be %v, but got %v", test.err, err)
		}

		if test.configInited {
			// TODO: validate cofig type (file, consul)
			require.NotNil(t, actualResult, "Expected config not nil")
		} else {
			require.Nil(t, actualResult, "Expected no config inited")
		}
	}
}

func TestGetKVPair(t *testing.T) {
	for _, test := range getKVPairTestCases {
		config = test.config
		actualResult, err := config.GetKVPair(test.key)

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected %v, but got %v", test.err, err)
		}

		if !reflect.DeepEqual(test.expected, actualResult) {
			t.Fatalf("Expected %v, but got %v", test.expected, actualResult)
		}
	}
}
