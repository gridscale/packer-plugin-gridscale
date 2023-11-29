# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Gridscale"
  description = "Create new (os) templates for use with gridscale."
  identifier = "packer/gridscale/gridscale"
  component {
    type = "builder"
    name = "gridscale"
    slug = "gridscale"
  }
}
