from flask import Flask, request, jsonify, abort
import requests

app = Flask(__name__)

FLAG = "FLAG{55R7_PWN3D_200OK}"

ALLOWED_IPS = ["127.0.0.1", "localhost"]  # Add internal server IPs if needed

@app.route("/")
def home():
    return "Welcome to Sneaky Request! Try to access the admin panel."

@app.route("/fetch", methods=["GET"])
def fetch():
    url = request.args.get("url")
    if not url:
        return jsonify({"error": "URL parameter is required"}), 400

    try:
        response = requests.get(url, timeout=3)  # Server makes an HTTP request
        return jsonify({"status": response.status_code, "content": response.text[:300]})
    except requests.exceptions.RequestException:
        return jsonify({"error": "Failed to fetch URL"}), 500

@app.route("/admin")
def admin():
    if request.remote_addr not in ALLOWED_IPS:
        abort(403)  # Forbidden

    return f"Admin Panel: The flag is {FLAG}"

if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0")
