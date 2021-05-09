const db = require("./firestore");

async function main() {
  const docs = await db
    .collection("conditions")
    .orderBy("createdAt", "desc")
    .limit(100)
    .get();
  console.log("Recent conditions:", docs.size);
  docs.forEach((doc) => {
    console.log(doc.id + ":", doc.data());
  });
}

main();
