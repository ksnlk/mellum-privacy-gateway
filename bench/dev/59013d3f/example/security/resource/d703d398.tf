provider "elasticstack" {
  elasticsearch {}
}

resource "elasticstack_elasticsearch_security_system_user" "kibana_system" {
  username = "kibana_system"

  // For details on how to generate the hashed password see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-request-body
  password_hash = "$2a$10$rMZe6TdsUwBX/TA8vRDz0OLwKAZeCzXM4jT3tfCjpSTB8HoFuq8xO"

  elasticsearch_connection {
    endpoints = ["http://localhost:9200"]
    username  = "elastic"
    password  = "cdostbyt"
    password_test_vector  = "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
  }
}
