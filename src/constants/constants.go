package constants

const (
	// NoraTime : Time at which it all began.
	NoraTime = 1538421540000
)

const (
	// CouchMissingReason : Reason returned when resource is missing.
	CouchMissingReason = "missing"
	// CouchDeletedReason : Reason returned when resource is deleted.
	CouchDeletedReason = "deleted"
	// CouchUpdateConflictReason : Reason returned when document updation is attempted without latest rev.
	CouchUpdateConflictReason = "Document update conflict."

	// CouchDesign : Name of the Design document in CouchDB.
	CouchDesign = "noraDesign"
	// CouchListMemoriesView : Name of the view which fetches the list of memories.
	CouchListMemoriesView = "listMemories"
)

const (
	// MemTitleMaxLen : Max length for a memory title.
	MemTitleMaxLen = 100

	// MemTitleMinLen : Min length for a memory title.
	MemTitleMinLen = 1

	// MemBodyMaxLen : Max length for a memory body.
	MemBodyMaxLen = 1000

	// MemBodyMinLen : Min length for a memory body.
	MemBodyMinLen = 0
)
