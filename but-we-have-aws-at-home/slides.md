---
# You can also start simply with 'default'
theme: seriph
# random image from a curated Unsplash collection by Anthony
# like them? see https://unsplash.com/collections/94734566/slidev
background: /images/taylor-vick-M5tzZtFCOfs-unsplash.jpg
# some information about your slides (markdown enabled)
title: But we have AWS at home
info: |
  It runs on optimism, compromises, and open source
# apply unocss classes to the current slide
class: text-center
# https://sli.dev/features/drawing
drawings:
  persist: false
# slide transition: https://sli.dev/guide/animations.html#slide-transitions
transition: slide-left
# enable MDC Syntax: https://sli.dev/features/mdc
mdc: true
# open graph
# seoMeta:
#  ogImage: https://cover.sli.dev

---
# But we have AWS at home

<div class="mt-12 py-1">
  It runs on optimism, compromises, and open source
</div>

<!--
The last comment block of each slide will be treated as slide notes. It will be visible and editable in Presenter Mode along with the slide. [Read more in the docs](https://sli.dev/guide/syntax.html#notes)
-->

---
layout: profile
image: /images/saKWC52.jpg
---

# Ryan Belgrave
Staff Engineer @ Confluent - WarpStream

~20 Years of technical experience, self-taught with Java in 2005

Ordered my first Ubuntu liveCD soon after and forever got stuck in vim

Running a homelab since 2012 when the first Raspberry Pi was released

Previously helped run the infrastructure automation for a few popular Minecraft servers

<br />

<carbon-logo-github /> [rmb938](https://github.com/rmb938)
<br />

<carbon-logo-linkedin /> [rbelgrave](https://www.linkedin.com/in/rbelgrave/)

---
---
# AWS At Home?
It's a home lab, not an outpost

![](/images/no_cloud_someone_else_computer.png)

---
layout: quote
---
# What is a Home Lab
Not someone else's computer

An infrastructure sandbox that you fully manage yourself running at home.

From Paspberry Pis, to mini-computers, used hardware, and real servers.

---
layout: two-cols
---

# Why Build One
Learning, Hosting, Because I want to

It can be a place to learn about hardware or software in a low risk environment.

Or maybe you want a cheap place to host your website and personal projects without the risk of a giant cloud bill.

Or maybe you are just a bit addicted to physical infrastructure and can't get enough of it at your job.

Big or small, always have a goal.

::right::

# Communities & Resources
There's places to get ideas

https://www.youtube.com/@CraftComputing

https://www.youtube.com/@HardwareHaven

https://www.youtube.com/@RaidOwl

https://www.youtube.com/@jeffsponaugle6339

https://www.reddit.com/r/homelab/

https://www.reddit.com/r/datahorder/

---
transition: fade
---
# Big or Small it's still a home lab
A single Raspberry Pi or multiple racks, if it's at home it's a home lab

<img style="max-height:80%" src="/images/jeff_geerling_pi_vs_n100.png"/>


[Jeff Geerling Is an Intel N100 a better value than a Raspberry Pi](https://www.jeffgeerling.com/blog/2025/intel-n100-better-value-raspberry-pi)

---
transition: fade
---
# Big or Small it's still a home lab
A single Raspberry Pi or multiple racks, if it's at home it's a home lab

<img style="max-height:80%" src="/images/jeff_geerling_project_mini_rack.png"/>


[Jeff Geerling Project Mini Rack](https://www.jeffgeerling.com/blog/2025/project-mini-rack-compact-and-portable-homelabs)

---
transition: fade
---
# Big or Small it's still a home lab
A single Raspberry Pi or multiple racks, if it's at home it's a home lab

<img style="max-height:80%" src="/images/jeff_cto.png"/>


[Jeff's CTO Laboratory](https://www.youtube.com/@jeffsponaugle6339)

---
transition: fade
---
# Big or Small it's still a home lab
A single Raspberry Pi or multiple racks, if it's at home it's a home lab

<img style="max-height:80%" src="/images/PXL_20250526_224801614.jpg"/>


Or part my own, yes it's dusty ☁️

---
---
# Staring Out and Trying new things

TODO: vmware, sandwich cloud, baremetal, ect...

---
layout: quote
---
# Building a Cloud at Home: Part 1
Requirements

## 1. Storage & Networking
## 2. Under-Cloud Compute
## 3. Security
## 4. Automation

---
---
# Part 1: Requirements - Storage & Networking
My requirements and hardware don't need to be yours

* Fast Shared Storage
  * Goal: 5GB/s+ Reads and Writes, faster is nice but expensive
  * <span style="color:green">Consumer SATA SSDs - happy wallet at $100 per 1TB for Samsung 870 EVOs</span>
  * <span style="color:green">Intel Optane Write Caches for low latency - sad wallet at $1k per 1TB used</span>
  * Used Enterprise U.3 SSDs - happy wallet at $100 per 1TB but sad wallet at expensive motherboards and CPUs

* Fast Networking
  * Goal: Don't bottleneck storage
  * 1 Gbs (100MB/s) is saturated by 1 hard drive (200MB/s)
  * 10 Gbs (1GB/s) is saturated by 2 SATA SSDs (500MB/s) - good for most
  * <span style="color:green">50 Gbs (5GB/s) is saturated by 10 SATA SSDs (500MB/s)</span>
  * 100 Gbs (10GB/s) is saturated by 1 PCI-E 5.0 NVMe SSD (13+ GB/s) - too expensive to scale

---
---
# Part 1: Requirements - Under-Cloud Compute
It's the cloud the runs the cloud

I'm using the [Proxmox](https://www.proxmox.com/en/) hypervisor, it runs KVM under the hood wrapped in a UI and API.

There are various terraform providers and packer plugins that can interface with Proxmox.

[XCP-ng](https://xcp-ng.org/) is a good alternative for a Xen based hypervisor but it needs an external orchestration service similar to VMWare VCenter.

I wanted a KVM based hypervisor because it is easier to develop on locally and I just have way more experience with it.

---
---
# Part 1: Requirements - Security
Security is easier then you think

SSH Keys, no usernames or passwords

All services should use TLS, preferably TLS 1.3

Prefer mTLS (mutual TLS) instead of usernames and passwords

Some sort of Secrets Management Service

Logging of all the things

---
layout: quote
---
# Part 1: Requirements - Automation

## Minimal manual configurating, deploying, and poking

---
layout: quote
---
# Building a Cloud at Home: Part 2
Implementing

## 1. TLS
## 2. Secrets Management
## 3. Service Discovery
## 4. Packer Images
## 5. Load Balancing

---
layout: quote
---
# Building a Cloud at Home: Part 2 - TLS.. the journey
This is really hard to do correctly without enterprise tools, this was the hardest part

---
layout: quote
---
# Smallstep - Step CA - smallstep.com/docs/step-ca
Open Source online Certificate Authority - ACME (Let's Encrypt), X.5C, ect..

---
---
# Step CA - Security Hardware
We need to security store the CA keys

The enterprise way is to store your Certificate Authority keys on an HSM (Hardware Security Module). The good ones are over $600, you need at least two, and they are hard to buy as a normal consumer.

There are cheap alternatives like the [NitroKey HSM](https://shop.nitrokey.com/shop/nkhs2-nitrokey-hsm-2-7#attr=) but documentation is poor.

What about Yubikeys? They are small, cheap, and have good documentation. They aren't made for this specific usecase but we can make it work and it'll still be secure. No keys should touch and filesystem.

```python {*|307|41-50|65-75|73-77|90-99|100-109|110-118}{maxHeight:'200px'}
#!/usr/bin/env python3

import datetime

import click
import PyKCS11
from cryptography import x509
from cryptography.hazmat.primitives import _serialization, hashes
from cryptography.hazmat.primitives.asymmetric import ec, utils
from cryptography.x509.oid import ExtendedKeyUsageOID, NameOID
from ykman import scripting as s
from yubikit.core.smartcard import ApduError
from yubikit.openpgp import DEFAULT_USER_PIN
from yubikit.piv import DEFAULT_MANAGEMENT_KEY, MANAGEMENT_KEY_TYPE, SLOT, PivSession

libykcs11_locations = [
    "/usr/lib64/libykcs11.so.2",
    "/usr/lib/x86_64-linux-gnu/libykcs11.so.2",
    "/usr/lib/aarch64-linux-gnu/libykcs11.so.2",
]


class YubiKeyECPrivateKey(ec.EllipticCurvePrivateKey):
    def __init__(self, serial: int, pin: str):
        self._pkcs11_lib = PyKCS11.PyKCS11Lib()

        lib_found = False
        for libykcs11_location in libykcs11_locations:
            print(f"Trying to load {libykcs11_location}")
            try:
                self._pkcs11_lib.load(libykcs11_location)
                lib_found = True
                print(f"Loaded {libykcs11_location}")
                break
            except PyKCS11.PyKCS11Error:
                continue

        if not lib_found:
            raise Exception("could not find libykcs11.so.2")

        for slot in self._pkcs11_lib.getSlotList(tokenPresent=True):
            token_info = self._pkcs11_lib.getTokenInfo(slot)
            if int(token_info.serialNumber) == serial:
                self._pkcs11_slot = slot
                break

        self._session = self._pkcs11_lib.openSession(
            self._pkcs11_slot, PyKCS11.CKF_SERIAL_SESSION | PyKCS11.CKF_RW_SESSION
        )
        self._session.login(pin)

    def exchange(
        self, algorithm: ec.ECDH, peer_public_key: ec.EllipticCurvePublicKey
    ) -> bytes:
        raise NotImplementedError()

    def public_key(self) -> ec.EllipticCurvePublicKey:
        """
        The EllipticCurvePublicKey for this private key.
        """
        return self.certificate.public_key()

    @property
    def certificate(self) -> x509.Certificate:
        objs = self._session.findObjects(
            [
                (PyKCS11.CKA_CLASS, PyKCS11.CKO_CERTIFICATE),
                (PyKCS11.CKA_LABEL, "X.509 Certificate for Retired Key 1"),
            ]
        )
        ca_cert_handle = objs[0]

        attributes = self._session.getAttributeValue(
            ca_cert_handle, [PyKCS11.CKA_VALUE], True
        )

        return x509.load_der_x509_certificate(bytes(attributes[0]))

    @property
    def curve(self) -> ec.EllipticCurve:
        """
        The EllipticCurve that this key is on.
        """
        return self.public_key().curve

    @property
    def key_size(self) -> int:
        return self.public_key().key_size

    def sign(
        self,
        data: bytes,
        signature_algorithm: ec.EllipticCurveSignatureAlgorithm,
    ) -> bytes:
        # TODO: what to do with signature_algorithm
        objs = self._session.findObjects(
            [
                (PyKCS11.CKA_CLASS, PyKCS11.CKO_PRIVATE_KEY),
                (PyKCS11.CKA_KEY_TYPE, PyKCS11.CKK_ECDSA),
                (PyKCS11.CKA_LABEL, "Private key for Retired Key 1"),
            ]
        )
        ca_private_key_handle = objs[0]

        signature = self._session.sign(
            ca_private_key_handle,
            data,
            PyKCS11.Mechanism(PyKCS11.CKM_ECDSA_SHA384),
        )

        key_size_bytes = (
            self.key_size // 8
        )  # Calculate key size in bytes (384 // 8 = 48)

        r = int.from_bytes(signature[:key_size_bytes], byteorder="big")
        s = int.from_bytes(signature[key_size_bytes:], byteorder="big")

        return utils.encode_dss_signature(r, s)

    def private_numbers(self) -> ec.EllipticCurvePrivateNumbers:
        raise NotImplementedError()

    def private_bytes(
        self,
        encoding: _serialization.Encoding,
        format: _serialization.PrivateFormat,
        encryption_algorithm: _serialization.KeySerializationEncryption,
    ) -> bytes:
        raise NotImplementedError()


def generate_intermediate_certificate(
    root_certificate: x509.Certificate,
    root_private_key: ec.EllipticCurvePrivateKey,
) -> tuple[x509.Certificate, ec.EllipticCurvePrivateKey]:
    print("Generating Intermediate Private Key")
    private_key = ec.generate_private_key(ec.SECP384R1())

    now = datetime.datetime.now()

    subject = x509.Name(
        [
            x509.NameAttribute(NameOID.COMMON_NAME, "Home Lab Intermediate CA"),
            x509.NameAttribute(NameOID.COUNTRY_NAME, "US"),
            x509.NameAttribute(NameOID.ORGANIZATION_NAME, "Home Lab"),
            x509.NameAttribute(NameOID.ORGANIZATIONAL_UNIT_NAME, "Step CA"),
        ]
    )

    builder = (
        x509.CertificateBuilder()
        .issuer_name(root_certificate.subject)
        .subject_name(subject)
        .not_valid_before(now)
        .not_valid_after(now + datetime.timedelta(days=(365 * 3)))  # 3 year validity
        .serial_number(x509.random_serial_number())
        .public_key(private_key.public_key())
        # Some examples of extensions to add, many more are possible:
        .add_extension(
            x509.BasicConstraints(ca=True, path_length=0),
            critical=True,
        )
        .add_extension(
            x509.KeyUsage(
                digital_signature=True,
                content_commitment=False,
                key_encipherment=False,
                data_encipherment=False,
                key_agreement=False,
                key_cert_sign=True,
                crl_sign=True,
                encipher_only=False,
                decipher_only=False,
            ),
            critical=True,
        )
        .add_extension(
            x509.ExtendedKeyUsage(
                [
                    ExtendedKeyUsageOID.CLIENT_AUTH,
                    ExtendedKeyUsageOID.SERVER_AUTH,
                ]
            ),
            critical=False,
        )
        .add_extension(
            x509.NameConstraints(
                permitted_subtrees=[
                    x509.DNSName(".rmb938.me"),
                    x509.DNSName(".tailnet-047c.ts.net"),
                ],
                excluded_subtrees=None,
            ),
            critical=True,
        )
        .add_extension(
            x509.SubjectKeyIdentifier.from_public_key(private_key.public_key()),
            critical=False,
        )
        .add_extension(
            x509.AuthorityKeyIdentifier.from_issuer_public_key(
                root_certificate.public_key()
            ),
            critical=False,
        )
        .add_extension(
            x509.CRLDistributionPoints(
                [
                    x509.DistributionPoint(
                        full_name=[
                            x509.UniformResourceIdentifier(
                                "http://step-ca.us-homelab1.hl.rmb938.me/1.0/crl"
                            ),
                        ],
                        relative_name=None,
                        reasons=None,
                        crl_issuer=None,
                    ),
                    x509.DistributionPoint(
                        full_name=[
                            x509.UniformResourceIdentifier(
                                "http://hl-us-homelab1-step-ca.tailnet-047c.ts.net/1.0/crl"
                            ),
                        ],
                        relative_name=None,
                        reasons=None,
                        crl_issuer=None,
                    ),
                ]
            ),
            critical=False,
        )
    )

    print("Creating Intermediate CA certificate")
    certificate = builder.sign(
        private_key=root_private_key,
        algorithm=hashes.SHA384(),
    )

    return certificate, private_key


def write_keys(
    yubikey: s.ScriptingDevice,
    root_certificate: x509.Certificate,
    certificate: x509.Certificate,
    private_key: ec.EllipticCurvePrivateKey,
    pin: str,
):
    # Slot 82
    slot = SLOT.RETIRED1

    root_slot = SLOT.RETIRED2

    # Establish PIV session
    piv = PivSession(yubikey.smart_card())

    # Check if a key is already stored, if so error
    try:
        piv.get_slot_metadata(slot)
        print(f"Key already exists in slot {slot}.")
        exit(1)
    except ApduError as e:
        pass

    piv.authenticate(
        key_type=MANAGEMENT_KEY_TYPE.AES256, management_key=bytes.fromhex(pin)
    )

    # Put the key and certificate
    print("Writting private key and public certificate")
    piv.put_key(slot, private_key)
    piv.put_certificate(slot, certificate)
    piv.put_certificate(root_slot, root_certificate)


def next_yubikey(serials: list[int]) -> tuple[s.ScriptingDevice, list[int]]:
    while True:
        click.echo("Remove Yubikey and insert another, press enter to continue...")
        input()
        print("Connecting to new Yubikey")
        yubikey = s.single()
        print(f"Connected to Yubikey with Serial {yubikey.info.serial}")

        # Serial is different so we have a different one inserted
        if yubikey.info.serial not in serials:
            serials.append(yubikey.info.serial)
            break

        # Serial is the same to log and continue
        click.echo("The same yubikey was inserted, please insert a different one.")

    return yubikey, serials


def main():
    print("Connecting to Root Yubikey")
    yubikey = s.single()
    serials = [yubikey.info.serial]
    print(f"Connected to Root Yubikey with Serial {yubikey.info.serial}")

    # Unlock with the user pin
    root_user_pin = click.prompt(
        "Enter user pin", default=DEFAULT_USER_PIN, hide_input=True
    )
    ybi_key = YubiKeyECPrivateKey(yubikey.info.serial, root_user_pin)

    intermediate_certificate, intermediate_private_key = (
        generate_intermediate_certificate(ybi_key.certificate, ybi_key)
    )

    yubikey, serials = next_yubikey(serials)
    piv = PivSession(yubikey.smart_card())

    # Unlock with the management key
    while True:
        try:
            intermediate_management_pin = click.prompt(
                "Enter management key",
                default=DEFAULT_MANAGEMENT_KEY.hex(),
                hide_input=True,
            )
            piv.authenticate(
                key_type=MANAGEMENT_KEY_TYPE.AES256,
                management_key=bytes.fromhex(intermediate_management_pin),
            )
            break
        except (ApduError, ValueError):
            print("Invalid Management Key, try again.")
    write_keys(
        yubikey,
        ybi_key.certificate,
        intermediate_certificate,
        intermediate_private_key,
        intermediate_management_pin,
    )


if __name__ == "__main__":
    main()
```

---
---
# Step CA - Compute Hardware
Raspberry Pi - With Encryption?

Raspberry Pi does not natively support any form of full disk encryption.

[SwissBit](https://www.swissbit.com/) sells a self encrypting SD card for the Pi that encrypts based off of some hardware attributes. Very little documentation on how this actually works.

[Raspberry Pi 4 UEFI Firmware](https://github.com/pftf/RPi4) to the rescue!

This allows installing a normal ARM64 based operating system in the same way you would with a normal computer, no more burning Linux Images.

Install Ubuntu normally with Full Disk encryption then setup a USB unlock
key for automated unlocks.

The system becomes unusable without the USB unlock key.

https://github.com/rmb938/ansible_step-ca/

---
---
# Building a Cloud at Home: Part 2 - TLS.. the journey reviewed!
Complicated pieces for simple outcome

Step CA as the certificate authority

Certificate Authority keys are stored securely on hardware tokens and never touched the filesystem

We have operating system with full disk encryption running on a Raspberry Pi

This gives us a fully secure Let's Encrypt-like certificate authority... at least if I kept the hardware keys holding the Root CA in a physical vault, which I don't.

---
---
# Building a Cloud at Home: Part 2 - Secrets Management &  Service Discovery
Our cloud needs secrets and the things running on the cloud needs secrets

In my opinion the go-to secrets management system in Hashicorp Vault or the MPL 2.0 Licensed OpenBao.

We can program our servers to request a certficate from Step CA via ACME and use that certificate to authenticate to Vault via PKI Auth.

Vault can be configured using Terraform to set policies so each server can only access the secrets it needs.

The servers then can use Vault to provision themselves onto a Hashicorp Consul cluster to start advertising and discovering services.

https://github.com/rmb938/ansible_hashicorp-vault

https://github.com/rmb938/tf_hashicorp_vault

https://github.com/rmb938/tf_hashicorp_vault_apps

---
---
# Building a Cloud at Home: Part 2 - Packer Images
No not the football team

Hashicorp Packer can plug into every cloud provider and most hypervisors to build Operating Systems Images.

Allows easy deployment of many identical VMs.

Can build images using shell scripts or in my case ansible.

https://github.com/hashicorp/packer-plugin-ansible

https://github.com/rmb938/proxmox-images

```hcl {*|2-6|33-42|45-48}{maxHeight:'200px'}
source "proxmox-clone" "ubuntu-noble-haproxy-t1" {
  username    = "${var.proxmox_username}"
  token       = "${var.proxmox_token}"
  proxmox_url = "https://${var.proxmox_hostname}:8006/api2/json"
  node        = "freenas-pm"
  clone_vm_id = "${var.clone_vm_id}"
  sockets     = 1
  cores       = 1
  memory      = 2048
  network_adapters {
    bridge = "vmbr0v52"
    model  = "virtio"
  }
  ipconfig {
    ip = "dhcp"
    ip6 = "auto"
  }
  bios    = "ovmf"
  machine = "q35"
  scsi_controller = "virtio-scsi-single"
  os           = "l26"
  ssh_username = "ubuntu"
  // Must have at least one tag, otherwise it shows up as `"tags": " "` in the api.
  tags = "packer"
  cloud_init              = true
  cloud_init_storage_pool = "freenas-nfs"

}

build {
  sources = ["source.proxmox-clone.ubuntu-noble-haproxy-t1"]
  // Packer setup
  provisioner "shell" {
    script = "../scripts/provisioner-shell-image-packer.sh"
  }
  provisioner "ansible" {
    playbook_file   = "ansible/site.yaml"
    user            = "ubuntu"
    extra_arguments = [
      "-v",
      "--diff"
    ]
    ansible_env_vars = ["ANSIBLE_FORCE_COLOR=1"]
  }
  // Trivy
  provisioner "shell" {
    script = "../scripts/provisioner-shell-image-trivy.sh"
  }
  // Cleanup
  provisioner "shell" {
    script = "../scripts/provisioner-shell-image-cleanup.sh"
  }
  post-processor "manifest" {}
}
```

---
layout: two-cols
---
# Building a Cloud at Home: Part 2 - Load Balancing
On no! My server died! Anyway...

Protect your services from (un-)planned interruptions by putting them behind a load balancer!

I am using [HAProxy](https://www.haproxy.org/) in a two tier setup for ease of use and no frills configuration.

Automatic wildcard certificates from Step CA

HAProxy is configured to load services from Consul so as new services come online and register as needing a load balancer they are automatically picked up.

::right::

<img src="/images/haproxy_t2.png"/>
