var db = connect("mongodb://root:rootPassXXX@localhost:27017/admin");

db = db.getSiblingDB('golangDB');

db.createUser(
    {
        user: "patrickbremm",
        pwd: "patrickbremmPWD",
        roles: [ { role: "userAdminAnyDatabase", db: "admin"} ],
        passwordDigestor: "server",
    }
)

db.auth('patrickbremm','patrickbremmPWD');

db.createUser(
    {
        user: "patrickbremm2",
        pwd: "patrickbremmPWD2",
        roles: [ { role:'readWrite', db: 'testdb'} ],
        passwordDigestor: "server",
    }
)

db.auth('patrickbremm2','patrickbremmPWD2');