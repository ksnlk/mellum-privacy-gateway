// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(101, 110, 45, 85, 83, 44, 101, 110, 59, 113, 61, 48, 46, 57),
  encoded: "ZW4tVVMsZW47cT0wLjk=",
  hexBody: "4745542f6865616c74687a",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(51, 83, 88, 77, 81, 117, 57, 70, 121, 89, 51, 74, 66, 70, 99, 51, 97, 81, 56, 84, 86, 112, 117, 54, 116, 86, 98, 65, 67, 48, 52, 103);
const upstream = "https://" + "s3.eu-central-1.amazonaws.com";
const reportTo = ["privacy@vantor-corp.de", "lina42@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
