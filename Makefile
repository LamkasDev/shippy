SHIPPYPLATFORM=shippy64
SHIPPYDIR=$(USERPROFILE)\Desktop\shippy

GOOS=windows
GOARCH=amd64
GOTAGS=$(SHIPPYPLATFORM),shippydebug

.PHONY: buildshippy build install runshippy run clean
buildshippy:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(SHIPPYPLATFORM)/shippy.exe -tags $(GOTAGS) cmd/client/main.go

build: buildshippy

install: build
	@if exist "$(SHIPPYDIR)\bin\$(SHIPPYPLATFORM)" rmdir /S /Q "$(SHIPPYDIR)\bin\$(SHIPPYPLATFORM)"
	@xcopy "build\$(SHIPPYPLATFORM)" "$(SHIPPYDIR)\bin\$(SHIPPYPLATFORM)\" /E /C /I >nul
	@if exist"$(SHIPPYDIR)\resources" rmdir /S /Q "$(SHIPPYDIR)\resources"
	@xcopy "resources" "$(SHIPPYDIR)\resources\" /E /C /I >nul

runshippy: buildshippy
	@if not exist "$(SHIPPYDIR)\bin\dev" mkdir "$(SHIPPYDIR)\bin\dev"
	@copy "build\$(SHIPPYPLATFORM)\shippy.exe" "$(SHIPPYDIR)\bin\dev\shippy.exe" >nul
	@cd "build\$(SHIPPYPLATFORM)" && .\shippy.exe

run: build
	@if exist "$(SHIPPYDIR)\bin\dev" rmdir /S /Q "$(SHIPPYDIR)\bin\dev"
	@xcopy "build\$(SHIPPYPLATFORM)" "$(SHIPPYDIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(SHIPPYPLATFORM)" && .\shippy.exe

clean:
	@if exist "build" rmdir /S /Q build