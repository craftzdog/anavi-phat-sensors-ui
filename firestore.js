const path = require("path");
const admin = require("firebase-admin");
const serviceAccount = require(path.join(
  __dirname,
  "service-account-key.json"
));

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount),
});

const db = admin.firestore();

module.exports = db;
