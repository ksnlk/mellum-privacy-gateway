// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(71, 69, 84, 47, 104, 101, 97, 108, 116, 104, 122),
  encoded: "R0VUL2hlYWx0aHo=",
  hexBody: "757365722d6167656e743a6375726c2f382e342e30",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(81, 110, 113, 56, 107, 69, 117, 54, 121, 100, 85, 115, 122, 75, 70, 90, 102, 70, 79, 78, 72, 69, 99, 69, 52, 89, 76, 108, 52, 57, 65, 51);
const upstream = "https://" + "s3.eu-central-1.amazonaws.com";
const reportTo = ["press@nordwind-corp.de", "j.becker@nordwind-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
