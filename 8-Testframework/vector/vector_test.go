//Test module for vector
package vector

import "testing"

func TestAddAt(t *testing.T) {
    v := New()
    expectedLen := 1
    expectedElem := 1
    v.Add(expectedElem)
    actualLen := v.Len()
    actualElem := v.At(0)
    if expectedLen != actualLen {
        t.Errorf("Expected v.Len() = %d, insted %d", expectedLen, actualLen)
    }
    if expectedElem != actualElem {
        t.Errorf("Expected v.At(0) = %d, insted %d", expectedElem, actualElem)
    }
}

func TestContains(t *testing.T) {
    v := New()
    expected := 2
    v.Add(expected)
    if !v.Contains(expected) {
        t.Errorf("Expected v.Contains(%d) = true, insted false", expected)
    }
    notExp := 1
    if v.Contains(notExp) {
        t.Errorf("Expected v.Contains(%d) = false, insted true", notExp)
    }
}
