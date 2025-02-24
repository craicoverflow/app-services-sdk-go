
package error

// registrymgmt error codes 
  
const (

  // Unspecified error
  ERROR_1 string = "SRS-MGMT-1"

  // Registry with id='?' not found
  ERROR_2 string = "SRS-MGMT-2"

  // Bad date or time format
  ERROR_3 string = "SRS-MGMT-3"

  // Invalid request content. Make sure the request conforms to the given JSON schema
  ERROR_4 string = "SRS-MGMT-4"

  // Bad request format - invalid JSON
  ERROR_5 string = "SRS-MGMT-5"

  // Required terms have not been accepted for account id='?'
  ERROR_6 string = "SRS-MGMT-6"

  // The maximum number of allowed Registry instances has been reached
  ERROR_7 string = "SRS-MGMT-7"

  // Error type with id='?' not found
  ERROR_8 string = "SRS-MGMT-8"

  // Data conflict. Make sure a Registry with the given name does not already exist
  ERROR_9 string = "SRS-MGMT-9"

  // Bad request format - unsupported media type
  ERROR_10 string = "SRS-MGMT-10"

)