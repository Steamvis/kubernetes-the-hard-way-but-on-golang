data "libvirt_network_dns_host_template" "hosts" {
    for_each = {
      for v in local.hosts : v.name => v
    }
    ip       = each.value.ip
    hostname = each.key
}

resource "libvirt_network" "common" {
  name = "common"

  # mode can be: "nat" (default), "none", "route", "open", "bridge"
  mode = "nat"

  domain = "local"

  addresses = [local.network-cidr]

  dns {
    enabled    = true
    local_only = true

    dynamic "hosts" {
        for_each = {
          for v in data.libvirt_network_dns_host_template.hosts : v.hostname => v
        }
        content {
            hostname = hosts.key
            ip = hosts.value.ip
        } 
    }
  }
}