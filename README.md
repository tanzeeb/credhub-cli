**Pivotal Cloud Foundry CredHub CLI helps you configure and interact with CredHub APIs.**

*Starred commands and parameters are planned, but not yet implemented*

```
Usage: credhub [<options>] <command> [<args>]
		--version				    	 		Show version of CLI and API
		-h, --help								Displays help menu


GETTING STARTED: 

 	api
		View or set the targeted CredHub API
		-s, --server 'URI'						URI of API server to target

	login
		Authenticates interactively with CredHub
		-u, --user 'USER'           			Authentication username
    	-p, --password 'PASSWORD'				Authentication password
		-s, --server 'URI'						URI of API server to target

	logout
		Discard authenticated user session


CREDENTIAL MANAGEMENT:

	get --name <cred name>
		Get the value and attributes of a credential
		-n, --name 'CRED'						Name of credential to retrieve

	set --type <cred type> --name <cred name> [set params]
		Set the value and attributes of a credential
		-t, --type ['value', 'certificate']		Sets the type of credential to store (Default: 'value')
		-n, --name 'CRED'						Selects the credential being set
		-O, --no-overwrite						Credential is not modified if stored value already exists

		Set parameters by [Type]
		-v, --value 'VALUE'						[Value] Sets the value for the credential
		-r, --root	<FILE>						[Certificate] Sets the root CA from file
		-c, --certificate <FILE>				[Certificate] Sets the certificate from file
		-p, --private <FILE>					[Certificate] Sets the private key from file
		-R, --root-string 'ROOT'				[Certificate] Sets the root CA from string input
		-C, --certificate-string 'CERT'       	[Certificate] Sets the certificate from string input
		-P, --private-string 'PRIVATE'			[Certificate] Sets the private key from string input

	generate --type <cred type> --name <cred name> [generate params]
		Generate and set a credential value based on generation parameters
		-t, --type ['value', 'certificate']		Sets the type of credential to generate (Default: 'value')
		-n, --name 'CRED'						Selects the credential being set
		-O, --no-overwrite						Credential is not modified if stored value already exists


		Generate parameters by [Type]
		-l, --length [4-200]					[Value] Length of generated value (Default: 20)
		-U, --exclude-upper 		            [Value] Exclude upper alpha characters from generated value
		-L, --exclude-lower 		            [Value] Exclude lower alpha characters from generated value
		-N, --exclude-number 		            [Value] Exclude numbers from generated value
		-S, --exclude-special 	  	            [Value] Exclude special characters from generated value
		--ca 'CA NAME'					    	[Certificate] Name of CA used to sign the generated certificate (Default: 'default')
		-d, --duration [1-3650]					[Certificate] Valid duration (in days) of the generated certificate (Default: 365)
		-k, --key-length [2048, 3072, 4096]		[Certificate] Bit length of the generated key (Default: 2048)
		-c, --common-name 'COMMON NAME'			[Certificate] Common name of the generated certificate
		-a, --alternative-name 'ALT NAME'		[Certificate] Alternative name(s) of the generated certificate
		-o, --organization 'ORG'				[Certificate] Organization of the generated certificate
		-u, --organization-unit 'ORG UNIT'		[Certificate] Organization unit of the generated certificate
		-i, --locality 'LOCALITY'				[Certificate] Locality/city of the generated certificate
		-s, --state	'ST'						[Certificate] State/province of the generated certificate
		-y, --country 'CC'						[Certificate] Country of the generated certificate

	delete --name <cred name>
		Delete a credential
		-n, --name 'CRED'						Name of credential to delete
		
CERTIFICATE AUTHORITY:

NOTE: CA with name 'default' will be used when generating a certificate credential without a named CA

	ca-get --name <ca name>
		Get the value and attributes of a CA
		-n, --name 'CA'							Name of CA to retrieve

	ca-set --type <ca type> --name <ca name> [set params]
		Set the value and attributes of a CA
		-t, --type ['root']						Sets the type of CA to store (Default: 'root')
		-n, --name 'CA'							Selects the CA being set

		Set parameters by [Type]
		-c, --certificate <FILE>				[Root] Sets the CA certificate from file
		-p, --private <FILE>					[Root] Sets the CA private key from file
		-C, --certificate-string 'CERT'       	[Root] Sets the CA certificate from string input
		-P, --private-string 'PRIVATE'			[Root] Sets the CA private key from string input

	ca-generate --type <ca type> --name <ca name> [generate params]
		Generate and set a credential value based on generation parameters
		-t, --type ['root']	       	     		Sets the type of CA to generate (Default: 'root')
		-n, --name 'CRED'						Selects the CA to generate

		Generate parameters by [Type]
		-d, --duration [1-3650]					[Root] Valid duration (in days) of the generated certificate (Default: 365)
		-k, --key-length [2048, 3072, 4096]		[Root] Bit length of the generated key (Default: 2048)
		-c, --common-name 'COMMON NAME'			[Root] Common name of the generated certificate
		-o, --organization 'ORG'				[Root] Organization of the generated certificate
		-u, --organization-unit 'ORG UNIT'		[Root] Organization unit of the generated certificate
		-i, --locality 'LOCALITY'				[Root] Locality/city of the generated certificate
		-s, --state	'ST'						[Root] State/province of the generated certificate
		-y, --country 'CC'						[Root] Country of the generated certificate
```
