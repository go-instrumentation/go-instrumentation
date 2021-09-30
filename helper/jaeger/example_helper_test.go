package jaeger

func ExampleSpan() {
	span, closer := Span("op", "service")
	if closer != nil {
		defer func() {
			closer.Close()
			IsGlobalRegistered = false
		}()
	}
	span.SetTag("tag", "tag")
	defer span.Finish()
}
