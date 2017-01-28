package braintree

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestCustomFieldsMarshalXML(t *testing.T) {
	t.Parallel()

	cf := CustomFields{"key1": "value1", "key2": "value2"}
	buf, err := xml.Marshal(cf)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	str := string(buf)
	if !strings.Contains(str, "<key1>value1</key1>") {
		t.Errorf("want contains <key1>value1</key1>, got: %s", str)
	}
	if !strings.Contains(str, "<key2>value2</key2>") {
		t.Errorf("want contains <key2>value2</key2>, got: %s", str)
	}
}

func TestCustomFieldsUnmarshalXML(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()

		str := "<custom-fields><key1>value1</key1><key2>value2</key2></custom-fields>"
		cf := CustomFields{}
		if err := xml.Unmarshal([]byte(str), &cf); err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		got := cf["key1"]
		if got != "value1" {
			t.Errorf("want: value1, got: %s", got)
		}
		got = cf["key2"]
		if got != "value2" {
			t.Errorf("want: value2, got: %s", got)
		}
	})

	t.Run("malformed", func(t *testing.T) {
		t.Parallel()

		str := "<custom-fields><key1></custom-fields>"
		cf := CustomFields{}
		if err := xml.Unmarshal([]byte(str), &cf); err == nil {
			t.Errorf("want: an error, got nil")
		}
	})
}
