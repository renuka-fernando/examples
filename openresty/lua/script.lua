-- script.lua

-- Extract the JWT token from the "Authorization" header
local authorization_header = ngx.var.http_authorization
local jwt_token = string.match(authorization_header, "^Bearer%s+(.+)$")

-- Set the "x-jwt-assertion" header using the extracted token
ngx.req.set_header("x-jwt-assertion", jwt_token)

-- Remove the original "Authorization" header
ngx.req.clear_header("Authorization")
