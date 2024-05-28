#!/usr/bin/env bash

# Function to check if Node.js is installed
check_node_installed() {
    if command -v node &> /dev/null; then
        return 0  # Node.js is installed
    else
        return 1  # Node.js is not installed
    fi
}

# Check if Node.js is installed
if check_node_installed; then
    echo "Node.js is installed."
else
    echo "Node.js is not installed. Please install it from https://nodejs.org/."
    exit 1
fi

#!/usr/bin/env bash

# Define the required Node.js major version
required_major_version="16"

# Function to check if Node.js major version meets the requirement
check_node_major_version() {
    local installed_version=$(node -v | cut -d "v" -f 2)
    local installed_major_version=$(echo "$installed_version" | cut -d "." -f 1)

    if [[ "$installed_major_version" -ge "$required_major_version" ]]; then
        return 0  # Node.js major version meets the requirement
    else
        return 1  # Node.js major version does not meet the requirement
    fi
}

# Check Node.js major version
if check_node_major_version; then
    echo "Node.js major version $required_major_version or higher is installed."
else
    echo "Node.js major version $required_major_version or higher is required. Please install it from https://nodejs.org/."
    exit 1
fi

# Use this script to install pkgs / build / run the React application easily

npm install

npm run dev