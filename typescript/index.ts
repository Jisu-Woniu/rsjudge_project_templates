import { createInterface } from "readline/promises";

const read = createInterface(process.stdin);

for await (const line of read) {
  console.log(line);
  throw new Error("Error");
}
