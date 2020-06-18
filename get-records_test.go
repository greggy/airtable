// Copyright © 2020 Mike Berezin
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package airtable

import (
	"testing"
)

func TestGetRecordsConfig_Do(t *testing.T) {
	table := testTable(t)
	table.client.baseURL = mockResponse("get_records_with_filter.json").URL
	sortQuery1 := struct {
		fieldName string
		direction string
	}{"Field1", "desc"}
	sortQuery2 := struct {
		fieldName string
		direction string
	}{"Field2", "asc"}

	records, err := table.GetRecords().
		FromView("view_1").
		WithFilterFormula("AND({Field1}='value_1',NOT({Field2}='value_2'))").
		WithSort(sortQuery1, sortQuery2).
		ReturnFields("Field1", "Field2").
		InStringFormat("Europe/Moscow", "ru").
		Do()
	if err != nil {
		t.Errorf("there should not be an err, but was: %v", err)
	}
	if len(records.Records) != 3 {
		t.Errorf("there should be 3 records, but was %v", len(records.Records))
	}
}
