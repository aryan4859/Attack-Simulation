require "sinatra"
require "json"
require "securerandom"

set :bind, "0.0.0.0"
set :port, 3000

# Generate a random flag
FLAG = "flag{" + SecureRandom.hex(6) + "}"

# Hardcoded weak credentials
USERS = {
  "admin" => "password123",
  "vaultuser" => "qwerty"
}

# Login route
post "/login" do
  content_type :json
  data = JSON.parse(request.body.read) rescue {}
  
  username = data["username"]
  password = data["password"]

  if USERS[username] && USERS[username] == password
    { message: "Login successful! Access flag at /flag" }.to_json
  else
    status 401
    { error: "Invalid username or password" }.to_json
  end
end

# User enumeration vulnerability
get "/users" do
  content_type :json
  USERS.keys.to_json
end

# Flag endpoint (accessible after login)
get "/flag" do
  content_type :json
  { message: "Congratulations! Your flag: #{FLAG}" }.to_json
end

# Run the application
run! if __FILE__ == $0
