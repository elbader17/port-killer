# Port Killer

`Port Killer` is a simple command-line tool to find and kill processes that are occupying (listening) on a specific TCP port.

It works by detecting the operating system and executing the necessary native commands (`lsof` on Linux/macOS and `netstat`/`taskkill` on Windows) to identify and kill the process.

---

## ⚡️ Usage

The tool requires elevated permissions to list and kill processes that may belong to other users or the system.

Assuming the `portKiller` binary (obtained from the repository) is in the current directory:

```bash
sudo ./portKiller
```

---

## 🚀 Create an Alias (Quick Access)

To avoid typing `sudo ./portKiller` (or the full path) every time, you can create an alias in your terminal.

> **Note:** The alias must use the absolute path to the `portKiller` binary. If the binary is in `/usr/local/bin`, use that path. If it is in the repository folder (e.g., `~/project/portKiller`), use that absolute path.

### For Zsh (modern macOS, some Linux distros)

1.  Open (or create) your Zsh configuration file:
    ```bash
    nano ~/.zshrc
    ```

2.  Add the following line to the end of the file (adjust the path according to where your binary is located):
    ```bash
    # Example if the binary is in the PATH (e.g., /usr/local/bin):
    alias freeport="sudo /usr/local/bin/portKiller"
    
    # Example if the binary is in the project folder:
    # alias freeport="sudo /home/your_user/project/portKiller/portKiller"
    ```

3.  Save the file and reload your configuration:
    ```bash
    source ~/.zshrc
    ```

### For Bash (common Linux, old macOS)

1.  Open your Bash configuration file:
    ```bash
    nano ~/.bashrc
    ```

2.  Add the following line to the end of the file (adjust the path):
    ```bash
    # Example if the binary is in the PATH (e.g., /usr/local/bin):
    alias freeport="sudo /usr/local/bin/portKiller"
    
    # Example if the binary is in the project folder:
    # alias freeport="sudo /home/your_user/project/portKiller/portKiller"
    ```

3.  Save the file and reload your configuration:
    ```bash
    source ~/.bashrc
    ```

### Alias Usage

Now, you can simply run the alias in your terminal:

```bash
freeport
```
