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
	golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
	golang.org/x/mobile v0.0.0-20220722155234-aaac322e2105 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
)
