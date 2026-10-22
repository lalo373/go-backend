// Package static embeds the frontend assets into the binary so the server
// can serve them regardless of which working directory it runs from
// (repo root, container, or Vercel function bundle).
package static

import _ "embed"

//go:embed index.html
var Index []byte
