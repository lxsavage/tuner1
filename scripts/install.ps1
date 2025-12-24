$ErrorActionPreference = "Stop"

# ---- 1. Repository info ----
$Repo    = "lxsavage/tuner1"
$Binary  = "tuner1"
$OS      = "windows"

# ---- 2. Determine architecture ----
$Arch = $env:PROCESSOR_ARCHITECTURE.ToLower()

switch ($Arch) {
    "amd64"    { $Arch = "amd64" }
    "arm64"    { $Arch = "amd64" } # No arm build yet, will just use the x64 one for now
    default    { throw "Unsupported architecture: $Arch" }
}

$Asset = "$Binary-$OS-$Arch.exe"

# ---- 3. Determine install directory ----
if (-not $env:INSTALL_DIR) {
    $InstallDir = Join-Path $env:USERPROFILE ".local\bin"
} else {
    $InstallDir = $env:INSTALL_DIR
}
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

# ---- 4. Get latest release asset URL from GitHub ----
$ApiUrl = "https://api.github.com/repos/$Repo/releases/latest"

try {
    $Response = Invoke-RestMethod -Uri $ApiUrl -UseBasicParsing
} catch {
    throw "Failed to fetch release info from $ApiUrl"
}

$AssetUrl = $Response.assets | Where-Object { $_.name -eq $Asset } | Select-Object -First 1 -ExpandProperty browser_download_url

if (-not $AssetUrl) {
    throw "Release asset $Asset not found in latest release"
}

# ---- 5. Download & install ----
$Tmp = New-Item -ItemType Directory -Path ([System.IO.Path]::GetTempPath() + [System.Guid]::NewGuid()) -Force
$TmpFile = Join-Path $Tmp.FullName $Asset

Write-Host "Downloading $Asset from $AssetUrl..."
Invoke-WebRequest -Uri $AssetUrl -OutFile $TmpFile -UseBasicParsing

$DestFile = Join-Path $InstallDir "$Binary.exe"
Move-Item -Force $TmpFile $DestFile

# ---- 6. Pull standards.txt into tuner1 config dir if not already there ----
$ConfigDir = Join-Path $env:APPDATA "tuner1"
New-Item -ItemType Directory -Force -Path $ConfigDir | Out-Null

$StandardsFile = Join-Path $ConfigDir "standards.txt"

if (-not (Test-Path $StandardsFile)) {
    $StandardsUrl = "https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/config/standards.txt"
    Write-Host "Downloading standards.txt from $StandardsUrl..."
    Invoke-WebRequest -Uri $StandardsUrl -OutFile $StandardsFile -UseBasicParsing
} else {
    Write-Host "Standards file already exists locally. Skipping download."
}

# ---- 7. Final message ----
Write-Host "`n$Binary installed to: $DestFile"
Write-Host "standards.txt located at: $StandardsFile"

# ---- 8. PATH check ----
$PathDirs = $env:Path -split ';'
if ($PathDirs -notcontains $InstallDir) {
    Write-Warning "`n$InstallDir is not in your PATH."
    Write-Host "To add it to your current session, run:"
    Write-Host "  setx PATH `"$($env:Path);$InstallDir`""
}
