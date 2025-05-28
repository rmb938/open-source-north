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
The want to be a cloud

It all started in 2014 with a old HP Proliant DL380 G5 and a few Cisco routers and switches. They were very loud and I didn't end up using them much.

Eventually I built a workstation and run VMWare vSphere ontop but it felt like something was missing... I wanted to automate something that wasn't ready to be automated.

I then learned about OpenStack through an internship but I was a poor college student and couldn't afford the hardware it needed to run it in a home lab.

<br />

That became my goal, to run a cloud at home. A goal I wouldn't fully achieve until late 2024.

---
---
# Building a Cloud with Sticks and Stones
You can't make a rock think... right?

VMWare in 2014-16 really did not want to be automated, there was no REST API and a barely documented SOAP API.

Packer for building cloud images didn't really exist and this new thing called Terraform would never take off.

The mainframe was being reinvented with containers and a friendly blue whale, who would ever use such a thing.

<br />

I wanted a cloud at home! But I had no idea what I was doing so I dealt with VMWare struggles.

---
layout: quote
---
# Building a Cloud With Smart Rocks
They are thinking now... maybe AI will exist one day.

VMWare 6.5 released late 2016 with a REST API! It wasn't good but it was something!

Time to build a cloud at home and name it... Sandwich Cloud... it was from the Sandwich alignment chart meme...

---
layout: image

image: /images/sandwiches.png

backgroundSize: contain
---

---
layout: quote
---
# It was funny at the time!
No really...

---
---
# Sandwiches in the Cloud! Maybe not such a great idea
Rocks and Sandwiches thinking? Maybe we'll have flying cars next!

It was fun to build and I learned alot about what not to do with a complex software project.

But it wasn't want I wanted, I had to build my own tooling for everything including a metadata service, terraform providers, packer plugins and it just wasn't fun.

<br />

Time to rip it all out and build a baremetal as a service... that wasn't fun either!

---
---
# Sticks to Rocks to Metal
Silicon is a metal? Does that count?

I quickly learned that consumer hardware didn't want to be automated. No matter how hard you tried it did not want automation.

UEFI as a BIOS replacement was brand new and super buggy.

Computers would randomly forget how to boot into an operating system.

Things felt manual again.

<br />

~5 iterations later I had something that worked 90% of the time but that
10% was frustraiting!

---
layout: quote
---
# The Dark Ages
No sticks, no sandwiches, no metal

And for many years I was stuck, I wanted a cloud at home but I could never acheieve it.

I bought more hardware and got rid of others hoping one day I would figure it out.

Tried many things... but a real OpenStack deployment always felt too far away...

---
---
# Finding a new Goal
The dark shadow of a cloud at home

I tried to find a new goal.

Automating Kubernetes check... but it didn't fufil the dream.

Running a hypervisor cluster with KVM or Proxmox or XCP-NG check... still no dream.

<br />

Wasting money double check.

Something had to change.

---
layout: quote
---
# A new Job, renewed hope!
Rebellions are built on hope

I was going to deploy OpenStack and I was going to do it right!

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
  * 1 Gbs (~100MB/s) is saturated by 1 hard drive (200MB/s)
  * 10 Gbs (~1GB/s) is saturated by 2 SATA SSDs (500MB/s) - good for most
  * <span style="color:green">50 Gbs (~5GB/s) is saturated by 10 SATA SSDs (500MB/s)</span>
  * 100 Gbs (~10GB/s) is saturated by 1 PCI-E 5.0 NVMe SSD (13+ GB/s) - too expensive to scale

---
---
# Part 1: Requirements - Under-Cloud Compute
It's the cloud the runs the cloud, it's not a sandwich!

I'm using the [Proxmox](https://www.proxmox.com/en/) hypervisor, it runs KVM under the hood wrapped in a UI and API.

There are various terraform providers and packer plugins that can interface with Proxmox.

