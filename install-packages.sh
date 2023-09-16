#!/bin/bash

# Packages for Debian/Ubuntu
DEB_PACKAGES=(rsync diffutils tar openssh-client cron adduser coreutils procps libx11-6 systemd fdupes libwebkit2gtk-4.0-37 kdialog)

# Packages for RedHat/Fedora/CentOS
RPM_PACKAGES=(rsync diffutils tar openssh-clients cronie shadow-utils coreutils procps-ng systemd fdupes webkit2gtk3 kdialog)

# Packages for Arch/Manjaro
ARCH_PACKAGES=(rsync diffutils tar openssh cronie shadow coreutils procps-ng systemd fdupes webkit2gtk kdialog)

# Function to install packages based on the package management system
install_packages() {
    if command -v apt > /dev/null 2>&1; then
        sudo apt-get update
        sudo apt-get install -y "${DEB_PACKAGES[@]}"
    elif command -v dnf > /dev/null 2>&1; then
        sudo dnf install -y "${RPM_PACKAGES[@]}"
    elif command -v yum > /dev/null 2>&1; then
        sudo yum install -y "${RPM_PACKAGES[@]}"
    elif command -v pacman > /dev/null 2>&1; then
        sudo pacman -Syu --noconfirm "${ARCH_PACKAGES[@]}"
    else
        echo "Your system's package manager is not supported by this script."
        exit 1
    fi
}

install_packages
