
# Create output folder
mkdir -force .\bin

# Build Plugin
go build -o .\bin\plugin.exe .\plugin\main.go

# Build Host
go build -o .\bin\main.exe .\main.go

# Run host
cd .\bin
.\main.exe
cd ..
