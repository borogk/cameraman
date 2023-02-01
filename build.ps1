$srcPath = Join-Path $PSScriptRoot 'src'
$srcAcsPath = Join-Path $srcPath 'acs'
$srcEditorPath = Join-Path $srcPath 'editor'
$srcPlayerPath = Join-Path $srcPath 'player'
$srcPs1Path = Join-Path $srcPath 'ps1'
$srcBashPath = Join-Path $srcPath 'bash'

$outPath = Join-Path $PSScriptRoot 'out'
$outEditorPath = Join-Path $outPath 'editor'
$outPlayerPath = Join-Path $outPath 'player'

# Start from clean "out" folder
Remove-Item $outPath -Recurse -Force -ErrorAction Ignore
New-Item -Path $PSScriptRoot -Name 'out' -ItemType directory

# Compile ACS scripts
Get-ChildItem $srcAcsPath -Filter *.acs | Foreach-Object {
    $name = $_.BaseName
    $srcFilePath = Join-Path $srcAcsPath "$name.acs"
    $destFilePath = Join-Path $outPath "$name.o"
    $accArgs = "`"$srcFilePath`"", "`"$destFilePath`""
    & 'acc' $accArgs
    if (!$?)
    {
        exit
    }
}

# Build "CameramanEditor.pk3"
New-Item -Path $outEditorPath -Name 'acs' -ItemType directory
Copy-Item (Join-Path $outPath '*.o') -Destination (Join-Path $outEditorPath 'acs')
Copy-Item (Join-Path $srcEditorPath '*') -Destination $outEditorPath -Recurse
Copy-Item (Join-Path $PSScriptRoot 'LICENSE') -Destination $outEditorPath
Compress-Archive -Path (Join-Path $outEditorPath '*') -DestinationPath (Join-Path $outPath 'CameramanEditor.zip') -Force
Rename-Item -Path (Join-Path $outPath 'CameramanEditor.zip') -NewName 'CameramanEditor.pk3'

# Build "CameramanPlayer.pk3"
New-Item -Path $outPlayerPath -Name 'acs' -ItemType directory
Copy-Item (Join-Path $outPath 'player.o') -Destination (Join-Path $outPlayerPath 'acs')
Copy-Item (Join-Path $outPath 'geometry.o') -Destination (Join-Path $outPlayerPath 'acs')
Copy-Item (Join-Path $outPath 'common.o') -Destination (Join-Path $outPlayerPath 'acs')
Copy-Item (Join-Path $srcPlayerPath '*') -Destination $outPlayerPath
Copy-Item (Join-Path $PSScriptRoot 'LICENSE') -Destination $outPlayerPath
Compress-Archive -Path (Join-Path $outPlayerPath '*') -DestinationPath (Join-Path $outPath 'CameramanPlayer.zip') -Force
Rename-Item -Path (Join-Path $outPath 'CameramanPlayer.zip') -NewName 'CameramanPlayer.pk3'

# Clean the output folder from the archive sources
Remove-Item (Join-Path $outPath '*') -Recurse -Exclude *.pk3

# Copy helper PowerShell and bash scripts over to output
Copy-Item (Join-Path $srcPs1Path '*') -Destination $outPath
Copy-Item (Join-Path $srcBashPath '*') -Destination $outPath

# Copy LICENSE to output
Copy-Item (Join-Path $PSScriptRoot 'LICENSE') -Destination $outPath
