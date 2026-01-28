Param(
    [string[]]$service
)

$composeSvc = "E:\develop\go\personal\lingxing-ipass\docker\docker-compose.lingxing-ipass.service.yml"
$project = "lingxing-ipass"

# Full service list
$allServices = @(
    "ipass",
    "ipass_admin"
)

# 兼容历史输入：ipass-admin -> ipass_admin
$serviceAliases = @{
    "ipass-admin" = "ipass_admin"
}

# 外部命令失败时，必须显式中止（PowerShell 对外部程序默认不会抛异常）
function Invoke-DockerCompose {
    param(
        [Parameter(Mandatory = $true)]
        [string[]]$composeArgs
    )

    # 这里用“数组参数自动展开”传参，避免不同 PowerShell 版本对 splatting 的兼容性差异
    & docker compose -p $project -f $composeSvc $composeArgs
    if ($LASTEXITCODE -ne 0) {
        Write-Error "docker compose failed (exit=$LASTEXITCODE): docker compose -p $project -f $composeSvc $($composeArgs -join ' ')"
        exit $LASTEXITCODE
    }
}

# Determine target services
if (-not $service -or $service.Count -eq 0) {
    $targets = $allServices
    Write-Host "No service specified. Restarting all services: $($targets -join ', ')"
} else {
    $service = $service | ForEach-Object { if ($serviceAliases.ContainsKey($_)) { $serviceAliases[$_] } else { $_ } }
    $unknown = $service | Where-Object { $_ -notin $allServices }
    if ($unknown.Count -gt 0) {
        Write-Error "Unknown services detected: $($unknown -join ', '). Valid service names: $($allServices -join ', ')"
        exit 1
    }
    $targets = $service
    Write-Host "Restarting specified services: $($targets -join ', ')"
}

Write-Host "Stopping services: $($targets -join ', ')"
Invoke-DockerCompose (@("stop") + $targets)

Write-Host "Removing stopped containers: $($targets -join ', ')"
Invoke-DockerCompose (@("rm", "-f") + $targets)

Write-Host "Starting services in detached mode: $($targets -join ', ')"
Invoke-DockerCompose (@("up", "-d") + $targets)

if ($targets.Count -eq $allServices.Count) {
    Write-Host "All backend services have been restarted."
} else {
    Write-Host "Services restarted: $($targets -join ', ')"
}

Write-Host "Done."
