$ErrorActionPreference = "Stop"

# ---- 1. Setup variables ----
$Binary = "tuner1.exe"

if (-not $env:INSTALL_DIR) {
    $InstallDir = Join-Path $env:USERPROFILE ".local\bin"
} else {
    $InstallDir = $env:INSTALL_DIR
}

$BinaryPath = Join-Path $InstallDir $Binary

# ---- 2. Check if tuner1 exists ----
if (-not (Test-Path $BinaryPath)) {
    Write-Host "$Binary not found in $InstallDir; nothing to uninstall."
    exit 0
}

# ---- 3. Remove the binary ----
try {
    Remove-Item -Force $BinaryPath
    Write-Host "$Binary removed from $InstallDir."
} catch {
    Write-Error "Failed to remove $BinaryPath. You may need admin privileges."
    exit 1
}

# ---- 4. Config directory info ----
$ConfigDir = Join-Path $env:APPDATA "tuner1"
$StandardsFile = Join-Path $ConfigDir "standards.txt"

Write-Host ""
Write-Host "To remove the standards.txt file as well, run:"
Write-Host "  Remove-Item -Force `"$StandardsFile`""
Write-Host ""
