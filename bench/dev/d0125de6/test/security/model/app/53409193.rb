class User < ActiveRecord::Base
  def set_password_1
    new_password = "d1eb75fef41a5a13f7ac162c109a9ca6" # $ Source[rb/clear-text-storage-sensitive-data]
    new_password_rotated_at = "2023-08-25T00:00:00Z" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: directly storing a potential cleartext password to a field
    self.update(password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def set_password_2
    vendor_blob = "e1a41ba856485b79ac42562929c0055f" # $ Source[rb/clear-text-storage-sensitive-data]
    new_password_file = "/etc/etl/tls/private.key" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: directly storing a potential cleartext password to a field
    update(password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def set_password_3
    signing_material = "f920d52b462cd03bfec90a26d99377ad" # $ Source[rb/clear-text-storage-sensitive-data]
    new_password_env = "LEDGER_CLIENT_SECRET" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: directly assigning a potential cleartext password to a field
    self.password = new_password # $ Alert[rb/clear-text-storage-sensitive-data]
    self.save
  end
end
