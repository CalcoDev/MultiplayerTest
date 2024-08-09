all: build_working_gdextension

WORKING_GDEXT_FILE := working-godot.dll

WORKING_GO_DIR := go/working-godot
WORKING_GDEXT_DIR := extensions/working-godot/

# TODO(calco): Make the build work on other OSs except Windows lmfao.
build_working_gdextension:
	cd $(WORKING_GO_DIR) && go build -o $(WORKING_GDEXT_FILE) -buildmode=c-shared
	mv $(WORKING_GO_DIR)/$(WORKING_GDEXT_FILE) $(WORKING_GDEXT_DIR)