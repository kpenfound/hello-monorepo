from http.server import BaseHTTPRequestHandler, HTTPServer
import socketserver
import subprocess

PORT = 8000

class MyHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            result = subprocess.run(['/usr/bin/go-uname'], capture_output=True, text=True)
            self.send_response(200)
            self.send_header("Content-type", "text/html")
            self.end_headers()
            self.wfile.write(bytes(result.stdout, "utf-8"))

with socketserver.TCPServer(("0.0.0.0", PORT), MyHandler) as httpd:
    print("serving at port", PORT)
    httpd.serve_forever()