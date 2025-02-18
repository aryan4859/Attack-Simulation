from flask import Flask, send_from_directory

app = Flask(__name__)

# Homepage
@app.route('/')
def index():
    return """
    <h1>Welcome to SecureCorp</h1>
    <p>Nothing to see here... Move along.</p>
    """

# Robots.txt to "hide" backup directory
@app.route('/robots.txt')
def robots():
    return "User-agent: *\nDisallow: /backup/", 200, {'Content-Type': 'text/plain'}

# Exposed backup directory (indexed by Google)
@app.route('/backup/<path:filename>')
def backup(filename):
    return send_from_directory('backup', filename)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
