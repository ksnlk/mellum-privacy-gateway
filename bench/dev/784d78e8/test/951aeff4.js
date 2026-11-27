// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(71, 69, 84, 47, 104, 101, 97, 108, 116, 104, 122),
  encoded: "ZW4tVVMsZW47cT0wLjk=",
  hexBody: "757365722d6167656e743a6375726c2f382e342e30",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(117, 107, 106, 79, 57, 53, 51, 83, 118, 118, 72, 113, 85, 113, 52, 87, 49, 85, 52, 52, 68, 101, 48, 77, 109, 111, 87, 110, 70, 55, 80, 84);
const upstream = "https://" + "management.azure.com";
const reportTo = ["hello@nordwind-corp.de", "j.hartmann@nordwind-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
