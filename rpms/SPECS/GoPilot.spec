Name:           GoPilot
Version:        0.5
Release:        1%{?dist}
Summary:        GoPilot

License:        GNU General Public License v3.0
URL:            https://github.com/miguelfdez99/GoPilot/
Source0:        GoPilot.tar.gz

Requires:       rsync, diffutils, tar, openssh-clients, cronie, shadow-utils, coreutils, procps-ng, systemd, fdupes, webkit2gtk3

%description
GoPilot is a desktop application that allows you to automates and manages your Linux system

%prep
%setup -q -n GoPilot

%install
rsync -a %buildroot/
install -d %buildroot/usr/bin/
install -m 755 usr/bin/GoPilot %buildroot/usr/bin/

%files
/usr/bin/GoPilot

%changelog
* Fri Sep 02 2023 Miguel Fernandez <miguelfdez99@gmail.com> - 0.5-1
- Initial RPM package

