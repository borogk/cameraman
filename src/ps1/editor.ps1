$baseArgs = '+freelook', '1', '+noclip2', '+notarget', '-file', "$PSScriptRoot\CameramanEditor.pk3", '+logfile', "$PSScriptRoot\editor.log"

if ($args[0] -eq 'load')
{
	$allArgs = $baseArgs + $args[2..($args.Length-1)]
	
	$level = 1
	foreach ($line in Get-Content $args[1]) 
	{
		$match = $line | Select-String -Pattern '^([\w\d]+) = (.*)$'
		$name = $match.Matches.Groups[1].Value
		$value = $match.Matches.Groups[2].Value
		
		if ($name -ne $null)
		{
			if ($name -eq 'level')
			{
				$level = $value
			}
			else
			{
				$allArgs += '"+cman_{0} {1}"' -f $name, $value
			}
		}
	}
	
	$allArgs += '-warp', $level, '+pukename', 'Cman_WarpToPath'
}
else
{
	$allArgs = $baseArgs + $args
}

$gzdoom = Start-Process -FilePath 'gzdoom.exe' -ArgumentList $allArgs -PassThru
Wait-Process -Id $gzdoom.Id

$exportingFile = ''
foreach ($line in Get-Content "$PSScriptRoot\editor.log") 
{
	if ($exportingFile -ne '')
	{
		if ($line -ne '--- END CAMERAMAN ---')
		{
			Add-Content -Path $exportingFile -Value $line
		}
		else
		{
			Write-Host "Exported $exportingFile"
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
		
		$exportingFile = "$PSScriptRoot\export-{0:0000}.cman" -f ($maxNum + 1)
    }
}

Remove-Item "$PSScriptRoot\editor.log" -Force
