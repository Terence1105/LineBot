db.createUser(
    {
        user: "sa",
        pwd: "pass.123",
        roles: [
            {
                role: "readWrite",
                db: "LineBot"
            }
        ]
    }
);

db.createCollection("MessageLog");

db.MessageLog.createIndex({"ts":1, "userId":1})