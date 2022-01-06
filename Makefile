setup:
	brew install goreleaser

release:
	goreleaser release --rm-dist
