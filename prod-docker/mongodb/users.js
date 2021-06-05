admin = db.getSiblingDB('admin')
if (admin.getUser("mongodb") === null) {
    admin.createUser(
        {
            user: "mongodb",
            pwd: "dbm6jwN6mEhbXxCu",
            roles: [{role: "userAdminAnyDatabase", db: "admin"}]
        }
    )
}

proxy = db.getSiblingDB("proxy")
if (proxy.getUser("proxy") === null) {
    proxy.createUser(
        {
            user: "proxy",
            pwd: "BGZendgysf9m2jr6",
            roles: [{
                role: "userAdminAnyDatabase",
                db: "admin"
            }]
        }
    )
}