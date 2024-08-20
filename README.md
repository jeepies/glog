# glog
(G)oLang (Log) is yet another GoLang logging package, but I added ✨colours✨

## Installation

```shell
go get -u github.com/jeepies/glog@v0.0.1
```

## Usage

**All logs will be trailed with a newline**

```go
// Prints a red error message
glog.Exception(err.Error())
// Out: [2024-08-20T09:19:42.759] [ERR] Failed to bind on port 1337

// Prints a green success message
glog.Success("Application is running on port %s", PORT)
// Out: [2024-08-20T09:19:42.758] [SUCC] Application is running on port 1337

// Prints a magenta information message
glog.Info("Attempting to bind to port %s", PORT)
// Out: [2024-08-20T09:19:42.759] [INFO] Attempting to bind to port 1337

// Prints a yellow warning message
glog.Warn("Port %s was already in use - Trying backup port")
// Out: [2024-08-20T09:19:42.759] [WARN] Port 1337 was already in use - Trying backup port

// Prints a generic white log message
glog.Log("GLog is pretty cool")
// Out: [2024-08-20T09:19:42.760] [LOG] GLog is pretty cool
```