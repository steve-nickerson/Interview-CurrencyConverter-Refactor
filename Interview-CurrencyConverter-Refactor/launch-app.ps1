# Use this script to restore / build / run the dotnet BFF application easily
# Function to check if the .NET SDK is installed
function Test-DotNetSDKInstalled {
    $dotnetExe = Get-Command "dotnet" -ErrorAction SilentlyContinue
    if ($dotnetExe) {
        return $true  # .NET SDK is installed
    } else {
        return $false  # .NET SDK is not installed
    }
}

# Check if .NET SDK is installed
Write-Host "Check if .NET SDK is installed..."
if (Test-DotNetSDKInstalled) {
    Write-Host ".NET SDK is installed.`n"
} else {
    Write-Host ".NET SDK is not installed. Please install it from https://dotnet.microsoft.com/download. Exiting...`n"
    return $false
}

# Function to check if .NET SDK version 6 or higher is installed
function Test-DotNetSDK6OrHigherInstalled {
    $dotnetVersion = dotnet --version
    if ($dotnetVersion -match "^6\..*") {
        return $true  # .NET SDK version 6 or higher is installed
    } else {
        return $false  # .NET SDK version 6 or higher is not installed
    }
}

# Check if .NET SDK version 6 or higher is installed
if (Test-DotNetSDK6OrHigherInstalled) {
    Write-Host ".NET SDK version 6 or higher is installed, , proceeding with launch.`n`n"
} else {
    Write-Host ".NET SDK version 6 or higher is not installed. Please install it from https://dotnet.microsoft.com/download/dotnet/6.0."
    return $false
}

Write-Host 'Restoring packages...'
dotnet restore Interview-CurrencyConverter-Refactor.sln

Write-Host "`n`nBuilding application"
dotnet build ./Interview-CurrencyConverter-Refactor/Interview-CurrencyConverter-Refactor.csproj

Write-Host "`n`nRunning application"
dotnet run --project ./Interview-CurrencyConverter-Refactor/Interview-CurrencyConverter-Refactor.csproj
