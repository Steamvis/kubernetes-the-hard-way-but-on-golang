variable "size" {
  default = 21474836480 # 20 GB
}

resource "libvirt_volume" "jump_host_volume" {
  name = "jump_host.qcow2"

  base_volume_id = resource.libvirt_volume.ubuntu-2404.id
  size = var.size
}

resource "libvirt_volume" "volumes" {
  for_each = {
    for v in local.hosts : v.name => v
  }
  name = "${each.key}.qcow2"
  base_volume_id = resource.libvirt_volume.ubuntu-2404.id
  size = var.size
}
