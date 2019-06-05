package main

// mysql database username
var dbUser = lookupEnvOrUseDefault("DB_USER", "")

// mysql database password
var dbPass = lookupEnvOrUseDefault("DB_PASSWORD", "")

// mysql database host
var dbHost = lookupEnvOrUseDefault("DB_HOST", "")

// mysql database port
var dbPort = lookupEnvOrUseDefault("DB_PORT", "3306")

// mysql database name
var dbName = lookupEnvOrUseDefault("DB_NAME", "")
