Type: `gridscale`
Artifact BuilderId: `packer.gridscale`

The `gridscale` Packer builder is used to create new (os) templates for use with [gridscale](https://www.gridscale.io/). The builder takes a template (in gridscale) or an iso-image, runs any provisioning necessary on the template/iso-image after launching it, then snapshots it into a reusable template. This reusable template can then be used as the foundation of new servers that are provisioned within gridscale user space.

The builder does _not_ manage templates. Once it creates a template, it is up to you
to use it or delete it.

## Installation

To install this plugin using Packer v1.7.0 or later, copy and paste this code
into your Packer configuration.

Then, run [`packer init`](https://www.packer.io/docs/commands/init).


```hcl
packer {
  required_plugins {
    gridscale = {
      version = ">= 0.0.1"
      source  = "github.com/gridscale/gridscale"
    }
  }
}
```

## Configuration Reference

There are many configuration options available for the builder. They are
segmented below into two categories: required and optional parameters. Within
each category, the available configuration keys are alphabetized.

### Required:

<!-- Code generated from the comments of the Config struct in builder/gridscale/config.go; DO NOT EDIT MANUALLY -->

- `api_token` (string) - The client TOKEN to use to access your account. Environment variable `GRIDSCALE_TOKEN` can be set instead.

- `api_key` (string) - The client KEY to use to access your account. Environment variable `GRIDSCALE_UUID` can be set instead.

- `server_cores` (int) - Number of server cores

- `server_memory` (int) - Server memory capacity (in GB)

- `storage_capacity` (int) - Storage capacity (in GB)

<!-- End of code generated from the comments of the Config struct in builder/gridscale/config.go; -->


### Optional:

<!-- Code generated from the comments of the Config struct in builder/gridscale/config.go; DO NOT EDIT MANUALLY -->

- `api_url` (string) - The server URL to use to access your account. Default: "https://api.gridscale.io". Environment variable `GRIDSCALE_URL` can be set instead.

- `api_request_headers` (string) - APIRequestHeaders is for debug purpose only. Format: "key1:val1,key2:val2"

- `template_name` (string) - The name of the new template.

- `hostname` (string) - Name of the host.

- `server_name` (string) - Name of the server used for producing the template.

- `secondary_storage` (bool) - SecondaryStorage is set to true when the server needs a secondary storage
  during producing template process.
  **NOTE**: If `secondary_storage=true`, the template will be built from the second storage.

- `base_template_uuid` (string) - A pre-built template UUID. This template is used to produce another template. E.g: Ubuntu template.
  **NOTE**: One of these fields has to be set: `isoimage_uuid`, `isoimage_url`, `base_template_uuid`.

- `isoimage_uuid` (string) - A pre-built ISO image is used by the given ISO image UUID. If IsoImageUUID is set, IsoImageURL is ignored.
  **NOTE**: One of these fields has to be set: `isoimage_uuid`, `isoimage_url`, `base_template_uuid`.

- `isoimage_url` (string) - An URL is used to download the image. If IsoImageUUID is set, IsoImageURL is ignored.
  **NOTE**: One of these fields has to be set: `isoimage_uuid`, `isoimage_url`, `base_template_uuid`.

- `boot_command` ([]string) - This is an array of commands to type when the server instance is first
  booted. The goal of these commands should be to type just enough to
  initialize the operating system installer. Special keys can be typed as
  well, and are covered in the section below on the boot command. If this
  is not specified, it is assumed the installer will start itself.

- `boot_wait` (duration string | ex: "1h5m2s") - The time to wait after booting the initial virtual machine before typing
  the `boot_command`. The value of this should be a duration. Examples are
  `5s` and `1m30s` which will cause Packer to wait five seconds and one
  minute 30 seconds, respectively. If this isn't specified, the default is
  `10s` or 10 seconds. To set boot_wait to 0s, use a negative number, such
  as "-1s"

- `boot_key_interval` (duration string | ex: "1h5m2s") - Time in ms to wait between each key press

- `files` ([]string) - A list of files' relative paths that need to be served on a HTTP server.
  Put this address ({{__HTTP__ADDRESS__}} is a placeholder, do not edit) http://{{__HTTP__ADDRESS__}}/path/to/file
  to `boot_command` to use http-served files in boot commands.

<!-- End of code generated from the comments of the Config struct in builder/gridscale/config.go; -->


## Basic Example

Here is a basic example. It is completely valid as soon as you enter your own `api_key` and `api_token` (or via environment variables `GRIDSCALE_UUID` and `GRIDSCALE_TOKEN`):

**HCL2**

```hcl
source "gridscale" "example" {
	base_template_uuid = "fd65f8ce-e2c6-40af-8fc3-92efa0d4eecb"
	hostname           = "test-hostname"
	ssh_password       = "testPassword"
	ssh_username       = "root"
	server_cores       = 2
	server_memory      = 4
	storage_capacity   = 10
	template_name      = "my-ubuntu20.04-template"
}

build {
	sources = ["source.gridscale.example"]
}
```

**JSON**

```json
{
	"builders": [
		{
			"type": "gridscale",
			"template_name": "my-ubuntu20.04-template",
			"password": "testPassword",
			"hostname": "test-hostname",
			"ssh_username": "root",
			"server_memory": 4,
			"server_cores": 2,
			"storage_capacity": 10,
			"base_template_uuid": "fd65f8ce-e2c6-40af-8fc3-92efa0d4eecb"
		}
	]
}

```


### Communicator Config

In addition to the builder options, a
[communicator](/docs/templates/legacy_json_templates/communicator) can be configured for this builder.

<!-- Code generated from the comments of the Config struct in communicator/config.go; DO NOT EDIT MANUALLY -->

- `communicator` (string) - Packer currently supports three kinds of communicators:
  
  -   `none` - No communicator will be used. If this is set, most
      provisioners also can't be used.
  
  -   `ssh` - An SSH connection will be established to the machine. This
      is usually the default.
  
  -   `winrm` - A WinRM connection will be established.
  
  In addition to the above, some builders have custom communicators they
  can use. For example, the Docker builder has a "docker" communicator
  that uses `docker exec` and `docker cp` to execute scripts and copy
  files.

- `pause_before_connecting` (duration string | ex: "1h5m2s") - We recommend that you enable SSH or WinRM as the very last step in your
  guest's bootstrap script, but sometimes you may have a race condition
  where you need Packer to wait before attempting to connect to your
  guest.
  
  If you end up in this situation, you can use the template option
  `pause_before_connecting`. By default, there is no pause. For example if
  you set `pause_before_connecting` to `10m` Packer will check whether it
  can connect, as normal. But once a connection attempt is successful, it
  will disconnect and then wait 10 minutes before connecting to the guest
  and beginning provisioning.

<!-- End of code generated from the comments of the Config struct in communicator/config.go; -->


<!-- Code generated from the comments of the SSH struct in communicator/config.go; DO NOT EDIT MANUALLY -->

- `ssh_host` (string) - The address to SSH to. This usually is automatically configured by the
  builder.

- `ssh_port` (int) - The port to connect to SSH. This defaults to `22`.

- `ssh_username` (string) - The username to connect to SSH with. Required if using SSH.

- `ssh_password` (string) - A plaintext password to use to authenticate with SSH.

- `ssh_ciphers` ([]string) - This overrides the value of ciphers supported by default by golang.
  The default value is [
    "aes128-gcm@openssh.com",
    "chacha20-poly1305@openssh.com",
    "aes128-ctr", "aes192-ctr", "aes256-ctr",
  ]
  
  Valid options for ciphers include:
  "aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com",
  "chacha20-poly1305@openssh.com",
  "arcfour256", "arcfour128", "arcfour", "aes128-cbc", "3des-cbc",

- `ssh_clear_authorized_keys` (bool) - If true, Packer will attempt to remove its temporary key from
  `~/.ssh/authorized_keys` and `/root/.ssh/authorized_keys`. This is a
  mostly cosmetic option, since Packer will delete the temporary private
  key from the host system regardless of whether this is set to true
  (unless the user has set the `-debug` flag). Defaults to "false";
  currently only works on guests with `sed` installed.

- `ssh_key_exchange_algorithms` ([]string) - If set, Packer will override the value of key exchange (kex) altorighms
  supported by default by golang. Acceptable values include:
  "curve25519-sha256@libssh.org", "ecdh-sha2-nistp256",
  "ecdh-sha2-nistp384", "ecdh-sha2-nistp521",
  "diffie-hellman-group14-sha1", and "diffie-hellman-group1-sha1".

- `ssh_certificate_file` (string) - Path to user certificate used to authenticate with SSH.
  The `~` can be used in path and will be expanded to the
  home directory of current user.

- `ssh_pty` (bool) - If `true`, a PTY will be requested for the SSH connection. This defaults
  to `false`.

- `ssh_timeout` (duration string | ex: "1h5m2s") - The time to wait for SSH to become available. Packer uses this to
  determine when the machine has booted so this is usually quite long.
  Example value: `10m`.

- `ssh_disable_agent_forwarding` (bool) - If true, SSH agent forwarding will be disabled. Defaults to `false`.

- `ssh_handshake_attempts` (int) - The number of handshakes to attempt with SSH once it can connect. This
  defaults to `10`.

- `ssh_bastion_host` (string) - A bastion host to use for the actual SSH connection.

- `ssh_bastion_port` (int) - The port of the bastion host. Defaults to `22`.

- `ssh_bastion_agent_auth` (bool) - If `true`, the local SSH agent will be used to authenticate with the
  bastion host. Defaults to `false`.

- `ssh_bastion_username` (string) - The username to connect to the bastion host.

- `ssh_bastion_password` (string) - The password to use to authenticate with the bastion host.

- `ssh_bastion_interactive` (bool) - If `true`, the keyboard-interactive used to authenticate with bastion host.

- `ssh_bastion_private_key_file` (string) - Path to a PEM encoded private key file to use to authenticate with the
  bastion host. The `~` can be used in path and will be expanded to the
  home directory of current user.

- `ssh_bastion_certificate_file` (string) - Path to user certificate used to authenticate with bastion host.
  The `~` can be used in path and will be expanded to the
  home directory of current user.

- `ssh_file_transfer_method` (string) - `scp` or `sftp` - How to transfer files, Secure copy (default) or SSH
  File Transfer Protocol.

- `ssh_proxy_host` (string) - A SOCKS proxy host to use for SSH connection

- `ssh_proxy_port` (int) - A port of the SOCKS proxy. Defaults to `1080`.

- `ssh_proxy_username` (string) - The optional username to authenticate with the proxy server.

- `ssh_proxy_password` (string) - The optional password to use to authenticate with the proxy server.

- `ssh_keep_alive_interval` (duration string | ex: "1h5m2s") - How often to send "keep alive" messages to the server. Set to a negative
  value (`-1s`) to disable. Example value: `10s`. Defaults to `5s`.

- `ssh_read_write_timeout` (duration string | ex: "1h5m2s") - The amount of time to wait for a remote command to end. This might be
  useful if, for example, packer hangs on a connection after a reboot.
  Example: `5m`. Disabled by default.

- `ssh_remote_tunnels` ([]string) - 

- `ssh_local_tunnels` ([]string) - 

<!-- End of code generated from the comments of the SSH struct in communicator/config.go; -->


- `ssh_private_key_file` (string) - Path to a PEM encoded private key file to use to authenticate with SSH.
  The `~` can be used in path and will be expanded to the home directory
  of current user.
