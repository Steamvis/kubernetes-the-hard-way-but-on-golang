resource "libvirt_volume" "ubuntu-2404" {
  name   = "ubuntu-2404"
  source = "${path.cwd}/images/ubuntu-2404.img"
}