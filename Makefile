VERSION := v1.4.0-alpha7

GO_BIN ?= go
ACC_BIN ?= acc
ACC_INCLUDE ?= /usr/local/share/acc/

# Primary recipe - performs clean build for multiple platforms
build: clean test
	$(MAKE) build-windows
	$(MAKE) build-linux
	$(MAKE) build-macos
	@echo
	@echo "Successfully built Cameraman ${VERSION}"

# Cleans build directory
clean:
	rm -rf ./build/

# Runs tests
test:
	$(GO_BIN) test ./launcher/

# Windows x64 build
build-windows: OUT_ZIP=cameraman-$(VERSION)-windows-amd64.zip
build-windows: OUT_BUILD_DIR=./build/windows
build-windows: OUT_BIN_EXT=.exe
build-windows: GOOS=windows
build-windows: GOARCH=amd64
build-windows: build-one

# Linux x64 build
build-linux: OUT_ZIP=cameraman-$(VERSION)-linux-amd64.zip
build-linux: OUT_BUILD_DIR=./build/linux
build-linux: OUT_BIN_EXT=
build-linux: GOOS=linux
build-linux: GOARCH=amd64
build-linux: build-one

# macOS ARM build
build-macos: OUT_ZIP=cameraman-$(VERSION)-macos-arm64.zip
build-macos: OUT_BUILD_DIR=./build/macos
build-macos: OUT_BIN_EXT=
build-macos: GOOS=darwin
build-macos: GOARCH=arm64
build-macos: build-one

# Builds for one platform (reused by specific platform recipes)
build-one: editor-pk3 player-pk3 cm-editor cm-player
	cp ./LICENSE $(OUT_BUILD_DIR)
	cp ./build/CameramanEditor.pk3 $(OUT_BUILD_DIR)
	cp ./build/CameramanPlayer.pk3 $(OUT_BUILD_DIR)
	cd $(OUT_BUILD_DIR); zip ../$(OUT_ZIP) *

# Cross-platform compile macro
define go-build
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	$(GO_BIN) build \
		-trimpath \
		-ldflags="-s -w" \
		-o $(OUT_BUILD_DIR)/$(1)$(OUT_BIN_EXT) \
		$(2)
endef

# Compiles cm-editor
cm-editor:
	mkdir -p $(OUT_BUILD_DIR)
	$(call go-build,cm-editor,./cmd/cm-editor/main.go)

# Compiles cm-player
cm-player:
	mkdir -p $(OUT_BUILD_DIR)
	$(call go-build,cm-player,./cmd/cm-player/main.go)

# Compiles ACS sources that go into PK3s
compile-acs:
	mkdir -p ./build/acs/
	$(ACC_BIN) -i $(ACC_INCLUDE) ./zdoom/acs/common.acs ./build/acs/common.o
	$(ACC_BIN) -i $(ACC_INCLUDE) ./zdoom/acs/geometry.acs ./build/acs/geometry.o
	$(ACC_BIN) -i $(ACC_INCLUDE) ./zdoom/acs/editor.acs ./build/acs/editor.o
	$(ACC_BIN) -i $(ACC_INCLUDE) ./zdoom/acs/player.acs ./build/acs/player.o

# Packages CameramanEditor.pk3
editor-pk3: compile-acs
	mkdir -p ./build/editor/acs
	cp ./build/acs/common.o ./build/editor/acs/
	cp ./build/acs/geometry.o ./build/editor/acs/
	cp ./build/acs/editor.o ./build/editor/acs/
	cp ./build/acs/player.o ./build/editor/acs/
	cp -r ./zdoom/editor/* ./build/editor/
	cp ./LICENSE ./build/editor/
	cd ./build/editor/; zip -r ../CameramanEditor.pk3 *

# Packages CameramanPlayer.pk3
player-pk3: compile-acs
	mkdir -p ./build/player/acs
	cp ./build/acs/common.o ./build/player/acs/
	cp ./build/acs/geometry.o ./build/player/acs/
	cp ./build/acs/player.o ./build/player/acs/
	cp ./zdoom/player/* ./build/player/
	cp ./LICENSE ./build/player/
	cd ./build/player/; zip -r ../CameramanPlayer.pk3 *
