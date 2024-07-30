package certstore

const (
	// NCRYPT Object Property Names
	NCRYPT_ALGORITHM_GROUP_PROPERTY        = "Algorithm Group"
	NCRYPT_ALGORITHM_PROPERTY              = "Algorithm Name"
	NCRYPT_BLOCK_LENGTH_PROPERTY           = "Block Length"
	NCRYPT_CERTIFICATE_PROPERTY            = "SmartCardKeyCertificate"
	NCRYPT_DH_PARAMETERS_PROPERTY          = BCRYPT_DH_PARAMETERS
	BCRYPT_DH_PARAMETERS                   = "DHParameters"
	NCRYPT_EXPORT_POLICY_PROPERTY          = "Export Policy"
	NCRYPT_IMPL_TYPE_PROPERTY              = "Impl Type"
	NCRYPT_KEY_TYPE_PROPERTY               = "Key Type"
	NCRYPT_KEY_USAGE_PROPERTY              = "Key Usage"
	NCRYPT_LAST_MODIFIED_PROPERTY          = "Modified"
	NCRYPT_LENGTH_PROPERTY                 = "Length"
	NCRYPT_LENGTHS_PROPERTY                = "Lengths"
	NCRYPT_MAX_NAME_LENGTH_PROPERTY        = "Max Name Length"
	NCRYPT_NAME_PROPERTY                   = "Name"
	NCRYPT_PIN_PROMPT_PROPERTY             = "SmartCardPinPrompt"
	NCRYPT_PIN_PROPERTY                    = "SmartCardPin"
	NCRYPT_PROVIDER_HANDLE_PROPERTY        = "Provider Handle"
	NCRYPT_READER_PROPERTY                 = "SmartCardReader"
	NCRYPT_ROOT_CERTSTORE_PROPERTY         = "SmartcardRootCertStore"
	NCRYPT_SECURE_PIN_PROPERTY             = "SmartCardSecurePin"
	NCRYPT_SECURITY_DESCR_PROPERTY         = "Security Descr"
	NCRYPT_SECURITY_DESCR_SUPPORT_PROPERTY = "Security Descr Support"
	NCRYPT_SMARTCARD_GUID_PROPERTY         = "SmartCardGuid"
	NCRYPT_UI_POLICY_PROPERTY              = "UI Policy"
	NCRYPT_UNIQUE_NAME_PROPERTY            = "Unique Name"
	NCRYPT_USE_CONTEXT_PROPERTY            = "Use Context"
	NCRYPT_USE_COUNT_ENABLED_PROPERTY      = "Enabled Use Count"
	NCRYPT_USE_COUNT_PROPERTY              = "Use Count"
	NCRYPT_USER_CERTSTORE_PROPERTY         = "SmartCardUserCertStore"
	NCRYPT_VERSION_PROPERTY                = "Version"
	NCRYPT_WINDOW_HANDLE_PROPERTY          = "HWND Handle"

	// BCRYPT BLOB Types
	BCRYPT_DH_PRIVATE_BLOB     = "DHPRIVATEBLOB"
	BCRYPT_DH_PUBLIC_BLOB      = "DHPUBLICBLOB"
	BCRYPT_DSA_PRIVATE_BLOB    = "DSAPRIVATEBLOB"
	BCRYPT_DSA_PUBLIC_BLOB     = "DSAPUBLICBLOB"
	BCRYPT_ECCPRIVATE_BLOB     = "ECCPRIVATEBLOB"
	BCRYPT_ECCPUBLIC_BLOB      = "ECCPUBLICBLOB"
	BCRYPT_PRIVATE_KEY_BLOB    = "PRIVATEBLOB"
	BCRYPT_PUBLIC_KEY_BLOB     = "PUBLICBLOB"
	BCRYPT_RSAFULLPRIVATE_BLOB = "RSAFULLPRIVATEBLOB"
	BCRYPT_RSAPRIVATE_BLOB     = "RSAPRIVATEBLOB"
	BCRYPT_RSAPUBLIC_BLOB      = "RSAPUBLICBLOB"

	// BCRYPT Algorithm Names
	BCRYPT_3DES_ALGORITHM              = "3DES"
	BCRYPT_AES_ALGORITHM               = "AES"
	BCRYPT_DES_ALGORITHM               = "DES"
	BCRYPT_DSA_ALGORITHM               = "DSA"
	BCRYPT_ECDH_P256_ALGORITHM         = "ECDH_P256"
	BCRYPT_ECDH_P384_ALGORITHM         = "ECDH_P384"
	BCRYPT_ECDSA_P256_ALGORITHM        = "ECDSA_P256"
	BCRYPT_ECDSA_P384_ALGORITHM        = "ECDSA_P384"
	BCRYPT_ECDSA_P521_ALGORITHM        = "ECDSA_P521"
	BCRYPT_MD2_ALGORITHM               = "MD2"
	BCRYPT_MD4_ALGORITHM               = "MD4"
	BCRYPT_MD5_ALGORITHM               = "MD5"
	BCRYPT_RC2_ALGORITHM               = "RC2"
	BCRYPT_RC4_ALGORITHM               = "RC4"
	BCRYPT_RNG_ALGORITHM               = "RNG"
	BCRYPT_RSA_ALGORITHM               = "RSA"
	BCRYPT_RSA_SIGN_ALGORITHM          = "RSA_SIGN"
	BCRYPT_SHA1_ALGORITHM              = "SHA1"
	BCRYPT_SHA256_ALGORITHM            = "SHA256"
	BCRYPT_SHA384_ALGORITHM            = "SHA384"
	BCRYPT_SHA512_ALGORITHM            = "SHA512"
	BCRYPT_SP800108_CTR_HMAC_ALGORITHM = "SP800_108_CTR_HMAC"
	BCRYPT_SP80056A_CONCAT_ALGORITHM   = "SP800_56A_CONCAT"
	BCRYPT_PBKDF2_ALGORITHM            = "PBKDF2"
	BCRYPT_ECDSA_ALGORITHM             = "ECDSA"
	BCRYPT_ECDH_ALGORITHM              = "ECDH"
	BCRYPT_XTS_AES_ALGORITHM           = "XTS-AES"
)