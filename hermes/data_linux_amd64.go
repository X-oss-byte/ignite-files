package hermes

import _ "embed" // embed is required for binary embedding.

//go:embed hermes-v1.6.0-x86_64-unknown-linux-gnu.tar.gz
var binaryCompressed []byte
