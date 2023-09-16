# GoPilot
Desktop app for automating/managing Linux tasks written in Go

## Installation

Download and execute the script that installs all the packages needed (Optional)

```
curl https://raw.githubusercontent.com/miguelfdez99/GoPilot/main/install-packages.sh -o install-packages.sh
chmod +x install-packages.sh
./install-packages.sh
```

Download the binary

```
curl -L https://github.com/miguelfdez99/GoPilot/releases/download/v0.6/GoPilot -o GoPilot
```

Give execution permissions

```
chmod +x GoPilot
```

Execute the binary

```
./GoPilot
```

You must execute the application with sudo for some functionalities to work

```
sudo ./GoPilot
```

If your system is using Wayland it will probably give you an error when running the application with sudo, in that case, try to execute the application like this:

```
sudo -EH ./GoPilot
```
### Packages

You can also install the application using the deb and rpm packages, you can download them from the release page.


## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
