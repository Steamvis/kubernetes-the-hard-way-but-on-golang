How usage

Download img from https://cloud-images.ubuntu.com/ in directory `./images`
edit images.tf

```
resource "libvirt_volume" "ubuntu-2404" {
  name   = "ubuntu-2404"
  source = "./images/ubuntu_2404.img"
}
```

and volumes.tf

```
resource "libvirt_volume" "volumes" {
  base_volume_id = resource.libvirt_volume.ubuntu-2404.id
}
```

open locals.tf and edit hosts variable

example

```
  network-cidr ="10.10.10.0/24"
   hosts = [
        {
            ip = cidrhost(local.network-cidr, 100) # equals 10.10.10.100
            name = "master"
        },
        {
            ip = cidrhost(local.network-cidr, 199) # equals  10.10.10.199
            name = "nfs"
        }
    ]
```

/etc/libvirt/qemu.conf
```
security_driver = "none"
```