[XCP-ng](https://xcp-ng.org/) is a good alternative for a Xen based hypervisor but it needs an external orchestration service similar to VMWare VCenter.

I wanted a KVM based hypervisor because it is easier to develop on locally and I just have way more experience with it through previous homelab experiments.

---
---
# Part 1: Requirements - Security
Security is easy right? The sticks and rocks think so! ... wait thinking sticks?

SSH Keys, no usernames or passwords

All services should use TLS, preferably TLS 1.3

Prefer mTLS (mutual TLS) instead of usernames and passwords

Some sort of Secrets Management Service

Centralized Logging of all the things

---
layout: quote
---
# Part 1: Requirements - Automation

## Minimal manual configuration, deploying, and poking

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
Enterprise security at home

This was really hard, and not for the reasons you may think

---
layout: quote
---
# Smallstep - Step CA - smallstep.com/docs/step-ca
Open Source online Certificate Authority - ACME (Let's Encrypt), X.5C, ect..

---
---
# Step CA - Security Hardware
We need to security store the CA keys. In a notepad, on a train, in a box with a fox?

The enterprise way is to store your Certificate Authority keys on an HSM (Hardware Security Module). The good ones are over $600, you need at least two, and they are hard to buy as a normal consumer.

There are cheap alternatives like the [NitroKey HSM](https://shop.nitrokey.com/shop/nkhs2-nitrokey-hsm-2-7#attr=) but documentation is poor and getting it in the USA can be difficult.

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
Raspberry Pi - With Encryption? You can't encrypt a Pi!

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

---
layout: quote
---
# It's time for OpenStack!
10 years later it's finally happening!

---
---
# Doing OpenStack my way!
Doing all the wrong things to make it better.

OpenStack is built for MySQL, so let's use PostgreSQL instead!

<br />

OpenStack's identity service, Keystone, uses RSA and Fernet keys for user authentication and credential encryption, let's make it use HashiCorp Vault Instead!

<br />

The VM image storage service, Glance, is made to use local storage or object storage, naa we will use NFS.

<br />

Database migrations, service credentials, project setup... all typically done manually, let's automate it!

---
---
# OpenStack & PostgreSQL
Like peanut butter and jelly... wait that's a sandwich!

For the most part this just works, it's undocumented and unsupported but it works.

<br />

Some OpenStack services need 1-5 lines of code patched but it was super easy until LBaaS.

<br />

The OpenStack LBaaS service would not run database migrations on PostgreSQL even though the application could run with it just fine.
I had to re-write a lot of it's database migrations to make it work

Patch Submitted https://review.opendev.org/c/openstack/octavia/+/948767, still not reviewed or merged 😔

---
---
# OpenStack Keystone - No Manual RSA and Fernet Keys
JWT Keys and Hashi Vault!

```python {*|5-23|33-50|56-62}{maxHeight:'350px'}
    issued_at = utils.isotime(subsecond=True)
    issued_at_int = self._convert_time_string_to_int(issued_at)
    expires_at_int = self._convert_time_string_to_int(expires_at)

    payload = {
        # public claims
        "sub": user_id,
        "iat": issued_at_int,
        "exp": expires_at_int,
        # private claims
        "openstack_methods": methods,
        "openstack_audit_ids": audit_ids,
        "openstack_system": system,
        "openstack_domain_id": domain_id,
        "openstack_project_id": project_id,
        "openstack_trust_id": trust_id,
        "openstack_group_ids": federated_group_ids,
        "openstack_idp_id": identity_provider_id,
        "openstack_protocol_id": protocol_id,
        "openstack_access_token_id": access_token_id,
        "openstack_app_cred_id": app_cred_id,
        "openstack_thumbprint": thumbprint,
    }

    # NOTE(lbragstad): Calling .items() on a dictionary in python 2 returns
    # a list but returns an iterable in python 3. Casting to a list makes
    # it safe to modify the dictionary while iterating over it, regardless
    # of the python version.
    for k, v in list(payload.items()):
        if v is None:
            payload.pop(k)

    header = (
        base64.urlsafe_b64encode(
            json.dumps(
                {
                    "alg": "ES256",
                    "typ": "JWT",
                }
            ).encode("utf-8")
        )
        .rstrip(b"=")
        .decode("utf-8")
    )

    jwt_payload = (
        base64.urlsafe_b64encode(json.dumps(payload).encode("utf-8"))
        .rstrip(b"=")
        .decode("utf-8")
    )

    message = base64.standard_b64encode(
        f"{header}.{jwt_payload}".encode("utf-8")
    ).decode("utf-8")

    client = create_vault_client()
    sign_resp = client.secrets.transit.sign_data(
        mount_point=CONF.token_hashicorp_vault.transit_mount_point,
        name=CONF.token_hashicorp_vault.transit_key_name,
        hash_algorithm="sha2-256",
        marshaling_algorithm="jws",
        hash_input=message,
    )

    signature = sign_resp["data"]["signature"].removeprefix("vault:v1:")

    token_id = f"{header}.{jwt_payload}.{signature}"

    return token_id, issued_at
```

https://github.com/rmb938/openstack-keystone-hashicorp-vault-provider

---
---
# OpenStack Glance on NFS
It's kinda supported and mostly works

1. Mount /var/lib/glance as an NFS mount
2. Upload an Image into Glance's API
3. ???
4. It just works

A bit surprising, I'm waiting for something to explode.

---
---
# OpenStack Database Migrations are Scary - Automate them anyway
No time to be scared, let's automate and hope nothing explodes

```python {*|33-48|74-82|121-136}{maxHeight:'300px'}
#! /usr/bin/python3

import subprocess
import sys

import consul


def has_cell(cell_name: str):
    list_cells_command = [
        "/usr/bin/nova-manage",
        "cell_v2",
        "list_cells",
    ]

    result = subprocess.run(list_cells_command, capture_output=True, text=True)

    if result.returncode != 0:
        print(f"Unknown Error code from list cells: {result}")
        sys.exit(1)
        return

    if f"nova-{cell_name}" in result.stdout:
        return True

    return False


def main():
    lock_kv_path = "openstack-nova-controller/migration-lock"
    api_sync_command = ["/usr/bin/nova-manage", "api_db", "sync"]

    map_cell0_command = [
        "/usr/bin/nova-manage",
        "cell_v2",
        "map_cell0",
        "--database_connection",
        "postgresql+psycopg://nova-cell0@primary.openstack-postgres.service.consul:7432/nova-cell0?sslmode=verify-full&sslrootcert=/etc/nova/postgres-server-ca.crt&sslcert=/etc/nova/postgres-user-nova-cell0.crt&sslkey=/etc/nova/postgres-user-nova-cell0.key",
    ]
    create_cell_cell1_command = [
        "/usr/bin/nova-manage",
        "cell_v2",
        "create_cell",
        "--name=cell1",
        "--verbose",
        "--database_connection",
        "postgresql+psycopg://nova-cell1@primary.openstack-postgres.service.consul:7432/nova-cell1?sslmode=verify-full&sslrootcert=/etc/nova/postgres-server-ca.crt&sslcert=/etc/nova/postgres-user-nova-cell1.crt&sslkey=/etc/nova/postgres-user-nova-cell1.key",
    ]

    sync_command = ["/usr/bin/nova-manage", "db", "sync"]
    online_data_migrations_command = [
        "/usr/bin/nova-manage",
        "db",
        "online_data_migrations",
    ]

    consul_client = consul.Consul(
        host="127.0.0.1",
        port=8500,
        consistency="consistent",
    )

    session_id = consul_client.session.create(ttl=60, behavior="delete")

    try:
        acquired = consul_client.kv.put(lock_kv_path, "locked", acquire=session_id)

        if not acquired:
            # Not exiting 1 here because it's normal
            print(f"Could not aquire consul lock at {lock_kv_path}")
            return

        try:
            result = subprocess.run(api_sync_command, capture_output=True, text=True)

            if result.returncode != 0:
                print(f"Unknown Error code from api_db sync: {result}")
                sys.exit(1)
                return

            print("API Database sync successfully")

            if not has_cell("cell0"):
                result = subprocess.run(
                    map_cell0_command, capture_output=True, text=True
                )

                if result.returncode != 0:
                    print(f"Unknown Error code from map_cell0: {result}")
                    sys.exit(1)
                    return

                print("mapped cell0 successfully")

            if not has_cell("cell1"):
                result = subprocess.run(
                    create_cell_cell1_command, capture_output=True, text=True
                )

                if result.returncode != 0:
                    print(f"Unknown Error code from create cell cell1: {result}")
                    sys.exit(1)
                    return

                print("created cell1 successfully")

            result = subprocess.run(sync_command, capture_output=True, text=True)

            if result.returncode != 0:
                print(f"Unknown Error code from sync: {result}")
                sys.exit(1)
                return

            print("Database sync successfully")

            online_data_migrations_return = None
            while (
                online_data_migrations_return is None
                or online_data_migrations_return.returncode == 1
            ):
                result = subprocess.run(
                    online_data_migrations_command, capture_output=True, text=True
                )

                online_data_migrations_return = result

                print("Database ran online migrations")

            if online_data_migrations_return.returncode != 0:
                print(
                    f"Unknown Error code from online migrations: {online_data_migrations_return}"
                )
                sys.exit(1)
                return

            print("Database migrations complete")
        finally:
            released = consul_client.kv.put(
                lock_kv_path, "released", release=session_id
            )
            print(f"Lock released: {released}")
    finally:
        destroyed = consul_client.session.destroy(session_id)
        print(f"Session {session_id} destroyed: {destroyed}")


if __name__ == "__main__":
    main()
```

https://github.com/rmb938/proxmox-images/blob/main/ubuntu-noble-lts-amd64-openstack-nova/ansible/roles/nova/files/usr/local/bin/nova-migrate

---
---
# OpenStack Service Credentials
Services need to authenticate to each other, more python!

```python {*|73-75|84-90|105-106|169-175|180-189}{maxHeight:'300px'}
#! /usr/bin/python3


import secrets
import sys
import time
from datetime import datetime, timedelta

import consul
import hvac
from openstack import connection


def main():
    openstack_region = "us-homelab1"
    openstack_domain_name = "Default"
    openstack_service_project_name = "service"

    lock_kv_path = "openstack-keystone/service-users-token-lock"

    vault_client = hvac.Client(url="http://127.0.0.1:8100")

    consul_client = consul.Consul(
        host="127.0.0.1",
        port=8500,
        consistency="consistent",
    )

    session_id = consul_client.session.create(ttl=120, behavior="delete")

    try:
        acquired = consul_client.kv.put(lock_kv_path, "locked", acquire=session_id)

        if not acquired:
            # Not exiting 1 here because it's normal
            print(f"Could not aquire consul lock at {lock_kv_path}")
            return

        try:
            admin_password_secret = vault_client.secrets.kv.v1.read_secret(
                path=f"openstack-keystone/admin-password",
            )
            admin_password = admin_password_secret["data"]["password"]

            conn = connection.Connection(
                region_name=openstack_region,
                auth={
                    "auth_url": "https://openstack-keystone.haproxy.us-homelab1.hl.rmb938.me",  # TODO: just use local hostname
                    "username": "admin",
                    "password": admin_password,
                    "project_name": "admin",
                    "project_domain_name": openstack_domain_name,
                    "user_domain_name": openstack_domain_name,
                },
                identity_interface="internal",
            )

            domain = conn.identity.find_domain(openstack_domain_name)
            if not domain:
                print(f"Could not find domain '{openstack_domain_name}'")
                sys.exit(1)
                return

            project = conn.identity.find_project(openstack_service_project_name)
            if not project:
                print(
                    f"Project '{openstack_service_project_name}' does not exist, creating..."
                )
                project = conn.identity.create_project(
                    name=openstack_service_project_name, domain_id=domain.id
                )

            response_expected_users = vault_client.secrets.kv.v1.list_secrets(
                path="openstack-keystone/expected-service-users/"
            )

            raw_expected_users = response_expected_users["data"]["keys"]
            for raw_expected_user in raw_expected_users:
                openstack_user_name = raw_expected_user

                user = conn.identity.find_user(openstack_user_name, domain_id=domain.id)
                if not user:
                    print(f"User '{openstack_user_name}' does not exist, creating...")
                    user = conn.identity.create_user(
                        name=openstack_user_name,
                        password=secrets.token_hex(128),
                        domain_id=domain.id,
                        enabled=True,
                    )

                # TODO: idk if we still need admin
                # openstack documentation isn't clear
                admin_role = conn.identity.find_role("admin")
                if not admin_role:
                    print(f"Could not find admin role")
                    sys.exit(1)
                    return

                service_role = conn.identity.find_role("service")
                if not service_role:
                    print(f"Could not find service role")
                    sys.exit(1)
                    return

                conn.identity.assign_project_role_to_user(project, user, admin_role)
                conn.identity.assign_project_role_to_user(project, user, service_role)

                now = datetime.now()

                # Doing 7 days for now, it's long enough to not rotate constantly
                # But short enough to notice any issues
                credential_valid_length = timedelta(days=7)

                # The minimum amount of length left before we cut a new cred
                credential_min_length = credential_valid_length / 3

                existing_application_credentials = (
                    conn.identity.application_credentials(user=user)
                )

                print(
                    f"Finding if we need to rotate application credentials for user '{openstack_user_name}'"
                )
                newest_expires_at = now
                for eac in existing_application_credentials:

                    # If eac has no expires at skip it
                    if eac.expires_at is None:
                        continue

                    eac_expires_at = datetime.fromisoformat(eac.expires_at)

                    # If eac expires before the newest skip it
                    if eac_expires_at < newest_expires_at:
                        continue

                    # We are the newest so set it
                    newest_expires_at = eac_expires_at

                # If we are not going to expire soon, don't create a new credential
                if (newest_expires_at - now) > credential_min_length:
                    print(
                        f"User '{openstack_user_name}' does not need a new credential yet."
                    )
                    continue

                # Change the password to something random while we are here
                print(f"Rotating Password for '{openstack_user_name}'")
                new_password = secrets.token_hex(128)
                conn.identity.update_user(
                    user,
                    password=new_password,
                )

                # Sleeping due to token revocation race with connect_as
                time.sleep(1)

                expires_at = now + credential_valid_length

                # Connect as because only the owning user can create app creds
                # Admin user can't create it for them.
                connect_as = conn.connect_as(
                    username=openstack_user_name,
                    password=new_password,
                    project_name="service",
                )

                print(f"Rotating application credentials for '{openstack_user_name}'")
                application_credential = (
                    connect_as.identity.create_application_credential(
                        user=user,
                        name=f"service_creds_{time.time()}",
                        expires_at=expires_at.strftime("%Y-%m-%dT%H:%M:%S"),
                    )
                )

                print(
                    f"Writting application credentials to vault for '{openstack_user_name}'"
                )
                vault_client.secrets.kv.v1.create_or_update_secret(
                    path=f"openstack-keystone/service-users/{openstack_user_name}",
                    secret={
                        "application_credential_id": application_credential.id,
                        "application_credential_secret": application_credential.secret,
                        "created_at": now.strftime("%Y-%m-%dT%H:%M:%S"),
                        "expires_at": expires_at.strftime("%Y-%m-%dT%H:%M:%S"),
                        "ttl": "3600",  # Hint for lease_duration to check every 1 hour
                    },
                )

        finally:
            released = consul_client.kv.put(
                lock_kv_path, "released", release=session_id
            )
            print(f"Lock released: {released}")
    finally:
        destroyed = consul_client.session.destroy(session_id)
        print(f"Session {session_id} destroyed: {destroyed}")


if __name__ == "__main__":
    main()
```

https://github.com/rmb938/proxmox-images/blob/main/ubuntu-noble-lts-amd64-openstack-keystone/ansible/roles/keystone/files/usr/local/bin/keystone-service-users-token

---
---
# Monitoring OpenStack
What monitoring? No such thing!

Monitoring PostgreSQL, HAProxy, RabbitMQ is pretty easy. Prometheus exporters exist and have lots of pre-made Grafana dashboards.

<br />

Monitoring OpenStack services themselves are hard. Most are missing health probes, there's little to no built in metrics. Logs are good but sometimes lack detail.

<br />

The best way to monitor OpenStack is to use canary services that constantly poke at the API and perform various actions. Which means more custom tools that I have not built yet!

---
---
# OpenStack Networking is a Nightmare
Software without documentation is not a fun time

These days the standard SDN (software defined network) tooling that OpenStack uses is OVS (OpenVSwitch) and OVN (Open Virtual Network).

OVS uses a custom database called OVSDB that clusters weirdly, natively supports mTLS but doesn't validate hostnames, and works fine with DNS hostnames until it doesn't.

If you are going to deploy OVSDB in a cluster use IP addresses not hostnames, only use IPv4 and use TLS but the CN and SANs don't matter.

ChatGPT is not helpful parsing the mailing lists or documentation and you will constantly not know why it's not working until you find the one undocumented configuration item.

Make sure you have at least 3 compute hosts, with only two network failover doesn't happen. You don't need an odd number, just more then 2.

---
layout: quote
---
# Screenshots!
This is actually real, it's not theory!

---
layout: image

image: /images/Screenshot_20250527_220915.png

backgroundSize: contain
---

---
layout: image

image: /images/Screenshot_20250527_220756.png

backgroundSize: contain
---

---
layout: image

image: /images/Screenshot_20250527_221050.png

backgroundSize: contain
---

---
layout: image

image: /images/Screenshot_20250527_221139.png

backgroundSize: contain
---

---
layout: image

image: /images/Screenshot_20250429_111132.png

backgroundSize: contain
---

---
---
# Learnings & Challanges
What did you learn at school today?

Bootstrapping a cloud at home can be hard if you don't know where to start. What tools and systems to use can be overwhelming and you will get stuck in decision paralysis.

OpenStack as a whole can be fairly straight forward to install and setup, where it is difficult is it's new features and limited documentation.

Networking is one of the harder problems within OpenStack, the community has a standard and preferred solution but it's badly documented even though it's been production ready for many years.

<br />

Building a cloud at home is fun, there are lots of interesting challanges and weird issues.

---
---
# What's Next?
I get to use my cloud for something, what is the something? Who knows!

I need a new goal, my major goal is now complete and I'm looking for something else to do

OpenStack is not a fully feature complete cloud, instances don't have IAM profiles or service accounts, so I may build something for that so they can easily get secrets and authenticate to services.

I need some sort of Object Storage, Openstack has an implimentation of one called Swift but it's not home lab friendly at all. Maybe I'll automate Minio in some way.

I would like a Kubernetes as a Service service to easily spin up and scale clusters within OpenStack.

There's a lot of misc tools and scripts I want to build to make some of the automation a bit more straight forward and simpler.

---
layout: end
---

# Thanks You!
Q & A

Slides: https://github.com/rmb938/open-source-north
