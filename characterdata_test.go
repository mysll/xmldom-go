package dom_test

import (
	"testing"
	"xml/dom"
)

func TestCharacterDataSubstringData(t *testing.T) {
  d, _ := dom.ParseStringXml(`<parent>mom</parent>`)
  r := d.DocumentElement()
  txt := r.ChildNodes().Item(0).(*dom.Text)

  if txt.String() != "mom" {
	t.Errorf( "Did not get the correct text for the character data (%v)", txt.String() )
  }
  if txt.SubstringData(0, 1) != "m" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
  if txt.SubstringData(0, 2) != "mo" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
  if txt.SubstringData(0, 3) != "mom" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
  if txt.SubstringData(0, 4) != "mom" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
  if txt.SubstringData(1, 1) != "o" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
  if txt.SubstringData(1, 2) != "om" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
  if txt.SubstringData(1, 3) != "om" {
	t.Errorf( "Did not get the correct text for the character data substring" )
  }
}

