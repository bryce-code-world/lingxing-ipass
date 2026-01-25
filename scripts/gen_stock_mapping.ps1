param(
  # WID -> DSCO warehouseCode, e.g. 26=YQN-CA,27=COOL-INTELOGICS
  [Parameter(Mandatory = $false)]
  [string[]]$WidToWarehouseCode,

  # LingXing SKU -> DSCO SKU, e.g. LXSKU-1=DSCOSKU-1
  [Parameter(Mandatory = $false)]
  [string[]]$SkuToDscoSku,

  # If present, read mappings from JSON files in the same directory as this script.
  # Default:
  # - wid_to_warehouse_code.json
  # - sku_to_dsco_sku.json
  [Parameter(Mandatory = $false)]
  [string]$WidToWarehouseCodeJsonPath = '',
  [Parameter(Mandatory = $false)]
  [string]$SkuToDscoSkuJsonPath = ''
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
if ([string]::IsNullOrWhiteSpace($WidToWarehouseCodeJsonPath)) {
  $WidToWarehouseCodeJsonPath = Join-Path $scriptDir 'wid_to_warehouse_code.json'
}
if ([string]::IsNullOrWhiteSpace($SkuToDscoSkuJsonPath)) {
  $SkuToDscoSkuJsonPath = Join-Path $scriptDir 'sku_to_dsco_sku.json'
}

function Parse-KeyValuePairs {
  param(
    [Parameter(Mandatory = $true)]
    [string[]]$Pairs,
    [Parameter(Mandatory = $true)]
    [string]$Name
  )

  $map = [ordered]@{}
  foreach ($p0 in $Pairs) {
    if ($null -eq $p0) { continue }

    # Allow passing a single argument that contains comma-separated pairs.
    $items = $p0.Split(',') | ForEach-Object { $_.Trim() } | Where-Object { $_ -ne '' }
    foreach ($p in $items) {
      if ($p -eq '') { continue }

      $parts = $p.Split('=', 2)
      if ($parts.Count -ne 2) {
        throw ("{0} invalid pair: {1} (expect key=value)" -f $Name, $p)
      }

      $k = $parts[0].Trim()
      $v = $parts[1].Trim()
      if ($k -eq '' -or $v -eq '') {
        throw ("{0} invalid pair: {1} (key/value must be non-empty)" -f $Name, $p)
      }

      $map[$k] = $v
    }
  }
  return $map
}

function Read-PairsIfMissing {
  param(
    [Parameter(Mandatory = $true)]
    [string]$Prompt
  )

  $line = Read-Host $Prompt
  if ([string]::IsNullOrWhiteSpace($line)) { return @() }
  return $line.Split(',') | ForEach-Object { $_.Trim() } | Where-Object { $_ -ne '' }
}

function Read-JsonMapFile {
  param(
    [Parameter(Mandatory = $true)]
    [string]$Path,
    [Parameter(Mandatory = $true)]
    [string]$Name
  )

  if (-not (Test-Path -LiteralPath $Path)) {
    return $null
  }

  $raw = Get-Content -LiteralPath $Path -Raw
  if ([string]::IsNullOrWhiteSpace($raw)) {
    return [ordered]@{}
  }

  $obj = $raw | ConvertFrom-Json
  if ($null -eq $obj) {
    return [ordered]@{}
  }
  if ($obj -isnot [pscustomobject] -and $obj -isnot [hashtable]) {
    throw ("{0} JSON must be an object map (string->string): {1}" -f $Name, $Path)
  }

  $map = [ordered]@{}
  foreach ($p in $obj.PSObject.Properties) {
    $k = [string]$p.Name
    $v = $p.Value
    if ($null -eq $v) { $v = '' }
    if ($v -isnot [string]) {
      $v = [string]$v
    }
    if ([string]::IsNullOrWhiteSpace($k) -or [string]::IsNullOrWhiteSpace($v)) {
      throw ("{0} JSON invalid entry: key/value must be non-empty (file={1}, key={2})" -f $Name, $Path, $k)
    }
    $map[$k] = $v
  }
  return $map
}

$widMap = Read-JsonMapFile -Path $WidToWarehouseCodeJsonPath -Name 'WidToWarehouseCode'
$skuMap = Read-JsonMapFile -Path $SkuToDscoSkuJsonPath -Name 'SkuToDscoSku'

if ($null -eq $widMap) {
  if (-not $WidToWarehouseCode) {
    $WidToWarehouseCode = Read-PairsIfMissing "Enter WID->warehouseCode pairs (comma-separated, e.g. 26=YQN-CA,27=COOL-INTELOGICS). Empty to skip"
  }
  $widMap = Parse-KeyValuePairs -Pairs $WidToWarehouseCode -Name 'WidToWarehouseCode'
}

if ($null -eq $skuMap) {
  if (-not $SkuToDscoSku) {
    $SkuToDscoSku = Read-PairsIfMissing "Enter SKU->DSCO SKU pairs (comma-separated, e.g. LXSKU-1=DSCOSKU-1). Empty to skip"
  }
  $skuMap = Parse-KeyValuePairs -Pairs $SkuToDscoSku -Name 'SkuToDscoSku'
}

$widJson = ($widMap | ConvertTo-Json -Compress)
$skuJson = ($skuMap | ConvertTo-Json -Compress)

Write-Host ""
Write-Host "# Copy these lines into .env (single quotes avoid escaping)"
if ($widMap.Count -gt 0) {
  Write-Host ("IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON='{0}'" -f $widJson)
} else {
  Write-Host "# IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON='{}'  # skipped"
}

if ($skuMap.Count -gt 0) {
  Write-Host ("IPASS_STOCK_SKU_TO_DSCO_SKU_JSON='{0}'" -f $skuJson)
} else {
  Write-Host "# IPASS_STOCK_SKU_TO_DSCO_SKU_JSON='{}'  # skipped"
}
