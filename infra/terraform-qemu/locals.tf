locals {
    network-cidr ="10.10.10.0/24"
    jump_host = {
        name = "jump-host"
        memory = 512
        cpu = 1
        ip = cidrhost(local.network-cidr, 199)
    }
    hosts = [
        {
            ip = cidrhost(local.network-cidr, 100)
            name = "master"
        },
        # {
        #     ip = cidrhost(local.network-cidr, 199)
        #     name = "nfs"
        # }
    ]
}