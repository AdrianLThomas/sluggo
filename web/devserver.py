"""Dev server for the web build.

Plain http.server sends no Cache-Control, so browsers fall back to heuristic
caching and a normal reload can keep serving stale CSS or JS. That hides edits
until a hard refresh, so disable caching here instead.
"""

import functools
import http.server
import pathlib
import sys


class NoCacheHandler(http.server.SimpleHTTPRequestHandler):
    def end_headers(self):
        self.send_header("Cache-Control", "no-store")
        super().end_headers()


def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 9090
    web_dir = pathlib.Path(__file__).resolve().parent
    handler = functools.partial(NoCacheHandler, directory=str(web_dir))
    with http.server.ThreadingHTTPServer(("", port), handler) as server:
        print(f"serving {web_dir} on http://localhost:{port} (caching disabled)")
        server.serve_forever()


if __name__ == "__main__":
    main()
