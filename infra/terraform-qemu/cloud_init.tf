resource "libvirt_cloudinit_disk" "commoninit" {
  name      = "commoninit.iso"
  user_data = data.template_file.user_data.rendered
}

data "template_file" "user_data" {
  template = templatefile("${path.module}/templates/cloud_init.tftpl", {
    user = {
      name = "dev"
      password_plain_text = "dev"
      ssh_keys = [
        file("~/.ssh/id_rsa.pub")
      ]
    }
  })
}