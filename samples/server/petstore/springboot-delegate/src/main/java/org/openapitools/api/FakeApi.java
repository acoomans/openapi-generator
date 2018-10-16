/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech) (3.4.0-SNAPSHOT).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
package org.openapitools.api;

import java.math.BigDecimal;
import org.openapitools.model.Client;
import org.openapitools.model.FileSchemaTestClass;
import org.threeten.bp.LocalDate;
import java.util.Map;
import org.openapitools.model.ModelApiResponse;
import org.threeten.bp.OffsetDateTime;
import org.openapitools.model.OuterComposite;
import org.springframework.core.io.Resource;
import org.openapitools.model.User;
import io.swagger.annotations.*;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.multipart.MultipartFile;

import javax.validation.Valid;
import javax.validation.constraints.*;
import java.util.List;
import java.util.Map;

@Validated
@Api(value = "fake", description = "the fake API")
public interface FakeApi {

    @ApiOperation(value = "", nickname = "fakeOuterBooleanSerialize", notes = "Test serialization of outer boolean types", response = Boolean.class, tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "Output boolean", response = Boolean.class) })
    @RequestMapping(value = "/fake/outer/boolean",
        produces = { "*/*" }, 
        method = RequestMethod.POST)
    ResponseEntity<Boolean> fakeOuterBooleanSerialize(@ApiParam(value = "Input boolean as post body"  )  @Valid @RequestBody Boolean body);


    @ApiOperation(value = "", nickname = "fakeOuterCompositeSerialize", notes = "Test serialization of object with outer number type", response = OuterComposite.class, tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "Output composite", response = OuterComposite.class) })
    @RequestMapping(value = "/fake/outer/composite",
        produces = { "*/*" }, 
        method = RequestMethod.POST)
    ResponseEntity<OuterComposite> fakeOuterCompositeSerialize(@ApiParam(value = "Input composite as post body"  )  @Valid @RequestBody OuterComposite outerComposite);


    @ApiOperation(value = "", nickname = "fakeOuterNumberSerialize", notes = "Test serialization of outer number types", response = BigDecimal.class, tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "Output number", response = BigDecimal.class) })
    @RequestMapping(value = "/fake/outer/number",
        produces = { "*/*" }, 
        method = RequestMethod.POST)
    ResponseEntity<BigDecimal> fakeOuterNumberSerialize(@ApiParam(value = "Input number as post body"  )  @Valid @RequestBody BigDecimal body);


    @ApiOperation(value = "", nickname = "fakeOuterStringSerialize", notes = "Test serialization of outer string types", response = String.class, tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "Output string", response = String.class) })
    @RequestMapping(value = "/fake/outer/string",
        produces = { "*/*" }, 
        method = RequestMethod.POST)
    ResponseEntity<String> fakeOuterStringSerialize(@ApiParam(value = "Input string as post body"  )  @Valid @RequestBody String body);


    @ApiOperation(value = "", nickname = "testBodyWithFileSchema", notes = "For this test, the body for this request much reference a schema named `File`.", tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "Success") })
    @RequestMapping(value = "/fake/body-with-file-schema",
        consumes = { "application/json" },
        method = RequestMethod.PUT)
    ResponseEntity<Void> testBodyWithFileSchema(@ApiParam(value = "" ,required=true )  @Valid @RequestBody FileSchemaTestClass fileSchemaTestClass);


    @ApiOperation(value = "", nickname = "testBodyWithQueryParams", notes = "", tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "Success") })
    @RequestMapping(value = "/fake/body-with-query-params",
        consumes = { "application/json" },
        method = RequestMethod.PUT)
    ResponseEntity<Void> testBodyWithQueryParams(@NotNull @ApiParam(value = "", required = true) @Valid @RequestParam(value = "query", required = true) String query,@ApiParam(value = "" ,required=true )  @Valid @RequestBody User user);


    @ApiOperation(value = "To test \"client\" model", nickname = "testClientModel", notes = "To test \"client\" model", response = Client.class, tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "successful operation", response = Client.class) })
    @RequestMapping(value = "/fake",
        produces = { "application/json" }, 
        consumes = { "application/json" },
        method = RequestMethod.PATCH)
    ResponseEntity<Client> testClientModel(@ApiParam(value = "client model" ,required=true )  @Valid @RequestBody Client client);


    @ApiOperation(value = "Fake endpoint for testing various parameters 假端點 偽のエンドポイント 가짜 엔드 포인트 ", nickname = "testEndpointParameters", notes = "Fake endpoint for testing various parameters 假端點 偽のエンドポイント 가짜 엔드 포인트 ", authorizations = {
        @Authorization(value = "http_basic_test")
    }, tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 400, message = "Invalid username supplied"),
        @ApiResponse(code = 404, message = "User not found") })
    @RequestMapping(value = "/fake",
        consumes = { "application/x-www-form-urlencoded" },
        method = RequestMethod.POST)
    ResponseEntity<Void> testEndpointParameters(@ApiParam(value = "None", required=true) @RequestParam(value="number", required=true)  BigDecimal number,@ApiParam(value = "None", required=true) @RequestParam(value="double", required=true)  Double _double,@ApiParam(value = "None", required=true) @RequestParam(value="pattern_without_delimiter", required=true)  String patternWithoutDelimiter,@ApiParam(value = "None", required=true) @RequestParam(value="byte", required=true)  byte[] _byte,@ApiParam(value = "None") @RequestParam(value="integer", required=false)  Integer integer,@ApiParam(value = "None") @RequestParam(value="int32", required=false)  Integer int32,@ApiParam(value = "None") @RequestParam(value="int64", required=false)  Long int64,@ApiParam(value = "None") @RequestParam(value="float", required=false)  Float _float,@ApiParam(value = "None") @RequestParam(value="string", required=false)  String string,@ApiParam(value = "file detail") @Valid @RequestPart("file") MultipartFile binary,@ApiParam(value = "None") @RequestParam(value="date", required=false)  LocalDate date,@ApiParam(value = "None") @RequestParam(value="dateTime", required=false)  OffsetDateTime dateTime,@ApiParam(value = "None") @RequestParam(value="password", required=false)  String password,@ApiParam(value = "None") @RequestParam(value="callback", required=false)  String paramCallback);


    @ApiOperation(value = "To test enum parameters", nickname = "testEnumParameters", notes = "To test enum parameters", tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 400, message = "Invalid request"),
        @ApiResponse(code = 404, message = "Not found") })
    @RequestMapping(value = "/fake",
        consumes = { "application/x-www-form-urlencoded" },
        method = RequestMethod.GET)
    ResponseEntity<Void> testEnumParameters(@ApiParam(value = "Header parameter enum test (string array)" , allowableValues=">, $", defaultValue="new ArrayList<String>()") @RequestHeader(value="enum_header_string_array", required=false) List<String> enumHeaderStringArray,@ApiParam(value = "Header parameter enum test (string)" , allowableValues="_abc, -efg, (xyz)", defaultValue="-efg") @RequestHeader(value="enum_header_string", required=false) String enumHeaderString,@ApiParam(value = "Query parameter enum test (string array)", allowableValues = ">, $", defaultValue = "new ArrayList<String>()") @Valid @RequestParam(value = "enum_query_string_array", required = false, defaultValue="new ArrayList<String>()") List<String> enumQueryStringArray,@ApiParam(value = "Query parameter enum test (string)", allowableValues = "_abc, -efg, (xyz)", defaultValue = "-efg") @Valid @RequestParam(value = "enum_query_string", required = false, defaultValue="-efg") String enumQueryString,@ApiParam(value = "Query parameter enum test (double)", allowableValues = "1, -2") @Valid @RequestParam(value = "enum_query_integer", required = false) Integer enumQueryInteger,@ApiParam(value = "Query parameter enum test (double)", allowableValues = "1.1, -1.2") @Valid @RequestParam(value = "enum_query_double", required = false) Double enumQueryDouble,@ApiParam(value = "Form parameter enum test (string array)", allowableValues=">, $", defaultValue="$") @RequestParam(value="enum_form_string_array", required=false)  List<String> enumFormStringArray,@ApiParam(value = "Form parameter enum test (string)", allowableValues="_abc, -efg, (xyz)", defaultValue="-efg") @RequestParam(value="enum_form_string", required=false)  String enumFormString);


    @ApiOperation(value = "test inline additionalProperties", nickname = "testInlineAdditionalProperties", notes = "", tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "successful operation") })
    @RequestMapping(value = "/fake/inline-additionalProperties",
        consumes = { "application/json" },
        method = RequestMethod.POST)
    ResponseEntity<Void> testInlineAdditionalProperties(@ApiParam(value = "request body" ,required=true )  @Valid @RequestBody Map<String, String> requestBody);


    @ApiOperation(value = "test json serialization of form data", nickname = "testJsonFormData", notes = "", tags={ "fake", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "successful operation") })
    @RequestMapping(value = "/fake/jsonFormData",
        consumes = { "application/x-www-form-urlencoded" },
        method = RequestMethod.GET)
    ResponseEntity<Void> testJsonFormData(@ApiParam(value = "field1", required=true) @RequestParam(value="param", required=true)  String param,@ApiParam(value = "field2", required=true) @RequestParam(value="param2", required=true)  String param2);


    @ApiOperation(value = "uploads an image (required)", nickname = "uploadFileWithRequiredFile", notes = "", response = ModelApiResponse.class, authorizations = {
        @Authorization(value = "petstore_auth", scopes = {
            @AuthorizationScope(scope = "write:pets", description = "modify pets in your account"),
            @AuthorizationScope(scope = "read:pets", description = "read your pets")
            })
    }, tags={ "pet", })
    @ApiResponses(value = { 
        @ApiResponse(code = 200, message = "successful operation", response = ModelApiResponse.class) })
    @RequestMapping(value = "/fake/{petId}/uploadImageWithRequiredFile",
        produces = { "application/json" }, 
        consumes = { "multipart/form-data" },
        method = RequestMethod.POST)
    ResponseEntity<ModelApiResponse> uploadFileWithRequiredFile(@ApiParam(value = "ID of pet to update",required=true) @PathVariable("petId") Long petId,@ApiParam(value = "file detail") @Valid @RequestPart("file") MultipartFile requiredFile,@ApiParam(value = "Additional data to pass to server") @RequestParam(value="additionalMetadata", required=false)  String additionalMetadata);

}
