resource "libvirt_domain" "jump_host" {
  name   = local.jump_host.name
  memory = local.jump_host.memory
  vcpu   = local.jump_host.cpu

  disk {
    volume_id = libvirt_volume.jump_host_volume.id
  }

  cloudinit = libvirt_cloudinit_disk.commoninit.id

  network_interface {
    network_id = libvirt_network.common.id
    addresses = [local.jump_host.ip]
    wait_for_lease = true
  }

  graphics {
    type        = "vnc"
    listen_type = "address"
  }
}


resource "libvirt_domain" "hosts" {
  for_each = {
    for v in local.hosts  : v.name => v
  }
  name   = "${each.key}-vm"
  memory = "2048"
  vcpu   = 2

  disk {
    volume_id = libvirt_volume.volumes[each.value.name].id
  }

  cloudinit = libvirt_cloudinit_disk.commoninit.id

  network_interface {
    network_id = libvirt_network.common.id
    addresses = [each.value.ip]
    wait_for_lease = true
  }

  graphics {
    type        = "vnc"
    listen_type = "address"
  }
}
