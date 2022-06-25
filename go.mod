module github.com/dayvillefire/qv-tones-player

go 1.18

replace (
	github.com/jbuchbinder/beep => ../../jbuchbinder/beep
	github.com/jbuchbinder/beep/generators => ../../jbuchbinder/beep/generators
	github.com/jbuchbinder/beep/speaker => ../../jbuchbinder/beep/speaker
)

require (
	github.com/jbuchbinder/beep v1.1.1
	github.com/jbuchbinder/beep/generators v1.1.1
	github.com/jbuchbinder/beep/speaker v1.1.1
)

require (
	github.com/hajimehoshi/oto v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20190306152737-a1d7652674e8 // indirect
	golang.org/x/image v0.0.0-20190227222117-0694c2d4d067 // indirect
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6 // indirect
	golang.org/x/sys v0.0.0-20190626150813-e07cf5db2756 // indirect
)
