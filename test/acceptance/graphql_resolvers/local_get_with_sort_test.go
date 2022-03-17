//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/semi-technologies/weaviate/test/acceptance/helper"
	"github.com/stretchr/testify/assert"
)

func gettingObjectsWithSort(t *testing.T) {
	t.Run("simple sort", func(t *testing.T) {
		query := `
		{
			Get {
				City(
					sort: [{
						path: ["%s"]
						order: %s
					}]
				) {
					name
				}
			}
		}
		`
		tests := []struct {
			name            string
			property, order string
			expected        []interface{}
		}{
			{
				name:     "sort by name asc",
				property: "name",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Rotterdam"},
				},
			},
			{
				name:     "sort by name desc",
				property: "name",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
				},
			},
			{
				name:     "sort by population asc",
				property: "population",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
				},
			},
			{
				name:     "sort by population desc",
				property: "population",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by isCapital asc",
				property: "isCapital",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by isCapital desc",
				property: "isCapital",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
				},
			},
			{
				name:     "sort by cityArea asc",
				property: "cityArea",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
				},
			},
			{
				name:     "sort by cityArea desc",
				property: "cityArea",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
				},
			},

			{
				name:     "sort by cityRights asc",
				property: "cityRights",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
				},
			},
			{
				name:     "sort by cityRights desc",
				property: "cityRights",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by timezones asc",
				property: "timezones",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
				},
			},
			{
				name:     "sort by timezones desc",
				property: "timezones",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by museums asc",
				property: "museums",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Amsterdam"},
				},
			},
			{
				name:     "sort by museums desc",
				property: "museums",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by history asc",
				property: "history",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Dusseldorf"},
				},
			},
			{
				name:     "sort by history desc",
				property: "history",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by phoneNumber asc",
				property: "phoneNumber",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
				},
			},
			{
				name:     "sort by phoneNumber desc",
				property: "phoneNumber",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name:     "sort by location asc",
				property: "location",
				order:    "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
				},
			},
			{
				name:     "sort by location desc",
				property: "location",
				order:    "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Berlin"},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := AssertGraphQL(t, helper.RootAuth, fmt.Sprintf(query, tt.property, tt.order))
				got := result.Get("Get", "City").AsSlice()
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("sort objects got = %v, want %v", got, tt.expected)
				}
			})
		}
	})

	t.Run("complex sort", func(t *testing.T) {
		query := `
		{
			Get {
				City(
					%s
				) {
					name
				}
			}
		}
		`
		buildSort := func(path []string, order string) string {
			pathArgs := make([]string, len(path))
			for i := range path {
				pathArgs[i] = fmt.Sprintf("\"%s\"", path[i])
			}
			return fmt.Sprintf("{path:[%s] order:%s}", strings.Join(pathArgs, ","), order)
		}
		buildSortFilter := func(sort []string) string {
			return fmt.Sprintf("sort:[%s]", strings.Join(sort, ","))
		}
		tests := []struct {
			name     string
			sort     []string
			expected []interface{}
		}{
			{
				name: "sort by name and population asc",
				sort: []string{
					buildSort([]string{"name", "population"}, "asc"),
				},
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Berlin"},
				},
			},
			{
				name: "sort by name asc and population desc",
				sort: []string{
					buildSort([]string{"name"}, "asc"),
					buildSort([]string{"population"}, "desc"),
				},
				expected: []interface{}{
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name: "sort by name and population desc",
				sort: []string{
					buildSort([]string{"name", "population"}, "desc"),
				},
				expected: []interface{}{
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
			{
				name: "sort by name and population and phoneNumber asc",
				sort: []string{
					buildSort([]string{"name", "population", "phoneNumber"}, "asc"),
				},
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Rotterdam"},
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Berlin"},
					map[string]interface{}{"name": "Amsterdam"},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := AssertGraphQL(t, helper.RootAuth, fmt.Sprintf(query, buildSortFilter(tt.sort)))
				got := result.Get("Get", "City").AsSlice()
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("sort objects got = %v, want %v", got, tt.expected)
				}
			})
		}
	})

	t.Run("sort with where", func(t *testing.T) {
		query := `
		{
			Get {
				City(
					sort: [{
						path: ["location"]
						order: %s
					}]
					where: {
						operator: Or,
						operands: [
							{valueString: "6ffb03f8-a853-4ec5-a5d8-302e45aaaf13", path: ["id"], operator: Equal},
							{valueString: "823abeca-eef3-41c7-b587-7a6977b08003", path: ["id"], operator: Equal}
					]}
				) {
					name
				}
			}
		}
		`
		tests := []struct {
			name     string
			order    string
			expected []interface{}
		}{
			{
				name:  "location asc",
				order: "asc",
				expected: []interface{}{
					map[string]interface{}{"name": "Null Island"},
					map[string]interface{}{"name": "Dusseldorf"},
				},
			},
			{
				name:  "location desc",
				order: "desc",
				expected: []interface{}{
					map[string]interface{}{"name": "Dusseldorf"},
					map[string]interface{}{"name": "Null Island"},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := AssertGraphQL(t, helper.RootAuth, fmt.Sprintf(query, tt.order))
				got := result.Get("Get", "City").AsSlice()
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("sort objects got = %v, want %v", got, tt.expected)
				}
			})
		}
	})

	t.Run("sort with where with non-existent-uuid", func(t *testing.T) {
		query := `
		{
			Get {
				City(
					sort: [{
						path: ["location"]
						order: asc
					}]
					where: {
						valueString: "non-existent-uuid", path: ["id"], operator: Equal
					}
				) {
					name
				}
			}
		}
		`
		result := AssertGraphQL(t, helper.RootAuth, query)
		got := result.Get("Get", "City").AsSlice()
		assert.Empty(t, got)
	})
}
