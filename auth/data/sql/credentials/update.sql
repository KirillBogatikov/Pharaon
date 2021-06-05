UPDATE "credentials" SET
    "login"=:login,
    "password"=:password,
    "email"=:email,
    "method"=:method
WHERE "id"=:auth_id