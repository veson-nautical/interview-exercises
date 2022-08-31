import "./styles.css";

function log(...values: any[]) {
  document.getElementById("app")!.innerHTML +=
    values.map((v) => v?.toString()).join(" ") + "<br /><br />";
}

const admins: any[] = [
  { type: "admin", name: "Jane Doe", age: 32, role: "Administrator" },
  { type: "admin", name: "Bruce Willis", age: 64, role: "World saver" },
  { type: "admin", name: "Steve", age: 40, role: "Steve" },
  { type: "admin", name: "Will Bruces", age: 30, role: "Overseer" },
  { type: "admin", name: "Superwoman", age: 28, role: "Customer support" }
];

const users: any[] = [
  {
    type: "user",
    name: "Max Mustermann",
    age: 25,
    occupation: "Chimney sweep"
  },
  { type: "user", name: "Kate Müller", age: 23, occupation: "Astronaut" },
  { type: "user", name: "Moses", age: 70, occupation: "Desert guide" },
  { type: "user", name: "Superman", age: 28, occupation: "Ordinary person" },
  { type: "user", name: "Inspector Gadget", age: 31, occupation: "Undercover" }
];

log("ages: ", admins.map((a) => a.age).join(", "));

// how should we type items?
function getMaxAge(items: any[]): number {
  // implement the logic
  return items.length; // clearly wrong :)
}

log("max ages: ", getMaxAge(admins), getMaxAge(users));

// -------------------------------------------------
// Basic Algorithms
// -------------------------------------------------

function getUserWithMaxAge(items: any[]): any {
  return items[0];
}

log("max age user: ", JSON.stringify(getUserWithMaxAge(admins)));

function getNameOfUserWithMaxAge(items: any[]): string {
  return items[0].name;
}

// -------------------------------------------------
// Typescript Static Typing
// -------------------------------------------------

log("max age user name: ", getNameOfUserWithMaxAge(admins));

function getFieldOfUserWithMaxAge(field: string, items: any[]): any {
  return items[0][field];
}

log("max age user [field=role]: ", getFieldOfUserWithMaxAge("role", admins));
