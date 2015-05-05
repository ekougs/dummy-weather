SHELL := /bin/bash
CURRENT_GO_PATH=$(shell ./make/find_go_path.sh)
BUILD_DIR="$(CURRENT_GO_PATH)"/bin/dummy_weather
BUILD_RESOURCES_DIR="$(BUILD_DIR)"/resources

all: pkg_dummy_weather clean

pkg_dummy_weather:dummy_weather pkg_resources
	@echo "Packaging application..."
	@mkdir -p $(BUILD_DIR)
	@cp dummy-weather $(BUILD_DIR)

pkg_resources:
	@echo "Packaging resources..."
	@mkdir -p $(BUILD_RESOURCES_DIR)
	@cp -r resources/* $(BUILD_RESOURCES_DIR)

dummy_weather: dummy_weather.go test
	@echo "Building exec..."
	@go build

test: check_init
	@echo "Running tests..."
	@go test -v ./...

clean:
	@rm dummy-weather

check_init:
	@if [ ! -d "$(CURRENT_GO_PATH)" ]; then echo "$(CURRENT_GO_PATH)"; exit 1; fi
