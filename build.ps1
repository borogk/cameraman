# Start from clean "out" folder
Remove-Item "$PSScriptRoot\out" -Recurse -Force
New-Item -Path $PSScriptRoot -Name out -ItemType directory

# Compile ACS scripts
Get-ChildItem "$PSScriptRoot\src\acs" -Filter *.acs | Foreach-Object {
    $name = $_.BaseName
    $args = "$PSScriptRoot\src\acs\$name.acs", "$PSScriptRoot\out\$name.o"
    & 'acc' $args
    if (!$?)
    {
        exit
    }
}

# Build "CameramanEditor.pk3"
New-Item -Path "$PSScriptRoot\out\editor" -Name acs -ItemType directory
Copy-Item "$PSScriptRoot\out\*.o" -Destination "$PSScriptRoot\out\editor\acs"
Copy-Item "$PSScriptRoot\src\editor\*" -Destination "$PSScriptRoot\out\editor" -Recurse
Compress-Archive -Path "$PSScriptRoot\out\editor\*" -DestinationPath "$PSScriptRoot\out\editor.zip" -Force
Rename-Item -Path "$PSScriptRoot\out\editor.zip" -NewName CameramanEditor.pk3

# Build "CameramanPlayer.pk3"
New-Item -Path "$PSScriptRoot\out\player" -Name acs -ItemType directory
Copy-Item "$PSScriptRoot\out\player.o" -Destination "$PSScriptRoot\out\player\acs"
Copy-Item "$PSScriptRoot\out\geometry.o" -Destination "$PSScriptRoot\out\player\acs"
Copy-Item "$PSScriptRoot\out\common.o" -Destination "$PSScriptRoot\out\player\acs"
Copy-Item "$PSScriptRoot\src\player\*" -Destination "$PSScriptRoot\out\player"
Compress-Archive -Path "$PSScriptRoot\out\player\*" -DestinationPath "$PSScriptRoot\out\player.zip" -Force
Rename-Item -Path "$PSScriptRoot\out\player.zip" -NewName CameramanPlayer.pk3

# Copy helper PowerShell scripts over to output
Copy-Item "$PSScriptRoot\src\ps1\*" -Destination "$PSScriptRoot\out"

# Cleanup the output folder
Remove-Item "$PSScriptRoot\out\*" -Recurse -Exclude *.pk3, *.ps1
