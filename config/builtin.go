package config

var _defulatConfig = `
application: "echo realworld!!!"
http_server:
  address: "0.0.0.0:8080"
  read_timeout: 1s
  write_timeout: 3s
  idle_timeout: 2s
logging:
  level: "debug"
database:
  host: "localhost"
  port: 3306
  user: "realworld"
  password: "realworld"
  dbname: "realworld"
  charset: "utf8mb4"
  collation: "utf8mb4_unicode_ci"
  parse_time: true
  location: "Asia/Tehran"
  max_life_time: 5m
  max_idel_time: 0s
  max_open_connection: 10
  max_idel_connection: 5
  logger:
    slow_threshold: 200ms
    level: "info"
    colorfule: true
    ignore_record_not_found_error: false`