package client

var stringPtr func(s string) *string = func(s string) *string {
	return &s
}

var boolPtr func(b bool) *bool = func(b bool) *bool {
	return &b
}

var int32Ptr func(i int32) *int32 = func(i int32) *int32 {
	return &i
}
