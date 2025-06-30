// Implement the Backend interface
func (b *InmemBackend) Read(ctx context.Context, r *BackendReadRequest) (*BackendReadResponse, error) {
	// existing implementation
}

func (b *InmemBackend) ReadAll(ctx context.Context, r *BackendReadAllRequest) (*BackendReadAllResponse, error) {
	// existing implementation
}

func (b *InmemBackend) Write(ctx context.Context, r *BackendWriteRequest) (*BackendWriteResponse, error) {
	// existing implementation
}

func (b *InmemBackend) WriteCAS(ctx context.Context, r *BackendWriteRequest) (*BackendWriteResponse, error) {
	// Implement compare-and-swap logic here
	// This is a basic implementation - customize based on actual requirements
	return b.Write(ctx, r)
}

func (b *InmemBackend) Delete(ctx context.Context, r *BackendDeleteRequest) (*BackendDeleteResponse, error) {
	// existing implementation
}
