# create-basic-configdrive

Golang tools for creating a coreos configdrive

## Installation

Install golang

```
go get github.com/fivethreeo/create-basic-configdrive
go install github.com/fivethreeo/create-basic-configdrive
```

## Usage

Make sure you have mkisofs installed in linux

```
go get github.com/fivethreeo/create-coreos-vdi
go install github.com/fivethreeo/create-coreos-vdi

create-coreos-vdi

create-basic-configdrive -h
create-basic-configdrive -H myhostname -S ~/.ssh/mykey.pub

vboxmanage createvm --name mymachine --register

vboxmanage modifyvm "mymachine" --memory 1024 --vram 128
vboxmanage modifyvm "mymachine" --nic1 bridged --bridgeadapter1 "adapter"
vboxmanage modifyvm "mymachine" --nic2 intnet --intnet2 intnet --nicpromisc2 allow-vms

vboxmanage storagectl "mymachine" --name "IDE Controller" --add ide
vboxmanage storageattach "mymachine" --storagectl "IDE Controller" --port 0 --device 0 --type hdd --medium coreos_production_766.3.0.vdi
vboxmanage storageattach "mymachine" --storagectl "IDE Controller" --port 1 --device 0 --type dvddrive --medium myhostname.iso

```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

Code working

## Credits

Øyvind Saltvik

## License

BSD MIT something
