$srcPath = Join-Path $PSScriptRoot 'src'
$srcAcsPath = Join-Path $srcPath 'acs'
$srcEditorPath = Join-Path $srcPath 'editor'
$srcPlayerPath = Join-Path $srcPath 'player'
$srcPs1Path = Join-Path $srcPath 'ps1'

$outPath = Join-Path $PSScriptRoot 'out'
$outEditorPath = Join-Path $outPath 'editor'
$outPlayerPath = Join-Path $outPath 'player'

# Start from clean "out" folder
Remove-Item $outPath -Recurse -Force -ErrorAction Ignore
New-Item -Path $PSScriptRoot -Name 'out' -ItemType directory

# Compile ACS scripts
Get-ChildItem $srcAcsPath -Filter *.acs | Foreach-Object {
    $name = $_.BaseName
    $accArgs = (Join-Path $srcAcsPath "$name.acs"), (Join-Path $outPath "$name.o")
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
Compress-Archive -Path (Join-Path $outEditorPath '*') -DestinationPath (Join-Path $outPath 'CameramanEditor.zip') -Force
Rename-Item -Path (Join-Path $outPath 'CameramanEditor.zip') -NewName 'CameramanEditor.pk3'

# Build "CameramanPlayer.pk3"
New-Item -Path $outPlayerPath -Name 'acs' -ItemType directory
Copy-Item (Join-Path $outPath 'player.o') -Destination (Join-Path $outPlayerPath 'acs')
Copy-Item (Join-Path $outPath 'geometry.o') -Destination (Join-Path $outPlayerPath 'acs')
Copy-Item (Join-Path $outPath 'common.o') -Destination (Join-Path $outPlayerPath 'acs')
Copy-Item (Join-Path $srcPlayerPath '*') -Destination $outPlayerPath
Compress-Archive -Path (Join-Path $outPlayerPath '*') -DestinationPath (Join-Path $outPath 'CameramanPlayer.zip') -Force
Rename-Item -Path (Join-Path $outPath 'CameramanPlayer.zip') -NewName 'CameramanPlayer.pk3'

# Copy helper PowerShell scripts over to output
Copy-Item (Join-Path $srcPs1Path '*') -Destination $outPath

# Cleanup the output folder
Remove-Item (Join-Path $outPath '*') -Recurse -Exclude *.pk3, *.ps1
