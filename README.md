# Cameraman

This is a mod for ZDoom, which helps to perform camerawork in-engine.

Download the latest version: [cameraman-1.3.zip](https://github.com/borogk/cameraman/releases/download/v1.3/cameraman-1.3.zip).

### What is Cameraman?

**Cameraman** is a mod for [ZDoom](https://zdoom.org/) family of source ports.

Use it to create camera profiles, play them back in-engine and capture as clips for your videos or movies!

Possible use cases:

1. Make trailers for Doom engine maps and mods
2. Create machinima in Doom engine
3. Capture Doom gameplay from 3rd person
4. _...other creative uses you may think of!_

### Manual

- [Chapter 1 - Installation](docs/ch01.installation.md)
- [Chapter 2 - Quick start](docs/ch02.quick-start.md)
- [Chapter 3 - Editor](docs/ch03.editor.md)
- [Chapter 4.1 - Linear mode](docs/ch04.01.linear.md)
- [Chapter 4.2 - Radial mode](docs/ch04.02.radial.md)
- [Chapter 4.3 - Bezier mode](docs/ch04.03.bezier.md)
- [Chapter 4.4 - Point and shoot mode](docs/ch04.04.point-and-shoot.md)
- [Chapter 5 - Player](docs/ch05.player.md)
- [Chapter 6 - For developers](docs/ch06.developers.md)
- [Appendix A - List of CVARs](docs/ap01.cvars.md)
- [Appendix B - Usage with bash](docs/ap02.bash.md)

### ZDoom engine compatibility

Cameraman is compatible with all major branches of ZDoom:

1. [UZDoom](https://github.com/UZDoom/UZDoom)
2. [GZDoom](https://doomwiki.org/wiki/GZDoom)
3. [LZDoom](https://doomwiki.org/wiki/LZDoom)
4. [ZDoom](https://doomwiki.org/wiki/ZDoom) (though discontinued, the last released version was confirmed to work)

> [!NOTE]
> Wide compatibility was possible thanks to relying on classic tech like ACS and DECORATE. Thus, any lesser known ZDoom forks are likely to be compatible as well.

### DSDA-Doom engine compatibility (via cm-doom)

[cm-doom](https://github.com/borogk/cm-doom) is a special fork of [dsda-doom](https://github.com/kraflab/dsda-doom) with Cameraman support in mind.

It doesn't have editor component like **Cameraman**, instead it should be used alongside it:

1. Setup camera in Cameraman's editor.
2. Export camera profile to a separate file.
3. Load exported file into **cm-doom** for playback.

This setup is more complicated, but offers following benefits thanks to inheriting DSDA features:

1. Strict demo compatibility.
2. Viddump allows to directly export video clips with consistent framerate, without relying on screen-capture software.
3. More faithful visuals thanks to DSDA rendering implementation.

> [!NOTE]
> See [cm-doom's README on Github](https://github.com/borogk/cm-doom/blob/cm-doom-v1/README.md) for more help on how to install and use it.

### Seeking help?

1. Read the manual
2. Create an issue right here on GitHub ([link to create](https://github.com/borogk/cameraman/issues/new))

### Author and contributors

Originally created by **borogk** in 2021.

Contributed by **[m0rb](https://github.com/m0rb)**.
