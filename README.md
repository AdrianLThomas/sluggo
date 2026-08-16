# Sluggo!
Snake but for Slugs. Written in Go.

## Dependencies
Ebitengine requires Linux system dependencies to build and test. On Debian/Ubuntu-based distributions, you can install them with:

```bash
make setup-debian-deps
```

For other platforms, see the [Ebitengine installation guide](https://ebitengine.org/en/documents/install.html).

## Development
- `make run` - Run the game
- `make build` - Build the desktop binary (`sluggo.bin`)
- `make build-web` - Build the WebAssembly version (`web/game.wasm`)
- `make test` - Run tests
- `make test-ci` - Run tests headlessly using `xvfb-run` (for CI / headless Linux)

# Credits
- [Textures (Ground)](https://kenney.nl/assets/retro-textures-fantasy)
- [Textures (Food)](https://kenney.nl/assets/tiny-farm)
- [Fonts](https://www.kenney.nl/assets/kenney-fonts)