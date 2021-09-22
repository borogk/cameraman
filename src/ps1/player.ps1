. (Join-Path $PSScriptRoot 'settings.ps1')

$playerPk3Path = Join-Path $PSScriptRoot 'CameramanPlayer.pk3'

$baseArgs = '-file', "`"$playerPk3Path`""

if ($args[0] -eq 'load')
{
    $allArgs = $baseArgs + $args[1..($args.Length-1)]

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

    $allArgs += '+pukename', 'Cman_PlayInPlayer'
}
else
{
    $allArgs = $baseArgs + $args
}

$gzdoom = Start-Process -FilePath $GzdoomPath -ArgumentList $allArgs -PassThru
Wait-Process -Id $gzdoom.Id
