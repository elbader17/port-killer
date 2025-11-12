# Port Killer

`Port Killer` is a simple command-line tool to find and kill processes that are occupying (listening) on a specific TCP port.

It works by detecting the operating system and executing the necessary native commands (`lsof` on Linux/macOS and `netstat`/`taskkill` on Windows) to identify and kill the process.

---

## ⚡️ Usage

The tool requires elevated permissions to list and kill processes that may belong to other users or the system.

Assuming the `portKiller` binary (obtained from the repository) is in the current directory:

```bash
sudo ./portKiller
