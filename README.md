# LINAM

> **Lin**ux **A**pplication **M**anager

Manage applications on Linux servers through a modern web interface using **systemd**, not containers.

## Why?

Most Linux control panels rely on Docker to deploy and manage applications. While containers are powerful, they also introduce an additional layer of complexity and abstraction that isn't always necessary.

**LINAM** takes a different approach.

Instead of running applications inside containers, LINAM deploys and manages them as native **systemd** services. This keeps applications closer to the operating system, reduces resource usage, and gives administrators full control over their servers using standard Linux tools.

LINAM doesn't aim to replace Docker—it offers an alternative for those who prefer a simpler, more native Linux experience.

## Features

> **Note:** This list represents the planned features and may change as the project evolves.

* 📊 Hardware resource monitoring (CPU, RAM, Disk, Network)
* 🚀 Deploy and manage applications as native systemd services
* ⚙️ Service management (start, stop, restart, enable, disable)
* 📜 Integrated log viewer powered by `journalctl`
* 🌐 Reverse proxy management *(planned)*
* 🔒 HTTPS certificate management *(planned)*
* 🖥️ Remote terminal access
* 🔄 Application updates
* 👥 User and permission management
* 🔌 Plugin-based runtime support *(planned — initial releases will support executable files only)*
