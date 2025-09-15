from http.server import ThreadingHTTPServer, BaseHTTPRequestHandler
from urllib.parse import urlparse, parse_qs
import re, os

# Load the flag
if os.path.exists("/flag"):
    with open("/flag") as f:
        FLAG = f.read().strip()
else:
    FLAG = os.environ.get("FLAG", "flag{this_is_a_fake_flag}")

# URL validation regex
URL_REGEX = re.compile(r"https?://[a-zA-Z0-9.]+(/[a-zA-Z0-9./?#]*)?")

class RequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # Serve flag if requested
        if self.path == "/flag":
            self.send_response(200)
            self.end_headers()
            self.wfile.write(FLAG.encode())
            return

        # Handle optional redirection
        query = parse_qs(urlparse(self.path).query)
        redir = query.get("redir", [None])[0]
        if redir and not URL_REGEX.match(redir):
            redir = None

        self.send_response(302 if redir else 200)
        if redir:
            self.send_header("Location", redir)
        self.end_headers()
        self.wfile.write(b"Hello world!")

    def do_HEAD(self):
        # Respond to HEAD requests without a body
        if self.path == "/flag":
            self.send_response(200)
        else:
            query = parse_qs(urlparse(self.path).query)
            redir = query.get("redir", [None])[0]
            if redir and not URL_REGEX.match(redir):
                redir = None
            self.send_response(302 if redir else 200)
            if redir:
                self.send_header("Location", redir)
        self.end_headers()

if __name__ == "__main__":
    # Use PORT env variable or default to 7777
    port = int(os.environ.get("PORT", 7777))
    server = ThreadingHTTPServer(("0.0.0.0", port), RequestHandler)
    server.allow_reuse_address = True
    print(f"Starting server on port {port}, use <Ctrl-C> to stop")
    server.serve_forever()
