. (Join-Path $PSScriptRoot 'settings.ps1')

$editorLogPath = Join-Path $PSScriptRoot 'editor.log'

$baseArgs = '+freelook', '1', '+noclip2', '+notarget', '-file', (Join-Path $PSScriptRoot 'CameramanEditor.pk3'), '+logfile', $editorLogPath

if ($args[0] -eq 'load')
{
    $allArgs = $baseArgs + $args[2..($args.Length-1)]

    foreach ($line in Get-Content $args[1])
    {
        $match = $line | Select-String -Pattern '^([\w\d]+) = (.*)$'
        $name = $match.Matches.Groups[1].Value
        $value = $match.Matches.Groups[2].Value

        if ($null -ne $name)
        {
            $allArgs += '"+cman_{0} {1}"' -f $name, $value
        }
    }

    $allArgs += '+pukename', 'Cman_WarpToPath'
}
else
{
    $allArgs = $baseArgs + $args
}

$gzdoom = Start-Process -FilePath $GzdoomPath -ArgumentList $allArgs -PassThru
Wait-Process -Id $gzdoom.Id

$exportingFile = ''
foreach ($line in Get-Content $editorLogPath)
{
    if ($exportingFile -ne '')
    {
        if ($line -ne '--- END CAMERAMAN ---')
        {
            Add-Content -Path $exportingFile -Value $line
        }
        else
        {
            Write-Output "Exported $exportingFile"
            $exportingFile = ''
        }
    }
    elseif ($line -eq '--- BEGIN CAMERAMAN ---')
    {
        $maxNum = 0
        Get-ChildItem $PSScriptRoot -Filter *.cman | Foreach-Object {
            $groups = ($_.BaseName | Select-String -Pattern '^export-(\d{4})$').Matches.Groups
            if ($groups -ne $null)
            {
                $num = [int]$groups[1].Value
                if ($num -gt $maxNum)
                {
                    $maxNum = $num
                }
            }
        }

        $exportingFile = Join-Path $PSScriptRoot ("export-{0:0000}.cman" -f ($maxNum + 1))
    }
}

Remove-Item $editorLogPath -Force
