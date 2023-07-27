# Function to check if Node.js is installed
function Check-NodeInstalled {
    if (Get-Command "node" -ErrorAction SilentlyContinue) {
        return $true  # Node.js is installed
    } else {
        return $false  # Node.js is not installed
    }
}

# Check if Node.js is installed
if (Check-NodeInstalled) {
    Write-Host "Node.js is installed."
} else {
    Write-Host "Node.js is not installed. Please install it from https://nodejs.org/."
    return $false
}

# Define the required Node.js major version
$required_major_version = "16"

# Function to check if Node.js major version meets the requirement
function Check-NodeMajorVersion {
    $installed_version = (node -v).Substring(1)
    $installed_major_version = $installed_version.Split(".")[0]

    if ($installed_major_version -ge $required_major_version) {
        return $true  # Node.js major version meets the requirement
    } else {
        return $false  # Node.js major version does not meet the requirement
    }
}

# Check Node.js major version
if (Check-NodeMajorVersion) {
    Write-Host "Node.js major version $required_major_version or higher is installed."
} else {
    Write-Host "Node.js major version $required_major_version or higher is required. Please install it from https://nodejs.org/."
    return $false
}

# Use this script to install pkgs / build / run the React application easily

npm install

npm run dev
