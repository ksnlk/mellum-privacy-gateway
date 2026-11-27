            STATUS_FLT_IO_COMPLETE = 0x001C0001,
            /////////////////////////////////////////////////////////////////////////
            // Standard Information values
            /////////////////////////////////////////////////////////////////////////

            /// <summary>{Object Exists} An attempt was made to create an object and the object name already existed.</summary>
            STATUS_OBJECT_NAME_EXISTS = 0x40000000,

            /// <summary>{Thread Suspended} A thread termination occurred while the thread was suspended. The thread was resumed, and termination proceeded.</summary>
            STATUS_THREAD_WAS_SUSPENDED = 0x40000001,

            /// <summary>{Working Set Range Error} An attempt was made to set the working set minimum or maximum to values which are outside of the allowable range.</summary>
            STATUS_WORKING_SET_LIMIT_RANGE = 0x40000002,

            /// <summary>{Image Relocated} An image file could not be mapped at the address specified in the image file. Local fixups must be performed on this image.</summary>
            STATUS_IMAGE_NOT_AT_BASE = 0x40000003,

            /// <summary>This informational level status indicates that a specified registry sub-tree transaction state did not yet exist and had to be created.</summary>
            STATUS_RXACT_STATE_CREATED = 0x40000004,

            /// <summary>{Segment Load} A virtual DOS machine (VDM) is loading, unloading, or moving an MS-DOS or Win16 program segment image.
            /// An exception is raised so a debugger can load, unload or track symbols and breakpoints within these 16-bit segments.</summary>
            STATUS_SEGMENT_NOTIFICATION = 0x40000005,

            /// <summary>{Local Session Key} A user session key was requested for a local RPC connection. The session key returned is a constant value and not unique to this connection.</summary>
            STATUS_LOCAL_USER_SESSION_KEY = 0x40000006,

            /// <summary>{Invalid Current Directory} The process cannot switch to the startup current directory %hs.
            /// Select OK to set current directory to %hs, or select CANCEL to exit.</summary>
            STATUS_BAD_CURRENT_DIRECTORY = 0x40000007,

            /// <summary>{Serial IOCTL Complete} A serial I/O operation was completed by another write to a serial port.
            /// (The IOCTL_SERIAL_XOFF_COUNTER reached zero.)</summary>
            STATUS_SERIAL_MORE_WRITES = 0x40000008,

            /// <summary>{Registry Recovery} One of the files containing the system's Registry data had to be recovered by use of a log or alternate copy. The recovery was successful.</summary>
            STATUS_REGISTRY_RECOVERED = 0x40000009,

            /// <summary>{Redundant Read} To satisfy a read request, the NT fault-tolerant file system successfully read the requested data from a redundant copy.
            /// This was done because the file system encountered a failure on a member of the fault-tolerant volume, but was unable to reassign the failing area of the device.</summary>
            STATUS_FT_READ_RECOVERY_FROM_BACKUP = 0x4000000A,

            /// <summary>{Redundant Write} To satisfy a write request, the NT fault-tolerant file system successfully wrote a redundant copy of the information.
            /// This was done because the file system encountered a failure on a member of the fault-tolerant volume, but was not able to reassign the failing area of the device.</summary>
            STATUS_FT_WRITE_RECOVERY = 0x4000000B,

            /// <summary>{Serial IOCTL Timeout} A serial I/O operation completed because the time-out period expired. (The IOCTL_SERIAL_XOFF_COUNTER had not reached zero.)</summary>
            STATUS_SERIAL_COUNTER_TIMEOUT = 0x4000000C,

            /// <summary>{Password Too Complex} The Windows password is too complex to be converted to a LAN Manager password. The LAN Manager password returned is a NULL string.</summary>
            STATUS_NULL_LM_PASSWORD = 0x4000000D,

            /// <summary>{Machine Type Mismatch} The image file %hs is valid, but is for a machine type other than the current machine. Select OK to continue, or CANCEL to fail the DLL load.</summary>
            STATUS_IMAGE_MACHINE_TYPE_MISMATCH = 0x4000000E,

            /// <summary>{Partial Data Received} The network transport returned partial data to its client. The remaining data will be sent later.</summary>
            STATUS_RECEIVE_PARTIAL = 0x4000000F,

            /// <summary>{Expedited Data Received} The network transport returned data to its client that was marked as expedited by the remote system.</summary>
            STATUS_RECEIVE_EXPEDITED = 0x40000010,

            /// <summary>{Partial Expedited Data Received} The network transport returned partial data to its client and this data was marked as expedited by the remote system. The remaining data will be sent later.</summary>
            STATUS_RECEIVE_PARTIAL_EXPEDITED = 0x40000011,

            /// <summary>{TDI Event Done} The TDI indication has completed successfully.</summary>
            STATUS_EVENT_DONE = 0x40000012,

            /// <summary>{TDI Event Pending} The TDI indication has entered the pending state.</summary>
            STATUS_EVENT_PENDING = 0x40000013,

            /// <summary>Checking file system on %wZ</summary>
            STATUS_CHECKING_FILE_SYSTEM = 0x40000014,

            /// <summary>{Fatal Application Exit} %hs</summary>
            STATUS_FATAL_APP_EXIT = 0x40000015,

            /// <summary>A requested read/write cannot be granted due to a conflicting file lock.</summary>
            STATUS_FILE_LOCK_CONFLICT = 0xC0000054,

            /// <summary>A requested file lock cannot be granted due to other existing locks.</summary>
            STATUS_LOCK_NOT_GRANTED = 0xC0000055,

            /// <summary>A non close operation has been requested of a file object with a delete pending.</summary>
            STATUS_DELETE_PENDING = 0xC0000056,

            /// <summary>An attempt was made to set the control attribute on a file. This attribute is not supported in the target file system.</summary>
            STATUS_CTL_FILE_NOT_SUPPORTED = 0xC0000057,

            /// <summary>Indicates a revision number encountered or specified is not one known by the service. It may be a more recent revision than the service is aware of.</summary>
            STATUS_UNKNOWN_REVISION = 0xC0000058,

            /// <summary>Indicates two revision levels are incompatible.</summary>
            STATUS_REVISION_MISMATCH = 0xC0000059,

            /// <summary>Indicates a particular Security ID may not be assigned as the owner of an object.</summary>
            STATUS_INVALID_OWNER = 0xC000005A,

            /// <summary>Indicates a particular Security ID may not be assigned as the primary group of an object.</summary>
            STATUS_INVALID_PRIMARY_GROUP = 0xC000005B,

            /// <summary>An attempt has been made to operate on an impersonation token by a thread that is not currently impersonating a client.</summary>
            STATUS_NO_IMPERSONATION_TOKEN = 0xC000005C,

            /// <summary>A mandatory group may not be disabled.</summary>
            STATUS_CANT_DISABLE_MANDATORY = 0xC000005D,

            /// <summary>There are currently no logon servers available to service the logon request.</summary>
            STATUS_NO_LOGON_SERVERS = 0xC000005E,

            /// <summary>A specified logon session does not exist. It may already have been terminated.</summary>
            STATUS_NO_SUCH_LOGON_SESSION = 0xC000005F,

            /// <summary>A specified privilege does not exist.</summary>
            STATUS_NO_SUCH_PRIVILEGE = 0xC0000060,

            /// <summary>A required privilege is not held by the client.</summary>
            STATUS_PRIVILEGE_NOT_HELD = 0xC0000061,

            /// <summary>The name provided is not a properly formed account name.</summary>
            STATUS_INVALID_ACCOUNT_NAME = 0xC0000062,

            /// <summary>The specified account already exists.</summary>
            STATUS_USER_EXISTS = 0xC0000063,

            /// <summary>The specified account does not exist.</summary>
            STATUS_NO_SUCH_USER = 0xC0000064,

            /// <summary>The specified group already exists.</summary>
            STATUS_GROUP_EXISTS = 0xC0000065,

            /// <summary>The specified group does not exist.</summary>
            STATUS_NO_SUCH_GROUP = 0xC0000066,

            /// <summary>The specified user account is already in the specified group account. Also used to indicate a group cannot be deleted because it contains a member.</summary>
            STATUS_MEMBER_IN_GROUP = 0xC0000067,

            /// <summary>The specified user account is not a member of the specified group account.</summary>
            STATUS_MEMBER_NOT_IN_GROUP = 0xC0000068,

            /// <summary>Indicates the requested operation would disable, delete or could prevent logon for an administration account. This is not allowed to prevent creating a situation in which the system cannot be administrated.</summary>
            STATUS_LAST_ADMIN = 0xC0000069,

            /// <summary>When trying to update a password, this return status indicates that the value provided as the current password is not correct.</summary>
            STATUS_WRONG_PASSWORD = 0xC000006A,

            /// <summary>When trying to update a password, this return status indicates that the value provided for the new password contains values that are not allowed in passwords.</summary>
            STATUS_ILL_FORMED_PASSWORD = 0xC000006B,

            /// <summary>When trying to update a password, this status indicates that some password update rule has been violated. For example, the password may not meet length criteria.</summary>
            STATUS_PASSWORD_RESTRICTION = 0xC000006C,

            /// <summary>The attempted logon is invalid. This is either due to a bad username or authentication information.</summary>
            STATUS_LOGON_FAILURE = 0xC000006D,

            /// <summary>Indicates a referenced user name and authentication information are valid, but some user account restriction has prevented successful authentication (such as time-of-day restrictions).</summary>
            STATUS_ACCOUNT_RESTRICTION = 0xC000006E,

            /// <summary>The user account has time restrictions and may not be logged onto at this time.</summary>
            STATUS_INVALID_LOGON_HOURS = 0xC000006F,

            /// <summary>The user account is restricted such that it may not be used to log on from the source workstation.</summary>
            STATUS_INVALID_WORKSTATION = 0xC0000070,

            /// <summary>The user account's password has expired.</summary>
            STATUS_PASSWORD_EXPIRED = 0xC0000071,

            /// <summary>The referenced account is currently disabled and may not be logged on to.</summary>
            STATUS_ACCOUNT_DISABLED = 0xC0000072,

            /// <summary>None of the information to be translated has been translated.</summary>
            STATUS_NONE_MAPPED = 0xC0000073,

            /// <summary>The number of LUIDs requested may not be allocated with a single allocation.</summary>
            STATUS_TOO_MANY_LUIDS_REQUESTED = 0xC0000074,

            /// <summary>Indicates there are no more LUIDs to allocate.</summary>
            STATUS_LUIDS_EXHAUSTED = 0xC0000075,

            /// <summary>Indicates the sub-authority value is invalid for the particular use.</summary>
            STATUS_INVALID_SUB_AUTHORITY = 0xC0000076,

            /// <summary>Indicates the ACL structure is not valid.</summary>
            STATUS_INVALID_ACL = 0xC0000077,

            /// <summary>Indicates the SID structure is not valid.</summary>
            STATUS_INVALID_SID = 0xC0000078,

            /// <summary>Indicates the SECURITY_DESCRIPTOR structure is not valid.</summary>
            STATUS_INVALID_SECURITY_DESCR = 0xC0000079,

            /// <summary>Indicates the specified procedure address cannot be found in the DLL.</summary>
            STATUS_PROCEDURE_NOT_FOUND = 0xC000007A,

            /// <summary>{Bad Image} %hs is either not designed to run on Windows or it contains an error. Try installing the program again using the original installation media or contact your system administrator or the software vendor for support. Error status 0x%08lx.</summary>
            STATUS_INVALID_IMAGE_FORMAT = 0xC000007B,

            /// <summary>An attempt was made to reference a token that doesn't exist. This is typically done by referencing the token associated with a thread when the thread is not impersonating a client.</summary>
            STATUS_NO_TOKEN = 0xC000007C,

            /// <summary>Indicates that an attempt to build either an inherited ACL or ACE was not successful. This can be caused by a number of things. One of the more probable causes is the replacement of a CreatorId with an SID that didn't fit into the ACE or ACL.</summary>
            STATUS_BAD_INHERITANCE_ACL = 0xC000007D,

            /// <summary>The range specified in NtUnlockFile was not locked.</summary>
            STATUS_RANGE_NOT_LOCKED = 0xC000007E,

            /// <summary>An operation failed because the disk was full. If this is a thinly provisioned volume the physical storage backing this volume has been exhausted.</summary>
            STATUS_DISK_FULL = 0xC000007F,

            /// <summary>The GUID allocation server is [already] disabled at the moment.</summary>
            STATUS_SERVER_DISABLED = 0xC0000080,

            /// <summary>The GUID allocation server is [already] enabled at the moment.</summary>
            STATUS_SERVER_NOT_DISABLED = 0xC0000081,

            /// <summary>Too many GUIDs were requested from the allocation server at once.</summary>
            STATUS_TOO_MANY_GUIDS_REQUESTED = 0xC0000082,

            /// <summary>The GUIDs could not be allocated because the Authority Agent was exhausted.</summary>
            STATUS_GUIDS_EXHAUSTED = 0xC0000083,

            /// <summary>The value provided was an invalid value for an identifier authority.</summary>
            STATUS_INVALID_ID_AUTHORITY = 0xC0000084,

            /// <summary>The address handle given to the transport was invalid.</summary>
            STATUS_INVALID_ADDRESS = 0xC0000141,

            /// <summary>{DLL Initialization Failed} Initialization of the dynamic link library %hs failed. The process is terminating abnormally.</summary>
            STATUS_DLL_INIT_FAILED = 0xC0000142,

            /// <summary>{Missing System File} The required system file %hs is bad or missing.</summary>
            STATUS_MISSING_SYSTEMFILE = 0xC0000143,

            /// <summary>{Application Error} The exception %s (0x%08lx) occurred in the application at location 0x%08lx.</summary>
            STATUS_UNHANDLED_EXCEPTION = 0xC0000144,

            /// <summary>{Application Error} The application was unable to start correctly (0x%lx). Click OK to close the application.</summary>
            STATUS_APP_INIT_FAILURE = 0xC0000145,

            /// <summary>{Unable to Create Paging File} The creation of the paging file %hs failed (%lx). The requested size was %ld.</summary>
            STATUS_PAGEFILE_CREATE_FAILED = 0xC0000146,

            /// <summary>{No Paging File Specified} No paging file was specified in the system configuration.</summary>
            STATUS_NO_PAGEFILE = 0xC0000147,

            /// <summary>{Incorrect System Call Level} An invalid level was passed into the specified system call.</summary>
            STATUS_INVALID_LEVEL = 0xC0000148,

            /// <summary>{Incorrect Password to LAN Manager Server} You specified an incorrect password to a LAN Manager 2.x or MS-NET server.</summary>
            STATUS_WRONG_PASSWORD_CORE = 0xC0000149,

            /// <summary>{EXCEPTION} A real-mode application issued a floating-point instruction and floating-point hardware is not present.</summary>
            STATUS_ILLEGAL_FLOAT_CONTEXT = 0xC000014A,

            /// <summary>The pipe operation has failed because the other end of the pipe has been closed.</summary>
            STATUS_PIPE_BROKEN = 0xC000014B,

            /// <summary>{The Registry Is Corrupt} The structure of one of the files that contains Registry data is corrupt, or the image of the file in memory is corrupt, or the file could not be recovered because the alternate copy or log was absent or corrupt.</summary>
            STATUS_REGISTRY_CORRUPT = 0xC000014C,

            /// <summary>An I/O operation initiated by the Registry failed unrecoverably. The Registry could not read in, or write out, or flush, one of the files that contain the system's image of the Registry.</summary>
            STATUS_REGISTRY_IO_FAILED = 0xC000014D,

            /// <summary>An event pair synchronization operation was performed using the thread specific client/server event pair object, but no event pair object was associated with the thread.</summary>
            STATUS_NO_EVENT_PAIR = 0xC000014E,

            /// <summary>The volume does not contain a recognized file system. Please make sure that all required file system drivers are loaded and that the volume is not corrupt.</summary>
            STATUS_UNRECOGNIZED_VOLUME = 0xC000014F,

            /// <summary>No serial device was successfully initialized. The serial driver will unload.</summary>
            STATUS_SERIAL_NO_DEVICE_INITED = 0xC0000150,

            /// <summary>The specified local group does not exist.</summary>
            STATUS_NO_SUCH_ALIAS = 0xC0000151,

            /// <summary>{Invalid ACE Condition} The specified file already has a notification GUID associated with it.</summary>
            STATUS_NOTIFICATION_GUID_ALREADY_DEFINED = 0xC00001A4,

            /// <summary>An invalid exception handler routine has been detected.</summary>
            STATUS_INVALID_EXCEPTION_HANDLER = 0xC00001A5,

            /// <summary>Duplicate privileges were specified for the token.</summary>
            STATUS_DUPLICATE_PRIVILEGES = 0xC00001A6,

            /// <summary>Requested action not allowed on a file system internal file.</summary>
            STATUS_NOT_ALLOWED_ON_SYSTEM_FILE = 0xC00001A7,

            /// <summary>A portion of the file system requires repair.</summary>
            STATUS_REPAIR_NEEDED = 0xC00001A8,

            /// <summary>Quota support is not enabled on the system.</summary>
            STATUS_QUOTA_NOT_ENABLED = 0xC00001A9,

            /// <summary>The operation failed because the application is not part of an application package.</summary>
            STATUS_NO_APPLICATION_PACKAGE = 0xC00001AA,

            /// <summary>A remote open failed because the network open restrictions were not satisfied.</summary>
            STATUS_NETWORK_OPEN_RESTRICTION = 0xC0000201,

            /// <summary>There is no user session key for the specified logon session.</summary>
            STATUS_NO_USER_SESSION_KEY = 0xC0000202,

            /// <summary>The remote user session has been deleted.</summary>
            STATUS_USER_SESSION_DELETED = 0xC0000203,

            /// <summary>Indicates the specified resource language ID cannot be found in the image file.</summary>
            STATUS_RESOURCE_LANG_NOT_FOUND = 0xC0000204,

            /// <summary>Insufficient server resources exist to complete the request.</summary>
            STATUS_INSUFF_SERVER_RESOURCES = 0xC0000205,

            /// <summary>The size of the buffer is invalid for the specified operation.</summary>
            STATUS_INVALID_BUFFER_SIZE = 0xC0000206,

            /// <summary>The transport rejected the network address specified as invalid.</summary>
            STATUS_INVALID_ADDRESS_COMPONENT = 0xC0000207,

            /// <summary>The transport rejected the network address specified due to an invalid use of a wildcard.</summary>
            STATUS_INVALID_ADDRESS_WILDCARD = 0xC0000208,

            /// <summary>The transport address could not be opened because all the available addresses are in use.</summary>
            STATUS_TOO_MANY_ADDRESSES = 0xC0000209,

            /// <summary>The transport address could not be opened because it already exists.</summary>
            STATUS_ADDRESS_ALREADY_EXISTS = 0xC000020A,

            /// <summary>{Unable to Retrieve Browser Server List} The list of servers for this workgroup is not currently available.</summary>
            STATUS_NO_BROWSER_SERVERS_FOUND = 0xC000021C,

            /// <summary>NTVDM encountered a hard error.</summary>
            STATUS_VDM_HARD_ERROR = 0xC000021D,

            /// <summary>{Cancel Timeout} The driver %hs failed to complete a cancelled I/O request in the allotted time.</summary>
            STATUS_DRIVER_CANCEL_TIMEOUT = 0xC000021E,

            /// <summary>{Reply Message Mismatch} An attempt was made to reply to an LPC message, but the thread specified by the client ID in the message was not waiting on that message.</summary>
            STATUS_REPLY_MESSAGE_MISMATCH = 0xC000021F,

            /// <summary>{Mapped View Alignment Incorrect} An attempt was made to map a view of a file, but either the specified base address or the offset into the file were not aligned on the proper allocation granularity.</summary>
            STATUS_MAPPED_ALIGNMENT = 0xC0000220,

            /// <summary>{Bad Image Checksum} The image %hs is possibly corrupt. The header checksum does not match the computed checksum.</summary>
            STATUS_IMAGE_CHECKSUM_MISMATCH = 0xC0000221,

            /// <summary>{Delayed Write Failed} Windows was unable to save all the data for the file %hs. The data has been lost. This error may be caused by a failure of your computer hardware or network connection. Please try to save this file elsewhere.</summary>
            STATUS_LOST_WRITEBEHIND_DATA = 0xC0000222,

            /// <summary>The parameter(s) passed to the server in the client/server shared memory window were invalid. Too much data may have been put in the shared memory window.</summary>
            STATUS_CLIENT_SERVER_PARAMETERS_INVALID = 0xC0000223,

            /// <summary>The user's password must be changed before signing in.</summary>
            STATUS_PASSWORD_MUST_CHANGE = 0xC0000224,

            /// <summary>The object was not found.</summary>
            STATUS_NOT_FOUND = 0xC0000225,

            /// <summary>The stream is not a tiny stream.</summary>
            STATUS_NOT_TINY_STREAM = 0xC0000226,

            /// <summary>A transaction recover failed.</summary>
            STATUS_RECOVERY_FAILURE = 0xC0000227,

            /// <summary>The request must be handled by the stack overflow code.</summary>
            STATUS_STACK_OVERFLOW_READ = 0xC0000228,

            /// <summary>A consistency check failed.</summary>
            STATUS_FAIL_CHECK = 0xC0000229,

            /// <summary>The attempt to insert the ID in the index failed because the ID is already in the index.</summary>
            STATUS_DUPLICATE_OBJECTID = 0xC000022A,

            /// <summary>The attempt to set the object's ID failed because the object already has an ID.</summary>
            STATUS_OBJECTID_EXISTS = 0xC000022B,

            /// <summary>Internal OFS status codes indicating how an allocation operation is handled. Either it is retried after the containing onode is moved or the extent stream is converted to a large stream.</summary>
            STATUS_CONVERT_TO_LARGE = 0xC000022C,

            /// <summary>{Invalid Service Callback Entrypoint} The %hs service is not written correctly. The stack pointer has been left in an inconsistent state. The callback entrypoint should be declared as WINAPI or STDCALL. Selecting OK will cause the service to continue operation. However, the service process may operate incorrectly.</summary>
            STATUS_BAD_SERVICE_ENTRYPOINT = 0xC0000252,

            /// <summary>The server received the messages but did not send a reply.</summary>
            STATUS_LPC_REPLY_LOST = 0xC0000253,

            /// <summary>There is an IP address conflict with another system on the network</summary>
            STATUS_IP_ADDRESS_CONFLICT1 = 0xC0000254,

            /// <summary>There is an IP address conflict with another system on the network</summary>
            STATUS_IP_ADDRESS_CONFLICT2 = 0xC0000255,

            /// <summary>{Low On Registry Space} The system has reached the maximum size allowed for the system part of the registry. Additional storage requests will be ignored.</summary>
            STATUS_REGISTRY_QUOTA_LIMIT = 0xC0000256,

            /// <summary>The contacted server does not support the indicated part of the DFS namespace.</summary>
            STATUS_PATH_NOT_COVERED = 0xC0000257,

            /// <summary>A callback return system service cannot be executed when no callback is active.</summary>
            STATUS_NO_CALLBACK_ACTIVE = 0xC0000258,

            /// <summary>The service being accessed is licensed for a particular number of connections. No more connections can be made to the service at this time because there are already as many connections as the service can accept.</summary>
            STATUS_LICENSE_QUOTA_EXCEEDED = 0xC0000259,

            /// <summary>The password provided is too short to meet the policy of your user account. Please choose a longer password.</summary>
            STATUS_PWD_TOO_SHORT = 0xC000025A,

            /// <summary>The policy of your user account does not allow you to change passwords too frequently. This is done to prevent users from changing back to a familiar, but potentially discovered, password. If you feel your password has been compromised then please contact your administrator immediately to have a new one assigned.</summary>
            STATUS_PWD_TOO_RECENT = 0xC000025B,

            /// <summary>You have attempted to change your password to one that you have used in the past. The policy of your user account does not allow this. Please select a password that you have not previously used.</summary>
            STATUS_PWD_HISTORY_CONFLICT = 0xC000025C,

            /// <summary>You have attempted to load a legacy device driver while its device instance had been disabled.</summary>
            STATUS_PLUGPLAY_NO_DEVICE = 0xC000025E,

            /// <summary>The specified compression format is unsupported.</summary>
            STATUS_UNSUPPORTED_COMPRESSION = 0xC000025F,

            /// <summary>The specified hardware profile configuration is invalid.</summary>
            STATUS_INVALID_HW_PROFILE = 0xC0000260,

            /// <summary>The specified Plug and Play registry device path is invalid.</summary>
            STATUS_INVALID_PLUGPLAY_DEVICE_PATH = 0xC0000261,

            /// <summary>{Driver Entry Point Not Found} The %hs device driver could not locate the ordinal %ld in driver %hs.</summary>
            STATUS_DRIVER_ORDINAL_NOT_FOUND = 0xC0000262,

            /// <summary>{Driver Entry Point Not Found} The %hs device driver could not locate the entry point %hs in driver %hs.</summary>
            STATUS_DRIVER_ENTRYPOINT_NOT_FOUND = 0xC0000263,

            /// <summary>{Application Error} The application attempted to release a resource it did not own. Click OK to terminate the application.</summary>
            STATUS_RESOURCE_NOT_OWNED = 0xC0000264,

            /// <summary>An attempt was made to create more links on a file than the file system supports.</summary>
            STATUS_TOO_MANY_LINKS = 0xC0000265,

            /// <summary>The validation process needs to continue on to the next step.</summary>
            STATUS_VALIDATE_CONTINUE = 0xC0000271,

            /// <summary>There was no match for the specified key in the index.</summary>
            STATUS_NO_MATCH = 0xC0000272,

            /// <summary>There are no more matches for the current index enumeration.</summary>
            STATUS_NO_MORE_MATCHES = 0xC0000273,

            /// <summary>The NTFS file or directory is not a reparse point.</summary>
            STATUS_NOT_A_REPARSE_POINT = 0xC0000275,

            /// <summary>The Windows I/O reparse tag passed for the NTFS reparse point is invalid.</summary>
            STATUS_IO_REPARSE_TAG_INVALID = 0xC0000276,

            /// <summary>The Windows I/O reparse tag does not match the one present in the NTFS reparse point.</summary>
            STATUS_IO_REPARSE_TAG_MISMATCH = 0xC0000277,

            /// <summary>The user data passed for the NTFS reparse point is invalid.</summary>
            STATUS_IO_REPARSE_DATA_INVALID = 0xC0000278,

            /// <summary>The layered file system driver for this IO tag did not handle it when needed.</summary>
            STATUS_IO_REPARSE_TAG_NOT_HANDLED = 0xC0000279,

            /// <summary>The password provided is too long to meet the policy of your user account. Please choose a shorter password.</summary>
            STATUS_PWD_TOO_LONG = 0xC000027A,

            /// <summary>An application-internal exception has occurred.</summary>
            STATUS_STOWED_EXCEPTION = 0xC000027B,

            /// <summary>The NTFS symbolic link could not be resolved even though the initial file name is valid.</summary>
            STATUS_REPARSE_POINT_NOT_RESOLVED = 0xC0000280,

            /// <summary>The NTFS directory is a reparse point.</summary>
            STATUS_DIRECTORY_IS_A_REPARSE_POINT = 0xC0000281,

            /// <summary>The range could not be added to the range list because of a conflict.</summary>
            STATUS_RANGE_LIST_CONFLICT = 0xC0000282,

            /// <summary>The specified medium changer source element contains no media.</summary>
            STATUS_SOURCE_ELEMENT_EMPTY = 0xC0000283,

            /// <summary>The specified medium changer destination element already contains media.</summary>
            STATUS_DESTINATION_ELEMENT_FULL = 0xC0000284,

            /// <summary>The specified medium changer element does not exist.</summary>
            STATUS_ILLEGAL_ELEMENT_ADDRESS = 0xC0000285,

            /// <summary>The specified element is contained within a magazine that is no longer present.</summary>
            STATUS_MAGAZINE_NOT_PRESENT = 0xC0000286,

            /// <summary>A global group cannot have a local group as a member.</summary>
            STATUS_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER = 0xC00002D7,

            /// <summary>A global group cannot have a universal group as a member.</summary>
            STATUS_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER = 0xC00002D8,

            /// <summary>A universal group cannot have a local group as a member.</summary>
            STATUS_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER = 0xC00002D9,

            /// <summary>A global group cannot have a cross domain member.</summary>
            STATUS_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER = 0xC00002DA,

            /// <summary>A local group cannot have another cross domain local group as a member.</summary>
            STATUS_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER = 0xC00002DB,

            /// <summary>Cannot change to security disabled group because of having primary members in this group.</summary>
            STATUS_DS_HAVE_PRIMARY_MEMBERS = 0xC00002DC,

            /// <summary>The WMI operation is not supported by the data block or method.</summary>
            STATUS_WMI_NOT_SUPPORTED = 0xC00002DD,

            /// <summary>There is not enough power to complete the requested operation.</summary>
            STATUS_INSUFFICIENT_POWER = 0xC00002DE,

            /// <summary>Security Account Manager needs to get the boot password.</summary>
            STATUS_SAM_NEED_BOOTKEY_PASSWORD = 0xC00002DF,

            /// <summary>Security Account Manager needs to get the boot key from floppy disk.</summary>
            STATUS_SAM_NEED_BOOTKEY_FLOPPY = 0xC00002E0,

            /// <summary>Directory Service cannot start.</summary>
            STATUS_DS_CANT_START = 0xC00002E1,

            /// <summary>Directory Services could not start because of the following error: %hs
            /// Error Status: 0x%x.
            /// Please shutdown this system and reboot into Directory Services Restore Mode, check the event log for more detailed information.</summary>
            STATUS_DS_INIT_FAILURE = 0xC00002E2,

            /// <summary>Security Accounts Manager initialization failed because of the following error: %hs
            /// Error Status: 0x%x.
            /// Please click OK to shutdown this system and reboot into Safe Mode, check the event log for more detailed information.</summary>
            STATUS_SAM_INIT_FAILURE = 0xC00002E3,

            /// <summary>The requested operation can be performed only on a global catalog server.</summary>
            STATUS_DS_GC_REQUIRED = 0xC00002E4,

            /// <summary>A local group can only be a member of other local groups in the same domain.</summary>
            STATUS_DS_LOCAL_MEMBER_OF_LOCAL_ONLY = 0xC00002E5,

            /// <summary>Foreign security principals cannot be members of universal groups.</summary>
            STATUS_DS_NO_FPO_IN_UNIVERSAL_GROUPS = 0xC00002E6,
            /// <summary>The Master File Table on the volume is too fragmented to complete this operation.</summary>
            STATUS_MFT_TOO_FRAGMENTED = 0xC0000304,

            /// <summary>Copy protection failure.</summary>
            STATUS_COPY_PROTECTION_FAILURE = 0xC0000305,

            /// <summary>Copy protection error - DVD CSS Authentication failed.</summary>
            STATUS_CSS_AUTHENTICATION_FAILURE = 0xC0000306,

            /// <summary>Copy protection error - The given sector does not contain a valid key.</summary>
            STATUS_CSS_KEY_NOT_PRESENT = 0xC0000307,

            /// <summary>Copy protection error - DVD session key not established.</summary>
            STATUS_CSS_KEY_NOT_ESTABLISHED = 0xC0000308,

            /// <summary>Copy protection error - The read failed because the sector is encrypted.</summary>
            STATUS_CSS_SCRAMBLED_SECTOR = 0xC0000309,

            /// <summary>Copy protection error - The given DVD's region does not correspond to the region setting of the drive.</summary>
            STATUS_CSS_REGION_MISMATCH = 0xC000030A,

            /// <summary>Copy protection error - The drive's region setting may be permanent.</summary>
            STATUS_CSS_RESETS_EXHAUSTED = 0xC000030B,

            /// <summary>EAS policy requires that the user change their password before this operation can be performed.</summary>
            STATUS_PASSWORD_CHANGE_REQUIRED = 0xC000030C,
            /*++

             MessageId's 0x030c - 0x031f (inclusive) are reserved for future **STORAGE**
             copy protection errors.

            --*/

            /// <summary>The Kerberos protocol encountered an error while validating the KDC certificate during smartcard Logon. There is more information in the system event log.</summary>
            STATUS_PKINIT_FAILURE = 0xC0000320,

            /// <summary>The Kerberos protocol encountered an error while attempting to utilize the smartcard subsystem.</summary>
            STATUS_SMARTCARD_SUBSYSTEM_FAILURE = 0xC0000321,

            /// <summary>The target server does not have acceptable Kerberos credentials.</summary>
            STATUS_NO_KERB_KEY = 0xC0000322,
            /*++

             MessageId's 0x0323 - 0x034f (inclusive) are reserved for other future copy
             protection errors.

            --*/

            /// <summary>The transport determined that the remote system is down.</summary>
            STATUS_HOST_DOWN = 0xC0000350,

            /// <summary>An unsupported preauthentication mechanism was presented to the Kerberos package.</summary>
            STATUS_UNSUPPORTED_PREAUTH = 0xC0000351,

            /// <summary>The encryption algorithm used on the source file needs a bigger key buffer than the one used on the destination file.</summary>
            STATUS_EFS_ALG_BLOB_TOO_BIG = 0xC0000352,

            /// <summary>An attempt to remove a process's DebugPort was made, but a port was not already associated with the process.</summary>
            STATUS_PORT_NOT_SET = 0xC0000353,

            /// <summary>Debugger Inactive: Windows may have been started without kernel debugging enabled.</summary>
            STATUS_DEBUGGER_INACTIVE = 0xC0000354,

            /// <summary>This version of Windows is not compatible with the behavior version of directory forest, domain or domain controller.</summary>
            STATUS_DS_VERSION_CHECK_FAILURE = 0xC0000355,

            /// <summary>The specified thread is not in a session.</summary>
            STATUS_THREAD_NOT_IN_SESSION = 0xC0000457,

            /// <summary>The specified weight is invalid.</summary>
            STATUS_INVALID_WEIGHT = 0xC0000458,

            /// <summary>The operation was paused.</summary>
            STATUS_REQUEST_PAUSED = 0xC0000459,

            /// <summary>No ranges for the specified operation were able to be processed.</summary>
            STATUS_NO_RANGES_PROCESSED = 0xC0000460,

            /// <summary>The physical resources of this disk have been exhausted.</summary>
            STATUS_DISK_RESOURCES_EXHAUSTED = 0xC0000461,

            /// <summary>The application cannot be started. Try reinstalling the application to fix the problem.</summary>
            STATUS_NEEDS_REMEDIATION = 0xC0000462,

            /// <summary>{Device Feature Not Supported} The device does not support the command feature.</summary>
            STATUS_DEVICE_FEATURE_NOT_SUPPORTED = 0xC0000463,

            /// <summary>{Source/Destination device unreachable} The device is unreachable.</summary>
            STATUS_DEVICE_UNREACHABLE = 0xC0000464,

            /// <summary>{Invalid Proxy Data Token} The token representing the data is invalid.</summary>
            STATUS_INVALID_TOKEN = 0xC0000465,

            /// <summary>The file server is temporarily unavailable.</summary>
            STATUS_SERVER_UNAVAILABLE = 0xC0000466,

            /// <summary>The file is temporarily unavailable.</summary>
            STATUS_FILE_NOT_AVAILABLE = 0xC0000467,

            /// <summary>{Device Insufficient Resources} The target device has insufficient resources to complete the operation.</summary>
            STATUS_DEVICE_INSUFFICIENT_RESOURCES = 0xC0000468,

            /// <summary>The application cannot be started because it is currently updating.</summary>
            STATUS_PACKAGE_UPDATING = 0xC0000469,

            /// <summary>The specified copy of the requested data could not be read.</summary>
            STATUS_NOT_READ_FROM_COPY = 0xC000046A,

            /// <summary>The specified data could not be written to any of the copies.</summary>
            STATUS_FT_WRITE_FAILURE = 0xC000046B,

            /// <summary>One or more copies of data on this device may be out of sync. No writes may be performed until a data integrity scan is completed.</summary>
            STATUS_FT_DI_SCAN_REQUIRED = 0xC000046C,

            /// <summary>This object is not externally backed by any provider.</summary>
            STATUS_OBJECT_NOT_EXTERNALLY_BACKED = 0xC000046D,

            /// <summary>The file must be checked out before saving changes.</summary>
            STATUS_CHECKOUT_REQUIRED = 0xC0000902,

            /// <summary>The file type being saved or retrieved has been blocked.</summary>
            STATUS_BAD_FILE_TYPE = 0xC0000903,

            /// <summary>The file size exceeds the limit allowed and cannot be saved.</summary>
            STATUS_FILE_TOO_LARGE = 0xC0000904,

            /// <summary>Access Denied. Before opening files in this location, you must first browse to the web site and select the option to login automatically.</summary>
            STATUS_FORMS_AUTH_REQUIRED = 0xC0000905,

            /// <summary>Operation did not complete successfully because the file contains a virus or potentially unwanted software.</summary>
            STATUS_VIRUS_INFECTED = 0xC0000906,

            /// <summary>This file contains a virus or potentially unwanted software and cannot be opened. Due to the nature of this virus or potentially unwanted software, the file has been removed from this location.</summary>
            STATUS_VIRUS_DELETED = 0xC0000907,

            /// <summary>The resources required for this device conflict with the MCFG table.</summary>
            STATUS_BAD_MCFG_TABLE = 0xC0000908,

            /// <summary>The operation did not complete successfully because it would cause an oplock to be broken. The caller has requested that existing oplocks not be broken.</summary>
            STATUS_CANNOT_BREAK_OPLOCK = 0xC0000909,

            /// <summary>Bad key.</summary>
            STATUS_BAD_KEY = 0xC000090A,

            /// <summary>Bad data.</summary>
            STATUS_BAD_DATA = 0xC000090B,

            /// <summary>Key does not exist.</summary>
            STATUS_NO_KEY = 0xC000090C,

            /// <summary>Access to the specified file handle has been revoked.</summary>
            STATUS_FILE_HANDLE_REVOKED = 0xC0000910,

            /// <summary>WOW Assertion Error.</summary>
            STATUS_WOW_ASSERTION = 0xC0009898,

            /// <summary>The cryptographic signature is invalid.</summary>
            STATUS_INVALID_SIGNATURE = 0xC000A000,

            /// <summary>The cryptographic provider does not support HMAC.</summary>
            STATUS_HMAC_NOT_SUPPORTED = 0xC000A001,

            /// <summary>The computed authentication tag did not match the input authentication tag.</summary>
            STATUS_AUTH_TAG_MISMATCH = 0xC000A002,

            /// <summary>The requested state transition is invalid and cannot be performed.</summary>
            STATUS_INVALID_STATE_TRANSITION = 0xC000A003,

            /// <summary>The supplied kernel information version is invalid.</summary>
            STATUS_INVALID_KERNEL_INFO_VERSION = 0xC000A004,

            /// <summary>The supplied PEP information version is invalid.</summary>
            STATUS_INVALID_PEP_INFO_VERSION = 0xC000A005,
            /*++
            /// <summary>The Boot Configuration Data (BCD) settings are not supported or have changed since BitLocker was enabled.</summary>
            STATUS_FVE_AUTH_INVALID_CONFIG = 0xC021001C,

            /// <summary>Boot debugging is enabled. Run bcdedit to turn it off.</summary>
            STATUS_FVE_DEBUGGER_ENABLED = 0xC021001D,

            /// <summary>The BitLocker encryption key could not be obtained.</summary>
            STATUS_FVE_DRY_RUN_FAILED = 0xC021001E,

            /// <summary>The metadata disk region pointer is incorrect.</summary>
            STATUS_FVE_BAD_METADATA_POINTER = 0xC021001F,

            /// <summary>The backup copy of the metadata is out of date.</summary>
            STATUS_FVE_OLD_METADATA_COPY = 0xC0210020,

            /// <summary>No action was taken as a system reboot is required.</summary>
            STATUS_FVE_REBOOT_REQUIRED = 0xC0210021,

            /// <summary>No action was taken as BitLocker Drive Encryption is in RAW access mode.</summary>
            STATUS_FVE_RAW_ACCESS = 0xC0210022,

            /// <summary>BitLocker Drive Encryption cannot enter raw access mode for this volume.</summary>
            STATUS_FVE_RAW_BLOCKED = 0xC0210023,

            /// <summary>The auto-unlock master key was not available from the operating system volume. Retry the operation using the BitLocker WMI interface.</summary>
            STATUS_FVE_NO_AUTOUNLOCK_MASTER_KEY = 0xC0210024,

            /// <summary>The system firmware failed to enable clearing of system memory on reboot.</summary>
            STATUS_FVE_MOR_FAILED = 0xC0210025,

            /// <summary>This feature of BitLocker Drive Encryption is not included with this version of Windows.</summary>
            STATUS_FVE_NO_FEATURE_LICENSE = 0xC0210026,

            /// <summary>Group policy does not permit turning off BitLocker Drive Encryption on roaming data volumes.</summary>
            STATUS_FVE_POLICY_USER_DISABLE_RDV_NOT_ALLOWED = 0xC0210027,

            /// <summary>Bitlocker Drive Encryption failed to recover from aborted conversion. This could be due to either all conversion logs being corrupted or the media being write-protected.</summary>
            STATUS_FVE_CONV_RECOVERY_FAILED = 0xC0210028,

            /// <summary>The requested virtualization size is too big.</summary>
            STATUS_FVE_VIRTUALIZED_SPACE_TOO_BIG = 0xC0210029,

            /// <summary>The management information stored on the drive contained an unknown type. If you are using an old version of Windows, try accessing the drive from the latest version.</summary>
            STATUS_FVE_INVALID_DATUM_TYPE = 0xC021002A,

            /// <summary>The drive is too small to be protected using BitLocker Drive Encryption.</summary>
            STATUS_FVE_VOLUME_TOO_SMALL = 0xC0210030,

            /// <summary>The BitLocker encryption key could not be obtained from the Trusted Platform Module (TPM) and enhanced PIN. Try using a PIN containing only numerals.</summary>
            STATUS_FVE_ENH_PIN_INVALID = 0xC0210031,


            /// <summary>Change Tracking for the virtual disk is not in a valid state to perform this request.  Change tracking could be discontinued or already in the requested state.</summary>
            STATUS_CTLOG_INVALID_TRACKING_STATE = 0xC03A0023,

            /// <summary>Change Tracking file for the virtual disk is not in a valid state.</summary>
            STATUS_CTLOG_INCONSISTENT_TRACKING_FILE = 0xC03A0024,

            /// <summary>There is not enough space in the virtual disk file for the provided metadata item.</summary>
            STATUS_VHD_METADATA_FULL = 0xC03A0028,

            #endregion

            #region Vhd warnings.

            /// <summary>The virtualization storage subsystem has generated an error.</summary>
            STATUS_QUERY_STORAGE_ERROR = 0x803A0001,

            #endregion

            #region Resume Key Filter (RKF) error codes.

            /// <summary>The Resume Key Filter could not find the resume key supplied for the operation.</summary>
            STATUS_RKF_KEY_NOT_FOUND = 0xC0400001,

            /// <summary>The Resume Key Filter found an existing resume key that matches the one supplied for the handle.</summary>
            STATUS_RKF_DUPLICATE_KEY = 0xC0400002,

            /// <summary>The Resume Key Filter data blob attached to the handle is full. No more space is available.</summary>
            STATUS_RKF_BLOB_FULL = 0xC0400003,

            /// <summary>The Resume Key Filter handle store is full. No more resume handles can be accepted.</summary>
            STATUS_RKF_STORE_FULL = 0xC0400004,

            /// <summary>The Resume Key Filter failed the operation because the file is temporarily blocked pending the resume of existing handles on the file.</summary>
            STATUS_RKF_FILE_BLOCKED = 0xC0400005,

            /// <summary>The Resume Key Filter found an existing resume key that matches the one supplied on a handle that's active/open. The operation requires an inactive/closed handle.</summary>
            STATUS_RKF_ACTIVE_KEY = 0xC0400006,

            #endregion

            #region RDBSS / MiniRdr internal error codes.

            /// <summary>The operation must be restarted by RDBSS.</summary>
            STATUS_RDBSS_RESTART_OPERATION = 0xC0410001,

            /// <summary>The operation must continue processing.</summary>
            STATUS_RDBSS_CONTINUE_OPERATION = 0xC0410002,

            /// <summary>The operation must be posted to a thread to be retried at passive IRQL.</summary>
            STATUS_RDBSS_POST_OPERATION = 0xC0410003,

            #endregion

            #region Bluetooth Attribute Protocol Warnings

            /// <summary>The attribute handle given was not valid on this server.</summary>
            STATUS_BTH_ATT_INVALID_HANDLE = 0xC0420001,

            /// <summary>The attribute cannot be read.</summary>
            STATUS_BTH_ATT_READ_NOT_PERMITTED = 0xC0420002,

            /// <summary>The attribute cannot be written.</summary>