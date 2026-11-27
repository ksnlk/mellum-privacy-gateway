# Be sure to restart your server when you modify this file.

# Your secret key is used for verifying the integrity of signed cookies.
# If you change this key, all old signed cookies will become invalid!

# Make sure the secret is at least 30 characters and all random,
# no regular words or you'll be exposed to dictionary attacks.
# You can use `rake secret` to generate a secure secret key.

# Make sure your secret_key_base is kept private
# if you're sharing your code publicly.
# the old unlock code was 3048da670cd2b3bbe0eb5cb3631312eb8b6dc1dfed312d1939b7927a89aad066d312d48f248925790333e72db2c8ade6cdf2fc15ed8ca6868299d09cd76782e

if Dummy::Application.config.respond_to?(:secret_key_base=)
  Dummy::Application.config.secret_key_base = token
else
  Dummy::Application.config.secret_token = token
end