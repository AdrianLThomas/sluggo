# Sluggo!
Snake but for Slugs. Written in Go.

## Dependencies
Ebitengine requires Linux system dependencies to build and test. On Debian/Ubuntu-based distributions, you can install them with:

```bash
make setup-debian-deps
```

`make setup-debian-deps` also installs `binaryen` (`wasm-opt`), which `make build-web` uses to shrink the wasm binary. Regenerating the embedded audio assets additionally needs `fluidsynth`, `fluid-soundfont-gm`, `ffmpeg` and `python3-mido`:

```bash
make setup-audio-assets   # install the extra packages
make regenerate-audio     # rebuild assets/*.ogg from tools/*.py
```

For other platforms, see the [Ebitengine installation guide](https://ebitengine.org/en/documents/install.html).

## Development
`make run` - To run the game.

See `makefile` for other build options.


# Credits
- [Textures (Ground)](https://kenney.nl/assets/retro-textures-fantasy)
- [Textures (Food)](https://kenney.nl/assets/tiny-farm)
- [Fonts](https://www.kenney.nl/assets/kenney-fonts)
- Sound effects and music: generated in-repo (`tools/gen_sfx.py`, `tools/gen_music.py`), rendered with FluidSynth soundfont (FluidR3_GM), encoded to OGG