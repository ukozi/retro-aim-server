Title: WebAPI foundation/core wiring

Why
- Establish a minimal, reviewable base for WebAPI without feature endpoints
- Break import cycles and set up deps (AMF0/AMF3) cleanly

Changes
- Minimal server/webapi/server.go with only root route; feature routes follow later
- Add local AMF3 module via internal/goAMF3 and replace in go.mod; require github.com/speps/go-amf
- Bring in essential state WebAPI types to support presence/bridges while keeping handlers out
- No buddyfeed/vanity/chat endpoints in this PR

Notes
- Subsequent PRs will add handlers and route wiring per feature
- See docs/open_api/webapi.yml for endpoint specs
