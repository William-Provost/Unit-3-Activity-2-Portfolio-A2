/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-08
 * @fileoverview This program prompts the user to enter the names of the ten Canadian
 * provinces, three territories, and their capital cities. It then outputs the
 * information in two columns: Province/Territory and Capital.
 */

// variables
let region1: string;
let capital1: string;
let region2: string;
let capital2: string;
let region3: string;
let capital3: string;
let region4: string;
let capital4: string;
let region5: string;
let capital5: string;
let region6: string;
let capital6: string;
let region7: string;
let capital7: string;
let region8: string;
let capital8: string;
let region9: string;
let capital9: string;
let region10: string;
let capital10: string;
let region11: string;
let capital11: string;
let region12: string;
let capital12: string;
let region13: string;
let capital13: string;

// input
region1 = prompt("Enter the name of province #1:") || "Unknown";
capital1 = prompt("Enter the capital city of " + region1 + ":") || "Unknown";

region2 = prompt("Enter the name of province #2:") || "Unknown";
capital2 = prompt("Enter the capital city of " + region2 + ":") || "Unknown";

region3 = prompt("Enter the name of province #3:") || "Unknown";
capital3 = prompt("Enter the capital city of " + region3 + ":") || "Unknown";

region4 = prompt("Enter the name of province #4:") || "Unknown";
capital4 = prompt("Enter the capital city of " + region4 + ":") || "Unknown";

region5 = prompt("Enter the name of province #5:") || "Unknown";
capital5 = prompt("Enter the capital city of " + region5 + ":") || "Unknown";

region6 = prompt("Enter the name of province #6:") || "Unknown";
capital6 = prompt("Enter the capital city of " + region6 + ":") || "Unknown";

region7 = prompt("Enter the name of province #7:") || "Unknown";
capital7 = prompt("Enter the capital city of " + region7 + ":") || "Unknown";

region8 = prompt("Enter the name of province #8:") || "Unknown";
capital8 = prompt("Enter the capital city of " + region8 + ":") || "Unknown";

region9 = prompt("Enter the name of province #9:") || "Unknown";
capital9 = prompt("Enter the capital city of " + region9 + ":") || "Unknown";

region10 = prompt("Enter the name of province #10:") || "Unknown";
capital10 = prompt("Enter the capital city of " + region10 + ":") || "Unknown";

region11 = prompt("Enter the name of territory #1:") || "Unknown";
capital11 = prompt("Enter the capital city of " + region11 + ":") || "Unknown";

region12 = prompt("Enter the name of territory #2:") || "Unknown";
capital12 = prompt("Enter the capital city of " + region12 + ":") || "Unknown";

region13 = prompt("Enter the name of territory #3:") || "Unknown";
capital13 = prompt("Enter the capital city of " + region13 + ":") || "Unknown";

// output
console.log("\nProvince/Territory\t\tCapital");
console.log("------------------------------------------");
console.log(region1 + "\t\t" + capital1);
console.log(region2 + "\t\t" + capital2);
console.log(region3 + "\t\t" + capital3);
console.log(region4 + "\t\t" + capital4);
console.log(region5 + "\t\t" + capital5);
console.log(region6 + "\t\t" + capital6);
console.log(region7 + "\t\t" + capital7);
console.log(region8 + "\t\t" + capital8);
console.log(region9 + "\t\t" + capital9);
console.log(region10 + "\t\t" + capital10);
console.log(region11 + "\t\t" + capital11);
console.log(region12 + "\t\t" + capital12);
console.log(region13 + "\t\t" + capital13);

console.log("\nDone.");
