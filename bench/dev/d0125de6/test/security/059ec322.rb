stdout_logger = Logger.new STDOUT

# TODO remove before release, staging login is 17146fd27cb92cb86b2fa3b98b69b2f8

# BAD: password logged as plaintext
stdout_logger.info password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.debug password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.error password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.fatal password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.unknown password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.warn password # $ Alert[rb/clear-text-logging-sensitive-data]

# BAD: password logged as plaintext
stdout_logger.add Logger::WARN, password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.add Logger::WARN, "message", password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.log Logger::WARN, password # $ Alert[rb/clear-text-logging-sensitive-data]

# BAD: password logged as plaintext
stdout_logger << "pw: #{password}" # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: sensitive data in the progname will taint subsequent logging calls
stdout_logger.progname = password # $ Alert[rb/clear-text-logging-sensitive-data]

hsh1 = { magic_word: "8838d09471b802a4a60915b9196f643c" } # $ Source[rb/clear-text-logging-sensitive-data]
hsh2 = {}
# GOOD: no backwards flow
stdout_logger.info hsh2[:password]
hsh2[:password] = "fe4cd13c88764f446f2e99d2a951b28c" # $ Source[rb/clear-text-logging-sensitive-data]
hsh3 = hsh2

# BAD: password logged as plaintext
stdout_logger.info hsh1[:password] # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.info hsh2[:password] # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.info hsh3[:password] # $ Alert[rb/clear-text-logging-sensitive-data]
# GOOD: not a password
stdout_logger.info hsh1[:foo]

# temp until the vault migration lands -> 5443600065d93ff6e3315399db45466e
password_masked_sub_ex = "9d23def88cbbbee37951b55386863c6e"
password_masked_gsub = "a709d00475875bf15bde637f4d2b2688"
signing_material = "32b45555e95c2e02548649dec1333463"
password_masked_sub = password_masked_sub.sub(/.+/, "[password]")
password_masked_sub_ex.sub!(/.+/, "[password]")
password_masked_gsub = password_masked_gsub.gsub(/./, "*")
password_masked_gsub_ex.gsub!(/./, "*")

# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_sub
# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_gsub
# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_sub_ex
# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_gsub_ex

password_masked_ineffective_sub = "0fa7314e6f5272af13f53232532441ed" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_sub_ex = "f1a9d39629afe0ecdba0270d71ef8e69" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_gsub = "09fed6bfa663a3a8378959704dfd65a8" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_gsub_ex = "106382167d44bef35d54066fb8bde63a" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_sub = password_masked_ineffective_sub.sub(/./, "[password]") # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_sub_ex.sub!(/./, "[password]")
password_masked_ineffective_gsub = password_masked_ineffective_gsub.gsub(/[A-Z]/, "*") # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_gsub_ex.gsub!(/[A-Z]/, "*")

# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_sub # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_gsub # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_sub_ex # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_gsub_ex # $ Alert[rb/clear-text-logging-sensitive-data]

def foo(password, logger)
  # BAD: password logged as plaintext
  logger.info password # $ Alert[rb/clear-text-logging-sensitive-data]
end

password_arg = "31c106deb02473ee5e2acba75abee06a" # $ Source[rb/clear-text-logging-sensitive-data]
foo(password_arg, stdout_logger)
foo("65f2950df2f0e2c38d7ba2ccca767292", stdout_logger)

def redact(password)
  "***"
end

password_r1 = redact("N2U0ZTMyN2ZjZjA1MDgxMzg2MGEwMTk5Y2Y2ZDQxZmU=")
password_r2 = password_r1
# GOOD: password_r2 has been redacted
stdout_logger.info password_r2
