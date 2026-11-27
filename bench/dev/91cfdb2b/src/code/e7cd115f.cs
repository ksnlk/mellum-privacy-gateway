        ERROR_FILE_EXISTS = 0x00000050,

        /// <summary>
        /// The directory or file cannot be created.
        /// </summary>
        ERROR_CANNOT_MAKE = 0x00000052,

        /// <summary>
        /// Fail on INT 24.
        /// </summary>
        ERROR_FAIL_I24 = 0x00000053,

        /// <summary>
        /// Storage to process this request is not available.
        /// </summary>
        ERROR_OUT_OF_STRUCTURES = 0x00000054,

        /// <summary>
        /// The local device name is already in use.
        /// </summary>
        ERROR_ALREADY_ASSIGNED = 0x00000055,

        /// <summary>
        /// The specified network password is not correct.
        /// </summary>
        ERROR_INVALID_PASSWORD = 0x00000056,

        /// <summary>
        /// The parameter is incorrect.
        /// </summary>
        ERROR_INVALID_PARAMETER = 0x00000057,

        /// <summary>
        /// A write fault occurred on the network.
        /// </summary>
        ERROR_NET_WRITE_FAULT = 0x00000058,

        /// <summary>
        /// The system cannot start another process at this time.
        /// </summary>
        ERROR_NO_PROC_SLOTS = 0x00000059,

        /// <summary>
        /// Cannot create another system semaphore.
        /// </summary>
        ERROR_TOO_MANY_SEMAPHORES = 0x00000064,

        /// <summary>
        /// The exclusive semaphore is owned by another process.
        /// </summary>
        ERROR_EXCL_SEM_ALREADY_OWNED = 0x00000065,
        ERROR_INVALID_LOCK_RANGE = 0x00000133,

        /// <summary>
        /// An invalid exception handler routine has been detected.
        /// </summary>
        ERROR_INVALID_EXCEPTION_HANDLER = 0x00000136,

        /// <summary>
        /// Duplicate privileges were specified for the token.
        /// </summary>
        ERROR_DUPLICATE_PRIVILEGES = 0x00000137,

        /// <summary>
        /// No ranges for the specified operation were able to be processed.
        /// </summary>
        ERROR_NO_RANGES_PROCESSED = 0x00000138,

        /// <summary>
        /// The physical resources of this disk have been exhausted.
        /// </summary>
        ERROR_DISK_RESOURCES_EXHAUSTED = 0x0000013A,

        /// <summary>
        /// The token representing the data is invalid.
        /// </summary>
        ERROR_INVALID_TOKEN = 0x0000013B,

        /// <summary>
        /// The system cannot find message text for message number 0x%1 in the message file for %2.
        /// </summary>
        ERROR_MR_MID_NOT_FOUND = 0x0000013D,

        /// <summary>
        /// The scope specified was not found.
        /// </summary>
        ERROR_SCOPE_NOT_FOUND = 0x0000013E,

        /// <summary>
        /// The Central Access Policy specified is not defined on the target machine.
        /// </summary>
        ERROR_UNDEFINED_SCOPE = 0x0000013F,

        /// <summary>
        /// The Central Access Policy obtained from Active Directory is invalid.
        /// </summary>
        ERROR_INVALID_CAP = 0x00000140,

        /// <summary>
        /// The device is unreachable.
        /// </summary>
        ERROR_DEVICE_UNREACHABLE = 0x00000141,
        /// <summary>
        /// There is an IP address conflict with another system on the network
        /// </summary>
        ERROR_IP_ADDRESS_CONFLICT1 = 0x00000263,

        /// <summary>
        /// There is an IP address conflict with another system on the network
        /// </summary>
        ERROR_IP_ADDRESS_CONFLICT2 = 0x00000264,

        /// <summary>
        /// {Low On Registry Space}
        /// The system has reached the maximum size allowed for the system part of the registry. Additional storage requests will be ignored.
        /// </summary>
        ERROR_REGISTRY_QUOTA_LIMIT = 0x00000265,

        /// <summary>
        /// A callback return system service cannot be executed when no callback is active.
        /// </summary>
        ERROR_NO_CALLBACK_ACTIVE = 0x00000266,

        /// <summary>
        /// The password provided is too short to meet the policy of your user account.
        /// Please choose a longer password.
        /// </summary>
        ERROR_PWD_TOO_SHORT = 0x00000267,

        /// <summary>
        /// The policy of your user account does not allow you to change passwords too frequently.
        /// This is done to prevent users from changing back to a familiar, but potentially discovered, password.
        /// If you feel your password has been compromised then please contact your administrator immediately to have a new one assigned.
        /// </summary>
        ERROR_PWD_TOO_RECENT = 0x00000268,

        /// <summary>
        /// You have attempted to change your password to one that you have used in the past.
        /// The policy of your user account does not allow this. Please select a password that you have not previously used.
        /// </summary>
        ERROR_PWD_HISTORY_CONFLICT = 0x00000269,

        /// <summary>
        /// The specified compression format is unsupported.
        /// </summary>
        ERROR_UNSUPPORTED_COMPRESSION = 0x0000026A,

        /// <summary>
        /// The specified hardware profile configuration is invalid.
        /// </summary>
        ERROR_INVALID_HW_PROFILE = 0x0000026B,

        /// <summary>
        /// The specified quota list is internally inconsistent with its descriptor.
        /// </summary>
        ERROR_QUOTA_LIST_INCONSISTENT = 0x0000026D,

        /// <summary>
        /// {Windows Evaluation Notification}
        /// The evaluation period for this installation of Windows has expired. This system will shutdown in 1 hour. To restore access to this installation of Windows, please upgrade this installation using a licensed distribution of this product.
        /// </summary>
        ERROR_EVALUATION_EXPIRATION = 0x0000026E,

        /// <summary>
        /// {Illegal System DLL Relocation}
        /// The system DLL %hs was relocated in memory. The application will not run properly.
        /// <summary>
        /// There was error [%2] processing the driver database.
        /// </summary>
        ERROR_DRIVER_DATABASE_ERROR = 0x0000028C,

        /// <summary>
        /// System hive size has exceeded its limit.
        /// </summary>
        ERROR_SYSTEM_HIVE_TOO_LARGE = 0x0000028D,

        /// <summary>
        /// {Volume Shadow Copy Service}
        /// Please wait while the Volume Shadow Copy Service prepares volume %hs for hibernation.
        /// </summary>
        ERROR_VOLSNAP_PREPARE_HIBERNATE = 0x0000028F,

        /// <summary>
        /// The system has failed to hibernate (The error code is %hs). Hibernation will be disabled until the system is restarted.
        /// </summary>
        ERROR_HIBERNATION_FAILURE = 0x00000290,

        /// <summary>
        /// The password provided is too long to meet the policy of your user account.
        /// Please choose a shorter password.
        /// </summary>
        ERROR_PWD_TOO_LONG = 0x00000291,

        /// <summary>
        /// The requested operation could not be completed due to a file system limitation
        /// </summary>
        ERROR_FILE_SYSTEM_LIMITATION = 0x00000299,

        /// <summary>
        /// An assertion failure has occurred.
        /// </summary>
        ERROR_ASSERTION_FAILURE = 0x0000029C,

        /// <summary>
        /// An error occurred in the ACPI subsystem.
        /// </summary>
        ERROR_ACPI_ERROR = 0x0000029D,

        /// <summary>
        /// WOW Assertion Error.
        /// </summary>
        ERROR_WOW_ASSERTION = 0x0000029E,

        /// <summary>
        /// A device is missing in the system BIOS MPS table. This device will not be used.
        /// Please contact your system vendor for system BIOS update.
        /// </summary>

        /// <summary>
        /// Invalid flags.
        /// </summary>
        ERROR_INVALID_FLAGS = 0x000003EC,

        /// <summary>
        /// The volume does not contain a recognized file system.
        /// Please make sure that all required file system drivers are loaded and that the volume is not corrupted.
        /// </summary>
        ERROR_UNRECOGNIZED_VOLUME = 0x000003ED,

        /// <summary>
        /// The volume for a file has been externally altered so that the opened file is no longer valid.
        /// </summary>
        ERROR_FILE_INVALID = 0x000003EE,

        /// <summary>
        /// The requested operation cannot be performed in full-screen mode.
        /// </summary>
        ERROR_FULLSCREEN_MODE = 0x000003EF,

        /// <summary>
        /// An attempt was made to reference a token that does not exist.
        /// </summary>
        ERROR_NO_TOKEN = 0x000003F0,

        /// <summary>
        /// The configuration registry database is corrupt.
        /// </summary>
        ERROR_BADDB = 0x000003F1,

        /// <summary>
        /// The configuration registry key is invalid.
        /// </summary>
        ERROR_BADKEY = 0x000003F2,

        /// <summary>
        /// The configuration registry key could not be opened.
        /// </summary>
        ERROR_CANTOPEN = 0x000003F3,

        /// <summary>
        /// The configuration registry key could not be read.
        /// </summary>
        ERROR_CANTREAD = 0x000003F4,

        /// <summary>
        /// The configuration registry key could not be written.
        /// </summary>
        ERROR_CANTWRITE = 0x000003F5,
        ERROR_INVALID_EVENTNAME = 0x000004BB,

        /// <summary>
        /// The format of the specified domain name is invalid.
        /// </summary>
        ERROR_INVALID_DOMAINNAME = 0x000004BC,

        /// <summary>
        /// The format of the specified service name is invalid.
        /// </summary>
        ERROR_INVALID_SERVICENAME = 0x000004BD,

        /// <summary>
        /// The format of the specified network name is invalid.
        /// </summary>
        ERROR_INVALID_NETNAME = 0x000004BE,

        /// <summary>
        /// The format of the specified share name is invalid.
        /// </summary>
        ERROR_INVALID_SHARENAME = 0x000004BF,

        /// <summary>
        /// The format of the specified password is invalid.
        /// </summary>
        ERROR_INVALID_PASSWORDNAME = 0x000004C0,

        /// <summary>
        /// The format of the specified message name is invalid.
        /// </summary>
        ERROR_INVALID_MESSAGENAME = 0x000004C1,

        /// <summary>
        /// The format of the specified message destination is invalid.
        /// </summary>
        ERROR_INVALID_MESSAGEDEST = 0x000004C2,

        /// <summary>
        /// The workgroup or domain name is already in use by another computer on the network.
        /// </summary>
        ERROR_DUP_DOMAINNAME = 0x000004C5,

        /// <summary>
        /// The network is not present or not started.
        /// </summary>
        ERROR_NO_NETWORK = 0x000004C6,

        /// <summary>
        /// The operation was canceled by the user.
        /// </summary>
        ERROR_CANCELLED = 0x000004C7,
        ERROR_APP_HANG = 0x00000512,

        /// <summary>
        /// Indicates a particular Security ID may not be assigned as the label of an object.
        /// </summary>
        ERROR_INVALID_LABEL = 0x00000513,

        /// <summary>
        /// Not all privileges or groups referenced are assigned to the caller.
        /// </summary>
        ERROR_NOT_ALL_ASSIGNED = 0x00000514,

        /// <summary>
        /// Some mapping between account names and security IDs was not done.
        /// </summary>
        ERROR_SOME_NOT_MAPPED = 0x00000515,

        /// <summary>
        /// No system quota limits are specifically set for this account.
        /// </summary>
        ERROR_NO_QUOTAS_FOR_ACCOUNT = 0x00000516,

        /// <summary>
        /// No encryption key is available. A well-known encryption key was returned.
        /// </summary>
        ERROR_LOCAL_USER_SESSION_KEY = 0x00000517,

        /// <summary>
        /// The password is too complex to be converted to a LAN Manager password. The LAN Manager password returned is a NULL string.
        /// </summary>
        ERROR_NULL_LM_PASSWORD = 0x00000518,

        /// <summary>
        /// The revision level is unknown.
        /// </summary>
        ERROR_UNKNOWN_REVISION = 0x00000519,

        /// <summary>
        /// Indicates two revision levels are incompatible.
        /// </summary>
        ERROR_REVISION_MISMATCH = 0x0000051A,

        /// <summary>
        /// This security ID may not be assigned as the owner of this object.
        /// </summary>
        ERROR_INVALID_OWNER = 0x0000051B,

        /// <summary>
        /// This security ID may not be assigned as the primary group of an object.
        /// </summary>
        ERROR_INVALID_PRIMARY_GROUP = 0x0000051C,

        /// <summary>
        /// An attempt has been made to operate on an impersonation token by a thread that is not currently impersonating a client.
        /// </summary>
        ERROR_NO_IMPERSONATION_TOKEN = 0x0000051D,

        /// <summary>
        /// The group may not be disabled.
        /// </summary>
        ERROR_CANT_DISABLE_MANDATORY = 0x0000051E,

        /// <summary>
        /// There are currently no logon servers available to service the logon request.
        /// </summary>
        ERROR_NO_LOGON_SERVERS = 0x0000051F,

        /// <summary>
        /// A specified logon session does not exist. It may already have been terminated.
        /// </summary>
        ERROR_NO_SUCH_LOGON_SESSION = 0x00000520,

        /// <summary>
        /// A specified privilege does not exist.
        /// </summary>
        ERROR_NO_SUCH_PRIVILEGE = 0x00000521,

        /// <summary>
        /// A required privilege is not held by the client.
        /// </summary>
        ERROR_PRIVILEGE_NOT_HELD = 0x00000522,
        ERROR_GROUP_EXISTS = 0x00000526,

        /// <summary>
        /// The specified group does not exist.
        /// </summary>
        ERROR_NO_SUCH_GROUP = 0x00000527,

        /// <summary>
        /// Either the specified user account is already a member of the specified group, or the specified group cannot be deleted because it contains a member.
        /// </summary>
        ERROR_MEMBER_IN_GROUP = 0x00000528,

        /// <summary>
        /// The specified user account is not a member of the specified group account.
        /// </summary>
        ERROR_MEMBER_NOT_IN_GROUP = 0x00000529,

        /// <summary>
        /// This operation is disallowed as it could result in an administration account being disabled, deleted or unable to logon.
        /// </summary>
        ERROR_LAST_ADMIN = 0x0000052A,

        /// <summary>
        /// Unable to update the password. The value provided as the current password is incorrect.
        /// </summary>
        ERROR_WRONG_PASSWORD = 0x0000052B,

        /// <summary>
        /// Unable to update the password. The value provided for the new password contains values that are not allowed in passwords.
        /// </summary>
        ERROR_ILL_FORMED_PASSWORD = 0x0000052C,

        /// <summary>
        /// Unable to update the password. The value provided for the new password does not meet the length, complexity, or history requirements of the domain.
        /// </summary>
        ERROR_PASSWORD_RESTRICTION = 0x0000052D,

        /// <summary>
        /// The user name or password is incorrect.
        /// </summary>
        ERROR_LOGON_FAILURE = 0x0000052E,

        /// <summary>
        /// Account restrictions are preventing this user from signing in. For example: blank passwords aren't allowed, sign-in times are limited, or a policy restriction has been enforced.
        /// </summary>
        ERROR_ACCOUNT_RESTRICTION = 0x0000052F,

        /// <summary>
        /// Your account has time restrictions that keep you from signing in right now.
        /// </summary>
        ERROR_INVALID_LOGON_HOURS = 0x00000530,

        /// <summary>
        /// This user isn't allowed to sign in to this computer.
        /// </summary>
        ERROR_INVALID_WORKSTATION = 0x00000531,

        /// <summary>
        /// The password for this account has expired.
        /// </summary>
        ERROR_PASSWORD_EXPIRED = 0x00000532,

        /// <summary>
        /// This user can't sign in because this account is currently disabled.
        /// </summary>
        ERROR_ACCOUNT_DISABLED = 0x00000533,

        /// <summary>
        /// No mapping between account names and security IDs was done.
        /// </summary>
        ERROR_NONE_MAPPED = 0x00000534,

        /// <summary>
        /// Too many local user identifiers (LUIDs) were requested at one time.
        /// </summary>
        ERROR_TOO_MANY_LUIDS_REQUESTED = 0x00000535,

        /// <summary>
        /// No more local user identifiers (LUIDs) are available.
        /// </summary>
        ERROR_LUIDS_EXHAUSTED = 0x00000536,

        /// <summary>
        /// The subauthority part of a security ID is invalid for this particular use.
        /// </summary>
        ERROR_INVALID_SUB_AUTHORITY = 0x00000537,
        ERROR_INVALID_MEMBER = 0x0000056C,

        /// <summary>
        /// Too many security IDs have been specified.
        /// </summary>
        ERROR_TOO_MANY_SIDS = 0x0000056D,

        /// <summary>
        /// Indicates an ACL contains no inheritable components.
        /// </summary>
        ERROR_NO_INHERITANCE = 0x0000056F,

        /// <summary>
        /// The file or directory is corrupted and unreadable.
        /// </summary>
        ERROR_FILE_CORRUPT = 0x00000570,

        /// <summary>
        /// The disk structure is corrupted and unreadable.
        /// </summary>
        ERROR_DISK_CORRUPT = 0x00000571,

        /// <summary>
        /// There is no user session key for the specified logon session.
        /// </summary>
        ERROR_NO_USER_SESSION_KEY = 0x00000572,

        /// <summary>
        /// The service being accessed is licensed for a particular number of connections. No more connections can be made to the service at this time because there are already as many connections as the service can accept.
        /// </summary>
        ERROR_LICENSE_QUOTA_EXCEEDED = 0x00000573,

        /// <summary>
        /// The target account name is incorrect.
        /// </summary>
        ERROR_WRONG_TARGET_NAME = 0x00000574,

        /// <summary>
        /// Mutual Authentication failed. The server's password is out of date at the domain controller.
        /// </summary>
        ERROR_MUTUAL_AUTH_FAILED = 0x00000575,

        /// <summary>
        /// There is a time and/or date difference between the client and server.
        /// </summary>
        ERROR_TIME_SKEW = 0x00000576,

        /// <summary>
        /// Invalid window handle.
        /// </summary>
        ERROR_INVALID_WINDOW_HANDLE = 0x00000578,
        ERROR_INVALID_FORM_NAME = 0x0000076E,

        /// <summary>
        /// The specified form size is invalid.
        /// </summary>
        ERROR_INVALID_FORM_SIZE = 0x0000076F,

        /// <summary>
        /// The specified printer handle is already being waited on
        /// </summary>
        ERROR_ALREADY_WAITING = 0x00000770,

        /// <summary>
        /// The specified printer has been deleted.
        /// </summary>
        ERROR_PRINTER_DELETED = 0x00000771,

        /// <summary>
        /// The state of the printer is invalid.
        /// </summary>
        ERROR_INVALID_PRINTER_STATE = 0x00000772,

        /// <summary>
        /// The user's password must be changed before signing in.
        /// </summary>
        ERROR_PASSWORD_MUST_CHANGE = 0x00000773,

        /// <summary>
        /// The referenced account is currently locked out and may not be logged on to.
        /// </summary>
        ERROR_ACCOUNT_LOCKED_OUT = 0x00000775,

        /// <summary>
        /// The object exporter specified was not found.
        /// </summary>
        OR_INVALID_OXID = 0x00000776,

        /// <summary>
        /// The object specified was not found.
        /// </summary>
        OR_INVALID_OID = 0x00000777,

        /// <summary>
        /// The object resolver set specified was not found.
        /// </summary>
        OR_INVALID_SET = 0x00000778,

        /// <summary>
        /// Some data remains to be sent in the request buffer.
        /// </summary>
        RPC_S_SEND_INCOMPLETE = 0x00000779,
        RPC_S_GRP_ELT_NOT_ADDED = 0x00000788,

        /// <summary>
        /// The group element could not be removed.
        /// </summary>
        RPC_S_GRP_ELT_NOT_REMOVED = 0x00000789,

        /// <summary>
        /// The printer driver is not compatible with a policy enabled on your computer that blocks NT 4.0 drivers.
        /// </summary>
        ERROR_KM_DRIVER_BLOCKED = 0x0000078A,

        /// <summary>
        /// The context has expired and can no longer be used.
        /// </summary>
        ERROR_CONTEXT_EXPIRED = 0x0000078B,

        /// <summary>
        /// Authentication failed because NTLM authentication has been disabled.
        /// </summary>
        ERROR_NTLM_BLOCKED = 0x00000791,

        /// <summary>
        /// Logon Failure: EAS policy requires that the user change their password before this operation can be performed.
        /// </summary>
        ERROR_PASSWORD_CHANGE_REQUIRED = 0x00000792,

        /// <summary>
        /// The pixel format is invalid.
        /// </summary>
        ERROR_INVALID_PIXEL_FORMAT = 0x000007D0,

        /// <summary>
        /// The specified driver is invalid.
        /// </summary>
        ERROR_BAD_DRIVER = 0x000007D1,

        /// <summary>
        /// The window style or class attribute is invalid for this operation.
        /// </summary>
        ERROR_INVALID_WINDOW_STYLE = 0x000007D2,

        /// <summary>
        /// The requested metafile operation is not supported.
        /// </summary>
        ERROR_METAFILE_NOT_SUPPORTED = 0x000007D3,

        /// <summary>
        /// The requested transformation operation is not supported.
        /// </summary>
        ERROR_TRANSFORM_NOT_SUPPORTED = 0x000007D4,
        ERROR_ICM_NOT_ENABLED = 0x000007E2,

        /// <summary>
        /// There was an error while deleting the color transform.
        /// </summary>
        ERROR_DELETING_ICM_XFORM = 0x000007E3,

        /// <summary>
        /// The specified color transform is invalid.
        /// </summary>
        ERROR_INVALID_TRANSFORM = 0x000007E4,

        /// <summary>
        /// The specified transform does not match the bitmap's color space.
        /// </summary>
        ERROR_COLORSPACE_MISMATCH = 0x000007E5,

        /// <summary>
        /// The specified named color index is not present in the profile.
        /// </summary>
        ERROR_INVALID_COLORINDEX = 0x000007E6,

        /// <summary>
        /// The network connection was made successfully, but the user had to be prompted for a password other than the one originally specified.
        /// </summary>
        ERROR_CONNECTED_OTHER_PASSWORD = 0x0000083C,

        /// <summary>
        /// The specified username is invalid.
        /// </summary>
        ERROR_BAD_USERNAME = 0x0000089A,

        /// <summary>
        /// This network connection does not exist.
        /// </summary>
        ERROR_NOT_CONNECTED = 0x000008CA,

        /// <summary>
        /// This network connection has files open or requests pending.
        /// </summary>
        ERROR_OPEN_FILES = 0x00000961,

        /// <summary>
        /// Active connections still exist.
        /// </summary>
        ERROR_ACTIVE_CONNECTIONS = 0x00000962,

        /// <summary>
        /// The device is in use by an active process and cannot be disconnected.
        /// </summary>
        ERROR_DEVICE_IN_USE = 0x00000964,
        ERROR_IPSEC_IKE_QUEUE_DROP_MM = 0x000035F3,

        /// <summary>
        /// Negotiation took too long
        /// </summary>
        ERROR_IPSEC_IKE_MM_DELAY_DROP = 0x000035F6,

        /// <summary>
        /// Negotiation took too long
        /// </summary>
        ERROR_IPSEC_IKE_QM_DELAY_DROP = 0x000035F7,

        /// <summary>
        /// Unknown error occurred
        /// </summary>
        ERROR_IPSEC_IKE_ERROR = 0x000035F8,

        /// <summary>
        /// Certificate Revocation Check failed
        /// </summary>
        ERROR_IPSEC_IKE_CRL_FAILED = 0x000035F9,

        /// <summary>
        /// IKE negotiation failed because the machine certificate used does not have a private key. IPsec certificates require a private key. Contact your Network Security administrator about replacing with a certificate that has a private key.
        /// </summary>
        ERROR_IPSEC_IKE_NO_PRIVATE_KEY = 0x000035FC,

        /// <summary>
        /// Failure in Diffie-Hellman computation
        /// </summary>
        ERROR_IPSEC_IKE_DH_FAIL = 0x000035FE,

        /// <summary>
        /// Invalid header
        /// </summary>
        ERROR_IPSEC_IKE_INVALID_HEADER = 0x00003600,

        /// <summary>
        /// No policy configured
        /// </summary>
        ERROR_IPSEC_IKE_NO_POLICY = 0x00003601,

        /// <summary>
        /// Failed to authenticate using Kerberos
        /// </summary>
        ERROR_IPSEC_IKE_KERBEROS_ERROR = 0x00003603,

        /// <summary>
        /// Peer's certificate did not have a public key
        /// </summary>
        ERROR_IPSEC_IKE_NO_PUBLIC_KEY = 0x00003604,

        /// <summary>
        /// Error processing error payload
        /// </summary>
        ERROR_IPSEC_IKE_PROCESS_ERR = 0x00003605,

        /// <summary>
        /// Error processing SA payload
        /// </summary>
        ERROR_IPSEC_IKE_PROCESS_ERR_SA = 0x00003606,

        /// <summary>
        /// Error processing KE payload
        /// </summary>
        ERROR_IPSEC_IKE_PROCESS_ERR_KE = 0x00003609,

        /// <summary>
        /// Error processing ID payload
        /// </summary>
        ERROR_IPSEC_IKE_PROCESS_ERR_ID = 0x0000360A,

        /// <summary>
        /// Error processing Signature payload
        /// </summary>
        ERROR_IPSEC_IKE_PROCESS_ERR_SIG = 0x0000360E,
        /// <summary>A remote API error occurred.</summary>
        NERR_RemoteErr = 2100 + 27,

        /// <summary>An error occurred when opening or reading the configuration file.</summary>
        NERR_LanmanIniError = 2100 + 31,

        /// <summary>A general network error occurred.</summary>
        NERR_NetworkError = 2100 + 36,

        /// <summary>The Workstation service is in an inconsistent state. Restart the computer before restarting the Workstation service.</summary>
        NERR_WkstaInconsistentState = 2100 + 37,

        /// <summary>The Workstation service has not been started.</summary>
        NERR_WkstaNotStarted = 2100 + 38,

        /// <summary>The requested information is not available.</summary>
        NERR_BrowserNotStarted = 2100 + 39,

        /// <summary>An internal Windows error occurred.</summary>
        NERR_InternalError = 2100 + 40,

        /// <summary>The server is not configured for transactions.</summary>
        NERR_BadTransactConfig = 2100 + 41,

        /// <summary>The requested API is not supported on the remote server.</summary>
        NERR_InvalidAPI = 2100 + 42,

        /// <summary>The event name is invalid.</summary>
        NERR_BadEventName = 2100 + 43,

        /// <summary>The computer name already exists on the network. Change it and restart the computer.</summary>
        NERR_DupNameReboot = 2100 + 44,

        /// <summary>The specified component could not be found in the configuration information.</summary>
        NERR_CfgCompNotFound = 2100 + 46,

        /// <summary>The specified parameter could not be found in the configuration information.</summary>
        NERR_CfgParamNotFound = 2100 + 47,

        /// <summary>A line in the configuration file is too long.</summary>
        NERR_LineTooLong = 2100 + 49,

        /// <summary>The printer does not exist.</summary>
        NERR_QNotFound = 2100 + 50,

        /// <summary>The print job does not exist.</summary>
        NERR_JobNotFound = 2100 + 51,

        /// <summary>The printer destination cannot be found.</summary>
        NERR_DestNotFound = 2100 + 52,

        /// <summary>The service ended abnormally.</summary>
        NERR_ServiceKillProc = 2100 + 90,

        /// <summary>The requested pause, continue, or stop is not valid for this service.</summary>
        NERR_ServiceCtlNotValid = 2100 + 91,

        /// <summary>The service control dispatcher could not find the service name in the dispatch table.</summary>
        NERR_NotInDispatchTbl = 2100 + 92,

        /// <summary>The service control dispatcher pipe read failed.</summary>
        NERR_BadControlRecv = 2100 + 93,

        /// <summary>A thread for the new service could not be created.</summary>
        NERR_ServiceNotStarting = 2100 + 94,

        /// <summary>This workstation is already logged on to the local-area network.</summary>
        NERR_AlreadyLoggedOn = 2100 + 100,

        /// <summary>The workstation is not logged on to the local-area network.</summary>
        NERR_NotLoggedOn = 2100 + 101,

        /// <summary>The user name or group name parameter is invalid.</summary>
        NERR_BadUsername = 2100 + 102,

        /// <summary>The password parameter is invalid.</summary>
        NERR_BadPassword = 2100 + 103,

        /// <summary>@W The logon processor did not add the message alias.</summary>
        NERR_UnableToAddName_W = 2100 + 104,

        /// <summary>The logon processor did not add the message alias.</summary>
        NERR_UnableToAddName_F = 2100 + 105,

        /// <summary>@W The logoff processor did not delete the message alias.</summary>
        NERR_UnableToDelName_W = 2100 + 106,

        /// <summary>The logoff processor did not delete the message alias.</summary>
        NERR_UnableToDelName_F = 2100 + 107,

        /// <summary>Network logons are paused.</summary>
        NERR_LogonsPaused = 2100 + 109,

        /// <summary>A centralized logon-server conflict occurred.</summary>
        NERR_LogonServerConflict = 2100 + 110,

        /// <summary>The server is configured without a valid user path.</summary>
        NERR_LogonNoUserPath = 2100 + 111,

        /// <summary>An error occurred while loading or running the logon script.</summary>
        NERR_LogonScriptError = 2100 + 112,

        /// <summary>This operation is not allowed on this special group.</summary>
        NERR_SpeGroupOp = 2100 + 134,

        /// <summary>This user is not cached in user accounts database session cache.</summary>
        NERR_NotInCache = 2100 + 135,

        /// <summary>The user already belongs to this group.</summary>
        NERR_UserInGroup = 2100 + 136,

        /// <summary>The user does not belong to this group.</summary>
        NERR_UserNotInGroup = 2100 + 137,

        /// <summary>This user account is undefined.</summary>
        NERR_AccountUndefined = 2100 + 138,

        /// <summary>This user account has expired.</summary>
        NERR_AccountExpired = 2100 + 139,

        /// <summary>The user is not allowed to log on from this workstation.</summary>
        NERR_InvalidWorkstation = 2100 + 140,

        /// <summary>The user is not allowed to log on at this time.</summary>
        NERR_InvalidLogonHours = 2100 + 141,

        /// <summary>The password of this user has expired.</summary>
        NERR_PasswordExpired = 2100 + 142,

        /// <summary>The password of this user cannot change.</summary>
        NERR_PasswordCantChange = 2100 + 143,

        /// <summary>This password cannot be used now.</summary>
        NERR_PasswordHistConflict = 2100 + 144,

        /// <summary>The password does not meet the password policy requirements. Check the minimum password length, password complexity and password history requirements.</summary>
        NERR_PasswordTooShort = 2100 + 145,

        /// <summary>The password of this user is too recent to change.</summary>
        NERR_PasswordTooRecent = 2100 + 146,

        /// <summary>The security database is corrupted.</summary>
        NERR_InvalidDatabase = 2100 + 147,

        /// <summary>No updates are necessary to this replicant network/local security database.</summary>
        NERR_DatabaseUpToDate = 2100 + 148,

        /// <summary>This replicant database is outdated; synchronization is required.</summary>
        NERR_SyncRequired = 2100 + 149,

        /// <summary>The network connection could not be found.</summary>
        NERR_UseNotFound = 2100 + 150,

        /// <summary>This asg_type is invalid.</summary>
        NERR_BadAsgType = 2100 + 151,

        /// <summary>This device is currently being shared.</summary>
        NERR_DeviceIsShared = 2100 + 152,

        /// <summary>The user name may not be same as computer name.</summary>
        NERR_SameAsComputerName = 2100 + 153,

        /// <summary>The computer name could not be added as a message alias.  The name may already exist on the network.</summary>
        NERR_NoComputerName = 2100 + 170,

        /// <summary>The source and destination paths are on different servers.</summary>
        NERR_DifferentServers = 2100 + 283,

        /// <summary>The Run server you requested is paused.</summary>
        NERR_RunSrvPaused = 2100 + 285,

        /// <summary>An error occurred when communicating with a Run server.</summary>
        NERR_ErrCommRunSrv = 2100 + 289,

        /// <summary>An error occurred when starting a background process.</summary>
        NERR_ErrorExecingGhost = 2100 + 291,

        /// <summary>The shared resource you are connected to could not be found.</summary>
        NERR_ShareNotFound = 2100 + 292,

        /// <summary>The LAN adapter number is invalid.</summary>
        NERR_InvalidLana = 2100 + 300,

        /// <summary>There are open files on the connection.   </summary>
        NERR_OpenFiles = 2100 + 301,

        /// <summary>Active connections still exist.</summary>
        NERR_ActiveConns = 2100 + 302,

        /// <summary>This share name or password is invalid.</summary>
        NERR_BadPasswordCore = 2100 + 303,

        /// <summary>The device is being accessed by an active process.</summary>
        NERR_DevInUse = 2100 + 304,

        /// <summary>The drive letter is in use locally.</summary>
        NERR_LocalDrive = 2100 + 305,

        /// <summary>The specified client is already registered for the specified event.</summary>
        NERR_AlertExists = 2100 + 330,

        /// <summary>The alert table is full.</summary>
        NERR_TooManyAlerts = 2100 + 331,

        /// <summary>An invalid or nonexistent alert name was raised.</summary>
        NERR_NoSuchAlert = 2100 + 332,

        /// <summary>The alert recipient is invalid.</summary>
        NERR_BadRecipient = 2100 + 333,

        /// <summary>A user's session with this server has been deleted because the user's logon hours are no longer valid.</summary>
        NERR_AcctLimitExceeded = 2100 + 334,

        /// <summary>The log file does not contain the requested record number.</summary>
        NERR_InvalidLogSeek = 2100 + 340,

        /// <summary>The user accounts database is not configured correctly.</summary>
        NERR_BadUasConfig = 2100 + 350,

        /// <summary>This operation is not permitted when the Netlogon service is running.</summary>
        NERR_InvalidUASOp = 2100 + 351,

        /// <summary>This operation is not allowed on the last administrative account.</summary>
        NERR_LastAdmin = 2100 + 352,

        /// <summary>Could not find domain controller for this domain.</summary>
        NERR_DCNotFound = 2100 + 353,

        /// <summary>Could not set logon information for this user.</summary>
        NERR_LogonTrackingError = 2100 + 354,

        /// <summary>The Netlogon service has not been started.</summary>
        NERR_NetlogonNotStarted = 2100 + 355,

        /// <summary>Unable to add to the user accounts database.</summary>
        NERR_CanNotGrowUASFile = 2100 + 356,

        /// <summary>This server's clock is not synchronized with the primary domain controller's clock.</summary>
        NERR_TimeDiffAtDC = 2100 + 357,

        /// <summary>A password mismatch has been detected.</summary>
        NERR_PasswordMismatch = 2100 + 358,

        /// <summary>The server identification does not specify a valid server.</summary>
        NERR_NoSuchServer = 2100 + 360,

        /// <summary>The session identification does not specify a valid session.</summary>
        NERR_NoSuchSession = 2100 + 361,

        /// <summary>The connection identification does not specify a valid connection.</summary>
        NERR_NoSuchConnection = 2100 + 362,

        /// <summary>There is no space for another entry in the table of available servers.</summary>
        NERR_TooManyServers = 2100 + 363,

        /// <summary>The server has reached the maximum number of sessions it supports.</summary>
        NERR_TooManySessions = 2100 + 364,

        /// <summary>The server has reached the maximum number of connections it supports.</summary>
        NERR_TooManyConnections = 2100 + 365,

        /// <summary>The server cannot open more files because it has reached its maximum number.</summary>
        NERR_TooManyFiles = 2100 + 366,

        /// <summary>There are no alternate servers registered on this server.</summary>
        NERR_NoAlternateServers = 2100 + 367,

        /// <summary>This machine is not currently joined to a domain.</summary>
        NERR_SetupNotJoined = 2100 + 592,

        /// <summary>This machine is a domain controller and cannot be unjoined from a domain.</summary>
        NERR_SetupDomainController = 2100 + 593,

        /// <summary>The destination domain controller does not support creating machine accounts in OUs.</summary>
        NERR_DefaultJoinRequired = 2100 + 594,

        /// <summary>The specified workgroup name is invalid.</summary>
        NERR_InvalidWorkgroupName = 2100 + 595,

        /// <summary>The specified computer name is incompatible with the default language used on the domain controller.</summary>
        NERR_NameUsesIncompatibleCodePage = 2100 + 596,

        /// <summary>The specified computer account could not be found. Contact an administrator to verify the account is in the domain. If the account has been deleted unjoin, reboot, and rejoin the domain.</summary>
        NERR_ComputerAccountNotFound = 2100 + 597,

        /// <summary>This version of Windows cannot be joined to a domain.</summary>
        NERR_PersonalSku = 2100 + 598,

        /// <summary>An attempt to resolve the DNS name of a domain controller in the domain being joined has failed.  Please verify this client is configured to reach a DNS server that can resolve DNS names in the target domain. For information about network troubleshooting, see Windows Help.</summary>
        NERR_SetupCheckDNSConfig = 2100 + 599,

        /// <summary>Password must change at next logon</summary>
        NERR_PasswordMustChange = 2100 + 601,

        /// <summary>Account is locked out</summary>
        NERR_AccountLockedOut = 2100 + 602,

        /// <summary>Password is too long</summary>
        NERR_PasswordTooLong = 2100 + 603,

        /// <summary>Password doesn't meet the complexity policy</summary>
        NERR_PasswordNotComplexEnough = 2100 + 604,

        /// <summary>Password doesn't meet the requirements of the filter dll's</summary>
        NERR_PasswordFilterError = 2100 + 605,

        /// <summary>Offline join completion information was not found.</summary>
        NERR_NoOfflineJoinInfo = 2100 + 609,

        /// <summary>The offline join completion information was bad.</summary>
        NERR_BadOfflineJoinInfo = 2100 + 610,

        /// <summary>Unable to create offline join information. Please ensure you have access to the specified path location and permissions to modify its contents. Running as an elevated administrator may be required.</summary>
        NERR_CantCreateJoinInfo = 2100 + 611,

        /// <summary>The domain join info being saved was incomplete or bad.</summary>
        NERR_BadDomainJoinInfo = 2100 + 612,

        /// <summary>Offline join operation successfully completed but a restart is needed.</summary>
        NERR_JoinPerformedMustRestart = 2100 + 613,

        /// <summary>There was no offline join operation pending.</summary>
        NERR_NoJoinPending = 2100 + 614,

        /// <summary>Unable to set one or more requested machine or domain name values on the local computer.</summary>
        NERR_ValuesNotSet = 2100 + 615,

        /// <summary>Could not verify the current machine's hostname against the saved value in the join completion information.</summary>
        NERR_CantVerifyHostname = 2100 + 616,
