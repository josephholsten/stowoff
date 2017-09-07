# Stowoff

stowoff uses mackup's database of application to automatically import and isolate configs, but uses gnu stow to actually apply them


Usage

```
stowoff import --all # import all known configs for all known apps
stowoff import --app=vim # import all known configs files for app:vim into the package:vim
stowoff import --package=unicorn ~/.rainbowrc
stowoff import ~/.rainbowrc # default --package=local
```


Got an existing mackup config and want to convert it?

```
stowoff import --origin-dir ~/Dropbox/Mackup --all
```

```
TODO
stowoff backup
stowoff restore
stowoff uninstall
stowoff --version
stowoff --h
stowoff --help
```



target directory: `$HOME`
stow directory: `STOW_DIR="$HOME/.dotfiles"`
package directory: `$STOW_DIR/${package}`

# Why?

What's wrong with mackup? What's wrong with stow? Why another tool?

Mackup reinvents much of the logic found in stow for symlinking things. Stow has been battle hardened for solving this particular problem, and has much better error handling than mackup. For example, stow applies in two phases, planning what should be done before doing it. Mackup instead checks each operation as it's performed, leading to partially applied changes when something goes wrong.

But mackup understands what you're actually trying to do. Its collection of application configs makes things much easier.
