package misc

import json "github.com/json-iterator/go"

var (
	// JSONCodec is a static JSON serializer and JSON deserializer handler.
	JSONCodec = json.ConfigFastest
)
