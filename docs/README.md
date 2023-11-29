The `gridscale` Packer builder is used to create new (os) templates for use with [gridscale](https://www.gridscale.io/). The builder takes a template (in gridscale) or an iso-image, runs any provisioning necessary on the template/iso-image after launching it, then snapshots it into a reusable template. This reusable template can then be used as the foundation of new servers that are provisioned within gridscale user space.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    gridscale = {
      version = "= 0.0.2"
      source  = "github.com/gridscale/gridscale"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/gridscale/gridscale
```


### Components

#### Builders

- [gridscale](/packer/integrations/gridscale/gridscale/latest/components/builder/gridscale) - The builder takes a template (in gridscale) or an iso-image, runs any provisioning necessary on the template/iso-image after launching it, then snapshots it into a reusable template. This reusable template can then be used as the foundation of new servers that are provisioned within gridscale user space.
