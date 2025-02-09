require 'sinatra'
require 'dotenv/load'
require 'securerandom' 

# Allow external access
set :bind, '0.0.0.0'
set :port, 4567

# Allow any host (especially for Sinatra 4.x)
set :public_folder, 'public'
set :protection, except: :http_origin
set :allow_origin, '*'

# Load users & passwords from .env
USERS = (ENV['USERS'] || '').split(',')
PASSWORDS = (ENV['PASSWORDS'] || '').split(',')

# Generate a random flag
FLAG = "flag{3NUMER8T10N&85UT3}"

# Route for login page
get '/' do
  erb :index
end

# Handle login
post '/login' do
  username = params[:username]
  password = params[:password]

  if USERS.include?(username) && PASSWORDS[USERS.index(username)] == password
    content_type :html
    "Login Successful! Your flag: <b>#{FLAG}</b>"
  else
    status 401
    "Invalid credentials! Try again."
  end
end
