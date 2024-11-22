# DevHosts

Hosts file manager for development environments.

Create a file that contains a map of the hostnames you want to use in your development environment to the IP address of the server you want to use.
Then run `devhosts` to update your hosts file.  
You can choose the syntax of the file.
Supported formats are `yaml`, `toml` and `json`.

## Example

`devhosts.yaml`

```yaml
---
example.com: 127.0.0.1
``` 
