import ballerina/http;
import ballerina/log;
import ballerina/mime;

service / on new http:Listener(9090) {

    resource function post submit(http:Request req) returns string|error {
        mime:Entity[] bodyParts = check req.getBodyParts();

        foreach mime:Entity item in bodyParts {
            string contentType = item.getContentType();
            log:printInfo("Content-Type: " + contentType);
            mime:ContentDisposition contentDisposition = item.getContentDisposition();

            if contentType == mime:APPLICATION_JSON {
                log:printInfo("Content-Disposition JSON filename: " + contentDisposition.fileName);
                log:printInfo("Content-Disposition JSON name: " + contentDisposition.name);
            } else {
                log:printInfo("Content-Disposition Other filename: " + contentDisposition.fileName);
                log:printInfo("Content-Disposition Other name: " + contentDisposition.name);
                if contentDisposition.name == "claim" {
                    json content = check item.getJson();
                    log:printInfo("JSON: " + content.toString());
                } else {
                    byte[] content = check item.getByteArray();
                    log:printInfo("File len: " + content.length().toString());
                }
            }
        }
        return "Successfully submitted the form data.";
    }
}