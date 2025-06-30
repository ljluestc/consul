type Backend interface {
	Read(ctx context.Context, r *BackendReadRequest) (*BackendReadResponse, error)
	ReadAll(ctx context.Context, r *BackendReadAllRequest) (*BackendReadAllResponse, error)
	Write(ctx context.Context, r *BackendWriteRequest) (*BackendWriteResponse, error)
	WriteCAS(ctx context.Context, r *BackendWriteRequest) (*BackendWriteResponse, error)
	Delete(ctx context.Context, r *BackendDeleteRequest) (*BackendDeleteResponse, error)
}
