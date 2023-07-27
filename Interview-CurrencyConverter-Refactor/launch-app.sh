#!/usr/bin/env bash

# Function to check if the .NET SDK is installed
check_dotnet_sdk() {
    if command -v dotnet &> /dev/null; then
        return 0  # .NET SDK is installed
    else
        return 1  # .NET SDK is not installed
    fi
}

# Check if .NET SDK is installed
echo -e "Check if .NET SDK is installed..."
if check_dotnet_sdk; then
    echo -e ".NET SDK is installed.\n"
else
    echo -e ".NET SDK is not installed. Please install it from https://dotnet.microsoft.com/download. Exiting...\n"
    exit 1
fi

# Function to check if .NET SDK version 6 or higher is installed
check_dotnet_sdk_6_or_higher_installed() {
    local installed_version=$(dotnet --version)
    if [[ "$installed_version" =~ ^6\..* ]]; then
        return 0  # .NET SDK version 6 or higher is installed
    else
        return 1  # .NET SDK version 6 or higher is not installed
    fi
}

# Check if .NET SDK version 6 or higher is installed
if check_dotnet_sdk_6_or_higher_installed; then
    echo -e ".NET SDK version 6 or higher is installed, proceeding with launch.\n\n"
else
    echo -e ".NET SDK version 6 or higher is not installed. Please install it from https://dotnet.microsoft.com/download/dotnet/6.0."
    exit 1
fi

# Use this script to restore / build / run the dotnet BFF application easily
echo -e "Restoring packages..."
dotnet restore Interview-CurrencyConverter-Refactor.sln

echo -e "\n\nBuilding application"
dotnet build ./Interview-CurrencyConverter-Refactor/Interview-CurrencyConverter-Refactor.csproj

echo -e "\n\nRunning application"
dotnet run --project ./Interview-CurrencyConverter-Refactor/Interview-CurrencyConverter-Refactor.csproj

