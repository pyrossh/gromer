package app

// import "testing"

// func TestRange(t *testing.T) {
// 	testUpdate(t, []updateTest{
// 		{
// 			scenario: "range slice is updated",
// 			a: Div().Body(
// 				Range([]string{"hello", "world"}).Slice(func(i int) UI {
// 					src := []string{"hello", "world"}
// 					return Text(src[i])
// 				}),
// 			),
// 			b: Div().Body(
// 				Range([]string{"hello", "maxoo"}).Slice(func(i int) UI {
// 					src := []string{"hello", "maxoo"}
// 					return Text(src[i])
// 				}),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},
// 				{
// 					Path:     TestPath(0),
// 					Expected: Text("hello"),
// 				},
// 				{
// 					Path:     TestPath(1),
// 					Expected: Text("maxoo"),
// 				},
// 			},
// 		},
// 		{
// 			scenario: "range slice is updated to be empty",
// 			a: Div().Body(
// 				Range([]string{"hello", "world"}).Slice(func(i int) UI {
// 					src := []string{"hello", "world"}
// 					return Text(src[i])
// 				}),
// 			),
// 			b: Div().Body(
// 				Range([]string{}).Slice(func(i int) UI {
// 					src := []string{"hello", "maxoo"}
// 					return Text(src[i])
// 				}),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},
// 				{
// 					Path:     TestPath(0),
// 					Expected: nil,
// 				},
// 				{
// 					Path:     TestPath(1),
// 					Expected: nil,
// 				},
// 			},
// 		},
// 		{
// 			scenario: "range map is updated",
// 			a: Div().Body(
// 				Range(map[string]string{"key": "value"}).Map(func(k string) UI {
// 					src := map[string]string{"key": "value"}
// 					return Text(src[k])
// 				}),
// 			),
// 			b: Div().Body(
// 				Range(map[string]string{"key": "value"}).Map(func(k string) UI {
// 					src := map[string]string{"key": "maxoo"}
// 					return Text(src[k])
// 				}),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},
// 				{
// 					Path:     TestPath(0),
// 					Expected: Text("maxoo"),
// 				},
// 			},
// 		},
// 	})
// }

// func TestCondition(t *testing.T) {
// 	testUpdate(t, []updateTest{
// 		{
// 			scenario: "if is interpreted",
// 			a: Div().Body(
// 				If(false,
// 					H1(),
// 				),
// 			),
// 			b: Div().Body(
// 				If(true,
// 					H1(),
// 				),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},

// 				{
// 					Path:     TestPath(0),
// 					Expected: H1(),
// 				},
// 			},
// 		},
// 		{
// 			scenario: "if is not interpreted",
// 			a: Div().Body(
// 				If(true,
// 					H1(),
// 				),
// 			),
// 			b: Div().Body(
// 				If(false,
// 					H1(),
// 				),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},
// 				{
// 					Path:     TestPath(0),
// 					Expected: nil,
// 				},
// 			},
// 		},
// 		{
// 			scenario: "else if is interpreted",
// 			a: Div().Body(
// 				If(true,
// 					H1(),
// 				).ElseIf(false,
// 					H2(),
// 				),
// 			),
// 			b: Div().Body(
// 				If(false,
// 					H1(),
// 				).ElseIf(true,
// 					H2(),
// 				),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},

// 				{
// 					Path:     TestPath(0),
// 					Expected: H2(),
// 				},
// 			},
// 		},
// 		{
// 			scenario: "else if is not interpreted",
// 			a: Div().Body(
// 				If(false,
// 					H1(),
// 				).ElseIf(true,
// 					H2(),
// 				),
// 			),
// 			b: Div().Body(
// 				If(false,
// 					H1(),
// 				).ElseIf(false,
// 					H2(),
// 				),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},

// 				{
// 					Path:     TestPath(0),
// 					Expected: nil,
// 				},
// 			},
// 		},
// 		{
// 			scenario: "else is interpreted",
// 			a: Div().Body(
// 				If(false,
// 					H1(),
// 				).ElseIf(true,
// 					H2(),
// 				).Else(
// 					H3(),
// 				),
// 			),
// 			b: Div().Body(
// 				If(false,
// 					H1(),
// 				).ElseIf(false,
// 					H2(),
// 				).Else(
// 					H3(),
// 				),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},

// 				{
// 					Path:     TestPath(0),
// 					Expected: H3(),
// 				},
// 			},
// 		},
// 		{
// 			scenario: "else is not interpreted",
// 			a: Div().Body(
// 				If(false,
// 					H1(),
// 				).ElseIf(true,
// 					H2(),
// 				).Else(
// 					H3(),
// 				),
// 			),
// 			b: Div().Body(
// 				If(true,
// 					H1(),
// 				).ElseIf(false,
// 					H2(),
// 				).Else(
// 					H3(),
// 				),
// 			),
// 			matches: []TestUIDescriptor{
// 				{
// 					Path:     TestPath(),
// 					Expected: Div(),
// 				},

// 				{
// 					Path:     TestPath(0),
// 					Expected: H1(),
// 				},
// 			},
// 		},
// 	})
// }
